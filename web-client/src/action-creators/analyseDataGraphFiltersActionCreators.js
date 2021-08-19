import { analyseDataGraphFilterActions } from "actions/analyseDataGraphActions";

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
