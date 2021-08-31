import { fromJS } from "immutable";
import {
  calibrationActions,
  pidProgressActions,
  pidActions,
  updateCalibrationActions,
  motorActions,
  commonDetailsActions,
  updateCommonDetailsActions,
  heaterProgressActions,
  updatePidDetailsActions,
  fetchPidDetailsActions,
  abortActions,
  createTipsTubesActions,
} from "actions/calibrationActions";
import { DECKNAME, PID_STATUS, HEATER_STATUS } from "appConstants";
import loginActions from "actions/loginActions";

const commonDetailsInitialState = fromJS({
  isLoading: false,
  error: null,
  isUpdateApi: null, // to distinguish between fetch and put API
  details: {},
});

export const commonDetailsReducer = (
  state = commonDetailsInitialState,
  action
) => {
  switch (action.type) {
    case commonDetailsActions.commonDetailsInitiated:
      return state.merge({
        isLoading: true,
        error: null,
        isUpdateApi: false,
      });
    case commonDetailsActions.commonDetailsSuccess:
      const res = action.payload.response;
      return state.merge({
        isLoading: false,
        error: false,
        isUpdateApi: false,
        details: res,
      });
    case commonDetailsActions.commonDetailsFailure:
      return state.merge({
        isLoading: false,
        isUpdateApi: false,
        error: true,
      });

    case updateCommonDetailsActions.updateCommonDetaislInitiated:
      return state.merge({
        isLoading: true,
        error: null,
        isUpdateApi: true,
      });
    case updateCommonDetailsActions.updateCommonDetaislSuccess:
      return state.merge({
        isLoading: false,
        error: false,
        isUpdateApi: true,
      });
    case updateCommonDetailsActions.updateCommonDetaislFailure:
      return state.merge({
        isLoading: false,
        error: true,
        isUpdateApi: true,
      });
    case commonDetailsActions.commonDetailsReset:
      return commonDetailsInitialState;

    case updateCommonDetailsActions.updateCommonDetaislReset:
      return commonDetailsInitialState;

    default:
      return state;
  }
};

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
    default:
      return state;
  }
};

const updateCalibrationInitialState = fromJS({
  isLoading: false,
  error: null,
});

export const updateCalibrationReducer = (
  state = updateCalibrationInitialState,
  action
) => {
  switch (action.type) {
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

const heaterProgressInitialState = fromJS({
  heaterInProgress: null,
  data: {},
});

export const heaterProgressReducer = (
  state = heaterProgressInitialState,
  action
) => {
  switch (action.type) {
    case heaterProgressActions.heaterProgressAction:
      return state.merge({
        heaterInProgress: HEATER_STATUS.progressing,
        data: action.payload.heaterData,
      });

    default:
      return state;
  }
};

const initialStateOfDecks = [
  {
    name: DECKNAME.DeckAShort,
    deckName: DECKNAME.DeckA,
  },
  {
    name: DECKNAME.DeckBShort,
    deckName: DECKNAME.DeckB,
  },
];

const pidProgressInitialState = fromJS({
  isLoading: false,
  error: null,
  configs: {},
  decks: initialStateOfDecks,
});

export const pidProgessReducer = (state = pidProgressInitialState, action) => {
  switch (action.type) {
    case pidProgressActions.pidProgressAction:
      const { progressDetails } = action.payload;

      //store payload details into appropriate deck object
      const updatedDeckStateInProgress = state.toJS().decks.map((deckObj) => {
        return deckObj.name === progressDetails.deck
          ? {
              ...deckObj,
              isActive: true,
              progressStatus: PID_STATUS.progressing,
              progress: progressDetails.progress,
              remainingTime: progressDetails.operation_details.remaining_time,
              totalTime: progressDetails.operation_details.total_time,
            }
          : {
              ...deckObj,
              isActive: false,
            };
      });
      return state.merge({
        decks: updatedDeckStateInProgress,
      });

    case pidProgressActions.pidProgressActionSuccess:
      const { progressSucceeded } = action.payload;

      //store payload details into appropriate deck object
      const updatedDeckStateSuccess = state.toJS().decks.map((deckObj) => {
        return deckObj.name === progressSucceeded.deck
          ? {
              ...deckObj,
              isActive: true,
              progressStatus: PID_STATUS.progressComplete,
              progress: progressSucceeded.progress,
              remainingTime: progressSucceeded.operation_details.remaining_time,
              totalTime: progressSucceeded.operation_details.total_time,
            }
          : {
              ...deckObj,
              isActive: false,
            };
      });
      return state.merge({
        decks: updatedDeckStateSuccess,
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
  pidStatus: null,
  configs: {},
  pidData: {},
  isPidUpdateApi: null,
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

    case pidActions.pidActionReset:
      return pidInitialState;

    case fetchPidDetailsActions.fetchPidActionInitiated:
      return state.merge({
        isLoading: true,
        error: false,
        isPidUpdateApi: false,
      });

    case fetchPidDetailsActions.fetchPidActionSuccess:
      const res = action.payload.response;
      return state.merge({
        isLoading: false,
        error: false,
        isPidUpdateApi: false,
        pidData: res,
      });

    case fetchPidDetailsActions.fetchPidActionFailed:
      return state.merge({
        isLoading: false,
        error: true,
        isPidUpdateApi: false,
      });

    case updatePidDetailsActions.updatePidActionInitiated:
      return state.merge({
        isLoading: true,
        error: false,
        isPidUpdateApi: true,
      });

    case updatePidDetailsActions.updatePidActionSuccess:
      return state.merge({
        isLoading: false,
        error: false,
        isPidUpdateApi: true,
      });

    case updatePidDetailsActions.updatePidActionFailed:
      return state.merge({
        isLoading: false,
        error: true,
        isPidUpdateApi: true,
      });

    case loginActions.loginReset:
      return pidProgressInitialState;

    default:
      return state;
  }
};

// abort reducer
const abortInitialState = fromJS({
  isLoading: false,
  error: null,
  abortStatus: null,
});

// common reducer: used to abort process for PID, heater and shaker
export const abortReducer = (state = abortInitialState, action) => {
  switch (action.type) {
    case abortActions.abortActionInitiated:
      return state.merge({
        isLoading: true,
      });

    case abortActions.abortActionSuccess:
      return state.merge({
        isLoading: false,
        error: false,
        abortStatus: "aborted",
      });

    case abortActions.abortActionFailed:
      return state.merge({
        isLoading: false,
        error: true,
        abortStatus: "abortFailed",
      });

    case abortActions.abortActionReset:
      return abortInitialState;

    case loginActions.loginReset:
      return abortInitialState;

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

const tipTubeInitialState = fromJS({
  isLoading: false,
  error: null,
});

export const createTipTubeReducer = (state = tipTubeInitialState, action) => {
  switch (action.type) {
    case createTipsTubesActions.initiateAction:
      return state.merge({
        isLoading: true,
        error: null,
      });
    case createTipsTubesActions.successAction:
      return state.merge({
        isLoading: false,
        error: false,
      });
    case createTipsTubesActions.failureAction:
      return state.merge({
        isLoading: false,
        error: true,
      });
    case createTipsTubesActions.resetAction:
      return state.merge({
        isLoading: false,
        error: null,
      });

    case loginActions.loginReset:
      return tipTubeInitialState;

    default:
      return state;
  }
};
