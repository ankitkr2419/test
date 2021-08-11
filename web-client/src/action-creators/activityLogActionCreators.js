import {
  activityLogActions,
  mailReportActions,
} from "actions/activityLogActions";

export const mailReportInitiated = ({ body, token }) => ({
  type: mailReportActions.mailReportInitiated,
  payload: { body, token },
});

export const mailReportFailed = ({ error }) => ({
  type: mailReportActions.mailReportFailure,
  payload: { error },
});

export const activityLogInitiated = (token) => ({
  type: activityLogActions.activityLogInitiated,
  payload: {
    token,
  },
});

export const activityLogFailed = ({ error }) => ({
  type: activityLogActions.activityLogFailure,
  payload: {
    error,
  },
});
