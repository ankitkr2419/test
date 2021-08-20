import {
  calibrationActions,
  updateCalibrationActions,
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
