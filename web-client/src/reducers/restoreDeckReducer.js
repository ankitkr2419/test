import { restoreDeckActions } from "actions/restoreDeckActions";

export const initialState = {
  isLoading: false,
  serverErrors: {},
  restoreDeckError: null,
};

export const restoreDeckReducer = (state = initialState, action = {}) => {
  switch (action.type) {
    case restoreDeckActions.restoreDeckInitiated:
      return { ...state, ...action.payload, isLoading: true };
    case restoreDeckActions.restoreDeckSuccess:
      return {
        ...state,
        ...action.payload,
        isLoading: false,
        restoreDeckError: false,
      };
    case restoreDeckActions.restoreDeckFailed:
      return {
        ...state,
        ...action.payload,
        isLoading: false,
        restoreDeckError: true,
      };
    case restoreDeckActions.restoreDeckReset:
      return {
        ...state,
        ...action.payload,
        isLoading: false,
        restoreDeckError: null,
      };
    default:
      return state;
  }
};
