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

import {
	listStagesReducer,
	createStageReducer,
	deleteStageReducer,
	updateStageReducer,
} from 'reducers/stageReducer';

import {
	listStepsReducer,
	createStepReducer,
	deleteStepReducer,
	updateStepReducer,
} from 'reducers/stepReducer';

const rootReducer = combineReducers({
	createTemplateReducer,
	listTemplatesReducer,
	listTargetReducer,
	listTargetByTemplateIDReducer,
	saveTargetReducer,
	listStagesReducer,
	createStageReducer,
	deleteStageReducer,
	updateStageReducer,
	listStepsReducer,
	createStepReducer,
	deleteStepReducer,
	updateStepReducer,
});

export default rootReducer;
