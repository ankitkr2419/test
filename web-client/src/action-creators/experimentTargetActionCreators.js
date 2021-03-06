import {
	createExperimentTargetActions,
	listExperimentTargetActions,
} from 'actions/experimentTargetActions';

export const createExperimentTarget = (body, experimentId, token) => ({
	type: createExperimentTargetActions.createAction,
	payload: {
		body,
		experimentId,
		token
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

export const fetchExperimentTargets = (experimentId, token) => ({
	type: listExperimentTargetActions.listAction,
	payload: {
		experimentId,
		token
	},
});

export const fetchExperimentTargetsFailed = errorResponse => ({
	type: listExperimentTargetActions.failureAction,
	payload: {
		...errorResponse,
		error: true,
	},
});

export const resetExperimentTargets = () => ({
	type: listExperimentTargetActions.resetExperimentTargets,
});

export const updateExperimentTargetFilters = (index, key, value) => ({
	type: listExperimentTargetActions.updateGraphFilters,
	payload: {
		index,
		key,
		value,
	},
});
