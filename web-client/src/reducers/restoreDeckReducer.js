import { restoreDeckActions } from "actions/restoreDeckActions";

export const initialState = {
  isLoading: false,
  serverErrors: {}
};

export const restoreDeckReducer = (state = initialState, action = {}) => {
  switch (action.type) {
    case restoreDeckActions.restoreDeckInitiated:
      return { ...state, ...action.payload, isLoading: true };
    case restoreDeckActions.restoreDeckSuccess:
      return { ...state, ...action.payload, isLoading: false };
    case restoreDeckActions.restoreDeckFailed:
      return { ...state, ...action.payload, isLoading: false };
    default:
      return state
  }
}
