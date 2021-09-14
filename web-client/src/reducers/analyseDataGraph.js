import { fromJS } from "immutable";
import {
  analyseDataGraphFilterActions,
  fetchAnalyseDataWithBaselineActions,
  fetchAnalyseDataWithThresholdActions,
} from "actions/analyseDataGraphActions";

const analyseDataGraphFiltersInitialState = fromJS({
  selectedTarget: null,
  isAutoThreshold: true,
  threshold: 0,
  isAutoBaseline: true,
  startCycle: 0,
  endCycle: 0,
});

// analyse data graph filters reducer
export const analyseDataGraphFiltersReducer = (
  state = analyseDataGraphFiltersInitialState,
  action
) => {
  switch (action.type) {
    case analyseDataGraphFilterActions.updateFilter:
      return state.merge({
        ...action.payload,
      });

    case analyseDataGraphFilterActions.resetAllFilters:
      return state.merge({
        isAutoThreshold: true,
        threshold: 0,
        isAutoBaseline: true,
        startCycle: 0,
        endCycle: 0,
      });

    case analyseDataGraphFilterActions.resetThresholdFilter:
      return state.merge({
        isAutoThreshold: true,
        threshold: 0,
      });

    case analyseDataGraphFilterActions.resetBaselineFilter:
      return state.merge({
        isAutoBaseline: true,
        startCycle: 0,
        endCycle: 0,
      });

    default:
      return state;
  }
};

/**
 * Fetch with threshold
 */
const analyseDataGraphThresholdInitialState = fromJS({
  isLoading: false,
  isThresholdApiError: null,
  thresholdApiData: null,
});

export const analyseDataGraphThresholdReducer = (
  state = analyseDataGraphThresholdInitialState,
  action
) => {
  switch (action.type) {
    case fetchAnalyseDataWithThresholdActions.initiateAction:
      return state.merge({
        isLoading: true,
        isThresholdApiError: null,
      });
    case fetchAnalyseDataWithThresholdActions.successAction:
      return state.merge({
        isLoading: false,
        isThresholdApiError: false,
        thresholdApiData: action.payload.response,
      });
    case fetchAnalyseDataWithThresholdActions.failureAction:
      return state.merge({
        isLoading: false,
        isThresholdApiError: true,
      });
    case fetchAnalyseDataWithThresholdActions.resetAction:
      return state.merge({
        isLoading: false,
        isThresholdApiError: null,
        thresholdApiData: null,
      });

    default:
      return state;
  }
};

/**
 * Fetch with Baseline
 */
const analyseDataGraphBaselineInitialState = fromJS({
  isLoading: false,
  isBaselineApiError: null,
  baselineApiData: null,
});

export const analyseDataGraphBaselineReducer = (
  state = analyseDataGraphBaselineInitialState,
  action
) => {
  switch (action.type) {
    case fetchAnalyseDataWithBaselineActions.initiateAction:
      return state.merge({
        isLoading: true,
        isBaselineApiError: null,
      });
    case fetchAnalyseDataWithBaselineActions.successAction:
      return state.merge({
        isLoading: false,
        isBaselineApiError: false,
        baselineApiData: action.payload.response,
      });
    case fetchAnalyseDataWithBaselineActions.failureAction:
      return state.merge({
        isLoading: false,
        isBaselineApiError: true,
      });
    case fetchAnalyseDataWithBaselineActions.resetAction:
      return state.merge({
        isLoading: false,
        isBaselineApiError: null,
        baselineApiData: null,
      });

    default:
      return state;
  }
};
