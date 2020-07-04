import React, { useEffect, useReducer } from 'react';
import PropTypes from 'prop-types';
import SidebarGraph from 'components/Plate/Sidebar/Graph/SidebarGraph';
import graphFilterState, {
	graphFilterInitialState,
	graphFilterActions,
} from 'components/Plate/Sidebar/Graph/graphFilterState';
import { useSelector } from 'react-redux';
import { getRunExperimentReducer } from 'selectors/runExperimentSelector';

const ExperimentGraphContainer = (props) => {
	const { experimentTargetsList } = props;

	const wellGraphReducer = useSelector(state => state.wellGraphReducer);
	const runExperimentReducer = useSelector(getRunExperimentReducer);
	const isExperimentRunning =    runExperimentReducer.get('experimentStatus') === 'running';

	// local state to save filter graph data
	const [filterState, updateGraphFilterState] = useReducer(
		graphFilterState,
		graphFilterInitialState,
	);
	const { isSidebarOpen } = filterState.toJS();
	// helper function to update local state
	const updateGraphFilterStateWrapper = (key, value) => {
		updateGraphFilterState({
			type: graphFilterActions.SET_GRAPH_FILTER_VALUES,
			key,
			value,
		});
	};

	useEffect(() => {
		if (experimentTargetsList !== null && experimentTargetsList.size !== 0) {
			updateGraphFilterStateWrapper('targets', experimentTargetsList);
		}
	}, [experimentTargetsList]);

	const toggleSideBar = () => {
		updateGraphFilterStateWrapper('isSidebarOpen', !isSidebarOpen);
	};

	const onThresholdChangeHandler = (threshold, index) => {
		updateGraphFilterState({
			type: graphFilterActions.UPDATE_TARGET_LIST,
			index,
			threshold: parseFloat(threshold),
		});
	};

	const toggleGraphFilterActive = (index, isActive) => {
		updateGraphFilterState({
			type: graphFilterActions.UPDATE_GRAPH_FILTER_ACTIVE_STATE,
			index,
			isActive: !isActive,
		});
	};

	return (
		<SidebarGraph
			isExperimentRunning={isExperimentRunning}
			wellGraphReducer={wellGraphReducer}
			experimentTargetsList={experimentTargetsList}
			isSidebarOpen={isSidebarOpen}
			toggleSideBar={toggleSideBar}
			filterState={filterState}
			onThresholdChangeHandler={onThresholdChangeHandler}
			toggleGraphFilterActive={toggleGraphFilterActive}
		/>
	);
};

ExperimentGraphContainer.propTypes = {
	experimentTargetsList: PropTypes.object.isRequired,
};

export { ExperimentGraphContainer };
