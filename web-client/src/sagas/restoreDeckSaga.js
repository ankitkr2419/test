import { takeEvery, put, call } from "redux-saga/effects";
import { callApi } from "apis/apiHelper";
import { restoreDeckActions } from "actions/restoreDeckActions";
import { API_ENDPOINTS, HTTP_METHODS } from "appConstants";
import { restoreDeckFailed as restoreDeckFailure } from "action-creators/restoreDeckActionCreators";

export function* restoreDeck(actions) {
  const {
    payload: {
      params: { deckName },
    },
  } = actions;
  const { restoreDeckSuccess, restoreDeckFailed } = restoreDeckActions;

  try {
    yield call(callApi, {
      payload: {
        method: HTTP_METHODS.GET,
        body: null,
        reqPath: `${API_ENDPOINTS.restoreDeck}/${deckName}`,
        successAction: restoreDeckSuccess,
        failureAction: restoreDeckFailed,
      },
    });
  } catch (error) {
    console.error("Error in restoring a deck", error);
    yield put(restoreDeckFailure(error));
  }
}

export function* restoreDeckSaga() {
  yield takeEvery(restoreDeckActions.restoreDeckInitiated, restoreDeck);
}
