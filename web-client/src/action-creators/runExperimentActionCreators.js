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

// abort experiment, Will call stop api call
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

// web socket received message => experiment completed
export const experimentedCompleted = data => ({
	type: experimentCompleteActions.success,
	payload: {
		data,
	},
});

// web socket received message => experiment failed
export const experimentedFailed = data => ({
	type: experimentCompleteActions.failed,
	payload: {
		data,
	},
});
