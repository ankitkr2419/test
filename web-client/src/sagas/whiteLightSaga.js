import { takeEvery, call } from "redux-saga/effects";
import { callApi } from "apis/apiHelper";
import {
  whiteLightDeckActions,
  whiteLightBothDeckActions,
} from "actions/whiteLightActions";
import { API_ENDPOINTS, HTTP_METHODS } from "appConstants";

export function* whiteLightDeck(actions) {
  const {
    payload: {
      params: { deck, lightStatus },
    },
  } = actions;
  const { successAction, failureAction } = whiteLightDeckActions;
  try {
    yield call(callApi, {
      payload: {
        method: HTTP_METHODS.GET,
        body: null,
        reqPath: `${API_ENDPOINTS.whiteLight}/${lightStatus}/${deck}`,
        successAction: successAction,
        failureAction: failureAction,
      },
    });
  } catch (error) {
    console.error("error while aborting", error);
  }
}

export function* whiteLightSaga() {
  yield takeEvery(whiteLightDeckActions.initiateAction, whiteLightDeck);
}
