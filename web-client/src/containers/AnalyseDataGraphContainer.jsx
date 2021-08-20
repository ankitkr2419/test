import React, { useEffect } from "react";
import { useSelector, useDispatch } from "react-redux";

import {
  fetchAnalyseDataThreshold,
  updateFilter,
} from "action-creators/analyseDataGraphActionCreators";
import { getExperimentGraphTargets } from "selectors/experimentTargetSelector";
import AnalyseDataGraphComponent from "components/AnalyseDataGraph";
import { generateTargetOptions } from "components/AnalyseDataGraph/helper";
import { getExperimentId } from "selectors/experimentSelector";

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
  const analyseDataGraphFilters = analyseDataGraphFiltersReducer.toJS();
  const { selectedTarget, isAutoThreshold, isAutoBaseline } =
    analyseDataGraphFilters;

  //current experiment id
  const experimentId = useSelector(getExperimentId);
  console.log("exp id: ", experimentId);//check

  //get login reducer details
  const loginReducer = useSelector((state) => state.loginReducer);
  const loginReducerData = loginReducer.toJS();
  let activeDeckObj = loginReducerData?.decks.find((deck) => deck.isActive);
  const { token } = activeDeckObj;

  //TODO get graph data from reducer
  const data = {};

  //fetch analyseDataGraph data
  useEffect(() => {
    const { selectedTarget, isAutoThreshold, isAutoBaseline } =
      analyseDataGraphFilters;
    let thresholdDataForApi = {
      target_id: selectedTarget?.value,
      auto_threshold: isAutoThreshold,
      threshold: 1.8, //TODO make it dynamic with filters
    };
    dispatch(
      fetchAnalyseDataThreshold({ token, experimentId, thresholdDataForApi })
    );
  }, [dispatch]);

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
