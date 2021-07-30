import { fromJS } from "immutable";
import {
  calibrationActions,
  pidProgressActions,
  pidActions,
  updateCalibrationActions,
  motorActions,
} from "actions/calibrationActions";
import { PID_STATUS } from "appConstants";
import loginActions from "actions/loginActions";

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

const pidProgressInitialState = fromJS({
  isLoading: false,
  error: null,
  configs: {},
});

export const pidProgessReducer = (state = pidProgressInitialState, action) => {
  switch (action.type) {
    case pidProgressActions.pidProgressAction:
      const { progressDetails } = action.payload;

      return state.merge({
        progressStatus: PID_STATUS.progressing,
        deckName: progressDetails.deck,
        progress: progressDetails.progress,
        remainingTime: progressDetails.operation_details.remaining_time,
        totalTime: progressDetails.operation_details.total_time,
      });

    case pidProgressActions.pidProgressActionSuccess:
      const { progressSucceeded } = action.payload;

      return state.merge({
        progressStatus: PID_STATUS.progressComplete,
        deckName: progressSucceeded.deck,
        progress: progressSucceeded.progress,
        remainingTime: progressSucceeded.operation_details.remaining_time,
        totalTime: progressSucceeded.operation_details.total_time,
      });

    case loginActions.loginReset:
      return pidProgressInitialState;

    default:
      return state;
  }
};

// reducer to initiate websocket
const pidInitialState = fromJS({
  isLoading: false,
  error: null,
  configs: {},
});

export const pidReducer = (state = pidInitialState, action) => {
  switch (action.type) {
    case pidActions.pidActionInitiated:
      return pidInitialState;

    case pidActions.pidActionSuccess:
      return state.merge({
        isLoading: false,
        pidStatus: PID_STATUS.running,
      });

    case pidActions.pidActionFailure:
      return state.merge({
        isLoading: false,
        pidStatus: PID_STATUS.runFailed,
      });

    case loginActions.loginReset:
      return pidProgressInitialState;

    default:
      return state;
  }
};

const motorInitialState = fromJS({
  isLoading: false,
  error: null,
  data: {},
});

export const motorReducer = (state = motorInitialState, action) => {
  switch (action.type) {
    case motorActions.motorActionInitiated:
      return state.merge({
        isLoading: true,
      });

    case motorActions.motorActionSuccess:
      return state.merge({
        isLoading: false,
        error: false,
        data: action.payload,
      });

    case motorActions.motorActionFailure:
      return state.merge({
        isLoading: false,
        error: true,
      });

    case loginActions.loginReset:
      return motorInitialState;

    default:
      return state;
  }
};
