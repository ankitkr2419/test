import { fromJS } from 'immutable';
import {
	createTemplateActions,
	listTemplateActions,
	updateTemplateActions,
	deleteTemplateActions,
} from 'actions/templateActions';

const listTemplateInitialState = fromJS({
	isLoading: true,
	list: [],
});

const createTemplateInitialState = {
	data: {},
};

const updateTemplateInitialState = {
	data: {},
};

const deleteTemplateInitialState = {
	data: {},
};

export const listTemplatesReducer = (
	state = listTemplateInitialState,
	action,
) => {
	switch (action.type) {
	case listTemplateActions.listAction:
		return state.setIn(['isLoading'], true);
	case listTemplateActions.successAction:
		return state.merge({ list: fromJS(action.payload.response), isLoading: false });
	case listTemplateActions.failureAction:
		return state.merge({
			error: fromJS(action.payload.error),
			isLoading: false,
		});
	default:
		return state;
	}
};

export const createTemplateReducer = (
	state = createTemplateInitialState,
	action,
) => {
	switch (action.type) {
	case createTemplateActions.createAction:
		return { ...state, isLoading: true };
	case createTemplateActions.successAction:
		return { ...state, ...action.payload, isLoading: false };
	case createTemplateActions.failureAction:
		return { ...state, ...action.payload, isLoading: false };
	default:
		return state;
	}
};

export const updateTemplateReducer = (
	state = updateTemplateInitialState,
	action,
) => {
	switch (action.type) {
	case updateTemplateActions.updateAction:
		return { ...state, isLoading: true };
	case updateTemplateActions.successAction:
		return { ...state, ...action.payload, isLoading: false };
	case updateTemplateActions.failureAction:
		return { ...state, ...action.payload, isLoading: false };
	default:
		return state;
	}
};

export const deleteTemplateReducer = (
	state = deleteTemplateInitialState,
	action,
) => {
	switch (action.type) {
	case deleteTemplateActions.deleteAction:
		return { ...state, isLoading: true };
	case deleteTemplateActions.successAction:
		return { ...state, ...action.payload, isLoading: false };
	case deleteTemplateActions.failureAction:
		return { ...state, ...action.payload, isLoading: false };
	default:
		return state;
	}
};
