export const runExperimentInProgressActions = {
  runExperimentProgressAction: "RUN_EXPERIMENT_IN_PROGRESS",
  runExperimentProgressSuccessAction: "RUN_PROGRESS_EXPERIMENT_SUCCEEDED",
};

export const runExperimentActions = {
  runExperiment: "RUN_EXPERIMENT_INITIATED",
  successAction: "RUN_EXPERIMENT_SUCCEEDED",
  failureAction: "RUN_EXPERIMENT_FAILURE",
};

export const stopExperimentActions = {
  stopExperiment: "STOP_EXPERIMENT_INITIATED",
  successAction: "STOP_EXPERIMENT_SUCCEEDED",
  failureAction: "STOP_EXPERIMENT_FAILURE",
};

// web socket actions for experiment
export const experimentCompleteActions = {
  success: "EXPERIMENT_COMPLETED",
  failed: "EXPERIMENT_FAILED",
  error: "EXPERIMENT_ERROR",
};
