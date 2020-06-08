import { combineReducers } from "redux";
import {
  createTemplateReducer,
  listTemplatesReducer,
} from "reducers/templateReducer";
import {
  listTargetReducer,
  listTargetByTemplateIDReducer,
} from "reducers/targetReducer";

const rootReducer = combineReducers({
  createTemplateReducer,
  listTemplatesReducer,
  listTargetReducer,
  listTargetByTemplateIDReducer,
});

export default rootReducer;
