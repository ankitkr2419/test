import { all } from "redux-saga/effects";
import {
  createTemplateSaga,
  fetchTemplatesSaga,
  createTemplateSuccessSaga,
  deleteTemplateSaga,
  updateTemplateSaga,
} from "sagas/templateSaga";
import {
  saveTargetSaga,
  fetchMasterTargetsSaga,
  fetchTargetsByTemplateIDSaga,
} from "sagas/targetSaga";

import { fetchStagesSaga, updateStageSaga } from "sagas/stageSaga";

import {
  fetchHoldStepsSaga,
  fetchCycleStepsSaga,
  addStepSaga,
  deleteStepSaga,
  updateStepSaga,
} from "sagas/stepSaga";

import { loginSaga } from "sagas/loginSaga";

import { createExperimentSaga } from "sagas/experimentSaga";

import {
  fetchExperimentTargetsSaga,
  createExperimentTargetSaga,
} from "sagas/experimentTargetSaga";

import { fetchSamplesSaga } from "./samplesSaga";
import { addWellsSaga, fetchWellsSaga } from "./wellSaga";
import { runExperimentSaga, stopExperimentSaga } from "./runExperimentSaga";
import { fetchActiveWellsSaga } from "./actionWellSaga";
import { operatorLoginModalSaga } from "./operatorLoginModalSaga";
import { homingActionSaga } from "./homingSaga";
import { recipeActionSaga } from "./recipeActionSaga";
import { restoreDeckSaga } from "./restoreDeckSaga";
import { discardDeckSaga } from "./discardDeckSaga";

const allSagas = [
  createTemplateSaga(),
  fetchTemplatesSaga(),
  createTemplateSuccessSaga(),
  saveTargetSaga(),
  fetchMasterTargetsSaga(),
  fetchTargetsByTemplateIDSaga(),
  deleteTemplateSaga(),
  updateTemplateSaga(),
  fetchStagesSaga(),
  updateStageSaga(),
  fetchHoldStepsSaga(),
  fetchCycleStepsSaga(),
  addStepSaga(),
  deleteStepSaga(),
  updateStepSaga(),
  loginSaga(),
  createExperimentSaga(),
  fetchExperimentTargetsSaga(),
  createExperimentTargetSaga(),
  fetchSamplesSaga(),
  addWellsSaga(),
  fetchWellsSaga(),
  runExperimentSaga(),
  stopExperimentSaga(),
  fetchActiveWellsSaga(),
  operatorLoginModalSaga(),
  homingActionSaga(),
  recipeActionSaga(),
  restoreDeckSaga(),
  discardDeckSaga(),
];

export default function* rootSaga() {
  yield all(allSagas);
}
