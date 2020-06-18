import { fromJS } from 'immutable';
import {
	addStepActions,
	listStepActions,
	updateStepActions,
	deleteStepActions,
} from 'actions/stepActions';

const listStepInitialState = fromJS({
	isLoading: true,
	list: [],
});

const createStepInitialState = {
	data: {},
	isStepSaved: false,
};

const updateStepInitialState = {
	data: {},
	isStepUpdated: false,
};

const deleteStepInitialState = {
	data: {},
	isStepDeleted: false,
};

export const listStepsReducer = (
	state = listStepInitialState,
	action,
) => {
	switch (action.type) {
	case listStepActions.listAction:
		return state.setIn(['isLoading'], true);
	case listStepActions.successAction:
		return state.merge({ list: fromJS(action.payload.response || []), isLoading: false });
	case listStepActions.failureAction:
		return state.merge({
			error: fromJS(action.payload.error),
			isLoading: false,
		});
	default:
		return state;
	}
};

export const createStepReducer = (
	state = createStepInitialState,
	action,
) => {
	switch (action.type) {
	case addStepActions.addAction:
		return { ...state, isLoading: true, isStepSaved: false };
	case addStepActions.successAction:
		return {
			...state, ...action.payload, isLoading: false, isStepSaved: true,
		};
	case addStepActions.failureAction:
		return {
			...state, ...action.payload, isLoading: false, isStepSaved: false,
		};
	case addStepActions.addStepReset:
		return createStepInitialState;
	default:
		return state;
	}
};

export const updateStepReducer = (
	state = updateStepInitialState,
	action,
) => {
	switch (action.type) {
	case updateStepActions.updateAction:
		return { ...state, isLoading: true, isStepUpdated: false };
	case updateStepActions.successAction:
		return {
			...state, ...action.payload, isLoading: false, isStepUpdated: true,
		};
	case updateStepActions.failureAction:
		return {
			...state, ...action.payload, isLoading: false, isStepUpdated: true,
		};
	case updateStepActions.updateStepReset:
		return updateStepInitialState;
	default:
		return state;
	}
};

export const deleteStepReducer = (
	state = deleteStepInitialState,
	action,
) => {
	switch (action.type) {
	case deleteStepActions.deleteAction:
		return { ...state, isLoading: true, isStepDeleted: false };
	case deleteStepActions.successAction:
		return { ...state, ...action.payload, isLoading: false, isStepDeleted: true };
	case deleteStepActions.failureAction:
		return { ...state, ...action.payload, isLoading: false, isStepDeleted: true };
	case deleteStepActions.deleteStepReset:
		return deleteStepInitialState;
	default:
		return state;
	}
};
