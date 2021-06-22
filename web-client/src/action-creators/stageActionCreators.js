import {
	listStageActions,
	updateStageActions,
} from 'actions/stageActions';

export const fetchStages = (templateID, token) => ({
	type: listStageActions.listAction,
	payload: {
		templateID,
		token
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
