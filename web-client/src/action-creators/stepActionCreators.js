import {
  addStepActions,
  listStepActions,
  updateStepActions,
  deleteStepActions,
} from "actions/stepActions";

export const addStep = (body) => ({
  type: addStepActions.addAction,
  payload: {
    body,
  },
});

export const addStepFailed = (errorResponse) => ({
  type: addStepActions.failureAction,
  payload: {
    ...errorResponse,
    error: true,
  },
});

export const fetchSteps = (stageID) => ({
  type: listStepActions.listAction,
  payload: {
    stageID,
  },
});

export const fetchStepsFailed = (errorResponse) => ({
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

export const updateStepFailed = (errorResponse) => ({
  type: updateStepActions.failureAction,
  payload: {
    ...errorResponse,
    error: true,
  },
});

export const deleteStep = (stepID) => ({
  type: deleteStepActions.deleteAction,
  payload: {
    stepID,
  },
});

export const deleteStepFailed = (errorResponse) => ({
  type: deleteStepActions.failureAction,
  payload: {
    ...errorResponse,
    error: true,
  },
});
