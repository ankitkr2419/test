import { fromJS } from "immutable";
import { whiteLightActions } from "actions/whiteLightActions";

const whiteLightInitialState = fromJS({
  isLoading: false,
  isError: null,
  isLightOn: null,
});

export const whiteLightReducer = (state = whiteLightInitialState, action) => {
  switch (action.type) {
    case whiteLightActions.initiateAction:
      return state.merge({ isLoading: true });

    case whiteLightActions.successAction:
      return state.merge({
        isLoading: false,
        isLightOn: true,
      });

    case whiteLightActions.failureAction:
      return state.merge({ isLoading: false, isError: true });

    case whiteLightActions.resetAction:
      return whiteLightInitialState;

    default:
      return state;
  }
};
