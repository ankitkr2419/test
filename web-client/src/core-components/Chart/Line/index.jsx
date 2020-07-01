import React from 'react';
import { Line } from 'react-chartjs-2';

const LineChart = (props) => {
	const { data, width, height } = props;
	const options = {};

	return <Line width={width} height={height} data={data} options={options} />;
};

LineChart.defaultProps = {
	height: 275,
	width: 830,
};

export default React.memo(LineChart);
