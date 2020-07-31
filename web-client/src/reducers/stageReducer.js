import { fromJS } from 'immutable';
import {
	listStageActions,
	updateStageActions,
} from 'actions/stageActions';
import { createTemplateActions } from 'actions/templateActions';

const listStageInitialState = fromJS({
	isLoading: true,
	list: [],
	// selectedStageId: null,
});

const updateStageInitialState = {
	data: {},
	isStageUpdated: false,
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
