import { createSelector } from 'reselect';
import { getExperimentTargets } from './experimentTargetSelector';
import { getWells, getWellsPosition } from './wellSelectors';

// default line configs
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

const getWellGraphData = state => state.wellGraphReducer;
const nullableCheck = arr => arr.every(item => item === 0);
const getThresholdLineData = (value, count) => {
	const arr = [];
	for (let x = 0; x <= count; x += 1) {
		arr.push(value);
	}
	return arr;
};

// get xaxis labels
export const getXAxis = createSelector(
	count => count,
	(count) => {
		const arr = [];
		for (let x = 0; x <= count; x += 1) {
			arr.push(x);
		}
		return arr;
	},
);

const getThresholdLine = (label, max_threshold, count, lineColor) => ({
	label,
	fill: false,
	borderWidth: 2,
	pointRadius: 0,
	borderColor: lineColor,
	pointBorderColor: '#a2ee95',
	borderDash: [10, 5],
	pointBackgroundColor: '#fff',
	pointBorderWidth: 0,
	pointHoverRadius: 0,
	pointHoverBackgroundColor: '#a2ee95',
	pointHoverBorderColor: '#a2ee95',
	pointHoverBorderWidth: 0,
	data: getThresholdLineData(max_threshold, count),
});

/**
 * getLineChartData will return chart data
 * getLineChartData is listening to below reducers
 * 	1> wellGraphReducer (socket data is populated)
 * 	2> listExperimentTargetsReducer ( to fetch experiment targets (filtered targets))
 *  3> wellListReducer ( to get selected wells and isMultiSelectionOptionOn flag)
 */
export const getLineChartData = createSelector(
	getWellGraphData,
	getExperimentTargets,
	getWells,
	(wellGraphReducer, experimentTargets, wells) => {
		// graphTargets contains updated graph target values(Filters)
		const graphTargets = experimentTargets.get('graphTargets');
		const selectedPositions = getWellsPosition(wells);
		let wellGraphData = wellGraphReducer.get('chartData');
		// Should apply filter if we have positions selected from viewing graph
		if (selectedPositions.size !== 0) {
			wellGraphData = wellGraphData.filter(ele => selectedPositions.includes(ele.get('well_position')));
		}
		// filtering active wells
		const activeTargets = graphTargets.filter(ele => ele.get('isActive') === true);
		// by default all targets are active so checking if any target is not active
		if (activeTargets.size !== graphTargets.size) {
			// getting active target ids
			const activeTargetsIds = activeTargets.map(ele => ele.get('target_id'));
			// filtering graph data w.r.t active target ids
			wellGraphData = wellGraphData.filter(ele => activeTargetsIds.includes(ele.get('target_id')));
		}

		const chartData = wellGraphData
			.map((ele, index) => {
				if (nullableCheck(ele.get('f_value')) === false) {
					// getting color of target
					const lineColorFiltered = graphTargets.filter(
						target => target.get('target_id') === ele.get('target_id'),
					);

					let borderColor = 'rgba(148,147,147,1)'; // default line color
					if (lineColorFiltered && lineColorFiltered.size !== 0) {
						// if line color is present assign it
						borderColor = lineColorFiltered.first().get('lineColor');
					}
					return {
						...lineConfigs, // chart line config
						label: `index-${index}`, // unique label per line
						borderColor, // line color
						data: ele.get('f_value'), // line data
						totalCycles: ele.get('total_cycles'), // cycle count to x-axis
						cycles: ele.get('cycle'), // cycles array with intermediate steps
					};
				}
				return null;
			})
			.filter(ele => ele !== null);

		// if we don't have data then no need to calculate threshold values
		if (chartData.size === 0) {
			return chartData;
		}

		// populating threshold line per targets
		const { totalCycles } = chartData.first();
		const thresholdArray = activeTargets.map(ele => getThresholdLine(
			ele.get('target_name'),
			ele.get('threshold'),
			totalCycles,
			ele.get('lineColor'),
		));
		return chartData.merge(thresholdArray);
	},
);
