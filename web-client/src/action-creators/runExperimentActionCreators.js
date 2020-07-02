import runActions from 'actions/runExperimentActions';

export const runExperiment = experimentId => ({
	type: runActions.runExperiment,
	payload: {
		experimentId,
	},
});

export const runExperimentFailed = errorResponse => ({
	type: runActions.failureAction,
	payload: {
		...errorResponse,
		error: true,
	},
});
