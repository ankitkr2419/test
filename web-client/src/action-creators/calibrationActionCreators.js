import {
  heaterProgressActions,
  calibrationActions,
  updateCalibrationActions,
  pidProgressActions,
  pidActions,
  motorActions,
  commonDetailsActions,
  updateCommonDetailsActions,
  updatePidDetailsActions,
  fetchPidDetailsActions,
  shakerActions,
  heaterActions,
  abortActions,
  updateMotorDetailsActions,
  createTipsTubesActions,
} from "actions/calibrationActions";

//fetch common details - name, email, roomTemperature
export const shakerInitiated = (payload) => ({
  type: shakerActions.shakerActionInitiated,
  payload,
});

export const shakerFailed = ({ error }) => ({
  type: shakerActions.shakerActionFailed,
  payload: {
    error,
  },
});

//fetch common details - name, email, roomTemperature
export const heaterInitiated = (payload) => ({
  type: heaterActions.heaterActionInitiated,
  payload: payload,
});

export const heaterFailed = ({ error }) => ({
  type: heaterActions.heaterActionFailed,
  payload: {
    error,
  },
});

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

//websocket heater action creators
export const heaterProgress = (heaterData) => ({
  type: heaterProgressActions.heaterProgressAction,
  payload: { heaterData },
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

export const runPidReset = (errorResponse) => ({
  type: pidActions.pidActionReset,
  payload: {},
});

// common abort for PID, heater, shaker
export const abort = (token, deckName) => ({
  type: abortActions.abortActionInitiated,
  payload: {
    token,
    deckName,
  },
});

export const abortFailed = (errorResponse) => ({
  type: abortActions.abortActionFailed,
  payload: {
    ...errorResponse,
    error: true,
  },
});

export const abortReset = () => ({
  type: abortActions.abortActionReset,
  payload: {},
});

// action creators for pid details fetch
export const fetchPidInitiated = (token) => ({
  type: fetchPidDetailsActions.fetchPidActionInitiated,
  payload: { token },
});

export const fetchPidFailed = ({ error }) => ({
  type: fetchPidDetailsActions.fetchPidActionFailed,
  payload: { error },
});

// action creators for pid details update
export const updatePidInitiated = (token, body) => ({
  type: updatePidDetailsActions.updatePidActionInitiated,
  payload: { token, body },
});

export const updatePidFailed = ({ error }) => ({
  type: updatePidDetailsActions.updatePidActionFailed,
  payload: { error },
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

// action creators for update motor details
export const updateMotorDetailsInitiated = (payload) => ({
  type: updateMotorDetailsActions.updateMotorDetaislInitiated,
  payload,
});

export const updateMotorDetailsFailed = ({ error }) => ({
  type: updateMotorDetailsActions.updateMotorDetaislFailure,
  payload: {
    error,
  },
});

// action creators for tips and tubes
export const createTipsOrTubesInitiated = (token, body) => ({
  type: createTipsTubesActions.initiateAction,
  payload: {
    token,
    body,
  },
});

export const createTipsOrTubesFailed = ({ error }) => ({
  type: createTipsTubesActions.failureAction,
  payload: { error },
});

export const resetCreatingTipsOrTubes = () => ({
  type: createTipsTubesActions.resetAction,
});
