import {
  createTemplateActions,
  listTemplateActions,
  updateTemplateActions,
  deleteTemplateActions,
} from "actions/templateActions";

export const createTemplate = (body) => ({
  type: createTemplateActions.createAction,
  payload: {
    body,
  },
});

export const createTemplateFailed = (errorResponse) => ({
  type: createTemplateActions.failureAction,
  payload: {
    ...errorResponse,
    error: true,
  },
});

export const fetchTemplates = (body) => ({
  type: listTemplateActions.listAction,
  payload: {
    body,
  },
});

export const fetchTemplatesFailed = (errorResponse) => ({
  type: listTemplateActions.failureAction,
  payload: {
    ...errorResponse,
    error: true,
  },
});

export const updateTemplate = (templateID, body) => ({
  type: updateTemplateActions.updateTemplate,
  payload: {
    templateID,
    body,
  },
});

export const updateTemplateFailed = (errorResponse) => ({
  type: updateTemplateActions.failureAction,
  payload: {
    ...errorResponse,
    error: true,
  },
});

export const deleteTemplate = (templateID) => ({
  type: deleteTemplateActions.deleteAction,
  payload: {
    templateID,
  },
});

export const deleteTemplateFailed = (errorResponse) => ({
  type: deleteTemplateActions.deleteTemplateFailed,
  payload: {
    ...errorResponse,
    error: true,
  },
});
