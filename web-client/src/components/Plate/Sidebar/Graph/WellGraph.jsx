import React from "react";
import { LineChart } from "core-components";
import { Text } from "shared-components";
import styled from "styled-components";
import PropTypes from "prop-types";
import {
  MIN_THRESHOLD,
  MAX_THRESHOLD
} from "components/Target/targetConstants";
import GraphFilters from "./GraphFilters";
import GraphRange from "./GraphRange";

const options = {
  legend: {
    display: false
  },
  scales: {
    xAxes: [
      {
        scaleLabel: {
          display: true,
          labelString: "Cycles",
          fontSize: 15,
          fontStyle: "bold",
          padding: 5
        },
        offset: true,
        ticks: {
          fontSize: 15,
          fontStyle: "bold"
        }
      }
    ],
    yAxes: [
      {
        scaleLabel: {
          display: true,
          labelString: "F-value",
          fontSize: 15,
          fontStyle: "bold",
          padding: 10
        },
        ticks: {
          fontSize: 15,
          fontStyle: "bold"
        }
      }
    ]
  }
  //TODO: will be added later after testing & studying about it properly.
  // This is for animation.

  // animation: {
  //   duration: 1000,
  //   easing: "linear",
  // },
};

const WellGraph = ({
  data,
  experimentGraphTargetsList,
  onThresholdChangeHandler,
  toggleGraphFilterActive,
  isThresholdInvalid,
  setThresholdError,
  resetThresholdError
}) => (
  <div>
    <GraphCard>
      <LineChart data={data} options={options} />
    </GraphCard>
    <GraphFilters
      targets={experimentGraphTargetsList}
      onThresholdChangeHandler={onThresholdChangeHandler}
      toggleGraphFilterActive={toggleGraphFilterActive}
      setThresholdError={setThresholdError}
      resetThresholdError={resetThresholdError}
    />
    <GraphRange />
    {isThresholdInvalid && (
      <Text Tag="p" size={14} className="text-danger px-2 mb-1">
        Threshold value should be between {MIN_THRESHOLD} - {MAX_THRESHOLD}
      </Text>
    )}
    {/* TODO: Un-comment after discussion with client/backend-team */}
    {/* <Text size={14} className="text-default text-center mb-0">
      Note: Click on the threshold number to change it.
    </Text> */}
  </div>
);

const GraphCard = styled.div`
  width: 960px;
  height: 326px;
  background: #ffffff 0% 0% no-repeat padding-box;
  border: 1px solid #707070;
  padding: 8px;
  margin: 0 0 16px 0;
`;

WellGraph.propTypes = {
  experimentGraphTargetsList: PropTypes.object.isRequired,
  onThresholdChangeHandler: PropTypes.func.isRequired,
  toggleGraphFilterActive: PropTypes.func.isRequired,
  data: PropTypes.object.isRequired
};

export default React.memo(WellGraph);
