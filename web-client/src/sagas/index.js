import { all } from "redux-saga/effects";
import { createTemplateSaga, fetchTemplatesSaga } from "./templateSaga";

const allSagas = [
  createTemplateSaga(),
  fetchTemplatesSaga(),
];

export default function* rootSaga() {
  yield all(allSagas);
}
