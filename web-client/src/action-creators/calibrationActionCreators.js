import {
  calibrationActions,
  updateCalibrationActions,
  pidProgressActions,
  pidActions,
} from "actions/calibrationActions";

//fetch calibration configurations
export const calibrationInitiated = (token) => ({
  type: calibrationActions.calibrationInitiated,
  payload: {
    token,
  },
});

export const calibrationFailed = ({ error }) => ({
  type: calibrationActions.calibrationFailure,
  payload: {
    error,
  },
});

//update calibration configurations
export const updateCalibrationInitiated = ({ token, data }) => ({
  type: updateCalibrationActions.updateCalibrationInitiated,
  payload: {
    token,
    data,
  },
});

export const updateCalibrationFailed = ({ error }) => ({
  type: updateCalibrationActions.updateCalibrationFailure,
  payload: {
    error,
  },
});

//websocket PID action creators

export const runPidInProgress = (progressDetails) => ({
  type: pidProgressActions.pidProgressAction,
  payload: { progressDetails },
});

export const runPidInSuccess = (progressSucceeded) => ({
  type: pidProgressActions.pidProgressActionSuccess,
  payload: { progressSucceeded },
});

export const runPid = (token, deckName) => ({
  type: pidActions.pidActionInitiated,
  payload: {
    token,
    deckName,
  },
});

export const runPidFailed = (errorResponse) => ({
  type: pidActions.pidActionFailure,
  payload: {
    ...errorResponse,
    error: true,
  },
});
