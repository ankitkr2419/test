import React, { useState } from 'react';
import PropTypes from 'prop-types';
import SidebarGraph from 'components/Plate/Sidebar/Graph/SidebarGraph';
import { useSelector, useDispatch } from 'react-redux';
import { getLineChartData } from 'selectors/wellGraphSelector';
import { getExperimentGraphTargets } from 'selectors/experimentTargetSelector';
import { updateExperimentTargetFilters } from 'action-creators/experimentTargetActionCreators';
import { EXPERIMENT_STATUS } from 'appConstants';

const ExperimentGraphContainer = ({ experimentStatus }) => {
	const dispatch = useDispatch();
	// get targets from experiment target reducer(graph : target filters)
	const experimentGraphTargetsList = useSelector(getExperimentGraphTargets);

	// Extracting graph data, Which is populated from websocket
	const lineChartData = useSelector(getLineChartData);

	const isExperimentRunning = experimentStatus === EXPERIMENT_STATUS.running;
	const isExperimentSucceeded = experimentStatus === EXPERIMENT_STATUS.success;

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
			isExperimentSucceeded={isExperimentSucceeded}
			lineChartData={lineChartData}
			experimentGraphTargetsList={experimentGraphTargetsList}
			isSidebarOpen={isSidebarOpen}
			toggleSideBar={toggleSideBar}
			onThresholdChangeHandler={onThresholdChangeHandler}
			toggleGraphFilterActive={toggleGraphFilterActive}
		/>
	);
};

ExperimentGraphContainer.propTypes = {
	experimentStatus: PropTypes.string,
};

export { ExperimentGraphContainer };
