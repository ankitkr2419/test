import React, { useEffect } from 'react';
import { w3cwebsocket as W3CWebSocket } from 'websocket';
import Sidebar from 'components/Sidebar';
import { LineChart } from 'core-components';
import { Text } from 'shared-components';
import styled from 'styled-components';
import GraphFilters from './GraphFilters';

let webSocket = null;

const SidebarGraph = (props) => {
	const {
		filterState,
		experimentId,
		isSidebarOpen,
		toggleSideBar,
		onThresholdChangeHandler,
	} = props;

	const socketOnOpen = () => {
		console.log('socketOnOpen:');
	};

	const socketOnMessage = (event) => {
		console.log('socketOnMessage:', event);
	};

	const socketOnClose = () => {
		console.log('socketOnClose:');
	};

	useEffect(() => {
		if (isSidebarOpen === true) {
			webSocket = new W3CWebSocket(`ws://localhost:33001/experiments/${experimentId}/monitor`);
			// webSocket = new W3CWebSocket('ws://localhost:8081/echo');
			webSocket.onopen = socketOnOpen;
			webSocket.onmessage = socketOnMessage;
			webSocket.onclose = socketOnClose;
		}
		return () => {
			if (webSocket) {
				webSocket.close();
			}
		};
	}, [isSidebarOpen, experimentId]);

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
			/>
			<Text size={14} className="text-default mb-0">
        Note: Click on the threshold number to change it.
			</Text>
		</Sidebar>
	);
};

const data = {
	labels: ['1', '2', '3', '4', '5', '6'],
	datasets: [
		{
			label: 'first',
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
			data: [0.65, 0.59, 0.80, 0.81, 0.56, 0.55],
		},
		{
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
			data: [0.80, 0.80, 0.80, 0.80, 0.80, 0.80],
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

SidebarGraph.propTypes = {};

export default SidebarGraph;
