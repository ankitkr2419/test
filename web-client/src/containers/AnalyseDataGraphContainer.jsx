import React, { useEffect } from "react";
import { useSelector, useDispatch } from "react-redux";

import {
  fetchAnalyseDataThreshold,
  fetchAnalyseDataBaseline,
  updateFilter,
} from "action-creators/analyseDataGraphActionCreators";
import { getExperimentGraphTargets } from "selectors/experimentTargetSelector";
import AnalyseDataGraphComponent from "components/AnalyseDataGraph";
import {
  generateTargetOptions,
  lineConfigs,
  lineConfigThreshold,
} from "components/AnalyseDataGraph/helper";
import { getExperimentId } from "selectors/experimentSelector";

const AnalyseDataGraphContainer = (props) => {
  const { isInsidePreviewModal } = props;
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

  //get login reducer details
  const loginReducer = useSelector((state) => state.loginReducer);
  const loginReducerData = loginReducer.toJS();
  let activeDeckObj = loginReducerData?.decks.find((deck) => deck.isActive);
  const { token } = activeDeckObj;

  // default values for graph
  let xAxisLabels = [];
  let chartData = [];

  //baseline data reducer
  const analyseDataGraphBaselineReducer = useSelector(
    (state) => state.analyseDataGraphBaselineReducer
  );
  const baselineApiResponse = analyseDataGraphBaselineReducer.toJS();
  const baselineData = baselineApiResponse?.baselineApiData?.data;

  //fetch analyseDataGraph data
  useEffect(() => {
    let thresholdDataForApi = {
      target_id: selectedTarget?.value,
      auto_threshold: isAutoThreshold,
      threshold: 1.8, //TODO make it dynamic with filters
    };
    dispatch(
      fetchAnalyseDataThreshold({ token, experimentId, thresholdDataForApi })
    );

    let baselineDataForApi = {
      auto_baseline: isAutoBaseline,
      start_cycle: 1, //TODO make it dynamic with filters
      end_cycle: 5, //TODO make it dynamic with filters
    };

    dispatch(
      fetchAnalyseDataBaseline({ token, experimentId, baselineDataForApi })
    );
  }, [dispatch]);

  const onTargetChanged = (value) => {
    dispatch(updateFilter({ selectedTarget: value }));
  };

  //create graph data
  if (baselineData?.length > 0) {
    //x axis data
    xAxisLabels = baselineData[0].cycle;

    //y axis data
    let borderColor = "rgba(148,147,147,1)"; // default line color

    //filter by target id
    let baselineDataForSelectedTarget = baselineData.filter(
      (obj) => obj.target_id === selectedTarget?.value
    );

    chartData = baselineDataForSelectedTarget.map((obj, index) => {
      return {
        ...lineConfigs,
        label: `index-${index}`,
        borderColor,
        data: obj.f_value || [],
        totalCycles: obj.total_cycles || 0,
        cycle: obj.cycle || [],
      };
    });

    // if we don't have chartData then no need to calculate threshold value
    if (chartData.length > 0) {
      let autoThresholdOfSelectedTarget = selectedTarget?.threshold;
      let thresholdValue = autoThresholdOfSelectedTarget; //TODO get actual threshold value once backend api is ready
      let apiObject = baselineDataForSelectedTarget[0];
      let thresholdBorderColor = `rgba(245,144,178,1)`; //TODO check if it is required to be dynamic

      let thresholdData = {
        ...lineConfigThreshold,
        label: selectedTarget?.label,
        totalCycles: apiObject.total_cycles || 0,
        data: apiObject.cycle.map(() => thresholdValue),
        borderColor: thresholdBorderColor,
      };
      //merge graph data and threshold data
      chartData = [...chartData, thresholdData];
    }
  }

  //create graph data
  let data = {
    labels: xAxisLabels,
    datasets: chartData,
  };

  return (
    <AnalyseDataGraphComponent
      data={data}
      targetOptions={targetOptions}
      selectedTarget={selectedTarget}
      onTargetChanged={onTargetChanged}
      analyseDataGraphFilters={analyseDataGraphFilters}
      isInsidePreviewModal={isInsidePreviewModal}
    />
  );
};

export default React.memo(AnalyseDataGraphContainer);
