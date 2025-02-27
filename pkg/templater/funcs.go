package templater

import (
	"errors"
	"fmt"
	"text/template"

	"github.com/treeverse/lakefs/pkg/auth"
	"github.com/treeverse/lakefs/pkg/permissions"
)

var (
	ErrNotAuthorized = errors.New("unauthorized")
	// ErrTemplateFailed is an error generated by the "fail" template function.
	ErrTemplateFailed = errors.New("")
)

type Credentials struct {
	Key, Secret string
}

var templateFuncs = template.FuncMap{
	"new_credentials": newCredentials,
	"contenttype":     contentType,
	"fail":            fail,
}

// newCredentials creates new credentials for the user and returns them.
// name is used to cache the generated new credentials during a single call
// (between prepare and expand phases, but also between different parts of
// the template).
func newCredentials(params *ControlledParams, name string) (*Credentials, error) {
	resp, err := params.Auth.Authorize(params.Ctx, &auth.AuthorizationRequest{
		Username: params.User.Username,
		RequiredPermissions: permissions.Node{
			Permission: permissions.Permission{
				Action:   permissions.CreateCredentialsAction,
				Resource: permissions.UserArn(params.User.Username),
			},
		},
	})
	if err == nil && resp.Error != nil {
		err = resp.Error
	}
	if err != nil {
		return nil, fmt.Errorf("create credentials for %+v: %w", params.User, err)
	}
	if !resp.Allowed {
		return nil, fmt.Errorf("create credentials for %+v: %w", params.User, ErrNotAuthorized)
	}

	var (
		generatedCredsI interface{}
		ok              bool
		generatedCreds  map[string]*Credentials
	)

	generatedCredsI, ok = params.Store["new_credentials"]
	if !ok {
		generatedCreds = make(map[string]*Credentials)
		params.Store["new_credentials"] = generatedCreds
	} else {
		generatedCreds = generatedCredsI.(map[string]*Credentials)
	}

	if retCreds, ok := generatedCreds[name]; ok {
		return retCreds, nil
	}

	// TODO(ariels): monitor!

	// TODO(ariels): Name these credentials (once auth allows named credentials...).
	credentials, err := params.Auth.CreateCredentials(params.Ctx, params.User.Username)
	if err != nil {
		return nil, fmt.Errorf("create credentials for %+v: %w", params.User, err)
	}

	retCreds := &Credentials{Key: credentials.AccessKeyID, Secret: credentials.SecretAccessKey}
	generatedCreds[name] = retCreds
	return retCreds, nil
}

// fail fails template expansion with the message passed in.
func fail(_ *ControlledParams, msg string) (string, error) {
	return msg, fmt.Errorf("%s%w", msg, ErrTemplateFailed)
}

// contentType sets the content type of the response.
func contentType(params *ControlledParams, contentType string) (string, error) {
	if params.Phase == PhasePrepare {
		params.Header.Add("Content-Type", contentType)
	}
	return "", nil
}

// TODO(ariels): config, object
