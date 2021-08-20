import React from "react";
import { useSelector } from "react-redux";
import { LineChart } from "core-components";
import { getTemperatureChartData } from "selectors/temperatureGraphSelector";
import styled from "styled-components";

const TemperatureGraphContainer = (props) => {
  // Extracting temperature graph data, Which is populated from websocket
  const temperatureChartData = useSelector(getTemperatureChartData);
  // console.log("temperatureChartData: ", temperatureChartData);
  return (
    <GraphCard>
      <LineChart data={temperatureChartData} />
    </GraphCard>
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
