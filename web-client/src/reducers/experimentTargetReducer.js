import { fromJS } from 'immutable';
import {
	createExperimentTargetActions,
	listExperimentTargetActions,
} from 'actions/experimentTargetActions';

const listExperimentTargetInitialState = fromJS({
	isLoading: true,
	list: [],
});

const createExperimentTargetInitialState = {
	data: {},
	isExperimentTargetSaved: false,
};


export const listExperimentTargetsReducer = (
	state = listExperimentTargetInitialState,
	action,
) => {
	switch (action.type) {
	case listExperimentTargetActions.listAction:
		return state.setIn(['isLoading'], true);
	case listExperimentTargetActions.successAction:
		return state.merge({ list: fromJS(action.payload.response || []), isLoading: false });
	case listExperimentTargetActions.failureAction:
		return state.merge({
			error: fromJS(action.payload.error),
			isLoading: false,
		});
	default:
		return state;
	}
};

export const createExperimentTargetReducer = (
	state = createExperimentTargetInitialState,
	action,
) => {
	switch (action.type) {
	case createExperimentTargetActions.createAction:
		return { ...state, isLoading: true, isExperimentTargetSaved: false };
	case createExperimentTargetActions.successAction:
		return {
			...state, ...action.payload.response, isLoading: false, isExperimentTargetSaved: true,
		};
	case createExperimentTargetActions.failureAction:
		return {
			...state, ...action.payload, isLoading: false, isExperimentTargetSaved: true,
		};
	case createExperimentTargetActions.createExperimentTargetReset:
		return {
			...state, isExperimentTargetSaved: false,
		};
	default:
		return state;
	}
};
