import { fromJS } from "immutable";
import { activityLogActions } from "actions/activityLogActions";

const activityLogInitialState = fromJS({
  isLoading: false,
  error: null,
  activityLogs: [],
});

export const activityLogReducer = (state = activityLogInitialState, action) => {
  switch (action.type) {
    case activityLogActions.activityLogInitiated:
      return state.merge({
        isLoading: true,
        error: null,
        activityLogs: []
      });
    case activityLogActions.activityLogSuccess:
      return state.merge({
        isLoading: false,
        error: false,
        activityLogs: action.payload.response,
      });
    case activityLogActions.activityLogFailure:
      return state.merge({
        isLoading: false,
        error: true,
      });
    case activityLogActions.activityLogReset:
      return state.merge({
        isLoading: false,
        error: null,
        activityLogs: [],
      });
    default:
      return state;
  }
};
