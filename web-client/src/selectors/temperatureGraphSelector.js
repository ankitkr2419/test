import { createSelector } from 'reselect';
import { getTimeDiff } from 'utils/helpers';

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

const getTemperatureGraphReducer = state => state.temperatureGraphReducer;

// sort the temperature graph reducer wrt created_at property
const getSortedTemperatureGraphReducer = createSelector(
	getTemperatureGraphReducer,
	temperatureGraphReducer => temperatureGraphReducer.updateIn(['temperatureData'],
		myList => myList?.sortBy(ele => ele.get('created_at'))),
);

// get starting time of temperature graph
const getStartTime = temperatureGraphData => new Date(temperatureGraphData.first().get('created_at'));

const getYaxisDataForTemp = createSelector(
	temperatureGraphData => temperatureGraphData,
	getStartTime,
	(temperatureGraphData, startTime) => temperatureGraphData.map(ele => {
		const currTime = new Date(ele.get('created_at'));
		return {
			x: getTimeDiff(startTime, currTime),
			y: ele.get('temp'),
		};
	}),
);

const getYaxisDataForLidTemp = createSelector(
	temperatureGraphData => temperatureGraphData,
	getStartTime,
	(temperatureGraphData, startTime) => temperatureGraphData.map(ele => {
		const currTime = new Date(ele.get('created_at'));
		return {
			x: getTimeDiff(startTime, currTime),
			y: ele.get('lid_temp'),
		};
	}),
);

// get chart data object for plotting temperature line chart
export const getTemperatureChartData = createSelector(
	getSortedTemperatureGraphReducer,
	(temperatureGraphReducer) => {
		const temperatureGraphData = temperatureGraphReducer.get('temperatureData');
		// if no data present return empty object
		if (temperatureGraphData === null || temperatureGraphData === undefined || temperatureGraphData?.size === 0) {
			return {};
		}
		const borderColor = 'rgba(148,147,147,1)'; // default line color
		const chartData = {
			datasets: [
				{
					...lineConfigs,
					label: 'Temperature Plot',
					borderColor,
					data: getYaxisDataForTemp(temperatureGraphData).toJS(),
				},
				{
					...lineConfigs,
					label: 'Lid Temperature Plot',
					borderColor: `rgba(245,144,178,1)`,
					data: getYaxisDataForLidTemp(temperatureGraphData).toJS(),
				},
			],
		};
		return chartData;
	},
);
