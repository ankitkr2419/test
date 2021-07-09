import { takeEvery, put, call } from "redux-saga/effects";
import { callApi } from "apis/apiHelper";
import { listSampleActions } from "actions/sampleActions";
import { fetchSamplesFailed } from "action-creators/sampleActionCreators";

export function* fetchSamples(actions) {
  const {
    payload: { searchText, token },
  } = actions;

  const { successAction, failureAction } = listSampleActions;
  try {
    yield call(callApi, {
      payload: {
        body: null,
        reqPath: `samples/?text=${searchText}`,
        successAction,
        failureAction,
        token,
      },
    });
  } catch (error) {
    console.error("error in fetch stages ", error);
    yield put(fetchSamplesFailed(error));
  }
}

export function* fetchSamplesSaga() {
  yield takeEvery(listSampleActions.listAction, fetchSamples);
}
