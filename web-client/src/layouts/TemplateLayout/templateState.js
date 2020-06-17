import { fromJS } from 'immutable';
import { wizardList } from './templateConstant';

// const action types
export const templateLayoutActions = {
	SET_ACTIVE_WIDGET: 'SET_ACTIVE_WIDGET',
	SET_TEMPLATE_ID: 'SET_TEMPLATE_ID',
	SET_STAGE_ID: 'SET_STAGE_ID',
};

// Initial state wrap with fromJS for immutability
export const templateInitialState = fromJS({
	// active wizard id
	activeWidgetID: 'stage',
	// Pre-filled template initial list with saved wizard list
	wizardList,
	templateID: 'df4914e9-41e1-4d51-8655-23de1bffdc86',
	stageID: null,
});

// getUpdateList will update all disabled to true and set false to selected wizard
const getUpdatedList = (state, selectedID) => {
	let updatedState = state.update('wizardList', item => item.map((keyValue) => {
		if (keyValue.get('id') === selectedID) {
			return keyValue.set('isDisable', false);
		}
		return keyValue;
	}));
	updatedState = updatedState.setIn(['activeWidgetID'], selectedID);
	return updatedState;
};

const templateLayoutReducer = (state, action) => {
	switch (action.type) {
	case templateLayoutActions.SET_ACTIVE_WIDGET:
		return getUpdatedList(state, action.value);
	case templateLayoutActions.SET_TEMPLATE_ID:
		return state.setIn(['templateID'], action.value);
	case templateLayoutActions.SET_STAGE_ID:
		return state.setIn(['stageID'], action.value);
	default:
		throw new Error('Invalid action type');
	}
};

export default templateLayoutReducer;
