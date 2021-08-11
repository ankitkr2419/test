import React from "react";
import PropTypes from "prop-types";
import { getXAxis } from "selectors/wellGraphSelector";
import TemperatureGraphContainer from "containers/TemperatureGraphContainer";
import WellGraph from "./WellGraph";

const SidebarGraph = (props) => {
  const {
    headerData,
    showTempGraph,
    lineChartData,
    onThresholdChangeHandler,
    toggleGraphFilterActive,
    experimentGraphTargetsList,
    setThresholdError,
    resetThresholdError,
    isThresholdInvalid,
    experimentStatus,
    handleRangeChangeBtn,
    isInsidePreviewModal,
  } = props;

  let cyclesCount = 0;
  // below case can happen if user selects all filter we might get empty chart data
  if (lineChartData.size !== 0) {
    cyclesCount = lineChartData.first().totalCycles;
  }

  const data = {
    labels: getXAxis(cyclesCount),
    datasets: lineChartData.toJS(),
  };
  return (
    <>
      {/* show the well data graph if showTempGraph flag is off */}
      {!showTempGraph && (
        <WellGraph
          data={data}
          headerData={headerData}
          experimentGraphTargetsList={experimentGraphTargetsList}
          onThresholdChangeHandler={onThresholdChangeHandler}
          toggleGraphFilterActive={toggleGraphFilterActive}
          setThresholdError={setThresholdError}
          resetThresholdError={resetThresholdError}
          isThresholdInvalid={isThresholdInvalid}
          handleRangeChangeBtn={handleRangeChangeBtn}
          experimentStatus={experimentStatus}
          isInsidePreviewModal={isInsidePreviewModal}
        />
      )}
      {/* show temperature graph if showTempGraph flag is on */}
      {showTempGraph && <TemperatureGraphContainer />}
    </>
  );
};

SidebarGraph.propTypes = {
  isExperimentRunning: PropTypes.bool.isRequired,
  lineChartData: PropTypes.object.isRequired,
  isSidebarOpen: PropTypes.bool.isRequired,
  toggleSideBar: PropTypes.func.isRequired,
  onThresholdChangeHandler: PropTypes.func.isRequired,
  toggleGraphFilterActive: PropTypes.func.isRequired,
  experimentGraphTargetsList: PropTypes.object.isRequired,
  isExperimentSucceeded: PropTypes.bool.isRequired,
};

export default React.memo(SidebarGraph);
