import { homingActions } from "actions/homingActions";

export const initialState = {
  isLoading: false,
  homingData: {},
  serverErrors: {},
  isHomingActionCompleted: false,
  homingActionInProgress: {},
  homingActionInCompleted: {}
};

export const homingReducer = (state = initialState, action = {}) => {
  switch (action.type) {
    case homingActions.homingActionInitiated:
      return {
        ...state,
        ...action.payload,
        isLoading: true,
        isHomingActionCompleted: true,
      };
    case homingActions.homingActionSuccess:
      return { ...state, ...action.payload, isLoading: false };
    case homingActions.homingActionFailed:
      return { ...state, ...action.payload, isLoading: false };
    case homingActions.homingActionInProgress:
      return {
        ...state,
        ...action.payload,
        isLoading: false,
        isHomingActionCompleted: true,
      };
    case homingActions.homingActionInCompleted:
      return {
        ...state,
        ...action.payload,
        isLoading: false,
        isHomingActionCompleted: false,
      };
    default:
      return state;
  }
};
