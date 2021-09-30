import { fromJS } from "immutable";
import {
  calibrationActions,
  pidProgressActions,
  pidActions,
  motorActions,
  commonDetailsActions,
  updateCommonDetailsActions,
  heaterProgressActions,
  updatePidDetailsActions,
  fetchPidDetailsActions,
  abortActions,
  createTipsTubesActions,
  fetchRtpcrConfigsActions,
  updateRtpcrConfigsActions,
  fetchTECConfigsActions,
  updateTECConfigsActions,
  startLidPidActions,
  lidPidProgressActions,
  abortLidPidActions,
  runDyeCalibrationActions,
  shakerRunProgressActions,
  heaterRunProgressActions,
  fetchToleranceActions,
  updateToleranceActions,
  fetchConsumableActions,
  updateConsumableActions,
  addConsumableActions,
  fetchCalibrationsDeckBActions,
  fetchCalibrationsDeckAActions,
} from "actions/calibrationActions";
import {
  DECKNAME,
  PID_STATUS,
  HEATER_STATUS,
  SHAKER_RUN_STATUS,
  HEATER_RUN_STATUS,
} from "appConstants";
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

const calibrationDeckAInitialState = fromJS({
  isLoading: false,
  error: null,
  data: [],
});

export const calibrationDeckAReducer = (
  state = calibrationDeckAInitialState,
  action
) => {
  switch (action.type) {
    case fetchCalibrationsDeckAActions.initiateAction:
      return state.merge({
        isLoading: true,
        error: null,
      });
    case fetchCalibrationsDeckAActions.successAction:
      const res = action.payload.response;
      return state.merge({
        isLoading: false,
        error: false,
        data: res,
      });

    case fetchCalibrationsDeckAActions.failureAction:
      return state.merge({
        isLoading: false,
        error: true,
      });
    case fetchCalibrationsDeckAActions.resetAction:
      return calibrationDeckAInitialState;

    default:
      return state;
  }
};

const calibrationDeckBInitialState = fromJS({
  isLoading: false,
  error: null,
  data: [],
});

