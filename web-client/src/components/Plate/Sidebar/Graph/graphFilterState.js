import { fromJS } from 'immutable';

// const action types
export const graphFilterActions = {
	UPDATE_GRAPH_FILTER_STATE: 'UPDATE_GRAPH_FILTER_VALUES',
	SET_GRAPH_FILTER_VALUES: 'SET_GRAPH_FILTER_VALUES',
	RESET_GRAPH_FILTER_STATE: 'RESET_GRAPH_FILTER_STATE',
};

export const stepStateInitialState = fromJS({
	isGraphBarOpen: false,
	target1: '',
	target2: '',
	target3: '',
	target4: '',
	target5: '',
	target6: '',
});

const stepStateReducer = (state, action) => {
	switch (action.type) {
	case graphFilterActions.SET_GRAPH_FILTER_VALUES:
		return state.setIn([action.key], action.value);
	case graphFilterActions.UPDATE_GRAPH_FILTER_STATE:
		return state.merge(action.value);
	case graphFilterActions.RESET_GRAPH_FILTER_STATE:
		return stepStateInitialState;
	default:
		throw new Error('Invalid action type');
	}
};

export default stepStateReducer;
