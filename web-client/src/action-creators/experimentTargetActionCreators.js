import {
	createExperimentTargetActions,
	listExperimentTargetActions,
} from 'actions/experimentTargetActions';

export const createExperimentTarget = (body, experimentId) => ({
	type: createExperimentTargetActions.createAction,
	payload: {
		body,
		experimentId,
	},
});

export const createExperimentTargetFailed = errorResponse => ({
	type: createExperimentTargetActions.failureAction,
	payload: {
		...errorResponse,
		error: true,
	},
});

export const createExperimentTargetReset = () => ({
	type: createExperimentTargetActions.createExperimentTargetReset,
});

export const fetchExperimentTargets = experimentId => ({
	type: listExperimentTargetActions.listAction,
	payload: {
		experimentId,
	},
});

export const fetchExperimentTargetsFailed = errorResponse => ({
	type: listExperimentTargetActions.failureAction,
	payload: {
		...errorResponse,
		error: true,
	},
});

export const updateExperimentTargetFilters = (index, key, value) => ({
	type: listExperimentTargetActions.updateGraphFilters,
	payload: {
		index,
		key,
		value,
	},
});
