import { fromJS } from "immutable";
import wellGraphActions, { updateGraphActions } from "actions/wellGraphActions";
import loginActions from "actions/loginActions";
// import graphData from '../mock-json/graphData.json';

const wellGraphInitialState = fromJS({
  chartData: [],
  // chartData: graphData,
});

export const wellGraphReducer = (state = wellGraphInitialState, action) => {
  switch (action.type) {
    case wellGraphActions.successAction:
      return state.setIn(["chartData"], fromJS(action.payload.data));
    case loginActions.loginReset:
      return wellGraphInitialState;
    default:
      return state;
  }
};

const updateGraphInitialState = fromJS({
  isLoading: false,
  error: null,
  data: [],
});

export const updateWellGraphReducer = (
  state = updateGraphInitialState,
  action
) => {
  switch (action.type) {
    case updateGraphActions.updateGraphInitiated:
      return state.merge({
        isLoading: true,
        error: false,
      });

    case updateGraphActions.updateGraphSucceeded:
      return state.merge({
        isLoading: false,
        error: false,
        data: action.payload.response,
      });

    case updateGraphActions.updateGraphFailure:
      return state.merge({
        isLoading: false,
        error: true,
      });

    case loginActions.loginReset:
      return updateGraphInitialState;
    default:
      return state;
  }
};
