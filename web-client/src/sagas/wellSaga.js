import { takeEvery, put, call } from "redux-saga/effects";
import { callApi } from "apis/apiHelper";
import { addWellActions, listWellActions } from "actions/wellActions";
import {
  addWellFailed,
  fetchWellsFailed,
  resetSelectedWells,
} from "action-creators/wellActionCreators";
import { HTTP_METHODS } from "appConstants";

export function* addWells(actions) {
  const {
    payload: { body, experimentId, token },
  } = actions;

  const { successAction, failureAction } = addWellActions;

  try {
    yield call(callApi, {
      payload: {
        method: HTTP_METHODS.POST,
        body,
        reqPath: `experiments/${experimentId}/wells`,
        successAction,
        failureAction,
        token,
      },
    });
    // on success of wells created will clear selected wells
    yield put(resetSelectedWells());
  } catch (error) {
    console.error("error in adding well ", error);
    yield put(addWellFailed(error));
  }
}

export function* fetchWells(actions) {
  const {
    payload: { experimentId, token },
  } = actions;

  console.log("Actions: ", actions);

  const { successAction, failureAction } = listWellActions;
  try {
    yield call(callApi, {
      payload: {
        body: null,
        reqPath: `experiments/${experimentId}/wells`,
        successAction,
        failureAction,
        token,
      },
    });
  } catch (error) {
    console.error("error in fetch stages ", error);
    yield put(fetchWellsFailed(error));
  }
}

export function* addWellsSaga() {
  yield takeEvery(addWellActions.addAction, addWells);
}

export function* fetchWellsSaga() {
  yield takeEvery(listWellActions.listAction, fetchWells);
}
