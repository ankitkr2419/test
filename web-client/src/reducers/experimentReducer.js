import { fromJS } from 'immutable';
import {
	createExperimentActions,
	listExperimentActions,
} from 'actions/experimentActions';

const listExperimentInitialState = fromJS({
	isLoading: true,
	list: [],
});

const createExperimentInitialState = fromJS({
	data: {},
	isExperimentSaved: false,
	// id: '5e006837-f163-4523-96e7-f39450472a6f',
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
		return state.merge({ isExperimentSaved: true, isLoading: true, isError: true });
	case createExperimentActions.createExperimentReset:
		return createExperimentInitialState.setIn(['id'], state.get('id'));
	default:
		return state;
	}
};
