import React from "react";
import { Line } from "react-chartjs-2";

const LineChart = (props) => {
  const { data, width, height, options } = props;

  return (
    <Line redraw width={width} height={height} data={data} options={options} />
  );
};

LineChart.defaultProps = {
  height: 272,
  width: 830,
};

export default React.memo(LineChart);
