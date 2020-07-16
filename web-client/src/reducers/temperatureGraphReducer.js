import temperatureGraphActions from 'actions/temperatureGraphActions';
import loginActions from 'actions/loginActions';

const { fromJS } = require('immutable');

const temperatureGraphInitialState = fromJS({
	temperatureData: [],
});

export const temperatureGraphReducer = (state = temperatureGraphInitialState, action) => {
	switch (action.type) {
	case temperatureGraphActions.successAction:
		return state.setIn(['temperatureData'], fromJS(action.payload.data));
	case loginActions.loginReset:
		return temperatureGraphInitialState;
	default:
		return state;
	}
};
