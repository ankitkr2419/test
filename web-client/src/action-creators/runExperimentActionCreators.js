import {
  runExperimentInProgressActions,
  runExperimentActions,
  stopExperimentActions,
  experimentCompleteActions,
} from "actions/runExperimentActions";

export const runExperimentInProgress = (progressDetails) => ({
  type: runExperimentInProgressActions.runExperimentProgressAction,
  payload: { progressDetails },
});

export const runExperimentSuccess = (progressSucceeded) => ({
  type: runExperimentInProgressActions.runExperimentProgressSuccessAction,
  payload: { progressSucceeded },
});

export const runExperiment = (experimentId, token) => ({
  type: runExperimentActions.runExperiment,
  payload: {
    experimentId,
    token,
  },
});

export const runExperimentFailed = (errorResponse) => ({
  type: runExperimentActions.failureAction,
  payload: {
    ...errorResponse,
    error: true,
  },
});

// abort experiment, Will call stop api call
export const stopExperiment = (experimentId, token) => ({
  type: stopExperimentActions.stopExperiment,
  payload: {
    experimentId,
    token,
  },
});

export const stopExperimentFailed = (errorResponse) => ({
  type: stopExperimentActions.failureAction,
  payload: {
    ...errorResponse,
    error: true,
  },
});

// web socket received message => experiment completed
export const experimentedCompleted = (data) => ({
  type: experimentCompleteActions.success,
  payload: {
    data,
  },
});

// web socket received message => experiment failed
export const experimentedFailed = (data) => ({
  type: experimentCompleteActions.failed,
  payload: {
    data,
  },
});
