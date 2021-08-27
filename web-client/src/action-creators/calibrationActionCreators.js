import {
  calibrationActions,
  updateCalibrationActions,
  pidProgressActions,
  pidActions,
  motorActions,
  commonDetailsActions,
  updateCommonDetailsActions,
} from "actions/calibrationActions";

//fetch common details - name, email, roomTemperature
export const commonDetailsInitiated = (token) => ({
  type: commonDetailsActions.commonDetailsInitiated,
  payload: {
    token,
  },
});

export const commonDetailsFailed = ({ error }) => ({
  type: commonDetailsActions.commonDetailsFailure,
  payload: {
    error,
  },
});

//update common details configurations
export const updateCommonDetailsInitiated = ({ token, data }) => ({
  type: updateCommonDetailsActions.updateCommonDetaislInitiated,
  payload: {
    token,
    data,
  },
});

export const updateCommonDetailsFailed = ({ error }) => ({
  type: updateCommonDetailsActions.updateCommonDetaislFailure,
  payload: {
    error,
  },
});

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

export const abortPid = (token, deckName) => ({
  type: pidActions.pidAbortActionInitiated,
  payload: {
    token,
    deckName,
  },
});

export const abortPidFailed = (errorResponse) => ({
  type: pidActions.pidAbortActionFailure,
  payload: {
    ...errorResponse,
    error: true,
  },
});

// action creators for motor
export const motorInitiated = (token, body) => ({
  type: motorActions.motorActionInitiated,
  payload: { token, body },
});

export const motorFailed = ({ error }) => ({
  type: motorActions.motorActionFailure,
  payload: { error },
});
