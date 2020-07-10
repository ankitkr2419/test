import { runExperimentActions, stopExperimentActions, experimentCompleteActions } from 'actions/runExperimentActions';

export const runExperiment = experimentId => ({
	type: runExperimentActions.runExperiment,
	payload: {
		experimentId,
	},
});

export const runExperimentFailed = errorResponse => ({
	type: runExperimentActions.failureAction,
	payload: {
		...errorResponse,
		error: true,
	},
});


export const stopExperiment = experimentId => ({
	type: stopExperimentActions.stopExperiment,
	payload: {
		experimentId,
	},
});

export const stopExperimentFailed = errorResponse => ({
	type: stopExperimentActions.failureAction,
	payload: {
		...errorResponse,
		error: true,
	},
});

export const experimentedCompleted = data => ({
	type: experimentCompleteActions.success,
	payload: {
		data,
	},
});
