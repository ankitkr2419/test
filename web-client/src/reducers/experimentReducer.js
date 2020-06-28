import { fromJS } from 'immutable';
import {
	createExperimentActions,
	listExperimentActions,
} from 'actions/experimentActions';

const listExperimentInitialState = fromJS({
	isLoading: true,
	list: [],
});

const createExperimentInitialState = {
	data: {},
	isExperimentSaved: false,
	id: '178aa3ff-6ee3-4dd2-8276-f19d9df0330a',
	description: null,
};


export const listExperimentsReducer = (
	state = listExperimentInitialState,
	action,
) => {
	switch (action.type) {
	case listExperimentActions.listAction:
		return state.setIn(['isLoading'], true);
	case listExperimentActions.successAction:
		return state.merge({ list: fromJS(action.payload.response || []), isLoading: false });
	case listExperimentActions.failureAction:
		return state.merge({
			error: fromJS(action.payload.error),
			isLoading: false,
		});
	default:
		return state;
	}
};

export const createExperimentReducer = (
	state = createExperimentInitialState,
	action,
) => {
	switch (action.type) {
	case createExperimentActions.createAction:
		return { ...state, isLoading: true, isExperimentSaved: false };
	case createExperimentActions.successAction:
		return {
			...state, ...action.payload.response, isLoading: false, isExperimentSaved: true,
		};
	case createExperimentActions.failureAction:
		return {
			...state, ...action.payload, isLoading: false, isExperimentSaved: true,
		};
	case createExperimentActions.createExperimentReset:
		return {
			...state, isExperimentSaved: false,
		};
	default:
		return state;
	}
};
