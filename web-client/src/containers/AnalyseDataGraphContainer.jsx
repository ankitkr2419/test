import React, { useState } from "react";
import { useSelector, useDispatch } from "react-redux";

import { updateFilter } from "action-creators/analyseDataGraphFiltersActionCreators";
import { getExperimentGraphTargets } from "selectors/experimentTargetSelector";
import AnalyseDataGraphComponent from "components/AnalyseDataGraph";
import { generateTargetOptions } from "components/AnalyseDataGraph/helper";

const AnalyseDataGraphContainer = (props) => {
  const dispatch = useDispatch();

  // get targets from experiment target reducer(graph : target filters)
  const experimentGraphTargetsList = useSelector(getExperimentGraphTargets);
  const targetsData = experimentGraphTargetsList.toJS();
  // transform targets into (label, value) instead (target_name, target_id) to use in dropdown
  const targetOptions = generateTargetOptions(targetsData);

  //access filters from redux
  const analyseDataGraphFiltersReducer = useSelector(
    (state) => state.analyseDataGraphFiltersReducer
  );
  const filters = analyseDataGraphFiltersReducer.toJS();
  const { selectedTarget, isAutoThreshold, isAutoBaseline } = filters;

  //TODO get graph data from reducer
  const data = {};

  const onTargetChanged = (value) => {
    dispatch(updateFilter({ selectedTarget: value }));
  };

  return (
    <AnalyseDataGraphComponent
      data={data}
      targetOptions={targetOptions}
      selectedTarget={selectedTarget}
      onTargetChanged={onTargetChanged}
    />
  );
};

export default React.memo(AnalyseDataGraphContainer);
