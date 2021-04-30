export const runCleanUpAction = {
  runCleanUpInitiated: "RUN_CLEAN_UP_INITIATED",
  runCleanUpSuccess: "RUN_CLEAN_UP_SUCCESS",
  runCleanUpFailed: "RUN_CLEAN_UP_FAILED",
  runCleanUpReset: "RUN_CLEAN_UP_RESET",
  runCleanUpInProgress: "PROGRESS_CLEANUP",
  runCleanUpInCompleted: "SUCCESS_CLEANUP",
};

export const pauseCleanUpAction = {
  pauseCleanUpInitiated: "PAUSE_CLEAN_UP_INITIATED",
  pauseCleanUpSuccess: "PAUSE_CLEAN_UP_SUCCESS",
  pauseCleanUpFailed: "PAUSE_CLEAN_UP_FAILED",
  pauseCleanUpReset: "PAUSE_CLEAN_UP_RESET",
};

export const resumeCleanUpAction = {
  resumeCleanUpInitiated: "RESUME_CLEAN_UP_INITIATED",
  resumeCleanUpSuccess: "RESUME_CLEAN_UP_SUCCESS",
  resumeCleanUpFailed: "RESUME_CLEAN_UP_FAILED",
  resumeCleanUpReset: "RESUME_CLEAN_UP_RESET",
  resumeCleanUpInProgress: "PROGRESS_RESUME_CLEAN_UP",
  resumeCleanUpInCompleted: "SUCCESS_RESUME_CLEAN_UP"
};

export const abortCleanUpAction = {
  abortCleanUpInitiated: "ABORT_CLEAN_UP_INITIATED",
  abortCleanUpSuccess: "ABORT_CLEAN_UP_SUCCESS",
  abortCleanUpFailed: "ABORT_CLEAN_UP_FAILED",
  abortCleanUpReset: "ABORT_CLEAN_UP_RESET",
};

export const setCleanUpHoursAction = {
  setHours: "SET_HOURS",
  resetHours: "RESET_HOURS",
};

export const setCleanUpMinsAction = {
  setMins: "SET_MINS",
  resetMins: "RESET_MINS",
};

export const setCleanUpSecsAction = {
  setSecs: "SET_SECS",
  resetSecs: "RESET_SECS",
};

export const setShowCleanUpAction = {
  setShowCleanUp: "SET_SHOW_CLEAN_UP",
  resetShowCleanUp: "RESET_SHOW_CLEAN_UP",
};
