import React from 'react';
import { Line } from 'react-chartjs-2';

const LineChart = (props) => {
	const { data, width, height } = props;

	const options = {
		legend: {
			labels: {
				filter(item, chart) {
					// Logic to remove a particular legend item goes here
					return !item.text.includes('index');
				},
			},
		},
	};

	return <Line width={width} height={height} data={data} options={options} />;
};

LineChart.defaultProps = {
	height: 275,
	width: 830,
};

export default React.memo(LineChart);
