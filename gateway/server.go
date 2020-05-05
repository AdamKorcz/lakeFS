package gateway

import (
	"context"
	"net/http"

	"github.com/treeverse/lakefs/logging"
	"github.com/treeverse/lakefs/stats"

	"github.com/treeverse/lakefs/httputil"

	"github.com/treeverse/lakefs/block"
	"github.com/treeverse/lakefs/gateway/utils"
	"github.com/treeverse/lakefs/index"

	"github.com/treeverse/lakefs/permissions"

	"github.com/treeverse/lakefs/auth"
	"github.com/treeverse/lakefs/db"
	"github.com/treeverse/lakefs/gateway/errors"
	"github.com/treeverse/lakefs/gateway/operations"
	"github.com/treeverse/lakefs/gateway/sig"
	"golang.org/x/xerrors"
)

type ServerContext struct {
	region           string
	bareDomain       string
	meta             index.Index
	multipartManager index.MultipartManager
	blockStore       block.Adapter
	authService      utils.GatewayAuthService
	stats            stats.Collector
}

func (c *ServerContext) WithContext(ctx context.Context) *ServerContext {
	return &ServerContext{
		region:           c.region,
		bareDomain:       c.bareDomain,
		meta:             c.meta.WithContext(ctx),
		multipartManager: c.multipartManager.WithContext(ctx),
		blockStore:       c.blockStore.WithContext(ctx),
		authService:      c.authService,
		stats:            c.stats,
	}
}

type Server struct {
	ctx        *ServerContext
	Server     *http.Server
	bareDomain string
}

func NewServer(
	region string,
	meta index.Index,
	blockStore block.Adapter,
	authService utils.GatewayAuthService,
	multipartManager index.MultipartManager,
	listenAddr, bareDomain string,
	stats stats.Collector,
) *Server {
	ctx := &ServerContext{
		meta:             meta,
		region:           region,
		bareDomain:       bareDomain,
		blockStore:       blockStore,
		authService:      authService,
		multipartManager: multipartManager,
		stats:            stats,
	}

	// setup routes
	var handler http.Handler
	handler = &Handler{
		BareDomain:         bareDomain,
		ctx:                ctx,
		NotFoundHandler:    http.HandlerFunc(notFound),
		ServerErrorHandler: nil,
	}
	handler = utils.RegisterRecorder(
		httputil.LoggingMiddleWare(
			"X-Amz-Request-Id", logging.Fields{"service_name": "s3_gateway"}, handler,
		),
	)

	// assemble Server
	return &Server{
		ctx:        ctx,
		bareDomain: bareDomain,
		Server: &http.Server{
			Handler: handler,
			Addr:    listenAddr,
		},
	}
}

func (s *Server) Listen() error {
	return s.Server.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	s.Server.SetKeepAlivesEnabled(false)
	return s.Server.Shutdown(ctx)
}

func getApiErrOrDefault(err error, defaultApiErr errors.APIErrorCode) errors.APIError {
	apiError, ok := err.(*errors.APIErrorCode)
	if ok {
		return apiError.ToAPIErr()
	} else {
		return defaultApiErr.ToAPIErr()
	}
}
func authenticateOperation(s *ServerContext, writer http.ResponseWriter, request *http.Request, action permissions.Action) *operations.AuthenticatedOperation {
	o := &operations.Operation{
		Request:        request,
		ResponseWriter: writer,
		Region:         s.region,
		FQDN:           s.bareDomain,

		Index:            s.meta,
		MultipartManager: s.multipartManager,
		BlockStore:       s.blockStore,
		Auth:             s.authService,
		Incr:             func(action string) { s.stats.Collect("s3_gateway", action) },
	}
	// authenticate
	authenticator := sig.ChainedAuthenticator(
		sig.NewV4Authenticator(request),
		sig.NewV2SigAuthenticator(request))

	authContext, err := authenticator.Parse()
	if err != nil {
		o.Log().WithError(err).WithFields(logging.Fields{
			"key": authContext.GetAccessKeyId(),
		}).Warn("error parsing signature")
		o.EncodeError(getApiErrOrDefault(err, errors.ErrAccessDenied))
		return nil
	}
	creds, err := s.authService.GetAPICredentials(authContext.GetAccessKeyId())
	if err != nil {
		if !xerrors.Is(err, db.ErrNotFound) {
			o.Log().WithError(err).WithField("key", authContext.GetAccessKeyId()).Warn("error getting access key")
			o.EncodeError(errors.Codes.ToAPIErr(errors.ErrInternalError))
		} else {
			o.Log().WithError(err).WithField("key", authContext.GetAccessKeyId()).Warn("could not find access key")
			o.EncodeError(errors.Codes.ToAPIErr(errors.ErrAccessDenied))
		}
		return nil
	}

	err = authenticator.Verify(creds, s.bareDomain)
	if err != nil {
		o.Log().WithError(err).WithFields(logging.Fields{
			"key":           authContext.GetAccessKeyId(),
			"authenticator": authenticator,
		}).Warn("error verifying credentials for key")
		o.EncodeError(getApiErrOrDefault(err, errors.ErrAccessDenied))
		return nil
	}

	// we are verified!
	op := &operations.AuthenticatedOperation{
		Operation:   o,
		SubjectId:   *creds.UserId,
		SubjectType: creds.Type,
	}

	// interpolate arn string
	arn := action.Arn

	// authorize
	authResp, err := s.authService.Authorize(&auth.AuthorizationRequest{
		UserID:     op.SubjectId,
		Permission: action.Permission,
		SubjectARN: arn,
	})
	if err != nil {
		o.Log().WithError(err).Error("failed to authorize")
		o.EncodeError(errors.Codes.ToAPIErr(errors.ErrInternalError))
		return nil
	}

	if authResp.Error != nil || !authResp.Allowed {
		o.Log().WithError(authResp.Error).WithField("key", authContext.GetAccessKeyId()).Warn("no permission")
		o.EncodeError(errors.Codes.ToAPIErr(errors.ErrAccessDenied))
		return nil
	}

	// authentication and authorization complete!
	return op
}

