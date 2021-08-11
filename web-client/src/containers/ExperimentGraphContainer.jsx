import React from "react";
import PropTypes from "prop-types";
import SidebarGraph from "components/Plate/Sidebar/Graph/SidebarGraph";
import { useSelector, useDispatch } from "react-redux";
import { getLineChartData } from "selectors/wellGraphSelector";
import { getExperimentGraphTargets } from "selectors/experimentTargetSelector";
import { updateExperimentTargetFilters } from "action-creators/experimentTargetActionCreators";
import { EXPERIMENT_STATUS } from "appConstants";
import { parseFloatWrapper } from "utils/helpers";
import { isAnyThresholdInvalid } from "components/Target/targetHelper";
import { getTemperatureChartData } from "selectors/temperatureGraphSelector";
import { expandLogInitiated } from "action-creators/activityLogActionCreators";

const ExperimentGraphContainer = (props) => {
  const {
    token,
    experimentId,
    headerData,
    showTempGraph,
    setIsSidebarOpen,
    isSidebarOpen,
    experimentStatus,
    isMultiSelectionOptionOn,
    resetSelectedWells,
    isInsidePreviewModal,
  } = props;
  const dispatch = useDispatch();
  // get targets from experiment target reducer(graph : target filters)
  const experimentGraphTargetsList = useSelector(getExperimentGraphTargets);

  // Extracting graph data, Which is populated from websocket
  const lineChartData = useSelector(getLineChartData);

  const toggleSideBar = () => {
    // console log on graph drawer handle click
    console.info("Graph Drawer Handle Clicked");
    // reset the selected wells while closing the sidebar
    if (isSidebarOpen && isMultiSelectionOptionOn === false) {
      resetSelectedWells();
    }
    // console log on graph drawer opened or close
    console.info(
      `Graph Drawer ${isSidebarOpen === true ? "Closed" : "Opened"}`
    );
    setIsSidebarOpen((toggleStateValue) => !toggleStateValue);
  };

  const onThresholdChangeHandler = (threshold, index) => {
    dispatch(
      updateExperimentTargetFilters(
        index,
        "threshold",
        parseFloatWrapper(threshold)
      )
    );
  };

  const toggleGraphFilterActive = (index, isActive) => {
    dispatch(updateExperimentTargetFilters(index, "isActive", !isActive));
  };

  // set threshold error true
  const setThresholdError = (index) => {
    dispatch(updateExperimentTargetFilters(index, "thresholdError", true));
  };

  // reset threshold error to false
  const resetThresholdError = (index) => {
    dispatch(updateExperimentTargetFilters(index, "thresholdError", false));
  };

  const handleRangeChangeBtn = ({ xMax, xMin, yMax, yMin }) => {
    const requestBody = {
      x_axis_min: xMin,
      x_axis_max: xMax,
      y_axis_min: yMin,
      y_axis_max: yMax,
    };

    //dispatch API
    // dispatch(
    //   expandLogInitiated({
    //     params: requestBody,
    //     experimentId: experimentId,
    //     token: token,
    //   })
    // );
  };

  return (
    <SidebarGraph
      headerData={headerData}
      showTempGraph={showTempGraph}
      experimentStatus={experimentStatus}
      lineChartData={lineChartData}
      experimentGraphTargetsList={experimentGraphTargetsList}
      isSidebarOpen={isSidebarOpen}
      toggleSideBar={toggleSideBar}
      onThresholdChangeHandler={onThresholdChangeHandler}
      toggleGraphFilterActive={toggleGraphFilterActive}
      setThresholdError={setThresholdError}
      resetThresholdError={resetThresholdError}
      isThresholdInvalid={isAnyThresholdInvalid(experimentGraphTargetsList)}
      handleRangeChangeBtn={handleRangeChangeBtn}
      isInsidePreviewModal={isInsidePreviewModal}
    />
  );
};

ExperimentGraphContainer.propTypes = {
  experimentStatus: PropTypes.string,
};

export { ExperimentGraphContainer };
