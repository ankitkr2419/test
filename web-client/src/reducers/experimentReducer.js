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
	// id: '98f26854-d52c-4a14-8d96-09bcc89c3f86',
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
