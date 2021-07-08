import React from "react";
import { LineChart } from "core-components";
import { Text } from "shared-components";
import styled from "styled-components";
import PropTypes from "prop-types";
import {
  MIN_THRESHOLD,
  MAX_THRESHOLD,
} from "components/Target/targetConstants";
import GraphFilters from "./GraphFilters";

const options = {
  legend: {
    display: false,
  },
  scales: {
    xAxes: [
      {
        offset: true,
      },
    ],
  },
  animation: {
    duration: 10000,
    easing: "linear",
  },
};

const line1 = [
  0.5415625, 0.4934375, 0.345645, 0.6537, 0.2341, 0.815632, 0.75345,
];
const line2 = [
  0.855, 0.5965625, 0.15625, 0.1234375, 0.345735, 0.75677, 0.14354,
];
const line3 = [
  0.895625, 0.59877, 0.5625, 0.9756256, 0.4657625, 0.664256, 0.625,
];
const threshold = [2, 2, 2, 2, 2, 2, 2];

const data1 = {
  labels: [1, 2, 3, 4, 5, 6, 7],
  datasets: [
    {
      fill: false,
      borderWidth: 2,
      pointRadius: 0,
      pointBorderColor: "rgba(148,147,147,1)",
      pointBackgroundColor: "#fff",
      pointBorderWidth: 0,
      pointHoverRadius: 0,
      pointHoverBackgroundColor: "rgba(148,147,147,1)",
      pointHoverBorderColor: "rgba(148,147,147,1)",
      pointHoverBorderWidth: 0,
      lineTension: 0.1,
      borderCapStyle: "butt",
      label: "index-0",
      borderColor: "#F590B2",
      data: line1,
      totalCycles: 6,
      cycles: [1, 2],
    },
    {
      fill: false,
      borderWidth: 2,
      pointRadius: 0,
      pointBorderColor: "rgba(148,147,147,1)",
      pointBackgroundColor: "#fff",
      pointBorderWidth: 0,
      pointHoverRadius: 0,
      pointHoverBackgroundColor: "rgba(148,147,147,1)",
      pointHoverBorderColor: "rgba(148,147,147,1)",
      pointHoverBorderWidth: 0,
      lineTension: 0.1,
      borderCapStyle: "butt",
      label: "index-1",
      borderColor: "#F590B2",
      data: line2,
      totalCycles: 6,
      cycles: [1, 2],
    },
    {
      fill: false,
      borderWidth: 2,
      pointRadius: 0,
      pointBorderColor: "rgba(148,147,147,1)",
      pointBackgroundColor: "#fff",
      pointBorderWidth: 0,
      pointHoverRadius: 0,
      pointHoverBackgroundColor: "rgba(148,147,147,1)",
      pointHoverBorderColor: "rgba(148,147,147,1)",
      pointHoverBorderWidth: 0,
      lineTension: 0.1,
      borderCapStyle: "butt",
      label: "index-2",
      borderColor: "#F590B2",
      data: line3,
      totalCycles: 6,
      cycles: [1, 2],
    },
    {
      label: "COVID",
      fill: false,
      borderWidth: 2,
      pointRadius: 0,
      borderColor: "#F590B2",
      pointBorderColor: "#a2ee95",
      borderDash: [10, 5],
      pointBackgroundColor: "#fff",
      pointBorderWidth: 0,
      pointHoverRadius: 0,
      pointHoverBackgroundColor: "#a2ee95",
      pointHoverBorderColor: "#a2ee95",
      pointHoverBorderWidth: 0,
      data: threshold,
    },
  ],
};

const WellGraph = ({
  data,
  experimentGraphTargetsList,
  onThresholdChangeHandler,
  toggleGraphFilterActive,
  isThresholdInvalid,
  setThresholdError,
  resetThresholdError,
}) => (
  <div>
    <Text size={20} className="text-default mb-4">
      Amplification plot
    </Text>
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
    {isThresholdInvalid && (
      <Text Tag="p" size={14} className="text-danger px-2 mb-1">
        Threshold value should be between {MIN_THRESHOLD} - {MAX_THRESHOLD}
      </Text>
    )}
    <Text size={14} className="text-default mb-0">
      Note: Click on the threshold number to change it.
    </Text>
  </div>
);

const GraphCard = styled.div`
  width: 830px;
  height: 344px;
  background: #ffffff 0% 0% no-repeat padding-box;
  border: 1px solid #707070;
  padding: 8px;
  margin: 0 0 32px 0;
`;

WellGraph.propTypes = {
  experimentGraphTargetsList: PropTypes.object.isRequired,
  onThresholdChangeHandler: PropTypes.func.isRequired,
  toggleGraphFilterActive: PropTypes.func.isRequired,
  data: PropTypes.object.isRequired,
};

export default React.memo(WellGraph);
