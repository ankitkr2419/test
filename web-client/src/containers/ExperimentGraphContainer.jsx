import React, { useState } from 'react';
import SidebarGraph from 'components/Plate/Sidebar/Graph/SidebarGraph';
import { useSelector, useDispatch } from 'react-redux';
import { getExperimentStatus } from 'selectors/runExperimentSelector';
import { getLineChartData } from 'selectors/wellGraphSelector';
import { getExperimentGraphTargets } from 'selectors/experimentTargetSelector';
import { updateExperimentTargetFilters } from 'action-creators/experimentTargetActionCreators';

const ExperimentGraphContainer = () => {
	const dispatch = useDispatch();
	// getExperimentStatus will return us current experiment status
	const experimentStatus = useSelector(getExperimentStatus);

	// get targets from experiment target reducer(graph : target filters)
	const experimentGraphTargetsList = useSelector(getExperimentGraphTargets);

	// Extracting graph data, Which is populated from websocket
	const lineChartData = useSelector(getLineChartData);

	const isExperimentRunning = experimentStatus === 'running';

	// local state to save filter graph data
	const [isSidebarOpen, setIsSidebarOpen] = useState(false);

	const toggleSideBar = () => {
		setIsSidebarOpen(toggleStateValue => !toggleStateValue);
	};

	const onThresholdChangeHandler = (threshold, index) => {
		dispatch(updateExperimentTargetFilters(index, 'threshold', parseFloat(threshold)));
	};

	const toggleGraphFilterActive = (index, isActive) => {
		dispatch(updateExperimentTargetFilters(index, 'isActive', !isActive));
	};

	return (
		<SidebarGraph
			isExperimentRunning={isExperimentRunning}
			lineChartData={lineChartData}
			experimentGraphTargetsList={experimentGraphTargetsList}
			isSidebarOpen={isSidebarOpen}
			toggleSideBar={toggleSideBar}
			onThresholdChangeHandler={onThresholdChangeHandler}
			toggleGraphFilterActive={toggleGraphFilterActive}
		/>
	);
};

ExperimentGraphContainer.propTypes = {};

export { ExperimentGraphContainer };
