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

export const createCartridgesActions = {
  initiateAction: "CREATE_CARTRIDGES_INITIATED",
  successAction: "CREATE_CARTRIDGES_SUCCESS",
  failureAction: "CREATE_CARTRIDGES_FAILURE",
  resetAction: "CREATE_CARTRIDGES_RESET",
};

export const deleteCartridgesActions = {
  initiateAction: "DELETE_CARTRIDGES_INITIATED",
  successAction: "DELETE_CARTRIDGES_SUCCESS",
  failureAction: "DELETE_CARTRIDGES_FAILURE",
  resetAction: "DELETE_CARTRIDGES_RESET",
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

// websocket for shaker
export const shakerRunProgressActions = {
  shakerRunProgressAction: "SHAKER_RUN_IN_PROGRESS",
  shakerRunProgressActionSuccess: "SHAKER_RUN_SUCCEEDED",
};

// websocket for heater
export const heaterRunProgressActions = {
  heaterRunProgressAction: "HEATER_RUN_IN_PROGRESS",
  heaterRunProgressActionSuccess: "HEATER_RUN_SUCCEEDED",
};

// tolerance values
export const fetchToleranceActions = {
  initiateAction: "FETCH_TOLERANCE_INITIATED",
  successAction: "FETCH_TOLERANCE_SUCCESS",
  failureAction: "FETCH_TOLERANCE_FAILURE",
  resetAction: "FETCH_TOLERANCE_RESET",
};

export const updateToleranceActions = {
  initiateAction: "UPDATE_TOLERANCE_INITIATED",
  successAction: "UPDATE_TOLERANCE_SUCCESS",
  failureAction: "UPDATE_TOLERANCE_FAILURE",
  resetAction: "UPDATE_TOLERANCE_RESET",
};

// consumable distances values
export const fetchConsumableActions = {
  initiateAction: "FETCH_CONSUMABLE_INITIATED",
  successAction: "FETCH_CONSUMABLE_SUCCESS",
  failureAction: "FETCH_CONSUMABLE_FAILURE",
  resetAction: "FETCH_CONSUMABLE_RESET",
};

export const updateConsumableActions = {
  initiateAction: "UPDATE_CONSUMABLE_INITIATED",
  successAction: "UPDATE_CONSUMABLE_SUCCESS",
  failureAction: "UPDATE_CONSUMABLE_FAILURE",
  resetAction: "UPDATE_CONSUMABLE_RESET",
};

export const addConsumableActions = {
  initiateAction: "UPDATE_CONSUMABLE_INITIATED",
  successAction: "UPDATE_CONSUMABLE_SUCCESS",
  failureAction: "UPDATE_CONSUMABLE_FAILURE",
  resetAction: "UPDATE_CONSUMABLE_RESET",
};
