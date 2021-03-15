import { discardDeckActions } from "actions/discardDeckActions";

export const initialState = {
  isLoading: false,
  serverErrors: {}
};

export const discardDeckReducer = (state = initialState, action = {}) => {
  switch (action.type) {
    case discardDeckActions.discardDeckInitiated:
      return { ...state, ...action.payload, isLoading: true };
    case discardDeckActions.discardDeckSuccess:
      return { ...state, ...action.payload, isLoading: false };
    case discardDeckActions.discardDeckFailed:
      return { ...state, ...action.payload, isLoading: false };
    default:
      return state;
  }
}
