import { combineReducers } from 'redux';
import { createTemplateReducer, listTemplateReducer } from './templateReducer';

const rootReducer = combineReducers({
  createTemplateReducer,
  listTemplateReducer,
});

export default rootReducer;
