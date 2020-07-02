import { combineReducers } from 'redux';
import {
	createTemplateReducer,
	listTemplatesReducer,
	deleteTemplateReducer,
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

import { loginReducer } from 'reducers/loginReducer';
import { wellListReducer, addWellsReducer } from 'reducers/wellReducer';
import { listExperimentsReducer, createExperimentReducer } from 'reducers/experimentReducer';
import { listExperimentTargetsReducer, createExperimentTargetReducer } from 'reducers/experimentTargetReducer';
import { listSamplesReducer } from 'reducers/samplesReducer';
import { runExperimentReducer } from 'reducers/runExperimentReducer';
import { activeWellReducer } from 'reducers/actionWellReducer';

const rootReducer = combineReducers({
	createTemplateReducer,
	listTemplatesReducer,
	deleteTemplateReducer,
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
	loginReducer,
	wellListReducer,
	listExperimentsReducer,
	createExperimentReducer,
	listExperimentTargetsReducer,
	createExperimentTargetReducer,
	listSamplesReducer,
	addWellsReducer,
	runExperimentReducer,
	activeWellReducer,
});

export default rootReducer;
