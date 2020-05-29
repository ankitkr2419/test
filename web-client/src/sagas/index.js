import { all } from "redux-saga/effects";

const allSagas = [];

export default function* rootSaga() {
  yield all(allSagas);
}
