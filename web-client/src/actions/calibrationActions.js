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
