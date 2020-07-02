import React, { useEffect, useReducer } from 'react';
import PropTypes from 'prop-types';
import SidebarGraph from 'components/Plate/Sidebar/Graph/SidebarGraph';
import graphFilterState, {
	graphFilterInitialState,
	graphFilterActions,
} from 'components/Plate/Sidebar/Graph/graphFilterState';

const ExperimentGraphContainer = (props) => {
	const { experimentTargetsList, experimentId } = props;

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

	return (
		<SidebarGraph
			experimentId={experimentId}
			experimentTargetsList={experimentTargetsList}
			isSidebarOpen={isSidebarOpen}
			toggleSideBar={toggleSideBar}
			filterState={filterState}
			onThresholdChangeHandler={onThresholdChangeHandler}
		/>
	);
};

ExperimentGraphContainer.propTypes = {
	experimentTargetsList: PropTypes.object.isRequired,
};

export { ExperimentGraphContainer };
