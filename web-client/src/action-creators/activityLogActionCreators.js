import { activityLogActions } from "actions/activityLogActions";

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
