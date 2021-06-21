import {
	createTemplateActions,
	listTemplateActions,
	updateTemplateActions,
	deleteTemplateActions,
} from 'actions/templateActions';

export const createTemplate = (body, token) => ({
	type: createTemplateActions.createAction,
	payload: {
		body,
		token
	},
});

export const createTemplateFailed = errorResponse => ({
	type: createTemplateActions.failureAction,
	payload: {
		...errorResponse,
		error: true,
	},
});

export const addTemplateReset = () => ({
	type: createTemplateActions.createTemplateReset,
});

export const fetchTemplates = (params) => ({
	type: listTemplateActions.listAction,
	payload: {
		...params
	}
});

export const fetchTemplatesFailed = errorResponse => ({
	type: listTemplateActions.failureAction,
	payload: {
		...errorResponse,
		error: true,
	},
});

export const updateTemplate = (templateID, body) => ({
	type: updateTemplateActions.updateAction,
	payload: {
		templateID,
		body,
	},
});

export const updateTemplateFailed = errorResponse => ({
	type: updateTemplateActions.failureAction,
	payload: {
		...errorResponse,
		error: true,
	},
});

export const updateTemplateReset = () => ({
	type: updateTemplateActions.updateTemplateReset,
});

export const deleteTemplate = templateID => ({
	type: deleteTemplateActions.deleteAction,
	payload: {
		templateID,
	},
});

export const deleteTemplateFailed = errorResponse => ({
	type: deleteTemplateActions.failureAction,
	payload: {
		...errorResponse,
		error: true,
	},
});

export const deleteTemplateReset = () => ({
	type: deleteTemplateActions.deleteTemplateReset,
});
