import { fromJS } from 'immutable';

// const action types
export const stageStateActions = {
	UPDATE_STAGE_STATE: 'UPDATE_STAGE_STATE',
	SET_STAGE_VALUES: 'SET_STAGE_VALUES',
	RESET_STAGE_VALUES: 'RESET_STAGE_VALUES',
};

export const stageStateInitialState = fromJS({
	stageType: '',
	stageRepeatCount: '',
	stageId: null,
	isCreateStageModalVisible: false,
});

const stageStateReducer = (state, action) => {
	switch (action.type) {
	case stageStateActions.SET_STAGE_VALUES:
		return state.setIn([action.key], action.value);
	case stageStateActions.UPDATE_STAGE_STATE:
		return state.merge(action.value);
	case stageStateActions.RESET_STAGE_VALUES:
		return stageStateInitialState;
	default:
		throw new Error('Invalid action type');
	}
};

export default stageStateReducer;