func OperationHandler(ctx *ServerContext, handler operations.AuthenticatedOperationHandler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		// structure operation
		action := handler.Action("", "", "")
		authOp := authenticateOperation(ctx.WithContext(request.Context()), writer, request, action)
		if authOp == nil {
			return
		}
		// run callback
		handler.Handle(authOp)
	})
}

func RepoOperationHandler(ctx *ServerContext, repoId string, handler operations.RepoOperationHandler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		// structure operation
		action := handler.Action(repoId, "", "")
		authOp := authenticateOperation(ctx.WithContext(request.Context()), writer, request, action)
		if authOp == nil {
			return
		}

		// validate repo exists
		repo, err := authOp.Index.GetRepo(repoId)
		if xerrors.Is(err, db.ErrNotFound) {
			authOp.Log().WithField("repository", repoId).Warn("the specified repo does not exist")
			authOp.EncodeError(errors.ErrNoSuchBucket.ToAPIErr())
			return
		} else if err != nil {
			authOp.EncodeError(errors.ErrInternalError.ToAPIErr())
			return
		}
		// run callback
		repoOperation := &operations.RepoOperation{
			AuthenticatedOperation: authOp,
			Repo:                   repo,
		}
		repoOperation.AddLogFields(logging.Fields{
			"repository": repo.Id,
		})
		handler.Handle(repoOperation)
	})
}

func PathOperationHandler(ctx *ServerContext, repoId, refId, path string, handler operations.PathOperationHandler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		// structure operation
		action := handler.Action(repoId, refId, path)
		authOp := authenticateOperation(ctx.WithContext(request.Context()), writer, request, action)
		if authOp == nil {
			return
		}

		// validate repo exists
		repo, err := authOp.Index.GetRepo(repoId)
		if xerrors.Is(err, db.ErrNotFound) {
			authOp.Log().WithField("repository", repoId).Warn("the specified repo does not exist")
			authOp.EncodeError(errors.Codes.ToAPIErr(errors.ErrNoSuchBucket))
			return
		} else if err != nil {
			authOp.EncodeError(errors.Codes.ToAPIErr(errors.ErrInternalError))
			return
		}

		// run callback
		operation := &operations.PathOperation{
			RefOperation: &operations.RefOperation{
				RepoOperation: &operations.RepoOperation{
					AuthenticatedOperation: authOp,
					Repo:                   repo,
				},
				Ref: refId,
			},
			Path: path,
		}
		operation.AddLogFields(logging.Fields{
			"repository": repo.Id,
			"ref":        refId,
			"path":       path,
		})
		handler.Handle(operation)
	})
}

func unsupportedOperationHandler() http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		o := &operations.Operation{
			Request:        request,
			ResponseWriter: writer,
		}
		o.EncodeError(errors.ERRLakeFSNotSupported.ToAPIErr())
		return
	})
}

func notFound(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
}
