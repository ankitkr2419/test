import { createTemplateActions, listTemplateActions } from "actions/templateActions";

export const createTemplate = (body) => ({
  type: createTemplateActions.createAction,
  payload: {
    body
  }
})

export const createTemplateFailed = (errorResponse) => ({
  type: createTemplateActions.failureAction,
  payload: {
    ...errorResponse,
    error: true
  }
});

export const fetchTemplates = (body) => ({
  type: listTemplateActions.listAction,
  payload: {
    body
  }
})

export const fetchTemplatesFailed = (errorResponse) => ({
  type: listTemplateActions.failureAction,
  payload: {
    ...errorResponse,
    error: true
  }
});
