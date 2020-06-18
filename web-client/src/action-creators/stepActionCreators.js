import {
	addStepActions,
	listStepActions,
	updateStepActions,
	deleteStepActions,
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

export const fetchSteps = stageId => ({
	type: listStepActions.listAction,
	payload: {
		stageId,
	},
});

export const fetchStepsFailed = errorResponse => ({
	type: listStepActions.failureAction,
	payload: {
		...errorResponse,
		error: true,
	},
});

export const updateStep = (stepID, body) => ({
	type: updateStepActions.updateAction,
	payload: {
		stepID,
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
