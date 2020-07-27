import {
	addStepActions,
	listStepActions,
	updateStepActions,
	deleteStepActions,
	listHoldStepActions,
	listCycleStepActions,
} from 'actions/stepActions';

export const addStep = body => ({
	type: addStepActions.addAction,
	payload: {
		body,
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

export const fetchHoldSteps = stageId => ({
	type: listHoldStepActions.listAction,
	payload: {
		stageId,
	},
});

export const fetchHoldStepsFailed = errorResponse => ({
	type: listHoldStepActions.failureAction,
	payload: {
		...errorResponse,
		error: true,
	},
});

export const fetchCycleSteps = stageId => ({
	type: listCycleStepActions.listAction,
	payload: {
		stageId,
	},
});

export const fetchCycleStepsFailed = errorResponse => ({
	type: listCycleStepActions.failureAction,
	payload: {
		...errorResponse,
		error: true,
	},
});

export const updateStep = (stepId, body) => ({
	type: updateStepActions.updateAction,
	payload: {
		stepId,
		body,
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

export const deleteStep = stepId => ({
	type: deleteStepActions.deleteAction,
	payload: {
		stepId,
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
