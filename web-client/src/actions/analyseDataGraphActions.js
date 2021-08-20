/**
 * TO MANIPULATE ANALYSE DATA GRAPH FILTERS
 */
export const analyseDataGraphFilterActions = {
  updateFilter: "UPDATE_ANALYSE_DATA_FILTER",
  resetAllFilters: "RESET_ALL_ANALYSE_DATA_FILTER",
  resetThresholdFilter: "RESET_THRESHOLD_FILTER",
  resetBaselineFilter: "RESET_BASELINE_FILTER",
};

/**
 * TO FETCH ANALYSE DATA GRAPH DATA WITH THRESHOLD VALUES
 */
export const fetchAnalyseDataWithThresholdActions = {
  initiateAction: "FETCH_ANALYSE_DATA_THRESHOLD_INITIATED",
  successAction: "FETCH_ANALYSE_DATA_THRESHOLD_SUCCESS",
  failureAction: "FETCH_ANALYSE_DATA_THRESHOLD_FAILURE",
  resetAction: "FETCH_ANALYSE_DATA_THRESHOLD_RESET",
};

/**
 * TO FETCH ANALYSE DATA GRAPH DATA WITH BASELINE VALUES
 */
export const fetchAnalyseDataWithBaselineActions = {
  initiateAction: "FETCH_ANALYSE_DATA_BASELINE_INITIATED",
  successAction: "FETCH_ANALYSE_DATA_BASELINE_SUCCESS",
  failureAction: "FETCH_ANALYSE_DATA_BASELINE_FAILURE",
  resetAction: "FETCH_ANALYSE_DATA_BASELINE_RESET",
};
