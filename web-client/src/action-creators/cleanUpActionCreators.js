import { cleanUpActions } from "actions/cleanUpActions";

export const cleanUpActionInitiated = (params) => ({
  type: cleanUpActions.cleanUpActionInitiated,
  payload: {
    params
  }
});

export const cleanUpActionSuccess = (cleanUpData) => ({
  type: cleanUpActions.cleanUpActionSuccess,
  payload: {
    cleanUpData
  }
});

export const cleanUpActionFailed = (serverErrors) => ({
  type: cleanUpActions.cleanUpActionFailed,
  payload: {
    serverErrors
  }
});

export const cleanUpActionInProgress = (cleanUpActionInProgress) => ({
  type: cleanUpActions.cleanUpActionInProgress,
  payload: {
    cleanUpActionInProgress
  }
});

export const cleanUpActionInCompleted = (cleanUpActionInCompleted) => ({
  type: cleanUpActions.cleanUpActionInCompleted,
  payload: {
    cleanUpActionInCompleted
  }
});
