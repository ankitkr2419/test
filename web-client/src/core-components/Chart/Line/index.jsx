import React from "react";
import { Line } from "react-chartjs-2";
import { TEMPERATURE_GRAPH_OPTIONS } from "appConstants";

const LineChart = (props) => {
  const { data, width, height, options, isDataFromAPI } = props;

  return (
    <Line
      redraw={isDataFromAPI}
      width={width}
      height={height}
      data={data}
      options={options}
    />
  );
};

LineChart.defaultProps = {
  height: 272,
  width: 830,
};

export default React.memo(LineChart);