export const calibrationDeckBReducer = (
  state = calibrationDeckBInitialState,
  action
) => {
  switch (action.type) {
    case fetchCalibrationsDeckBActions.initiateAction:
      return state.merge({
        isLoading: true,
        error: null,
      });
    case fetchCalibrationsDeckBActions.successAction:
      const res = action.payload.response;
      return state.merge({
        isLoading: false,
        error: false,
        data: res,
      });

    case fetchCalibrationsDeckBActions.failureAction:
      return state.merge({
        isLoading: false,
        error: true,
      });
    case fetchCalibrationsDeckBActions.resetAction:
      return calibrationDeckBInitialState;

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

    case pidActions.pidActionProgressing:
      return state.merge({
        isLoading: false,
        pidStatus: PID_STATUS.running,
      });

    case pidActions.pidActionProgressSuccess:
      return state.merge({
        isLoading: false,
        pidStatus: PID_STATUS.progressComplete,
      });

    case pidActions.pidActionProgressAbort:
      return state.merge({
        isLoading: false,
        pidStatus: PID_STATUS.aborted,
      });

    case pidActions.pidActionFailure:
      return state.merge({
        isLoading: false,
        pidStatus: PID_STATUS.runFailed,
      });

    case abortActions.abortActionReset:
      return pidInitialState;

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
        abortStatus: "aborting",
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

const rtpcrConfigsInitialState = fromJS({
  isLoading: false,
  error: null,
  isUpdateApi: null, // to distinguish between fetch and put API
  details: {},
});

export const rtpcrConfigsReducer = (
  state = rtpcrConfigsInitialState,
  action
) => {
  switch (action.type) {
    case fetchRtpcrConfigsActions.initiateAction:
      return state.merge({
        isLoading: true,
        error: null,
        details: null,
        isUpdateApi: false,
      });
    case fetchRtpcrConfigsActions.successAction:
      return state.merge({
        isLoading: false,
        error: false,
        details: action.payload.response,
        isUpdateApi: false,
      });
    case fetchRtpcrConfigsActions.failureAction:
      return state.merge({
        isLoading: false,
        error: true,
        isUpdateApi: false,
      });
    case fetchRtpcrConfigsActions.resetAction:
      return rtpcrConfigsInitialState;

    case updateRtpcrConfigsActions.initiateAction:
      return state.merge({
        isLoading: true,
        error: null,
        isUpdateApi: true,
      });
    case updateRtpcrConfigsActions.successAction:
      return state.merge({
        isLoading: false,
        error: false,
        isUpdateApi: true,
      });

    case updateRtpcrConfigsActions.failureAction:
      return state.merge({
        isLoading: false,
        error: true,
        isUpdateApi: true,
      });

    case updateRtpcrConfigsActions.resetAction:
      return rtpcrConfigsInitialState;

    case loginActions.loginReset:
      return rtpcrConfigsInitialState;

    default:
      return state;
  }
};

const tecConfigsInitialState = fromJS({
  isLoading: false,
  error: null,
  isUpdateApi: null, // to distinguish between fetch and put API
  details: {},
});
export const tecConfigsReducer = (state = tecConfigsInitialState, action) => {
  switch (action.type) {
    case fetchTECConfigsActions.initiateAction:
      return state.merge({
        isLoading: true,
        error: null,
        details: null,
        isUpdateApi: false,
      });
    case fetchTECConfigsActions.successAction:
      return state.merge({
        isLoading: false,
        error: false,
        details: action.payload.response,
        isUpdateApi: false,
      });
    case fetchTECConfigsActions.failureAction:
      return state.merge({
        isLoading: false,
        error: true,
        isUpdateApi: false,
      });
    case fetchTECConfigsActions.resetAction:
      return tecConfigsInitialState;

    case updateTECConfigsActions.initiateAction:
      return state.merge({
        isLoading: true,
        error: null,
        isUpdateApi: true,
      });
    case updateTECConfigsActions.successAction:
      return state.merge({
        isLoading: false,
        error: false,
        isUpdateApi: true,
      });

    case updateTECConfigsActions.failureAction:
      return state.merge({
        isLoading: false,
        error: true,
        isUpdateApi: true,
      });

    case updateTECConfigsActions.resetAction:
      return tecConfigsInitialState;

    case loginActions.loginReset:
      return tecConfigsInitialState;

    default:
      return state;
  }
};

//lid pid tuning reducer
const lidPidInitialState = fromJS({
  isLoading: false,
  error: null,
  lidPidStatus: null,
});

export const lidPidReducer = (state = lidPidInitialState, action) => {
  switch (action.type) {
    case startLidPidActions.initiateAction:
      return state.merge({
        isLoading: true,
        error: null,
        lidPidStatus: null,
      });
    case startLidPidActions.successAction:
      return state.merge({
        isLoading: false,
        error: false,
        lidPidStatus: PID_STATUS.running,
      });
    case startLidPidActions.failureAction:
      return state.merge({
        isLoading: false,
        error: true,
        lidPidStatus: PID_STATUS.runFailed,
      });
    case startLidPidActions.resetAction:
      return lidPidInitialState;

    case abortLidPidActions.initiateAction:
      return state.merge({
        isLoading: true,
        error: null,
        lidPidStatus: PID_STATUS.aborting,
      });
    case abortLidPidActions.successAction:
      return state.merge({
        isLoading: false,
        error: false,
        lidPidStatus: PID_STATUS.stopped,
      });
    case abortLidPidActions.failureAction:
      return state.merge({
        isLoading: false,
        error: true,
        lidPidStatus: PID_STATUS.abortFailed,
      });
    case abortLidPidActions.resetAction:
      return lidPidInitialState;

    case lidPidProgressActions.lidPidProgressAction:
      return state.merge({
        isLoading: false,
        error: null,
        lidPidStatus: PID_STATUS.progressing,
      });
    case lidPidProgressActions.lidPidProgressActionSuccess:
      return state.merge({
        isLoading: false,
        error: false,
        lidPidStatus: PID_STATUS.progressComplete,
      });

    case loginActions.loginReset:
      return lidPidInitialState;

    default:
      return state;
  }
};

const shakerRunProgressInitialState = fromJS({
  shakerRunStatus: null,
});

// websocket ShakerRun
export const shakerRunProgessReducer = (
  state = shakerRunProgressInitialState,
  action
) => {
  switch (action.type) {
    case shakerRunProgressActions.shakerRunProgressAction:
      return state.merge({
        shakerRunStatus: SHAKER_RUN_STATUS.progressing,
      });

    case shakerRunProgressActions.shakerRunProgressActionSuccess:
      return state.merge({
        shakerRunStatus: SHAKER_RUN_STATUS.progressComplete,
      });

    case shakerRunProgressActions.shakerRunProgressActionAborted:
      return state.merge({
        shakerRunStatus: SHAKER_RUN_STATUS.progressAborted,
      });

    case abortActions.abortActionReset:
      return shakerRunProgressInitialState;

    case loginActions.loginReset:
      return shakerRunProgressInitialState;

    default:
      return state;
  }
};

const heaterRunProgressInitialState = fromJS({
  heaterRunStatus: null,
});

// websocket HeaterRun
export const heaterRunProgessReducer = (
  state = heaterRunProgressInitialState,
  action
) => {
  switch (action.type) {
    case heaterRunProgressActions.heaterRunProgressAction:
      return state.merge({
        heaterRunStatus: HEATER_RUN_STATUS.progressing,
      });

    case heaterRunProgressActions.heaterRunProgressActionSuccess:
      return state.merge({
        heaterRunStatus: HEATER_RUN_STATUS.progressComplete,
      });

    case heaterRunProgressActions.heaterRunProgressActionAborted:
      return state.merge({
        heaterRunStatus: HEATER_RUN_STATUS.progressAborted,
      });

    case abortActions.abortActionReset:
      return heaterRunProgressInitialState;

    case loginActions.loginReset:
      return heaterRunProgressInitialState;

    default:
      return state;
  }
};

const toleranceInitialState = fromJS({
  isLoading: false,
  error: null,
  data: [],
  isUpdateApi: null, // to distinguish between fetch and put API
});
export const toleranceReducer = (state = toleranceInitialState, action) => {
  switch (action.type) {
    case fetchToleranceActions.initiateAction:
      return state.merge({
        isLoading: true,
        error: null,
        isUpdateApi: false,
      });
    case fetchToleranceActions.successAction:
      return state.merge({
        isLoading: false,
        error: false,
        data: action.payload.response,
        isUpdateApi: false,
      });
    case fetchToleranceActions.failureAction:
      return state.merge({
        isLoading: false,
        error: true,
        isUpdateApi: false,
      });
    case fetchToleranceActions.resetAction:
      return toleranceInitialState;

    case updateToleranceActions.initiateAction:
      return state.merge({
        isLoading: true,
        error: null,
        isUpdateApi: true,
      });
    case updateToleranceActions.successAction:
      return state.merge({
        isLoading: false,
        error: false,
        isUpdateApi: true,
      });

    case updateToleranceActions.failureAction:
      return state.merge({
        isLoading: false,
        error: true,
        isUpdateApi: true,
      });

    case updateToleranceActions.resetAction:
      return toleranceInitialState;

    case loginActions.loginReset:
      return toleranceInitialState;

    default:
      return state;
  }
};

//run dye calibration reducer
const dyeCalibrationInitialState = fromJS({
  isLoading: false,
  error: null,
  dyeCalibrationStatus: null,
});

export const dyeCalibrationReducer = (
  state = dyeCalibrationInitialState,
  action
) => {
  switch (action.type) {
    case runDyeCalibrationActions.initiateAction:
      return state.merge({
        isLoading: true,
        error: null,
        dyeCalibrationStatus: PID_STATUS.running,
      });
    case runDyeCalibrationActions.successAction:
      return state.merge({
        isLoading: false,
        error: false,
        dyeCalibrationStatus: PID_STATUS.running,
      });
    case runDyeCalibrationActions.failureAction:
      return state.merge({
        isLoading: false,
        error: true,
        dyeCalibrationStatus: PID_STATUS.runFailed,
      });
    case runDyeCalibrationActions.resetAction:
      return dyeCalibrationInitialState;
    case runDyeCalibrationActions.progressAction:
      return state.merge({
        isLoading: false,
        error: false,
        dyeCalibrationStatus: PID_STATUS.progressing,
      });
    case runDyeCalibrationActions.completedAction:
      return state.merge({
        isLoading: false,
        error: false,
        dyeCalibrationStatus: PID_STATUS.progressComplete,
      });

    case loginActions.loginReset:
      return dyeCalibrationInitialState;

    default:
      return state;
  }
};

const consumableInitialState = fromJS({
  isLoading: false,
  error: null,
  data: [],
  isUpdateApi: null, // to distinguish between fetch and put API
});

export const consumableReducer = (state = consumableInitialState, action) => {
  switch (action.type) {
    case fetchConsumableActions.initiateAction:
      return state.merge({
        isLoading: true,
        error: null,
        isUpdateApi: false,
      });
    case fetchConsumableActions.successAction:
      return state.merge({
        isLoading: false,
        error: false,
        data: action.payload.response,
        isUpdateApi: false,
      });
    case fetchConsumableActions.failureAction:
      return state.merge({
        isLoading: false,
        error: true,
        isUpdateApi: false,
      });
    case fetchConsumableActions.resetAction:
      return consumableInitialState;

    case updateConsumableActions.initiateAction:
      return state.merge({
        isLoading: true,
        error: null,
        isUpdateApi: true,
      });
    case updateConsumableActions.successAction:
      return state.merge({
        isLoading: false,
        error: false,
        isUpdateApi: true,
      });

    case updateConsumableActions.failureAction:
      return state.merge({
        isLoading: false,
        error: true,
        isUpdateApi: true,
      });

    case updateConsumableActions.resetAction:
      return consumableInitialState;

    case addConsumableActions.initiateAction:
      return state.merge({
        isLoading: true,
        error: null,
        isUpdateApi: true,
      });
    case addConsumableActions.successAction:
      return state.merge({
        isLoading: false,
        error: false,
        isUpdateApi: true,
      });

    case addConsumableActions.failureAction:
      return state.merge({
        isLoading: false,
        error: true,
        isUpdateApi: true,
      });

    case addConsumableActions.resetAction:
      return consumableInitialState;

    case loginActions.loginReset:
      return consumableInitialState;

    default:
      return state;
  }
};
