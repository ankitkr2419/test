import {
  homingActions,
  discardTipAndHomingActions,
  homingInProgressActionsForDeckB,
  homingInProgressActionsForDeckA,
} from "actions/homingActions";

import { HOMING_STATUS } from "appConstants";

export const initialState = {
  isLoading: false,
  showHomingModal: true,
  homingData: {},
  serverErrors: {},
  homingAllDeckCompletionPercentage: 0,
  isHomingActionCompleted: false,
  homingActionInProgress: {},
  homingActionInCompleted: {},
  homingStatus: null,
};

export const homingReducer = (state = initialState, action = {}) => {
  switch (action.type) {
    case homingActions.homingActionInitiated:
      return {
        ...state,
        ...action.payload,
        isLoading: true,
        isHomingActionCompleted: false,
        homingStatus: HOMING_STATUS.progressStarted,
      };

    // case homingActions.homingActionSuccess:
    //   return { ...state, ...action.payload, isLoading: false };

    // case homingActions.homingActionFailed:
    //   return { ...state, ...action.payload, isLoading: false };

    case homingActions.homingActionInProgress:
      return {
        ...state,
        ...action.payload,
        isLoading: false,
        homingAllDeckCompletionPercentage: action.payload.progress,
        isHomingActionCompleted: false,
        homingStatus: HOMING_STATUS.progressing,
      };

    case homingActions.homingActionInCompleted:
      return {
        ...state,
        ...action.payload,
        isLoading: false,
        isHomingActionCompleted: true,
        homingStatus: HOMING_STATUS.progressComplete,
      };

    case homingActions.hideHomingModaal:
      return {
        ...initialState,
        showHomingModal: false,
      };

    case homingActions.showHomingModaal:
      return {
        ...state,
        showHomingModal: true,
      };

    default:
      return state;
  }
};

const discardTipHomingInitialState = {
  isLoading: false,
  homingData: {},
  serverErrors: {},
  discardTipAndHomingError: null,
};

export const discardTipAndHomingReducer = (
  state = discardTipHomingInitialState,
  action
) => {
  switch (action.type) {
    case discardTipAndHomingActions.discardTipAndHomingActionInitiated:
      return {
        ...state,
        ...action.payload,
        isLoading: true,
      };
    case discardTipAndHomingActions.discardTipAndHomingActionSuccess:
      return {
        ...state,
        ...action.payload,
        isLoading: false,
        discardTipAndHomingError: false,
      };
    case discardTipAndHomingActions.discardTipAndHomingActionFailed:
      return {
        ...state,
        ...action.payload,
        isLoading: false,
        discardTipAndHomingError: true,
      };
    case discardTipAndHomingActions.discardTipAndHomingActionReset:
      return { ...state, discardTipAndHomingError: null };

    default:
      return state;
  }
};

const homingInitialStateForDeckA = {
  isLoading: false,
  homingData: {},
  serverErrors: {},
  isHomingCompletedForDeckA: false,
  homingActionInProgress: {},
  homingActionInCompleted: {},
  homingStatusForDeckA: null,
};

export const homingReducerForDeckA = (
  state = homingInitialStateForDeckA,
  action
) => {
  switch (action.type) {
    case homingInProgressActionsForDeckA.homingActionInProgressForDeckA:
      return {
        ...state,
        ...action.payload,
        isHomingCompletedForDeckA: false,
        homingStatusForDeckA: HOMING_STATUS.progressing,
      };

    case homingInProgressActionsForDeckA.homingActionInSuccessForDeckA:
      return {
        ...state,
        ...action.payload,
        isHomingCompletedForDeckA: true,
        homingStatusForDeckA: HOMING_STATUS.progressComplete,
      };

    //reset state
    case homingInProgressActionsForDeckA.resetHomingStateForDeckA:
      return {
        ...homingInitialStateForDeckA,
      };

    // resets homing state
    case homingActions.hideHomingModaal:
      return {
        ...homingInitialStateForDeckA,
      };

    default:
      return state;
  }
};

const homingInitialStateForDeckB = {
  isLoading: false,
  homingData: {},
  serverErrors: {},
  isHomingCompletedForDeckB: false,
  homingActionInProgress: {},
  homingActionInCompleted: {},
  homingStatusForDeckB: null,
};

export const homingReducerForDeckB = (
  state = homingInitialStateForDeckB,
  action
) => {
  switch (action.type) {
    case homingInProgressActionsForDeckB.homingActionInProgressForDeckB:
      return {
        ...state,
        ...action.payload,
        isHomingCompletedForDeckB: false,
        homingStatusForDeckB: HOMING_STATUS.progressing,
      };

    case homingInProgressActionsForDeckB.homingActionInSuccessForDeckB:
      return {
        ...state,
        ...action.payload,
        isHomingCompletedForDeckB: true,
        homingStatusForDeckB: HOMING_STATUS.progressComplete,
      };

    //reset state
    case homingInProgressActionsForDeckB.resetHomingStateForDeckB:
      return {
        ...homingInitialStateForDeckB,
      };

    // resets homing state
    case homingActions.hideHomingModaal:
      return {
        ...homingInitialStateForDeckB,
      };

    default:
      return state;
  }
};
