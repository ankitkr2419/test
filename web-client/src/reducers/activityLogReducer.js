import { fromJS } from "immutable";
import {
  activityLogActions,
  mailReportActions,
} from "actions/activityLogActions";

const mailReportInitialState = fromJS({
  isLoading: false,
  error: null,
});

// reducer to send report via mail
export const mailReportReducer = (state = mailReportInitialState, action) => {
  switch (action.type) {
    case mailReportActions.mailReportInitiated:
      return state.merge({
        isLoading: true,
      });

    case mailReportActions.mailReportSuccess:
      return state.merge({
        isLoading: false,
        error: false,
        response: action.payload.response,
      });

    case mailReportActions.mailReportFailure:
      return state.merge({
        isLoading: false,
        error: true,
        errorResponse: action.payload.reponse,
      });

    case mailReportActions.mailReportReset:
      return state.merge({
        isLoading: false,
        error: null,
      });

    default:
      return state;
  }
};

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
        activityLogs: [],
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
