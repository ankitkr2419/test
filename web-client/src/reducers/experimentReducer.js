import { fromJS } from 'immutable';
import {
	createExperimentActions,
	listExperimentActions,
} from 'actions/experimentActions';
import loginActions from 'actions/loginActions';

const listExperimentInitialState = fromJS({
	isLoading: true,
	list: [],
});

const createExperimentInitialState = fromJS({
	data: {},
	isExperimentSaved: false,
	isExpanded: false, // isExpanded: boolean -> Determines whether this page is redirect via normal flow or by expanding
	id: null,
	description: null,
});


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
		return state.merge({ isExperimentSaved: false, isLoading: true });
	case createExperimentActions.successAction:
		return state.merge({ isExperimentSaved: true, isLoading: false, ...action.payload.response });
	case createExperimentActions.failureAction:
		return state.merge({ isExperimentSaved: true, isLoading: false, isError: true, id:null });
	case createExperimentActions.createExperimentReset:
		return state.setIn(['isExperimentSaved'], false);
	case loginActions.loginReset:
		return createExperimentInitialState;
	default:
		return state;
	}
};
