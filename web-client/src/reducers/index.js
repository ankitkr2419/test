import { combineReducers } from 'redux';
import { createTemplateReducer, listTemplatesReducer } from './templateReducer';

const rootReducer = combineReducers({
  createTemplateReducer,
  listTemplatesReducer,
});

export default rootReducer;
