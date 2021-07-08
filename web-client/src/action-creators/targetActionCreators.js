import {
	saveTargetActions,
	listTargetActions,
	listTargetByTemplateIDActions,
} from 'actions/targetActions';

export const saveTarget = (templateID, body, token) => ({
	type: saveTargetActions.saveAction,
	payload: {
		templateID,
		body,
		token
	},
});

export const saveTargetFailed = errorResponse => ({
	type: saveTargetActions.failureAction,
	payload: {
		...errorResponse,
		error: true,
	},
});

export const resetSaveTarget = () => ({
	type: saveTargetActions.saveTargetReset,
});

export const fetchMasterTargets = (token) => ({
	type: listTargetActions.listAction,
	payload: {
		token
	}
});

export const fetchMasterTargetsFailed = errorResponse => ({
	type: listTargetActions.failureAction,
	payload: {
		...errorResponse,
		error: true,
	},
});

export const fetchTargetsByTemplateID = (templateID, token) => ({
	type: listTargetByTemplateIDActions.listAction,
	payload: {
		templateID,
		token
	},
});

export const fetchTargetsByTemplateIDFailed = errorResponse => ({
	type: listTargetByTemplateIDActions.failureAction,
	payload: {
		...errorResponse,
		error: true,
	},
});
