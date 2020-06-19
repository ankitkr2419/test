import {
	addStageActions,
	listStageActions,
	updateStageActions,
	deleteStageActions,
} from 'actions/stageActions';

export const addStage = body => ({
	type: addStageActions.addAction,
	payload: {
		body,
	},
});

export const addStageFailed = errorResponse => ({
	type: addStageActions.failureAction,
	payload: {
		...errorResponse,
		error: true,
	},
});

export const addStageReset = () => ({
	type: addStageActions.addStageReset,
});

export const fetchStages = templateID => ({
	type: listStageActions.listAction,
	payload: {
		templateID,
	},
});

export const fetchStagesFailed = errorResponse => ({
	type: listStageActions.failureAction,
	payload: {
		...errorResponse,
		error: true,
	},
});

export const updateStage = (stageId, body) => ({
	type: updateStageActions.updateAction,
	payload: {
		stageId,
		body,
	},
});

export const updateStageReset = () => ({
	type: updateStageActions.updateStageReset,
});

export const updateStageFailed = errorResponse => ({
	type: updateStageActions.failureAction,
	payload: {
		...errorResponse,
		error: true,
	},
});

export const deleteStage = stageId => ({
	type: deleteStageActions.deleteAction,
	payload: {
		stageId,
	},
});

export const deleteStageReset = () => ({
	type: deleteStageActions.deleteStageReset,
});

export const deleteStageFailed = errorResponse => ({
	type: deleteStageActions.failureAction,
	payload: {
		...errorResponse,
		error: true,
	},
});
