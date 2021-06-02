import { takeEvery, put, call } from "redux-saga/effects";
import { callApi } from "apis/apiHelper";
import {
  homingActions,
  deckHomingActions,
  discardTipAndHomingActions,
} from "actions/homingActions";
import { API_ENDPOINTS, HTTP_METHODS } from "appConstants";
import {
  homingActionFailed as homingActionFailure,
  deckHomingActionFailed as deckHomingActionFailure,
  discardTipAndHomingActionFailed as discardTipAndHomingActionFailure,
} from "action-creators/homingActionCreators";

export function* homingAction() {
  const { homingActionSuccess, homingActionFailed } = homingActions;
  try {
    yield call(callApi, {
      payload: {
        method: HTTP_METHODS.GET,
        body: null,
        reqPath: API_ENDPOINTS.homing,
        successAction: homingActionSuccess,
        failureAction: homingActionFailed,
        // showPopupSuccessMessage: true,
        showPopupFailureMessage: true,
      },
    });
  } catch (error) {
    console.error("error while homing confirmation ", error);
    yield put(homingActionFailure(error));
  }
}

export function* deckHoming(actions) {
  const {
    payload: {
      params: { deckName, token },
    },
  } = actions;
  const { deckHomingActionSuccess, deckHomingActionFailed } = deckHomingActions;

  try {
    yield call(callApi, {
      payload: {
        method: HTTP_METHODS.GET,
        body: null,
        reqPath: `${API_ENDPOINTS.homing}/${deckName}`,
        successAction: deckHomingActionSuccess,
        failureAction: deckHomingActionFailed,
        // showPopupSuccessMessage: true,
        showPopupFailureMessage: true,
        token
      },
    });
  } catch (error) {
    console.error("error while deck homing confirmation", error);
    yield put(deckHomingActionFailure(error));
  }
}

export function* discardTipAndHoming(actions) {

  const {
    payload: {
      params: { discardTip, deckName, token },
    },
  } = actions;
  const {
    discardTipAndHomingActionSuccess,
    discardTipAndHomingActionFailed,
  } = discardTipAndHomingActions;

  try {
    yield call(callApi, {
      payload: {
        method: HTTP_METHODS.GET,
        body: null,
        reqPath: `${API_ENDPOINTS.discardTipAndHoming}/${discardTip}/${deckName}`,
        successAction: discardTipAndHomingActionSuccess,
        failureAction: discardTipAndHomingActionFailed,
        // showPopupSuccessMessage: true,
        showPopupFailureMessage: true,
        token
      },
    });
  } catch (error) {
    console.error("error while discard tip and homing confirmation", error);
    yield put(discardTipAndHomingActionFailure(error));
  }
}

export function* homingActionSaga() {
  yield takeEvery(homingActions.homingActionInitiated, homingAction);
  yield takeEvery(deckHomingActions.deckHomingActionInitiated, deckHoming);
  yield takeEvery(
    discardTipAndHomingActions.discardTipAndHomingActionInitiated,
    discardTipAndHoming
  );
}
