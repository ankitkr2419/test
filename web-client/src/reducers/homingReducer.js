import {
  homingActions,
  discardTipAndHomingActions,
} from "actions/homingActions";

export const initialState = {
  isLoading: false,
  homingData: {},
  serverErrors: {},
  homingAllDeckCompletionPercentage: 0,
  isHomingActionCompleted: false,
  homingActionInProgress: {},
  homingActionInCompleted: {},
};

export const homingReducer = (state = initialState, action = {}) => {
  switch (action.type) {
    case homingActions.homingActionInitiated:
      return {
        ...state,
        ...action.payload,
        isLoading: true,
        isHomingActionCompleted: false,
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
      };

    case homingActions.homingActionInCompleted:
      return {
        ...state,
        ...action.payload,
        isLoading: false,
        isHomingActionCompleted: true,
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
