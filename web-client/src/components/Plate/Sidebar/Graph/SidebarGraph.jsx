import React from 'react';
import Sidebar from 'components/Sidebar';
import PropTypes from 'prop-types';
import { LineChart } from 'core-components';
import { Text } from 'shared-components';
import styled from 'styled-components';
import { getXAxis } from 'selectors/wellGraphSelector';
import GraphFilters from './GraphFilters';

const SidebarGraph = (props) => {
	const {
		isExperimentRunning,
		lineChartData,
		isSidebarOpen,
		toggleSideBar,
		onThresholdChangeHandler,
		toggleGraphFilterActive,
		experimentGraphTargetsList,
	} = props;

	let noOfCycles = 0;
	// below case can happen if user selects all filter we might get empty chart data
	if (lineChartData.size !== 0) {
		noOfCycles = lineChartData.first().totalCycles;
	}
	const data = {
		labels: getXAxis(noOfCycles),
		datasets: lineChartData.toJS(),
	};

	if (isExperimentRunning === true) {
		return (
			<Sidebar
				isOpen={isSidebarOpen}
				toggleSideBar={toggleSideBar}
				className='graph'
				bodyClassName='py-4'
				handleIcon='graph'
				handleIconSize={56}
			>
				<Text size={20} className='text-default mb-4'>
					Amplification plot
				</Text>
				<GraphCard>
					<LineChart data={data} />
				</GraphCard>
				<GraphFilters
					targets={experimentGraphTargetsList}
					onThresholdChangeHandler={onThresholdChangeHandler}
					toggleGraphFilterActive={toggleGraphFilterActive}
				/>
				<Text size={14} className='text-default mb-0'>
					Note: Click on the threshold number to change it.
				</Text>
			</Sidebar>
		);
	}
	return null;
};

const GraphCard = styled.div`
	width: 830px;
	height: 344px;
	background: #ffffff 0% 0% no-repeat padding-box;
	border: 1px solid #707070;
	padding: 8px;
	margin: 0 0 32px 0;
`;

SidebarGraph.propTypes = {
	isExperimentRunning: PropTypes.bool.isRequired,
	lineChartData: PropTypes.object.isRequired,
	isSidebarOpen: PropTypes.bool.isRequired,
	toggleSideBar: PropTypes.func.isRequired,
	onThresholdChangeHandler: PropTypes.func.isRequired,
	toggleGraphFilterActive: PropTypes.func.isRequired,
	experimentGraphTargetsList: PropTypes.object.isRequired,
};

export default React.memo(SidebarGraph);
