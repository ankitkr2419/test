import { fromJS } from "immutable";
import {
  calibrationActions,
  updateCalibrationActions,
} from "actions/calibrationActions";

const calibrationInitialState = fromJS({
  isLoading: false,
  error: null,
  configs: {},
});

export const calibrationReducer = (state = calibrationInitialState, action) => {
  switch (action.type) {
    case calibrationActions.calibrationInitiated:
      return state.merge({
        isLoading: true,
        error: null,
        configs: calibrationInitialState.configs,
      });
    case calibrationActions.calibrationSuccess:
      const res = action.payload.response;
      return state.merge({
        isLoading: false,
        error: false,
        configs: res,
      });
    case calibrationActions.calibrationFailure:
      return state.merge({
        isLoading: false,
        error: true,
      });
    case calibrationActions.calibrationReset:
      return state.merge({
        isLoading: false,
        error: null,
        configs: calibrationInitialState.configs,
      });
    case updateCalibrationActions.updateCalibrationInitiated:
      return state.merge({
        isLoading: true,
        error: null,
      });
    case updateCalibrationActions.updateCalibrationSuccess:
      return state.merge({
        isLoading: false,
        error: false,
      });
    case updateCalibrationActions.updateCalibrationFailure:
      return state.merge({
        isLoading: false,
        error: true,
      });
    default:
      return state;
  }
};
