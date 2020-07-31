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
	holdTimeError: false,
	rampRateError: false,
	repeatCountError: false,
	targetTemperatureError: false,
	dataCapture: false,
	isCreateStepModalVisible: false,
});

export const stepStateReducer = (state, action) => {
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

export const repeatCounterStateActions = {
	SET_VALUES: 'SET_VALUES',
	RESET_VALUES: 'RESET_VALUES',
};

export const repeatCounterInitialState = fromJS({
	repeatCount: '',
	repeatCountError: false,
});

export const repeatCounterStateReducer = (state, action) => {
	switch (action.type) {
	case repeatCounterStateActions.SET_VALUES:
		return state.setIn([action.key], action.value);
	case repeatCounterStateActions.RESET_VALUES:
		return repeatCounterInitialState;
	default:
		throw new Error('Invalid action type');
	}
};
