import { fromJS } from 'immutable';
import { addWellActions, listWellActions } from 'actions/wellActions';
import loginActions from 'actions/loginActions';
import {
	getDefaultWellsList,
	setSelectedToList,
	setMultiSelectedToList,
	resetWellDefaultList,
	resetMultiWellDefaultList,
	updateWellListSelector,
	setActiveWells,
} from 'selectors/wellSelectors';
import activeWellActions from 'actions/activeWellActions';

const listWellInitialState = fromJS({
	isLoading: true,
	isMultiSelectionOptionOn: false,
	defaultList: getDefaultWellsList(),
	list: [],
	isWellFilled: false,
});

export const wellListReducer = (state = listWellInitialState, action) => {
	switch (action.type) {
	case listWellActions.listAction:
		return state.setIn(['isLoading'], true);
	case listWellActions.successAction:
		return updateWellListSelector(state, action);
	case listWellActions.failureAction:
		return state.merge({
			error: fromJS(action.payload.error),
			isLoading: false,
		});
	case listWellActions.setSelectedWell:
		return setSelectedToList(state, action.payload);
	case listWellActions.resetSelectedWell:
		return resetWellDefaultList(state);
	case listWellActions.setMultiSelectedWell:
		return setMultiSelectedToList(state, action.payload);
	case listWellActions.resetMultiSelectedWell:
		return resetMultiWellDefaultList(state);
	case listWellActions.toggleMultiSelectOption:
		if (state.get('isMultiSelectionOptionOn') === true) {
			// if group selection is set to false then clear group selection
			return resetMultiWellDefaultList(state).setIn(['isMultiSelectionOptionOn'], !state.get('isMultiSelectionOptionOn'));
		}
		// clear selected wells if any and update toggle value
		return resetWellDefaultList(state).setIn(['isMultiSelectionOptionOn'], !state.get('isMultiSelectionOptionOn'));

	// Update wells list when new wells are added
	case addWellActions.successAction:
		return updateWellListSelector(state, action);

	// Update wells list when new wells are added
	case activeWellActions.successAction:
		return setActiveWells(state, action);
	case loginActions.loginReset:
		return listWellInitialState;
	default:
		return state;
	}
};

const addWellInitialState = fromJS({
	data: {},
	isWellSaved: false,
});

export const addWellsReducer = (state = addWellInitialState, action) => {
	switch (action.type) {
	case addWellActions.addAction:
		return state.merge({
			isLoading: true,
			isWellSaved: false,
		});
	case addWellActions.successAction:
		return state.merge({
			isLoading: false,
			isWellSaved: true,
			data: action.payload.response,
		});
	case addWellActions.failureAction:
		return state.merge({
			isLoading: false,
			isWellSaved: true,
			error: action.payload.error,
		});
	case addWellActions.addWellReset:
		return addWellInitialState;
	default:
		return state;
	}
};
