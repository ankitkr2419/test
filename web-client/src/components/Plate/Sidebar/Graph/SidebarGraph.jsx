import React, { useState } from 'react';
import Sidebar from 'components/Sidebar';
import { Text } from 'shared-components';
import styled from 'styled-components';
import GraphFilters from './GraphFilters';

const GraphCard = styled.div`
  width: 830px;
  height: 275px;
  background: #ffffff 0% 0% no-repeat padding-box;
  border: 1px solid #707070;
  padding: 8px;
  margin: 0 0 40px 0;
`;

const SidebarGraph = (props) => {
	const [isSidebarOpen, setSideBarState] = useState(false);
	const toggleSideBar = () => {
		setSideBarState(isOpen => !isOpen);
	};

	return (
		<Sidebar
			isOpen={isSidebarOpen}
			isClose
			toggleSideBar={toggleSideBar}
			className="graph"
			handleIcon="graph"
			handleIconSize={56}
		>
			<Text size={20} className="text-default mb-5">
        Amplification plot
			</Text>
			<GraphCard></GraphCard>
			<GraphFilters />
			<Text size={14} className="text-default mb-0">
        Note: Click on the threshold number to change it.
			</Text>
		</Sidebar>
	);
};

SidebarGraph.propTypes = {};

export default SidebarGraph;
