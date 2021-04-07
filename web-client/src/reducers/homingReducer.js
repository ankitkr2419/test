import { homingActions } from "actions/homingActions";

export const initialState = {
  isLoading: false,
  homingData: {},
  serverErrors: {},
  homingAllDeckCompletionPercentage: 50,
  isHomingActionCompleted: false,
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
      console.log("PROGRESS----------->", action.payload);
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
