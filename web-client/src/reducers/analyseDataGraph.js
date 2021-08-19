import { fromJS } from "immutable";
import { analyseDataGraphFilterActions } from "actions/analyseDataGraphActions";

const analyseDataGraphFiltersInitialState = fromJS({
  selectedTarget: null,
  isAutoThreshold: true,
  isAutoBaseline: true,
});

// reducer to send report via mail
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
        isAutoBaseline: true,
      });

    case analyseDataGraphFilterActions.resetThresholdFilter:
      return state.merge({
        isAutoThreshold: true,
      });

    case analyseDataGraphFilterActions.resetBaselineFilter:
      return state.merge({
        isAutoBaseline: true,
      });

    default:
      return state;
  }
};
