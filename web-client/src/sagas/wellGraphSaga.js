import { takeEvery, put, call } from "redux-saga/effects";
import { callApi } from "apis/apiHelper";
import { API_ENDPOINTS, HTTP_METHODS } from "appConstants";
import {
  resetGraphActions,
  updateGraphActions,
} from "actions/wellGraphActions";
import { updateGraphFailed } from "action-creators/wellGraphActionCreators";

export function* updateGraph(actions) {
  const {
    payload: { query, token, experimentId },
  } = actions;

  const { updateGraphSucceeded, updateGraphFailure } = updateGraphActions;

  try {
    yield call(callApi, {
      payload: {
        method: HTTP_METHODS.GET,
        // body: body,
        reqPath: `${API_ENDPOINTS.graphUpdate}/${experimentId}?${query}`,
        successAction: updateGraphSucceeded,
        failureAction: updateGraphFailure,
        showPopupSuccessMessage: true,
        showPopupFailureMessage: true,
        token,
      },
    });
  } catch (error) {
    console.error("Error updating graph", error);
    yield put(updateGraphFailed({ error }));
  }
}

export function* resetGraph(actions) {
  const {
    payload: { token, experimentId },
  } = actions;

  const { resetGraphSucceeded, resetGraphFailure } = resetGraphActions;

  try {
    yield call(callApi, {
      payload: {
        method: HTTP_METHODS.GET,
        reqPath: `${API_ENDPOINTS.experiments}/${experimentId}/${API_ENDPOINTS.emission}`,
        successAction: resetGraphSucceeded,
        failureAction: resetGraphFailure,
        showPopupSuccessMessage: true,
        showPopupFailureMessage: true,
        token,
      },
    });
  } catch (error) {
    console.error("Error updating graph", error);
    yield put(updateGraphFailed({ error }));
  }
}

export function* wellGraphSaga() {
  yield takeEvery(updateGraphActions.updateGraphInitiated, updateGraph);
  yield takeEvery(resetGraphActions.resetGraphInitiated, resetGraph);
}
