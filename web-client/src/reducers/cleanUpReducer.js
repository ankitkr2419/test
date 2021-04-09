import { cleanUpActions } from "actions/cleanUpActions";

export const initialState = {
  isLoading: false,
  cleanUpData: {},
  serverErrors: {},
  isCleanUpActionCompleted: false,
};

export const cleanUpReducer = (state = initialState, action = {}) => {
  switch(action.type) {
    case cleanUpActions.cleanUpActionInitiated:
      return {
        ...state,
        ...action.payload,
        isLoading: true,
        isCleanUpActionCompleted: true,
        cleanUpActionInProgress: {},
        cleanUpActionInCompleted: {}
      };
    case cleanUpActions.cleanUpActionSuccess:
      return { ...state, ...action.payload, isLoading: false };
    case cleanUpActions.cleanUpActionFailed:
      return { ...state, ...action.payload, isLoading: false };
    case cleanUpActions.cleanUpActionInProgress:
      return { ...state, ...action.payload, isLoading: false, isCleanUpActionCompleted: true };
    case cleanUpActions.cleanUpActionInCompleted:
      return { ...state, ...action.payload, isLoading: false, isCleanUpActionCompleted: false };
    default:
      return state;
  }
};
