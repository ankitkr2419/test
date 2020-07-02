import { fromJS } from 'immutable';
import {
	listTargetActions,
	saveTargetActions,
	listTargetByTemplateIDActions,
} from 'actions/targetActions';

const listTargetInitialState = fromJS({
	list: [],
	error: false,
});

export const listTargetReducer = (state = listTargetInitialState, action) => {
	switch (action.type) {
	case listTargetActions.listAction:
		return state.setIn(['isLoading'], true);
	case listTargetActions.successAction:
		return state.merge({ list: fromJS(action.payload.response || []), isLoading: false });
	case listTargetActions.failureAction:
		return state.merge({
			error: fromJS(action.payload.error),
			isLoading: false,
		});
	default:
		return state;
	}
};

const listTargetByTemplateIDInitialState = fromJS({
	selectedTargets: [],
	error: false,
	isLoading: true,
});

export const listTargetByTemplateIDReducer = (
	state = listTargetByTemplateIDInitialState,
	action,
) => {
	switch (action.type) {
	case listTargetByTemplateIDActions.listAction:
		return state.setIn(['isLoading'], true);
	case listTargetByTemplateIDActions.successAction:
		return state.merge({ selectedTargets: fromJS(action.payload.response || []), isLoading: false });
	case listTargetByTemplateIDActions.failureAction:
		return state.merge({
			error: fromJS(action.payload.error),
			isLoading: false,
		});
	default:
		return state;
	}
};

const saveTargetInitialState = {
	data: {},
	isTargetSaved: false,
};

export const saveTargetReducer = (state = saveTargetInitialState, action) => {
	switch (action.type) {
	case saveTargetActions.createAction:
		return {
			...state, ...action.payload, isLoading: true, isTargetSaved: false,
		};
	case saveTargetActions.successAction:
		return {
			...state, ...action.payload, isLoading: false, isTargetSaved: true,
		};
	case saveTargetActions.failureAction:
		return {
			...state, ...action.payload, isLoading: false, isTargetSaved: false,
		};
	case saveTargetActions.saveTargetReset:
		return saveTargetInitialState;
	default:
		return state;
	}
};
