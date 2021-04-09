import { takeEvery, put, call } from "redux-saga/effects";
import { callApi } from "apis/apiHelper";
import { cleanUpActions } from "actions/cleanUpActions";
import { API_ENDPOINTS, HTTP_METHODS } from "appConstants";
import {
  cleanUpActionFailed as cleanUpActionFailure,
} from "action-creators/cleanUpActionCreators";

export function* UVCleaning(actions) {
  const {
    payload: {
      params: { time, deckName },
    },
  } = actions;
  const { cleanUpActionSuccess, cleanUpActionFailed } = cleanUpActions;

  try {
    yield call(callApi, {
      payload: {
        method: HTTP_METHODS.GET,
        body: null,
        reqPath: `${API_ENDPOINTS.cleanUp}/${time}/${deckName}`,
        cleanUpActionSuccess,
        cleanUpActionFailed,
      },
    });
  } catch (error) {
    console.error("error while deck cleaning", error);
    yield put(cleanUpActionFailure(error));
  }
}

export function* cleanUpSaga() {
  yield takeEvery(cleanUpActions.cleanUpActionInitiated, UVCleaning);
}
