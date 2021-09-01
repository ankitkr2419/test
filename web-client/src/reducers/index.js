import { combineReducers } from "redux";
import {
  createTemplateReducer,
  listTemplatesReducer,
  deleteTemplateReducer,
  updateTemplateReducer,
  finishCreateTemplateReducer,
} from "reducers/templateReducer";
import {
  listTargetReducer,
  listTargetByTemplateIDReducer,
  saveTargetReducer,
} from "reducers/targetReducer";

import { listStagesReducer, updateStageReducer } from "reducers/stageReducer";

import {
  listHoldStepsReducer,
  listCycleStepsReducer,
  createStepReducer,
  deleteStepReducer,
  updateStepReducer,
} from "reducers/stepReducer";

import { loginReducer } from "reducers/loginReducer";
import { wellListReducer, addWellsReducer } from "reducers/wellReducer";
import {
  listExperimentsReducer,
  createExperimentReducer,
} from "reducers/experimentReducer";
import {
  listExperimentTargetsReducer,
  createExperimentTargetReducer,
} from "reducers/experimentTargetReducer";
import { listSamplesReducer } from "reducers/samplesReducer";
import {
  runExpProgressReducer,
  runExperimentReducer,
} from "reducers/runExperimentReducer";
import { activeWellReducer } from "reducers/actionWellReducer";
import {
  wellGraphReducer,
  updateWellGraphReducer,
} from "reducers/wellGraphReducer";
import { socketReducer } from "reducers/socketReducer";
import { modalReducer } from "reducers/modalReducer";
import { temperatureGraphReducer } from "reducers/temperatureGraphReducer";
import { operatorLoginModalReducer } from "reducers/operatorLoginModalReducer";
import {
  homingReducer,
  discardTipAndHomingReducer,
} from "reducers/homingReducer";
import { recipeActionReducer } from "reducers/recipeActionReducer";
import { restoreDeckReducer } from "reducers/restoreDeckReducer";
import { discardDeckReducer } from "reducers/discardDeckReducer";
import { cleanUpReducer } from "reducers/cleanUpReducer";
import { updateRecipeDetailsReducer } from "reducers/updateRecipeDetailsReducer";
import { processListReducer } from "reducers/processListReducer";
import { editProcessReducer } from "reducers/editProcessReducer";
import { appInfoReducer } from "./appInfoReducer";
import { processesReducer } from "reducers/processesReducer";
import {
  activityLogReducer,
  mailReportReducer,
} from "reducers/activityLogReducer";
import {
  heaterProgressReducer,
  commonDetailsReducer,
  calibrationReducer,
  updateCalibrationReducer,
  pidProgessReducer,
  pidReducer,
  createTipTubeReducer,
  rtpcrConfigsReducer,
} from "./calibrationReducer";
import { reportReducer } from "reducers/reportReducer";
import {
  analyseDataGraphFiltersReducer,
  analyseDataGraphThresholdReducer,
  analyseDataGraphBaselineReducer,
} from "./analyseDataGraph";

const rootReducer = combineReducers({
  createTemplateReducer,
  listTemplatesReducer,
  deleteTemplateReducer,
  updateTemplateReducer,
  listTargetReducer,
  listTargetByTemplateIDReducer,
  saveTargetReducer,
  listStagesReducer,
  updateStageReducer,
  listHoldStepsReducer,
  listCycleStepsReducer,
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
  runExpProgressReducer,
  activeWellReducer,
  wellGraphReducer,
  updateWellGraphReducer,
  socketReducer,
  modalReducer,
  temperatureGraphReducer,
  operatorLoginModalReducer,
  homingReducer,
  recipeActionReducer,
  restoreDeckReducer,
  discardDeckReducer,
  discardTipAndHomingReducer,
  cleanUpReducer,
  updateRecipeDetailsReducer,
  processListReducer,
  editProcessReducer,
  appInfoReducer,
  processesReducer,
  activityLogReducer,
  mailReportReducer,
  commonDetailsReducer,
  calibrationReducer,
  updateCalibrationReducer,
  heaterProgressReducer,
  pidProgessReducer,
  pidReducer,
  finishCreateTemplateReducer,
  reportReducer,
  analyseDataGraphFiltersReducer,
  analyseDataGraphThresholdReducer,
  analyseDataGraphBaselineReducer,
  createTipTubeReducer,
  rtpcrConfigsReducer,
});

export default rootReducer;
