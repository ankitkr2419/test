import React from "react";
import PropTypes from "prop-types";
import { getXAxis } from "selectors/wellGraphSelector";
import TemperatureGraphContainer from "containers/TemperatureGraphContainer";
import WellGraph from "./WellGraph";
import { graphs } from "components/Plate/plateConstant";
import AnalyseDataGraphContainer from "containers/AnalyseDataGraphContainer";

const SidebarGraph = (props) => {
  const {
    headerData,
    activeGraph,
    lineChartData,
    onThresholdChangeHandler,
    toggleGraphFilterActive,
    experimentGraphTargetsList,
    setThresholdError,
    resetThresholdError,
    isThresholdInvalid,
    experimentStatus,
    handleRangeChangeBtn,
    handleResetBtn,
    isInsidePreviewModal,
    isDataFromAPI,
    isExpanded,
    options,
  } = props;

  let cyclesCount = 0;
  // default values : just to make graph look good.
  let xAxisLabels = [1, 2, 3, 4, 5, 6];

  let chartData = lineChartData.toJS();

  // below case can happen if user selects all filter we might get empty chart data
  if (lineChartData.size !== 0) {
    cyclesCount = lineChartData.first().totalCycles;
  }
  if (cyclesCount > 0) {
    xAxisLabels = getXAxis(cyclesCount);
  }

  // if data is fetched from API, then keep xAxis labels same as fetched.
  // Also we hide threshold, that is, remove the last objects from data array.
  if (lineChartData && lineChartData.size !== 0 && isDataFromAPI === true) {
    xAxisLabels = lineChartData?.first().cycles.toJS();

    // Removes threshold line
    // chartData = chartData.filter(
    //   (dataObj, index) => dataObj.label === `index-${index}`
    // );
  }

  const data = {
    labels: xAxisLabels,
    datasets: chartData,
  };

  return (
    <>
      {/* show the well data graph */}
      {activeGraph === graphs.Amplification && (
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
          handleResetBtn={handleResetBtn}
          experimentStatus={experimentStatus}
          isInsidePreviewModal={isInsidePreviewModal}
          isExpanded={isExpanded}
          options={options}
          isDataFromAPI={isDataFromAPI}
        />
      )}
      {/* show temperature graph */}
      {activeGraph === graphs.Temperature && <TemperatureGraphContainer />}

      {/** show analyse data graph */}
      {activeGraph === graphs.AnalyseData && (
        <>
          {(experimentStatus === EXPERIMENT_STATUS.success ||
            experimentStatus === EXPERIMENT_STATUS.stopped ||
            isExpanded === true) && (
            <AnalyseDataGraphContainer
              experimentStatus={experimentStatus}
              isExpanded={isExpanded}
            />
          )}
        </>
      )}
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
  showTempGraph: PropTypes.bool,
};

export default React.memo(SidebarGraph);
