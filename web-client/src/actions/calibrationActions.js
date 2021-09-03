export const abortActions = {
  abortActionInitiated: "ABORT_INITIATED",
  abortActionSuccess: "ABORT_SUCCEEDED",
  abortActionFailed: "ABORT_FAILED",
  abortActionReset: "ABORT_RESET",
};

export const shakerActions = {
  shakerActionInitiated: "SHAKER_INITIATED",
  shakerActionSuccess: "SHAKER_SUCCEEDED",
  shakerActionFailed: "SHAKER_FAILED",
};

export const heaterActions = {
  heaterActionInitiated: "HEATER_INITIATED",
  heaterActionSuccess: "HEATER_SUCCEEDED",
  heaterActionFailed: "HEATER_FAILED",
};

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
  pidActionReset: "PID_START_RESET",
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

export const updateMotorDetailsActions = {
  updateMotorDetaislInitiated: "UPDATE_MOTOR_DETAILS_INITIATED",
  updateMotorDetaislSuccess: "UPDATE_MOTOR_DETAILS_SUCCESS",
  updateMotorDetaislFailure: "UPDATE_MOTOR_DETAILS_FAILURE",
  updateMotorDetaislReset: "UPDATE_MOTOR_DETAILS_RESET",
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

export const fetchTECConfigsActions = {
  initiateAction: "FETCH_TEC_CONFIGS_INITIATED",
  successAction: "FETCH_TEC_CONFIGS_SUCCESS",
  failureAction: "FETCH_TEC_CONFIGS_FAILURE",
  resetAction: "FETCH_TEC_CONFIGS_RESET",
};

export const updateTECConfigsActions = {
  initiateAction: "UPDATE_TEC_CONFIGS_INITIATED",
  successAction: "UPDATE_TEC_CONFIGS_SUCCESS",
  failureAction: "UPDATE_TEC_CONFIGS_FAILURE",
  resetAction: "UPDATE_TEC_CONFIGS_RESET",
};

export const startLidPidActions = {
  initiateAction: "START_LID_PID_INITIATED",
  successAction: "START_LID_PID_SUCCESS",
  failureAction: "START_LID_PID_FAILURE",
  resetAction: "START_LID_PID_RESET",
};

export const lidPidProgressActions = {
  lidPidProgressAction: "LID_PID_IN_PROGRESS",
  lidPidProgressActionSuccess: "LID_PID_SUCCEEDED",
};

export const abortLidPidActions = {
  initiateAction: "ABORT_LID_PID_INITIATED",
  successAction: "ABORT_LID_PID_SUCCESS",
  failureAction: "ABORT_LID_PID_FAILURE",
  resetAction: "ABORT_LID_PID_RESET",
};

export const resetTECActions = {
  initiateAction: "RESET_TEC_INITIATED",
  successAction: "RESET_TEC_SUCCESS",
  failureAction: "RESET_TEC_FAILURE",
  resetAction: "RESET_TEC_RESET",
};

export const autoTuneTECActions = {
  initiateAction: "AUTOTUNE_TEC_INITIATED",
  successAction: "AUTOTUNE_TEC_SUCCESS",
  failureAction: "AUTOTUNE_TEC_FAILURE",
  resetAction: "AUTOTUNE_TEC_RESET",
};
