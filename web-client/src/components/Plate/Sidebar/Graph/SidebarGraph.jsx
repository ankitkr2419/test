import React, { useState } from 'react';
import Sidebar from 'components/Sidebar';
import { LineChart } from 'core-components';
import { Text } from 'shared-components';
import styled from 'styled-components';
import GraphFilters from './GraphFilters';

const data = {
	labels: ['target1', 'target2', 'target3', 'target4', 'target5', 'target6'],
	datasets: [
		{
			label: 'first',
			fill: false,
			borderWidth:  2,
			pointRadius	: 0,
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
			data: [65, 59, 80, 81, 56, 55],
		},
		{
			label: 'threshold',
			fill: false,
			borderWidth:  2,
			pointRadius	: 0,
			borderColor: '#a2ee95',
			pointBorderColor: '#a2ee95',
			pointBackgroundColor: '#fff',
			pointBorderWidth: 0,
			pointHoverRadius: 0,
			pointHoverBackgroundColor: '#a2ee95',
			pointHoverBorderColor: '#a2ee95',
			pointHoverBorderWidth: 0,
			data: [80, 80, 80, 80, 80, 80],
		},
	],
};


const GraphCard = styled.div`
	width: 830px;
	height: 275px;
	background: #ffffff 0% 0% no-repeat padding-box;
	border: 1px solid #707070;
	padding: 8px;
	margin: 0 0 40px 0;
`;

const SidebarGraph = (props) => {
	const [isSidebarOpen, setSideBarState] = useState(true);
	const toggleSideBar = () => {
		setSideBarState(isOpen => !isOpen);
	};

	return (
		<Sidebar
			isOpen={isSidebarOpen}
			toggleSideBar={toggleSideBar}
			className='graph'
			handleIcon='graph'
			handleIconSize={56}
		>
			<Text size={20} className='text-default mb-5'>
				Amplification plot
			</Text>
			<GraphCard>
				<LineChart data={data}/>
			</GraphCard>
			<GraphFilters />
			<Text size={14} className='text-default mb-0'>
				Note: Click on the threshold number to change it.
			</Text>
		</Sidebar>
	);
};

SidebarGraph.propTypes = {};

export default SidebarGraph;
