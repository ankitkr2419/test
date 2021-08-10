import { saveReportActions } from "actions/reportActions";

export const saveReportInitiated = (token, experimentId, data) => ({
  type: saveReportActions.saveReportInitiated,
  payload: {
    token,
    experimentId,
    data,
  },
});

export const saveReportFailed = ({ error }) => ({
  type: saveReportActions.saveReportFailure,
  payload: {
    error,
  },
});
