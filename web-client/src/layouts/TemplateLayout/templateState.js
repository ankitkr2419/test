import { fromJS } from 'immutable';
import { wizardList } from './templateConstant';

// const action types
export const templateLayoutActions = {
  SET_ACTIVE_WIDGET: 'SET_ACTIVE_WIDGET',
};

// Initial state wrap with fromJS for immutability
export const templateInitialState = fromJS({
  activeWidgetID: 0,
  // Pre-filled template initial list with saved wizard list
  wizardList,
});

// getUpdateList will update all disabled to true and set false to selected wizard
const getUpdateList = (state, selectedID) => {
  const updatedState = state.update('wizardList', item => item.map((keyValue) => {
    if (keyValue.get('id') === selectedID) {
      return keyValue.set('isDisable', false);
    }
    return keyValue.set('isDisable', true);
  }));
  return updatedState;
};

const templateLayoutReducer = (state, action) => {
  switch (action.type) {
    case templateLayoutActions.SET_ACTIVE_WIDGET:
      return getUpdateList(state, action.value);
    default:
      throw new Error('Invalid action type');
  }
};

export default templateLayoutReducer;
