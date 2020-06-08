import {
  saveTargetActions,
  listTargetActions,
  listTargetByTemplateIDActions,
} from "actions/targetActions";

export const saveTarget = (templateID, body) => ({
  type: saveTargetActions.saveAction,
  payload: {
    templateID,
    body,
  },
});

export const saveTargetFailed = (errorResponse) => ({
  type: saveTargetActions.failureAction,
  payload: {
    ...errorResponse,
    error: true,
  },
});

export const fetchMasterTargets = () => ({
  type: listTargetActions.listAction,
});

export const fetchMasterTargetsFailed = (errorResponse) => ({
  type: listTargetActions.failureAction,
  payload: {
    ...errorResponse,
    error: true,
  },
});

export const fetchTargetsByTemplateID = (templateID) => ({
  type: listTargetByTemplateIDActions.listAction,
  payload: {
    templateID,
  },
});

export const fetchTargetsByTemplateIDFailed = (errorResponse) => ({
  type: listTargetByTemplateIDActions.failureAction,
  payload: {
    ...errorResponse,
    error: true,
  },
});
