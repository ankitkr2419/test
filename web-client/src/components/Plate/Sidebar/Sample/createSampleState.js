import { fromJS, List } from 'immutable';

// const action types
export const createSampleActions = {
	UPDATE_STATE: 'UPDATE_SAMPLE_STATE',
	SET_VALUES: 'UPDATE_SAMPLE_VALUES',
	RESET_VALUES: 'RESET_SAMPLE_VALUES',
};

export const createSampleInitialState = fromJS({
	isSideBarOpen: false,
	sample: null,
	targets: List([]),
	task: null,
	isEdit: false,
	position: -1,
});

export const validate = (state) => {
	const { sample, targets, task } = state.toJS();
	if (sample === undefined || targets.length === 0 || task === null) {
		return false;
	}
	return true;
};

export const getSampleRequestData = (state, positions) => {
	const {
		sample, targets, task, isEdit, position,
	} = state.toJS();
	const requestObject = {};
	if (sample.label === sample.value) {
		requestObject.sample = {
			name: sample.label,
		};
	} else {
		requestObject.sample = {
			id: sample.value,
			name: sample.label,
		};
	}

	requestObject.targets = targets.map(ele => ele.target_id);
	requestObject.task = task.value;
	requestObject.position = isEdit === true ? [position] : positions;


	return requestObject;
};

const createSampleStateReducer = (state, action) => {
	switch (action.type) {
	case createSampleActions.SET_VALUES:
		return state.setIn([action.key], action.value);
	case createSampleActions.UPDATE_STATE:
		return state.merge(action.value);
	case createSampleActions.deleteTarget:
		return state.updateIn(['targets'], targetList => targetList.filter(ele => ele.get('target_id') !== action.value));
	case createSampleActions.RESET_VALUES:
		if (state.get('isEdit') === true) {
			return createSampleInitialState.updateIn(['targets'], () => state.get('targets'));
		}
		return createSampleInitialState;
	default:
		throw new Error('Invalid action type');
	}
};

export default createSampleStateReducer;
