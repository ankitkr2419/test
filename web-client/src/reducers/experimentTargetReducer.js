import { fromJS } from 'immutable';
import {
	createExperimentTargetActions,
	listExperimentTargetActions,
} from 'actions/experimentTargetActions';
import { addIsActiveFlag } from 'selectors/experimentTargetSelector';

const listExperimentTargetInitialState = fromJS({
	isLoading: true,
	list: [],
	graphTargets: [],
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
		return state.merge({
			list: fromJS(action.payload.response || []),
			graphTargets: fromJS(addIsActiveFlag(action.payload.response || [])),
			isLoading: false,
		});
	case listExperimentTargetActions.updateGraphFilters:
		return state.setIn(
			['graphTargets', action.payload.index, action.payload.key],
			action.payload.value,
		);
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
			...state,
			...action.payload.response,
			isLoading: false,
			isExperimentTargetSaved: true,
		};
	case createExperimentTargetActions.failureAction:
		return {
			...state,
			...action.payload,
			isLoading: false,
			isExperimentTargetSaved: true,
		};
	case createExperimentTargetActions.createExperimentTargetReset:
		return {
			...state,
			isExperimentTargetSaved: false,
		};
	default:
		return state;
	}
};
