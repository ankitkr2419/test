import { saveReportActions } from "actions/reportActions";

export const saveReportInitiated = (token, data) => ({
  type: saveReportActions.saveReportInitiated,
  payload: {
    token,
    data,
  },
});

export const saveReportFailed = ({ error }) => ({
  type: saveReportActions.saveReportFailure,
  payload: {
    error,
  },
});
