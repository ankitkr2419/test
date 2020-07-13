import { combineReducers } from 'redux';
import {
	createTemplateReducer,
	listTemplatesReducer,
	deleteTemplateReducer,
	updateTemplateReducer,
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
import { wellGraphReducer } from 'reducers/wellGraphReducer';
import { socketReducer } from 'reducers/socketReducer';

const rootReducer = combineReducers({
	createTemplateReducer,
	listTemplatesReducer,
	deleteTemplateReducer,
	updateTemplateReducer,
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
	wellGraphReducer,
	socketReducer,
});

export default rootReducer;
