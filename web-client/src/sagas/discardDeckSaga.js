import { takeEvery, put, call } from 'redux-saga/effects';
import { callApi } from 'apis/apiHelper';
import { discardDeckActions } from "actions/discardDeckActions";
import { API_ENDPOINTS, HTTP_METHODS } from "appConstants";
import {
  discardDeckFailed as discardDeckFailure
} from "action-creators/discardDeckActionCreators";

export function* discardDeck(actions){
  const { payload: { params : { deckName } }} = actions;
  const { discardDeckSuccess, discardDeckFailed } = discardDeckActions;

  try {
    yield call(callApi, {
      payload: {
        method: HTTP_METHODS.GET,
        body: null,
        reqPath: `${API_ENDPOINTS.discardDeck}/${deckName}`,
        discardDeckSuccess,
        discardDeckFailed
      },
    });
  } catch (error) {
    console.error("Error in discarding a deck", error);
    yield put(discardDeckFailure(error));
  }
}

export function* discardDeckSaga() {
  yield takeEvery(discardDeckActions.discardDeckInitiated, discardDeck)
}
