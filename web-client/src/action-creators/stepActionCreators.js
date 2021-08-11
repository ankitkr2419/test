import {
	addStepActions,
	updateStepActions,
	deleteStepActions,
	listHoldStepActions,
	listCycleStepActions,
} from 'actions/stepActions';

export const addStep = (body, token) => ({
	type: addStepActions.addAction,
	payload: {
		body,
		token
	},
});

export const addStepFailed = errorResponse => ({
	type: addStepActions.failureAction,
	payload: {
		...errorResponse,
		error: true,
	},
});

export const addStepReset = () => ({
	type: addStepActions.addStepReset,
});

export const fetchHoldSteps = (stageId, token) => ({
	type: listHoldStepActions.listAction,
	payload: {
		stageId,
		token
	},
});

export const fetchHoldStepsFailed = errorResponse => ({
	type: listHoldStepActions.failureAction,
	payload: {
		...errorResponse,
		error: true,
	},
});

export const fetchCycleSteps = (stageId, token) => ({
	type: listCycleStepActions.listAction,
	payload: {
		stageId,
		token
	},
});

export const fetchCycleStepsFailed = errorResponse => ({
	type: listCycleStepActions.failureAction,
	payload: {
		...errorResponse,
		error: true,
	},
});

export const updateStep = (stepId, body, token) => ({
	type: updateStepActions.updateAction,
	payload: {
		stepId,
		body,
		token
	},
});

export const updateStepFailed = errorResponse => ({
	type: updateStepActions.failureAction,
	payload: {
		...errorResponse,
		error: true,
	},
});

export const updateStepReset = () => ({
	type: updateStepActions.updateStepReset,
});

export const deleteStep = (stepId, token) => ({
	type: deleteStepActions.deleteAction,
	payload: {
		stepId,
		token
	},
});

export const deleteStepFailed = errorResponse => ({
	type: deleteStepActions.failureAction,
	payload: {
		...errorResponse,
		error: true,
	},
});

export const deleteStepReset = () => ({
	type: deleteStepActions.deleteStepReset,
});
