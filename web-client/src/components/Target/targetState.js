import { fromJS, List } from 'immutable';
import { getTargetOption } from './targetHelper';

// const action types
export const targetStateActions = {
	ADD_TARGET_ID: 'ADD_TARGET_ID',
	ADD_THRESHOLD_VALUE: 'ADD_THRESHOLD_VALUE',
	SET_CHECKED_STATE: 'SET_CHECKED_STATE',
	UPDATE_LIST: 'UPDATE_LIST',
};

// Initial state wrap with fromJS for immutability
export const targetInitialState = fromJS({
	// getTargetOption will return fixed size array with object initialize with isChecked false
	targetList: List(getTargetOption()),
	originalTargetList: List([]),
});

// isCheckable will validate weather target and threshold value is present for given index
export const isCheckable = (state, index) => {
	const targetList = state.get('targetList');
	if (targetList.get(index) === undefined) {
		return false;
	}
	const ele = targetList.get(index);
	if (ele.selectedTarget === undefined || ele.threshold === undefined) {
		return false;
	}
	return true;
};

// function will set selectedTarget flag w.r.t index
const addTargetId = (state, { targetId, index }) => state.setIn(['targetList', index, 'selectedTarget'], targetId);

// function will set threshold flag w.r.t index
const addThresholdValue = (state, { threshold, index }) => state.setIn(['targetList', index, 'threshold'], threshold);

// function will set isChecked flag w.r.t index
const setCheckedState = (state, { checked, index }) => state.setIn(['targetList', index, 'isChecked'], checked);

export const getCheckedTargets = (list, templateID) => {
	const targetList = list.filter(ele => ele.isChecked === true).toJSON();
	// if we don't found any items selected will return empty list
	// Used to set save button disabled on target page
	if (targetList.length === 0) {
		return targetList;
	}
	// below object manipulation is done for sending data to server
	return targetList.map(ele => ({
		template_id: templateID,
		target_id: ele.selectedTarget.value,
		threshold: parseFloat(ele.threshold),
	}));
};

export const getCheckedExperimentTargets = (list) => {
	const targetList = list.filter(ele => ele.isChecked === true).toJSON();
	// if we don't found any items selected will return empty list
	// Used to set save button disabled on target page
	if (targetList.length === 0) {
		return targetList;
	}
	// below object manipulation is done for sending data to server
	return targetList.map(ele => ({
		experiment_id: ele.experiment_id,
		template_id: ele.template_id,
		target_id: ele.selectedTarget.value,
		threshold: parseFloat(ele.threshold),
	}));
};

// function will replace the existing local target list with new
const updateList = (state, list) => state.merge({ targetList: List(list), originalTargetList: List(list) });

export const isTargetListUpdated = (state) => {
	const { targetList, originalTargetList } = state.toJS();
	return JSON.stringify(targetList) !== JSON.stringify(originalTargetList);
};

const targetStateReducer = (state, action) => {
	switch (action.type) {
	case targetStateActions.ADD_TARGET_ID:
		return addTargetId(state, action.value);
	case targetStateActions.ADD_THRESHOLD_VALUE:
		return addThresholdValue(state, action.value);
	case targetStateActions.SET_CHECKED_STATE:
		return setCheckedState(state, action.value);
	case targetStateActions.UPDATE_LIST:
		return updateList(state, action.value);
	default:
		throw new Error('Invalid action type');
	}
};

export default targetStateReducer;
