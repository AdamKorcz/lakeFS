import {Wizard} from "./wizard";
import React, {useState} from "react";
import {useAPI} from "../../../lib/hooks/api";
import {config} from "../../../lib/api";
import {Loading, ProgressSpinner} from "../../../lib/components/controls";
import Alert from "react-bootstrap/Alert";
import RepositoryCreateForm from "../../../lib/components/repositoryCreateForm";
import ImportDataStep from "./import_data_wizard_step";
import {SparkConfigStep} from "./spark_config_step";
import {useRouter} from "../../../lib/hooks/router";

const RepositoryCreationPhase = {
    NotStarted: 0,
    InProgress: 1,
    Completed: 2,
    Failed: 3,
}

const RepositoryCreationStep = ({repoCreationError, createRepo, onCancel, onComplete}) => {
    const {response, error: err, loading} = useAPI(() => config.getStorageConfig());
    const [repoName, setRepoName] = useState('');
    const [repoCreationPhase, setRepoCreationPhase] = useState(RepositoryCreationPhase.NotStarted);

    const onSubmit = async (repo) => {
        setRepoCreationPhase(RepositoryCreationPhase.InProgress);
        const success = await createRepo(repo, false);
        if (success) {
            setRepoCreationPhase(RepositoryCreationPhase.Completed);
            onComplete({ 'branch': repo.default_branch, 'namespace': repo.storage_namespace, 'repoId': repo.name },);
            setRepoName(repo.name);
        }
        else {
            setRepoCreationPhase(RepositoryCreationPhase.Failed);
        }
    }
    const showError = repoCreationError ? repoCreationError : err;
    let content;
    if (loading) {
        content = <Loading/>;
    } else if (repoCreationPhase === RepositoryCreationPhase.InProgress) {
        content = <ProgressSpinner text={'Creating repository...'} />;
    } else if (repoCreationPhase === RepositoryCreationPhase.Completed) {
        content = <Alert variant="info" className={"mt-3"}>Repository <span className={"font-weight-bold"}>{repoName}</span> created successfully</Alert>;
    }
    else {
        content = <RepositoryCreateForm config={response} error={showError} onSubmit={onSubmit} onCancel={onCancel} sm={8}/>;
    }
    return (
        <>
            <h3 className={"wizard-step-header"}>Create a repository</h3>
            {content}
        </>
    );
}

export const SparkQuickstart = ({onExit, createRepo, repoCreationError}) => {
    const [state, setState] = useState({
        branch: '',
        namespace: '',
        repoId: '',
    });
    const router = useRouter();
    const [completed, setCompleted] = useState(new Set());
    const completedStep = (vals = {}, stepNum) => {
        setState({...state, ...vals});
        setCompleted(currentCompleted => new Set(currentCompleted).add(stepNum));
    }
    const onComplete = async () => {
        // TODO(jonathan): generate README.md file, upload it
        // TODO(jonathan): commit README.md file to repo
        router.push({pathname: `/repositories/:repoId/objects`, params: {repoId: state.repoId}});
        onExit();
    }
    const steps = [
        {
            label: 'Create Repository',
            component: <RepositoryCreationStep
                stepNum
                repoCreationError={repoCreationError}
                createRepo={createRepo}
                onCancel={onExit}
                onComplete={(values) => {
                    completedStep(values, 0);
                }} />,
        },
        {
            label: 'Import Data',
            optional: true,
            component: <ImportDataStep
                repoId={state.repoId}
                onComplete={(values) => {
                    completedStep(values, 1);
                }}
                branchName={state.branch} />,
        },
        {
            label: 'Spark Configurations',
            component: <SparkConfigStep onComplete={(values) => { completedStep(values, 2); }}/>
        },
    ];
    return (
        <>
            <Wizard steps={steps} isShowBack={false} completed={completed} onDone={onComplete}/>
        </>
    );
}
