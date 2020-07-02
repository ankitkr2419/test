import { fromJS, List } from 'immutable';

// const action types
export const graphFilterActions = {
	UPDATE_GRAPH_FILTER_STATE: 'UPDATE_GRAPH_FILTER_VALUES',
	SET_GRAPH_FILTER_VALUES: 'SET_GRAPH_FILTER_VALUES',
	RESET_GRAPH_FILTER_STATE: 'RESET_GRAPH_FILTER_STATE',
	UPDATE_TARGET_LIST: 'UPDATE_TARGET_LIST',
};

export const graphFilterInitialState = fromJS({
	isSidebarOpen: false,
	targets:  List([]),
});

const graphFilterState = (state, action) => {
	switch (action.type) {
	case graphFilterActions.SET_GRAPH_FILTER_VALUES:
		return state.setIn([action.key], action.value);
	case graphFilterActions.UPDATE_GRAPH_FILTER_STATE:
		return state.merge(action.value);
	case graphFilterActions.UPDATE_TARGET_LIST:
		return state.setIn(['targets', action.index, 'threshold'], action.threshold);
	case graphFilterActions.RESET_GRAPH_FILTER_STATE:
		return graphFilterInitialState;
	default:
		throw new Error('Invalid action type');
	}
};

export default graphFilterState;
