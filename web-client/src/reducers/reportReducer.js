import { fromJS } from "immutable";
import { saveReportActions } from "actions/reportActions";

const reportInitialState = fromJS({
  isLoading: false,
  error: null,
});

export const reportReducer = (state = reportInitialState, action) => {
  switch (action.type) {
    case saveReportActions.saveReportInitiated:
      return state.merge({
        isLoading: true,
        error: null,
      });
    case saveReportActions.saveReportSuccess:
      return state.merge({
        isLoading: false,
        error: false,
      });
    case saveReportActions.saveReportFailure:
      return state.merge({
        isLoading: false,
        error: true,
      });
    case saveReportActions.saveReportReset:
      return state.merge({
        isLoading: false,
        error: null,
      });
    default:
      return state;
  }
};
