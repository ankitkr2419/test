import { fromJS } from 'immutable';
import { addWellActions, listWellActions } from 'actions/wellActions';
import {
	getDefaultPlatesList,
	setSelectedToList,
} from 'selectors/wellSelectors';

const listWellInitialState = fromJS({
	isLoading: true,
	defaultList: getDefaultPlatesList(),
	list: [],
});

export const wellListReducer = (state = listWellInitialState, action) => {
	switch (action.type) {
	case listWellActions.listAction:
		return state.setIn(['isLoading'], true);
	case listWellActions.successAction:
		return state.merge({
			list: fromJS(action.payload.response || []),
			isLoading: false,
		});
	case listWellActions.failureAction:
		return state.merge({
			error: fromJS(action.payload.error),
			isLoading: false,
		});
	case listWellActions.setWellSelected:
		return setSelectedToList(state, action.payload);
	default:
		return state;
	}
};

const addWellInitialState = {
	data: {},
	isStageSaved: false,
};

export const createStageReducer = (state = addWellInitialState, action) => {
	switch (action.type) {
	case addWellActions.addAction:
		return { ...state, isLoading: true, isStageSaved: false };
	case addWellActions.successAction:
		return {
			...state,
			...action.payload,
			isLoading: false,
			isStageSaved: true,
		};
	case addWellActions.failureAction:
		return {
			...state,
			...action.payload,
			isLoading: false,
			isStageSaved: false,
		};
	case addWellActions.addStageReset:
		return addWellInitialState;
	default:
		return state;
	}
};
