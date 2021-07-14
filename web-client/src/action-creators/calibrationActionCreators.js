import { calibrationActions } from "actions/calibrationActions";

export const calibrationInitiated = (token) => ({
  type: calibrationActions.calibrationInitiated,
  payload: {
    token,
  },
});

export const calibrationFailed = ({ error }) => ({
  type: calibrationActions.calibrationFailure,
  payload: {
    error,
  },
});
