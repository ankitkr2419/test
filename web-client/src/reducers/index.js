import { combineReducers } from 'redux';
import {
	createTemplateReducer,
	listTemplatesReducer,
} from 'reducers/templateReducer';
import {
	listTargetReducer,
	listTargetByTemplateIDReducer,
	saveTargetReducer,
} from 'reducers/targetReducer';

const rootReducer = combineReducers({
	createTemplateReducer,
	listTemplatesReducer,
	listTargetReducer,
	listTargetByTemplateIDReducer,
	saveTargetReducer,
});

export default rootReducer;
