import { all } from "redux-saga/effects";
import { createTemplateSaga, fetchTemplatesSaga } from "sagas/templateSaga";
import { saveTargetSaga, fetchMasterTargetsSaga, fetchTargetsByTemplateIDSaga } from "sagas/targetSaga";

const allSagas = [
  createTemplateSaga(),
  fetchTemplatesSaga(),
  saveTargetSaga(),
  fetchMasterTargetsSaga(),
  fetchTargetsByTemplateIDSaga(),
];

export default function* rootSaga() {
  yield all(allSagas);
}
