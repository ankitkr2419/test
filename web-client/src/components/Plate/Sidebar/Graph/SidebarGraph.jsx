import React, { useState } from "react";
import Sidebar from "components/Sidebar";
import PropTypes from "prop-types";
import { getXAxis } from "selectors/wellGraphSelector";
import TemperatureGraphContainer from "containers/TemperatureGraphContainer";
import { Switch } from "core-components";
import { SwitchWrapper } from "shared-components/SwitchWrapper";
import WellGraph from "./WellGraph";

const SidebarGraph = (props) => {
  const {
    showTempGraph,
    isExperimentRunning,
    lineChartData,
    onThresholdChangeHandler,
    toggleGraphFilterActive,
    experimentGraphTargetsList,
    isExperimentSucceeded,
    setThresholdError,
    resetThresholdError,
    isThresholdInvalid,
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

  if (isExperimentRunning === true || isExperimentSucceeded === true) {
    return (
      <>
        {/* show the well data graph if showTempGraph flag is off */}
        {!showTempGraph && (
          <WellGraph
            data={data}
            experimentGraphTargetsList={experimentGraphTargetsList}
            onThresholdChangeHandler={onThresholdChangeHandler}
            toggleGraphFilterActive={toggleGraphFilterActive}
            setThresholdError={setThresholdError}
            resetThresholdError={resetThresholdError}
            isThresholdInvalid={isThresholdInvalid}
          />
        )}
        {/* show temperature graph if showTempGraph flag is on */}
        {showTempGraph && <TemperatureGraphContainer />}
      </>
    );
  }
  return null;
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
