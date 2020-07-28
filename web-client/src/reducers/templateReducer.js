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
	isLoading: true,
	isTemplateCreated: false,
};

const updateTemplateInitialState = {
	data: {},
	isTemplateUpdated: false,
};

const deleteTemplateInitialState = {
	data: {},
	isTemplateDeleted: false,
};

// eslint-disable-next-line arrow-body-style
export const getTemplateById = (state, templateId) => {
	const result = state.get('list').filter(ele => ele.get('id') === templateId);
	if (result !== null && result.size !== 0) {
		return result.get(0);
	}
	return null;
};

export const listTemplatesReducer = (
	state = listTemplateInitialState,
	action,
) => {
	switch (action.type) {
	case listTemplateActions.listAction:
		return state.setIn(['isLoading'], true);
	case listTemplateActions.successAction:
		return state.merge({
			list: fromJS(action.payload.response || []),
			isLoading: false,
		});
	case listTemplateActions.failureAction:
		return state.merge({
			error: fromJS(action.payload.error),
			isLoading: false,
		});
		// appending the template list with created template
	case createTemplateActions.successAction:
		return state.updateIn(['list'], list => list.push(fromJS(action.payload.response.template)));
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
		return { ...state, isLoading: true, isTemplateCreated: false };
	case createTemplateActions.successAction:
		return {
			...state,
			...action.payload,
			isLoading: false,
			isTemplateCreated: true,
		};
	case createTemplateActions.failureAction:
		return {
			...state,
			...action.payload,
			isLoading: false,
			isTemplateCreated: true,
		};
	case createTemplateActions.createTemplateReset:
		return deleteTemplateInitialState;
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
		return {
			...state,
			...action.payload,
			isLoading: false,
			isTemplateUpdated: true,
		};
	case updateTemplateActions.failureAction:
		return {
			...state,
			...action.payload,
			isLoading: false,
			isTemplateUpdated: true,
		};
	case updateTemplateActions.updateTemplateReset:
		return updateTemplateInitialState;
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
		return { ...state, isLoading: true, isTemplateDeleted: false };
	case deleteTemplateActions.successAction:
		return {
			...state,
			...action.payload,
			isLoading: false,
			isTemplateDeleted: true,
		};
	case deleteTemplateActions.failureAction:
		return {
			...state,
			...action.payload,
			isLoading: false,
			isTemplateDeleted: true,
		};
	case deleteTemplateActions.deleteTemplateReset:
		return deleteTemplateInitialState;
	default:
		return state;
	}
};
