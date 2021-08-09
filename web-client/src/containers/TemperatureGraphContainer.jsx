import React from "react";
import { useSelector } from "react-redux";
import { LineChart } from "core-components";
import { getTemperatureChartData } from "selectors/temperatureGraphSelector";
import styled from "styled-components";

const options = {
  legend: {
    // display: false
    labels: {
      fontSize: 15,
      fontStyle: "bold",
    },
  },
  scales: {
    xAxes: [
      {
        scaleLabel: {
          display: true,
          labelString: "Time (minutes)",
          fontSize: 15,
          fontStyle: "bold",
          padding: 10,
        },
        offset: true,
        type: "linear",
        ticks: {
          source: "data",
          beginAtZero: true,
          suggestedMin: 0,
          min: 0,
          fontSize: 15,
          fontStyle: "bold",
        },
      },
    ],
    yAxes: [
      {
        scaleLabel: {
          display: true,
          labelString: "Temperature (°C)",
          fontSize: 15,
          fontStyle: "bold",
          padding: 10,
        },
        ticks: {
          fontSize: 15,
          fontStyle: "bold",
        },
      },
    ],
  },
};

const TemperatureGraphContainer = (props) => {
  // Extracting temperature graph data, Which is populated from websocket
  const temperatureChartData = useSelector(getTemperatureChartData);
  return (
    <GraphCard>
      <LineChart data={temperatureChartData} options={options} />
    </GraphCard>
  );
};

const GraphCard = styled.div`
  width: 960px;
  height: 326px;
  background: #ffffff 0% 0% no-repeat padding-box;
  border: 1px solid #707070;
  padding: 8px;
  margin: 0 0 16px 0;
`;

export default React.memo(TemperatureGraphContainer);
