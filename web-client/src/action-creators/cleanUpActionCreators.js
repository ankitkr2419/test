import {
  runCleanUpAction,
  pauseCleanUpAction,
  resumeCleanUpAction,
  abortCleanUpAction,
} from "actions/cleanUpActions";

export const runCleanUpActionInitiated = (params) => ({
  type: runCleanUpAction.runCleanUpInitiated,
  payload: {
    params,
  },
});

export const runCleanUpActionSuccess = (cleanUpData) => ({
  type: runCleanUpAction.runCleanUpSuccess,
  payload: {
    cleanUpData,
  },
});

export const runCleanUpActionFailed = (serverErrors) => ({
  type: runCleanUpAction.runCleanUpFailed,
  payload: {
    serverErrors,
  },
});

export const runCleanUpActionReset = () => ({
  type: runCleanUpAction.runCleanUpReset,
  payload: {},
});

export const runCleanUpActionInProgress = (cleanUpActionInProgress) => ({
  type: runCleanUpAction.runCleanUpInProgress,
  payload: {
    cleanUpActionInProgress,
  },
});

export const runCleanUpActionInCompleted = (cleanUpActionInCompleted) => ({
  type: runCleanUpAction.runCleanUpInCompleted,
  payload: {
    cleanUpActionInCompleted,
  },
});

export const pauseCleanUpActionInitiated = (params) => ({
  type: pauseCleanUpAction.pauseCleanUpInitiated,
  payload: {
    params,
  },
});

export const pauseCleanUpActionSuccess = (pauseCleanUpResponse) => ({
  type: pauseCleanUpAction.pauseCleanUpSuccess,
  payload: {
    pauseCleanUpResponse,
  },
});

export const pauseCleanUpActionFailed = (serverErrors) => ({
  type: pauseCleanUpAction.pauseCleanUpFailed,
  payload: {
    serverErrors,
  },
});

export const pauseCleanUpActionReset = () => ({
  type: pauseCleanUpAction.pauseCleanUpReset,
  payload: {},
});

export const resumeCleanUpActionInitiated = (params) => ({
  type: resumeCleanUpAction.resumeCleanUpInitiated,
  payload: {
    params,
  },
});

export const resumeCleanUpActionSuccess = (resumeCleanUpResponse) => ({
  type: resumeCleanUpAction.resumeCleanUpSuccess,
  payload: {
    resumeCleanUpResponse,
  },
});

export const resumeCleanUpActionFailed = (serverErrors) => ({
  type: resumeCleanUpAction.resumeCleanUpFailed,
  payload: {
    serverErrors,
  },
});

export const resumeCleanUpActionReset = () => ({
  type: resumeCleanUpAction.resumeCleanUpReset,
  payload: {},
});

export const abortCleanUpActionInitiated = (params) => ({
  type: abortCleanUpAction.abortCleanUpInitiated,
  payload: {
    params,
  },
});

export const abortCleanUpActionSuccess = (abortCleanUpResponse) => ({
  type: abortCleanUpAction.abortCleanUpSuccess,
  payload: {
    abortCleanUpResponse,
  },
});

export const abortCleanUpActionFailed = (serverErrors) => ({
  type: abortCleanUpAction.abortCleanUpFailed,
  payload: {
    serverErrors,
  },
});

export const abortCleanUpActionReset = () => ({
  type: abortCleanUpAction.abortCleanUpReset,
  payload: {},
});
