import {
  expandLogActions,
  activityLogActions,
  mailReportActions,
} from "actions/activityLogActions";

export const expandLogInitiated = ({ params, experimentId, token }) => ({
  type: expandLogActions.expandLogInitiated,
  payload: { params, experimentId, token },
});

export const expandLogFailed = ({ error }) => ({
  type: expandLogActions.expandLogFailure,
  payload: { error },
});

export const mailReportInitiated = ({ token, experimentId }) => ({
  type: mailReportActions.mailReportInitiated,
  payload: { token, experimentId },
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
