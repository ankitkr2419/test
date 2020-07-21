import React from 'react';
import PropTypes from 'prop-types';
import SidebarGraph from 'components/Plate/Sidebar/Graph/SidebarGraph';
import { useSelector, useDispatch } from 'react-redux';
import { getLineChartData } from 'selectors/wellGraphSelector';
import { getExperimentGraphTargets } from 'selectors/experimentTargetSelector';
import { updateExperimentTargetFilters } from 'action-creators/experimentTargetActionCreators';
import { EXPERIMENT_STATUS } from 'appConstants';
import { parseFloatWrapper } from 'utils/helpers';
import { isAnyThresholdInvalid } from 'components/Target/targetHelper';

const ExperimentGraphContainer = (props) => {
	const {
		setIsSidebarOpen,
		isSidebarOpen,
		experimentStatus,
		isMultiSelectionOptionOn,
		resetSelectedWells,
	} = props;
	const dispatch = useDispatch();
	// get targets from experiment target reducer(graph : target filters)
	const experimentGraphTargetsList = useSelector(getExperimentGraphTargets);

	// Extracting graph data, Which is populated from websocket
	const lineChartData = useSelector(getLineChartData);

	const isExperimentRunning = experimentStatus === EXPERIMENT_STATUS.running;
	const isExperimentSucceeded = experimentStatus === EXPERIMENT_STATUS.success;

	const toggleSideBar = () => {
		// reset the selected wells while closing the sidebar
		if (isSidebarOpen && isMultiSelectionOptionOn === false) {
			resetSelectedWells();
		}
		setIsSidebarOpen(toggleStateValue => !toggleStateValue);
	};

	const onThresholdChangeHandler = (threshold, index) => {
		dispatch(updateExperimentTargetFilters(index, 'threshold', parseFloatWrapper(threshold)));
	};

	const toggleGraphFilterActive = (index, isActive) => {
		dispatch(updateExperimentTargetFilters(index, 'isActive', !isActive));
	};

	// set threshold error true
	const setThresholdError = (index) => {
		dispatch(updateExperimentTargetFilters(index, 'thresholdError', true));
	};

	// reset threshold error to false
	const resetThresholdError = (index) => {
		dispatch(updateExperimentTargetFilters(index, 'thresholdError', false));
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
			setThresholdError={setThresholdError}
			resetThresholdError={resetThresholdError}
			isThresholdInvalid={isAnyThresholdInvalid(experimentGraphTargetsList)}
		/>
	);
};

ExperimentGraphContainer.propTypes = {
	experimentStatus: PropTypes.string,
};

export { ExperimentGraphContainer };
