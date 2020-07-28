import { fromJS } from 'immutable';
import {
	addStageActions,
	listStageActions,
	updateStageActions,
	deleteStageActions,
} from 'actions/stageActions';
import { createTemplateActions } from 'actions/templateActions';

const listStageInitialState = fromJS({
	isLoading: true,
	list: [],
	// selectedStageId: null,
});

const createStageInitialState = {
	data: {},
	isStageSaved: false,
};

const updateStageInitialState = {
	data: {},
	isStageUpdated: false,
};

const deleteStageInitialState = {
	data: {},
	isStageDeleted: false,
};

export const listStagesReducer = (
	state = listStageInitialState,
	action,
) => {
	switch (action.type) {
	case listStageActions.listAction:
		return state.setIn(['isLoading'], true);
	case listStageActions.successAction:
		return state.merge({ list: fromJS(action.payload.response || []), isLoading: false });
	case listStageActions.failureAction:
		return state.merge({
			error: fromJS(action.payload.error),
			isLoading: false,
		});
	// Add the stages on create template
	case createTemplateActions.successAction:
		return state.merge({ list : fromJS(action.payload.response.stages || []), isLoading: false });
	// Add the created stage to list to avoid an extra get api call
	case addStageActions.successAction:
		return state.updateIn(['list'], list => list.push(fromJS(action.payload.response || {})));
	// case listStageActions.setSelectedStageId:
	// 	return state.setIn(['selectedStageId'], action.payload.stageId);
	default:
		return state;
	}
};

export const createStageReducer = (
	state = createStageInitialState,
	action,
) => {
	switch (action.type) {
	case addStageActions.addAction:
		return { ...state, isLoading: true, isStageSaved: false };
	case addStageActions.successAction:
		return {
			...state, ...action.payload, isLoading: false, isStageSaved: true,
		};
	case addStageActions.failureAction:
		return {
			...state, ...action.payload, isLoading: false, isStageSaved: false,
		};
	case addStageActions.addStageReset:
		return createStageInitialState;
	default:
		return state;
	}
};

export const updateStageReducer = (
	state = updateStageInitialState,
	action,
) => {
	switch (action.type) {
	case updateStageActions.updateAction:
		return { ...state, isLoading: true, isStageUpdated: false };
	case updateStageActions.successAction:
		return {
			...state, ...action.payload, isLoading: false, isStageUpdated: true,
		};
	case updateStageActions.failureAction:
		return {
			...state, ...action.payload, isLoading: false, isStageUpdated: true,
		};
	case updateStageActions.updateStageReset:
		return updateStageInitialState;
	default:
		return state;
	}
};

export const deleteStageReducer = (
	state = deleteStageInitialState,
	action,
) => {
	switch (action.type) {
	case deleteStageActions.deleteAction:
		return { ...state, isLoading: true, isStageDeleted: false };
	case deleteStageActions.successAction:
		return {
			...state, ...action.payload, isLoading: false, isStageDeleted: true,
		};
	case deleteStageActions.failureAction:
		return {
			...state, ...action.payload, isLoading: false, isStageDeleted: true,
		};
	case deleteStageActions.deleteStageReset:
		return deleteStageInitialState;
	default:
		return state;
	}
};
