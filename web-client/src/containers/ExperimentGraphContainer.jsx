import React, { useEffect } from "react";
import PropTypes from "prop-types";
import SidebarGraph from "components/Plate/Sidebar/Graph/SidebarGraph";
import { useSelector, useDispatch } from "react-redux";
import { getLineChartData } from "selectors/wellGraphSelector";
import { getExperimentGraphTargets } from "selectors/experimentTargetSelector";
import { updateExperimentTargetFilters } from "action-creators/experimentTargetActionCreators";
import { parseFloatWrapper } from "utils/helpers";
import { isAnyThresholdInvalid } from "components/Target/targetHelper";
import { wellGraphSucceeded } from "action-creators/wellGraphActionCreators";

const ExperimentGraphContainer = (props) => {
  const {
    headerData,
    activeGraph,
    setIsSidebarOpen,
    isSidebarOpen,
    experimentStatus,
    isMultiSelectionOptionOn,
    resetSelectedWells,
    isInsidePreviewModal,
    isExpanded,
    handleRangeChangeBtn,
    handleResetBtn,
    options,
    isDataFromAPI,
  } = props;

  const dispatch = useDispatch();

  // get targets from experiment target reducer(graph : target filters)
  const experimentGraphTargetsList = useSelector(getExperimentGraphTargets);

  // Extracting graph data, Which is populated from websocket
  const lineChartData = useSelector(getLineChartData);

  // Update Well Graph Reducer
  const updateWellGraphReducer = useSelector(
    (state) => state.updateWellGraphReducer
  );
  const updateWellGraphReducerData = updateWellGraphReducer.toJS();
  const { isLoading, error, data } = updateWellGraphReducerData;

  useEffect(() => {
    if (isLoading === false && error === false) {
      dispatch(wellGraphSucceeded(data));
    }
  }, [isLoading, error]);

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

  return (
    <SidebarGraph
      headerData={headerData}
      activeGraph={activeGraph}
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
      handleResetBtn={handleResetBtn}
      isInsidePreviewModal={isInsidePreviewModal}
      isDataFromAPI={isDataFromAPI}
      isExpanded={isExpanded}
      options={options}
    />
  );
};

ExperimentGraphContainer.propTypes = {
  experimentStatus: PropTypes.string,
};

export { ExperimentGraphContainer };
