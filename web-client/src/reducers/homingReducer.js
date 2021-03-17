import { homingActions } from "actions/homingActions";

export const initialState = {
  isLoading: false,
  homingData: {},
  serverErrors: {}
};

export const homingReducer = (state = initialState, action = {}) => {
  switch (action.type) {
    case homingActions.homingActionInitiated:
      return { ...state, ...action.payload, isLoading: true };
    case homingActions.homingActionSuccess:
      return { ...state, ...action.payload, isLoading: false };
    case homingActions.homingActionFailed:
      return { ...state, ...action.payload, isLoading: false };
    default:
      return state;
  }
}
