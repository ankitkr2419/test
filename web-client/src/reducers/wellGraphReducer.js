import { fromJS } from 'immutable';
import webSocketActions from 'actions/webSocketActions';
import loginActions from 'actions/loginActions';
// import graphData from '../mock-json/graphData.json';

const wellGraphInitialState = fromJS({
	isOpened: false,
	isClosed: false,
	isError: false,
	chartData: fromJS({
		data: [],
		max_threshold: 0,
	}),
	// chartData: fromJS(graphData),
});

export const wellGraphReducer = (state = wellGraphInitialState, action) => {
	switch (action.type) {
	case webSocketActions.onOpen:
		return state.setIn(['isOpened'], true);
	case webSocketActions.onClose:
		return state.setIn(['isClosed'], true);
	case webSocketActions.onError:
		return state.setIn(['isError'], true);
	case webSocketActions.onMessage:
		return state.setIn(['chartData'], fromJS(action.payload.data));
	case loginActions.loginReset:
		return wellGraphInitialState;
	default:
		return state;
	}
};
