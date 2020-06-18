import { fromJS } from 'immutable';

// const action types
export const stepStateActions = {
	UPDATE_STATE: 'UPDATE_STATE',
	SET_VALUES: 'SET_VALUES',
	RESET_VALUES: 'RESET_VALUES',
};

export const stepStateInitialState = fromJS({
	stepId: null,
	rampRate: '',
	targetTemperature: '',
	holdTime: '',
	dataCapture: '',
	isCreateStepModalVisible: false,
});

const stepStateReducer = (state, action) => {
	switch (action.type) {
	case stepStateActions.SET_VALUES:
		return state.setIn([action.key], action.value);
	case stepStateActions.UPDATE_STATE:
		return state.merge(action.value);
	case stepStateActions.RESET_VALUES:
		return stepStateInitialState;
	default:
		throw new Error('Invalid action type');
	}
};

export default stepStateReducer;
