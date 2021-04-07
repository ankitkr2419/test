import { discardDeckActions, discardTipActions } from "actions/discardDeckActions";

export const initialState = {
  isLoading: false,
  serverErrors: {},
  discardTipInProgress: {},
  discardTipInCompleted: {}
};

export const discardDeckReducer = (state = initialState, action = {}) => {
  switch (action.type) {
    case discardDeckActions.discardDeckInitiated:
      return { ...state, ...action.payload, isLoading: true };
    case discardDeckActions.discardDeckSuccess:
      return { ...state, ...action.payload, isLoading: false };
    case discardDeckActions.discardDeckFailed:
      return { ...state, ...action.payload, isLoading: false };
    case discardTipActions.discardTipInitiated:
      return { ...state, ...action.payload, isLoading: true };
    case discardTipActions.discardTipSuccess:
      return { ...state, ...action.payload, isLoading: false };
    case discardTipActions.discardTipFailed:
      return { ...state, ...action.payload, isLoading: false };
    case discardTipActions.discardTipInProgress:
      return {
        ...state, ...action.payload, isLoading: false
      };
    case discardTipActions.discardTipInCompleted:
      return {
        ...state, ...action.payload, isLoading: false
      }
    default:
      return state;
  }
}
