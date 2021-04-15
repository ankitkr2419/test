import { takeEvery, put, call } from 'redux-saga/effects';
import { callApi } from 'apis/apiHelper';
import { discardDeckActions, discardTipActions } from "actions/discardDeckActions";
import { API_ENDPOINTS, HTTP_METHODS } from "appConstants";
import {
  discardDeckFailed as discardDeckFailure,
  discardTipFailed as discardTipFailure
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
        successAction: discardDeckSuccess,
        failureAction: discardDeckFailed
      },
    });
  } catch (error) {
    console.error("Error in discarding a deck", error);
    yield put(discardDeckFailure(error));
  }
}

export function* discardTip(actions){
  const { payload: { params : { deckName } }} = actions;
  const { discardTipSuccess, discardTipFailed } = discardTipActions;

  try {
    yield call(callApi, {
      payload: {
        method: HTTP_METHODS.GET,
        body: null,
        reqPath: `/discard-tip-and-home/true/${deckName}`, /**static api route**/
        successAction: discardTipSuccess,
        failureAction: discardTipFailed
      },
    });
  } catch (error) {
    console.error("Error in discarding a tip", error);
    yield put(discardTipFailure(error));
  }
}

export function* discardDeckSaga() {
  yield takeEvery(discardDeckActions.discardDeckInitiated, discardDeck)
}
