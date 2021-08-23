import React from "react";
import { Line } from "react-chartjs-2";
import { TEMPERATURE_GRAPH_OPTIONS } from "appConstants";

const LineChart = (props) => {
  const { data, width, height } = props;

  return (
    <Line
      width={width}
      height={height}
      data={data}
      options={TEMPERATURE_GRAPH_OPTIONS}
    />
  );
};

LineChart.defaultProps = {
  height: 272,
  width: 830,
};

export default React.memo(LineChart);
