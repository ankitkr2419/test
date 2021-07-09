import React, { useState } from 'react';
import Sidebar from 'components/Sidebar';
import PropTypes from 'prop-types';
import { getXAxis } from 'selectors/wellGraphSelector';
import TemperatureGraphContainer from 'containers/TemperatureGraphContainer';
import { Switch } from 'core-components';
import { SwitchWrapper } from 'shared-components/SwitchWrapper';
import WellGraph from './WellGraph';

const SidebarGraph = (props) => {
	const {
		isExperimentRunning,
		lineChartData,
		isSidebarOpen,
		toggleSideBar,
		onThresholdChangeHandler,
		toggleGraphFilterActive,
		experimentGraphTargetsList,
		isExperimentSucceeded,
		setThresholdError,
		resetThresholdError,
		isThresholdInvalid,
	} = props;

	// local state to toggle between emission graph and temperature graph
	const [showTempGraph, setShowTempGraph] = useState(false);

	let cyclesCount = 0;
	// below case can happen if user selects all filter we might get empty chart data
	if (lineChartData.size !== 0) {
		cyclesCount = lineChartData.first().totalCycles;
	}

	const data = {
		labels: getXAxis(cyclesCount),
		datasets: lineChartData.toJS(),
	};

	// helper function to toggle the graphs
	const toggleTempGraphSwitch = () => {
		setShowTempGraph(value => !value);
	};

	if (isExperimentRunning === true || isExperimentSucceeded === true) {
		return (
			<>
				{/* show the well data graph if showTempGraph flag is off */}
				{!showTempGraph
					&& <WellGraph
						data={data}
						experimentGraphTargetsList={experimentGraphTargetsList}
						onThresholdChangeHandler={onThresholdChangeHandler}
						toggleGraphFilterActive={toggleGraphFilterActive}
						setThresholdError={setThresholdError}
						resetThresholdError={resetThresholdError}
						isThresholdInvalid={isThresholdInvalid}
					/>
				}
				{/* show temperature graph if showTempGraph flag is on */}
				{showTempGraph
					&& <TemperatureGraphContainer />
				}
			</>

			// {/** TODO: remove after tested */}
			// <Sidebar
			// 	isOpen={isSidebarOpen}
			// 	toggleSideBar={toggleSideBar}
			// 	className="graph"
			// 	bodyClassName="py-4"
			// 	handleIcon="graph"
			// 	handleIconSize={56}
			// >
			// 	<SwitchWrapper>
			// 		<Switch
			// 			id="temperature"
			// 			name="temperature"
			// 			label="Show temperature graph"
			// 			value={showTempGraph}
			// 			onChange={toggleTempGraphSwitch}
			// 		/>
			// 	</SwitchWrapper>
			// 	{/* show the well data graph if showTempGraph flag is off */}
			// 	{!showTempGraph
			// 		&& <WellGraph
			// 			data={data}
			// 			experimentGraphTargetsList={experimentGraphTargetsList}
			// 			onThresholdChangeHandler={onThresholdChangeHandler}
			// 			toggleGraphFilterActive={toggleGraphFilterActive}
			// 			setThresholdError={setThresholdError}
			// 			resetThresholdError={resetThresholdError}
			// 			isThresholdInvalid={isThresholdInvalid}
			// 		/>
			// 	}
			// 	{/* show temperature graph if showTempGraph flag is on */}
			// 	{showTempGraph
			// 		&& <TemperatureGraphContainer />
			// 	}
			// </Sidebar>
		);
	}
	return null;
};

SidebarGraph.propTypes = {
	isExperimentRunning: PropTypes.bool.isRequired,
	lineChartData: PropTypes.object.isRequired,
	isSidebarOpen: PropTypes.bool.isRequired,
	toggleSideBar: PropTypes.func.isRequired,
	onThresholdChangeHandler: PropTypes.func.isRequired,
	toggleGraphFilterActive: PropTypes.func.isRequired,
	experimentGraphTargetsList: PropTypes.object.isRequired,
	isExperimentSucceeded: PropTypes.bool.isRequired,
};

export default React.memo(SidebarGraph);
