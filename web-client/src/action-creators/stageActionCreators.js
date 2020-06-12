import {
  addStageActions,
  listStageActions,
  updateStageActions,
  deleteStageActions,
} from "actions/stageActions";

export const addStage = (body) => ({
  type: addStageActions.addAction,
  payload: {
    body,
  },
});

export const addStageFailed = (errorResponse) => ({
  type: addStageActions.failureAction,
  payload: {
    ...errorResponse,
    error: true,
  },
});

export const fetchStages = (body) => ({
  type: listStageActions.listAction,
  payload: {
    body,
  },
});

export const fetchStagesFailed = (errorResponse) => ({
  type: listStageActions.failureAction,
  payload: {
    ...errorResponse,
    error: true,
  },
});

export const updateStage = (stageID, body) => ({
  type: updateStageActions.updateAction,
  payload: {
    stageID,
    body,
  },
});

export const updateStageFailed = (errorResponse) => ({
  type: updateStageActions.failureAction,
  payload: {
    ...errorResponse,
    error: true,
  },
});

export const deleteStage = (templateID) => ({
  type: deleteStageActions.deleteAction,
  payload: {
    templateID,
  },
});

export const deleteStageFailed = (errorResponse) => ({
  type: deleteStageActions.failureAction,
  payload: {
    ...errorResponse,
    error: true,
  },
});
