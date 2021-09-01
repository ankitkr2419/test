import { fromJS } from "immutable";
import {
  temperatureDataSuccess,
  temperatureGraphActions,
} from "actions/temperatureGraphActions";
import loginActions from "actions/loginActions";

const temperatureGraphInitialState = fromJS({
  isLoading: false,
  error: true,
  temperatureData: [],
});

export const temperatureGraphReducer = (
  state = temperatureGraphInitialState,
  action
) => {
  switch (action.type) {
    case temperatureDataSuccess.successAction:
      return state.setIn(["temperatureData"], fromJS(action.payload.data));

    case temperatureGraphActions.temperatureGraphInitated:
      return state.merge({
        isLoading: true,
      });

    case temperatureGraphActions.temperatureGraphSuccess:
      return state.merge({
        temperatureData: fromJS(action.payload.response.data),
        isLoading: false,
        error: false,
      });

    case temperatureGraphActions.temperatureGraphFailed:
      return state.merge({
        errorMsg: action.payload.error,
        error: true,
      });

    case loginActions.loginReset:
      return temperatureGraphInitialState;
    default:
      return state;
  }
};
