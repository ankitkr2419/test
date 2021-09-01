export const heaterProgressActions = {
  heaterProgressAction: "HEATER_IN_PROGRESS",
};

export const pidProgressActions = {
  pidProgressAction: "PID_IN_PROGRESS",
  pidProgressActionSuccess: "PID_SUCCEEDED",
};

export const pidActions = {
  pidActionInitiated: "PID_START_INITIATED",
  pidActionSuccess: "PID_START_SUCCEEDED",
  pidActionFailure: "PID_START_FAILURE",
  pidAbortActionInitiated: "PID_ABORT_INITIATED",
  pidAbortActionSuccess: "PID_ABORT_SUCCESS",
  pidAbortActionFailure: "PID_ABORT_FAILURE",
};

export const fetchPidDetailsActions = {
  fetchPidActionInitiated: "FETCH_PID_START_INITIATED",
  fetchPidActionSuccess: "FETCH_PID_START_SUCEEDED",
  fetchPidActionFailed: "FETCH_PID_START_FAILED",
};

export const updatePidDetailsActions = {
  updatePidActionInitiated: "UPDATE_PID_START_INITIATED",
  updatePidActionSuccess: "UPDATE_PID_START_SUCEEDED",
  updatePidActionFailed: "UPDATE_PID_START_FAILED",
};

export const motorActions = {
  motorActionInitiated: "MOTOR_INITIATED",
  motorActionSuccess: "MOTOR_SUCCEEDED",
  motorActionFailure: "MOTOR_FAILURE",
};

export const calibrationActions = {
  calibrationInitiated: "CALIBRATION_INITIATED",
  calibrationSuccess: "CALIBRATION_SUCCESS",
  calibrationFailure: "CALIBRATION_FAILURE",
  calibrationReset: "CALIBRATION_RESET",
};

export const updateCalibrationActions = {
  updateCalibrationInitiated: "UPDATE_CALIBRATION_INITIATED",
  updateCalibrationSuccess: "UPDATE_CALIBRATION_SUCCESS",
  updateCalibrationFailure: "UPDATE_CALIBRATION_FAILURE",
  updateCalibrationReset: "UPDATE_CALIBRATION_RESET",
};

export const commonDetailsActions = {
  commonDetailsInitiated: "COMMON_DETAILS_INITIATED",
  commonDetailsSuccess: "COMMON_DETAILS_SUCCESS",
  commonDetailsFailure: "COMMON_DETAILS_FAILURE",
  commonDetailsReset: "COMMON_DETAILS_RESET",
};

export const updateCommonDetailsActions = {
  updateCommonDetaislInitiated: "UPDATE_COMMON_INITIATED",
  updateCommonDetaislSuccess: "UPDATE_COMMON_SUCCESS",
  updateCommonDetaislFailure: "UPDATE_COMMON_FAILURE",
  updateCommonDetaislReset: "UPDATE_COMMON_RESET",
};

export const createTipsTubesActions = {
  initiateAction: "CREATE_TIPS_TUBES_INITIATED",
  successAction: "CREATE_TIPS_TUBES_SUCCESS",
  failureAction: "CREATE_TIPS_TUBES_FAILURE",
  resetAction: "CREATE_TIPS_TUBES_RESET",
};

export const fetchRtpcrConfigsActions = {
  initiateAction: "FETCH_RTPCR_CONFIGS_INITIATED",
  successAction: "FETCH_RTPCR_CONFIGS_SUCCESS",
  failureAction: "FETCH_RTPCR_CONFIGS_FAILURE",
  resetAction: "FETCH_RTPCR_CONFIGS_RESET",
};

export const updateRtpcrConfigsActions = {
  initiateAction: "UPDATE_RTPCR_CONFIGS_INITIATED",
  successAction: "UPDATE_RTPCR_CONFIGS_SUCCESS",
  failureAction: "UPDATE_RTPCR_CONFIGS_FAILURE",
  resetAction: "UPDATE_RTPCR_CONFIGS_RESET",
};
