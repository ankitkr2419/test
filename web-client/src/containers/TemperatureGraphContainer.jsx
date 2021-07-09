import React from "react";
import { useSelector } from "react-redux";
import { LineChart } from "core-components";
import { Text } from "shared-components";
import { getTemperatureChartData } from "selectors/temperatureGraphSelector";
import styled from "styled-components";

const options = {
  legend: {
    display: false
  },
  scales: {
    xAxes: [
      {
        scaleLabel: {
          display: true,
          labelString: "minutes"
        },
        type: "linear",
        ticks: {
          source: "data",
          beginAtZero: true,
          suggestedMin: 0,
          min: 0
        }
      }
    ],
    yAxes: [
      {
        scaleLabel: {
          display: true,
          labelString: "Temperature"
        }
      }
    ]
  }
};

const TemperatureGraphContainer = (props) => {
  // Extracting temperature graph data, Which is populated from websocket
  const temperatureChartData = useSelector(getTemperatureChartData);
  return (
    <div>
      <Text size={20} className="text-default mb-4">
        Temperature Plot
      </Text>
      <GraphCard>
        <LineChart data={temperatureChartData} options={options} />
      </GraphCard>
    </div>
  );
};

const GraphCard = styled.div`
  width: 830px;
  height: 280px;
  background: #ffffff 0% 0% no-repeat padding-box;
  border: 1px solid #707070;
  padding: 8px;
  margin: 0 0 32px 0;
`;

export default React.memo(TemperatureGraphContainer);
