import { fromJS } from 'immutable';
import wellGraphActions from 'actions/wellGraphActions';
import loginActions from 'actions/loginActions';
// import graphData from '../mock-json/graphData.json';

const wellGraphInitialState = fromJS({
	chartData: [],
	// chartData: graphData,
});

export const wellGraphReducer = (state = wellGraphInitialState, action) => {
	switch (action.type) {
	case wellGraphActions.successAction:
		return state.setIn(['chartData'], fromJS(action.payload.data));
	case loginActions.loginReset:
		return wellGraphInitialState;
	default:
		return state;
	}
};
