import { createSelector } from 'reselect';
import { formatTime } from 'utils/helpers';

const lineConfigs = {
	fill: false,
	borderWidth: 2,
	pointRadius: 0,
	pointBorderColor: 'rgba(148,147,147,1)',
	pointBackgroundColor: '#fff',
	pointBorderWidth: 0,
	pointHoverRadius: 0,
	pointHoverBackgroundColor: 'rgba(148,147,147,1)',
	pointHoverBorderColor: 'rgba(148,147,147,1)',
	pointHoverBorderWidth: 0,
	lineTension: 0.1,
	borderCapStyle: 'butt',
};

const getTemperatureGraphData = state => state.temperatureGraphReducer;

// get x-axis data, which will be array with time values in hh:mm:ss format
const getXaxisData = createSelector(
	temperatureGraphData => temperatureGraphData.get('time'),
	(timeData) => {
		const timeDataJS = timeData.toJS();
		const xAxisData = timeDataJS.map(ele => formatTime(ele));
		return xAxisData;
	},
);

const getYaxisData = createSelector(
	temperatureGraphData => temperatureGraphData,
	(temperatureGraphData) => temperatureGraphData.map(ele => {
		const start_time = new Date(temperatureGraphData.first().get('created_at'));
		const curr_time = new Date(ele.get('created_at'));
		let time_diff = curr_time.getMinutes() - start_time.getMinutes();
		time_diff = time_diff >= 0 ? time_diff : time_diff + 60;
		return {
			x: time_diff,
			y: ele.get('temp'),
		};
	}),
);

// get chart data object for plotting temperature line chart
export const getTemperatureChartData = createSelector(
	getTemperatureGraphData,
	(temperatureGraphReducer) => {
		const temperatureGraphData = temperatureGraphReducer.get('temperatureData');
		// if no data present return empty object
		if (temperatureGraphData.size === 0) {
			return {};
		}
		const borderColor = 'rgba(148,147,147,1)'; // default line color
		const chartData = {
			datasets: [
				{
					...lineConfigs,
					label: 'Temperature Plot',
					borderColor,
					data: getYaxisData(temperatureGraphData).toJS(),
				},
			],
		};
		return chartData;
	},
);
