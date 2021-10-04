import React, { useEffect } from "react";
import { useSelector, useDispatch } from "react-redux";

import {
  fetchAnalyseDataThreshold,
  fetchAnalyseDataBaseline,
  updateFilter,
  resetThresholdFilter,
  resetBaselineFilter,
  resetAllFiltersOfAnalyseDataGraph,
} from "action-creators/analyseDataGraphActionCreators";
import { getExperimentGraphTargets } from "selectors/experimentTargetSelector";
import AnalyseDataGraphComponent from "components/AnalyseDataGraph";
import {
  generateTargetOptions,
  lineConfigs,
  lineConfigThreshold,
} from "components/AnalyseDataGraph/helper";
import { getExperimentId } from "selectors/experimentSelector";
import { GRAY_COLOR, PINK_COLOR } from "appConstants";

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
  const {
    selectedTarget,
    isAutoThreshold,
    threshold,
    isAutoBaseline,
    startCycle,
    endCycle,
  } = analyseDataGraphFilters;

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

  //threshold data reducer
  const analyseDataGraphThresholdReducer = useSelector(
    (state) => state.analyseDataGraphThresholdReducer
  );
  let thresholdApiResponse = analyseDataGraphThresholdReducer.toJS();

  //fetch analyseDataGraph data
  useEffect(() => {
    dispatch(
      fetchAnalyseDataThreshold({
        token,
        experimentId,
        target_id: selectedTarget?.value,
        auto_threshold: isAutoThreshold,
        threshold: threshold,
        start_cycle: startCycle,
        end_cycle: endCycle,
      })
    );
    dispatch(
      fetchAnalyseDataBaseline({
        token,
        experimentId,
        auto_baseline: isAutoBaseline,
        start_cycle: startCycle,
        end_cycle: endCycle,
      })
    );
  }, [dispatch, threshold, startCycle, endCycle]);

  const onFiltersChanged = (changedFilters) => {
    dispatch(updateFilter(changedFilters));
  };

  const onResetThresholdFilter = () => {
    dispatch(resetThresholdFilter());
  };

  const onResetBaselineFilter = () => {
    dispatch(resetBaselineFilter());
  };

  const onTargetChanged = (value) => {
    dispatch(resetAllFiltersOfAnalyseDataGraph());
    onFiltersChanged({ selectedTarget: value });
  };

  const getThresholdValue = () => {
    //use default threshold value
    let thresholdValue = threshold;

    //if auto threshold is true get threshold value from api
    if (isAutoThreshold === true) {
      let targetsFromThresholdApi =
        (thresholdApiResponse?.thresholdApiData !== null &&
          thresholdApiResponse.thresholdApiData[0]?.targets) ||
        [];
      let targetObjFromApi = targetsFromThresholdApi.find(
        (ele) => ele?.target_id === selectedTarget?.value
      );
      thresholdValue = targetObjFromApi?.threshold;
    }

    return thresholdValue;
  };

  //create graph data
  if (baselineData?.length > 0) {
    //x axis data
    xAxisLabels = baselineData[0].cycle;

    //y axis data
    let borderColor = GRAY_COLOR; // default line color

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
      let thresholdValue = getThresholdValue();
      let apiObject = baselineDataForSelectedTarget[0];
      let thresholdBorderColor = PINK_COLOR;

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
      onFiltersChanged={onFiltersChanged}
      onResetThresholdFilter={onResetThresholdFilter}
      onResetBaselineFilter={onResetBaselineFilter}
    />
  );
};

export default React.memo(AnalyseDataGraphContainer);
