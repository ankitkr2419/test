import React from 'react';
import Sidebar from 'components/Sidebar';
import { LineChart } from 'core-components';
import { Text } from 'shared-components';
import styled from 'styled-components';
import GraphFilters from './GraphFilters';

const getXAxis = (count) => {
	const arr = [];
	for (let x = 1; x <= count; x += 1) {
		arr.push(x);
	}
	return arr;
};

const getThresholdLineData = (value, count) => {
	const arr = [];
	for (let x = 1; x <= count; x += 1) {
		arr.push(value);
	}
	return arr;
};

const nullableCheck = arr => arr.every(item => item === 0);

const getThresholdLine = (max_threshold, count) => ({
	label: 'threshold',
	fill: false,
	borderWidth: 2,
	pointRadius: 0,
	borderColor: '#a2ee95',
	pointBorderColor: '#a2ee95',
	pointBackgroundColor: '#fff',
	pointBorderWidth: 0,
	pointHoverRadius: 0,
	pointHoverBackgroundColor: '#a2ee95',
	pointHoverBorderColor: '#a2ee95',
	pointHoverBorderWidth: 0,
	data: getThresholdLineData(max_threshold, count),
});

const SidebarGraph = (props) => {
	const {
		isExperimentRunning,
		wellGraphReducer,
		filterState,
		isSidebarOpen,
		toggleSideBar,
		onThresholdChangeHandler,
		toggleGraphFilterActive,
	} = props;

	const { data: tableData = [], max_threshold = 0 } = wellGraphReducer
		.get('data')
		.toJS();
	const { total_cycles: totalCycles = 0 } =    tableData.length !== 0 && tableData[0];

	const datasets = tableData
		.map((ele, index) => {
			if (nullableCheck(ele.f_value) === false) {
				return {
					label: `index-${index}`,
					fill: false,
					borderWidth: 2,
					pointRadius: 0,
					borderColor: 'rgba(148,147,147,1)',
					pointBorderColor: 'rgba(148,147,147,1)',
					pointBackgroundColor: '#fff',
					pointBorderWidth: 0,
					pointHoverRadius: 0,
					pointHoverBackgroundColor: 'rgba(148,147,147,1)',
					pointHoverBorderColor: 'rgba(148,147,147,1)',
					pointHoverBorderWidth: 0,
					lineTension: 0.1,
					borderCapStyle: 'butt',
					data: ele.f_value,
				};
			}
			return null;
		})
		.filter(ele => ele !== null);
	datasets.push(getThresholdLine(max_threshold, totalCycles));

	const data = {
		labels: getXAxis(totalCycles),
		datasets,
	};

	if (isExperimentRunning === true) {
		return (
			<Sidebar
				isOpen={isSidebarOpen}
				toggleSideBar={toggleSideBar}
				className="graph"
				handleIcon="graph"
				handleIconSize={56}
			>
				<Text size={20} className="text-default mb-5">
          Amplification plot
				</Text>
				<GraphCard>
					<LineChart data={data} />
				</GraphCard>
				<GraphFilters
					targets={filterState.get('targets')}
					onThresholdChangeHandler={onThresholdChangeHandler}
					toggleGraphFilterActive={toggleGraphFilterActive}
				/>
				<Text size={14} className="text-default mb-0">
          Note: Click on the threshold number to change it.
				</Text>
			</Sidebar>
		);
	}
	return null;
};

const GraphCard = styled.div`
  width: 830px;
  height: 275px;
  background: #ffffff 0% 0% no-repeat padding-box;
  border: 1px solid #707070;
  padding: 8px;
  margin: 0 0 40px 0;
`;

SidebarGraph.propTypes = {};

export default SidebarGraph;
