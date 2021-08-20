import {
  analyseDataGraphFilterActions,
  fetchAnalyseDataWithBaselineActions,
  fetchAnalyseDataWithThresholdActions,
} from "actions/analyseDataGraphActions";

/**
 * Filters
 */
export const updateFilter = (payload) => ({
  type: analyseDataGraphFilterActions.updateFilter,
  payload: payload,
});

export const resetAllFiltersOfAnalyseDataGraph = () => ({
  type: analyseDataGraphFilterActions.resetAllFilters,
});

export const resetThresholdFilter = () => ({
  type: analyseDataGraphFilterActions.resetThresholdFilter,
});

export const resetBaselineFilter = () => ({
  type: analyseDataGraphFilterActions.resetBaselineFilter,
});

/**
 * Fetch with threshold
 */
export const fetchAnalyseDataThreshold = (payload) => ({
  type: fetchAnalyseDataWithThresholdActions.initiateAction,
  payload: payload,
});

export const fetchAnalyseDataThresholdFailed = ({ error }) => ({
  type: fetchAnalyseDataWithThresholdActions.failureAction,
  payload: { error },
});
