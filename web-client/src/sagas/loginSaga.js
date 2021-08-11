import { takeEvery, call, put } from "redux-saga/effects";
// Purposefully commented code, will remvove once api is integrated
import { callApi } from "apis/apiHelper";
import loginActions, { logoutActions } from "actions/loginActions";
// import { loginFailed } from 'action-creators/loginActionCreators';
import { API_ENDPOINTS, DECKNAME, HTTP_METHODS } from "../appConstants";
import { logoutFailure } from "action-creators/loginActionCreators";

export function* login(actions) {
  const {
    payload: {
      body: { email, password, role, deckName, showToast },
    },
  } = actions;

  let deck = "";
  if (deckName !== "") {
    deck =
      deckName === DECKNAME.DeckA ? DECKNAME.DeckAShort : DECKNAME.DeckBShort;
  }

  const { successAction, failureAction } = loginActions;
  try {
    yield call(callApi, {
      payload: {
        body: {
          username: email,
          password: password,
          role: role,
        },
        reqPath: `${API_ENDPOINTS.login}/${deck}`,
        method: HTTP_METHODS.POST,
        successAction,
        failureAction,
        showPopupSuccessMessage: showToast,
        showPopupFailureMessage: showToast,
      },
    });
  } catch (error) {
    console.log("error while login: ", error);
  }
}

export function* logout(actions) {
  const {
    payload: { deckName, token, showToast },
  } = actions;

  let deck = "";
  if (deckName !== "") {
    deck =
      deckName === DECKNAME.DeckA ? DECKNAME.DeckAShort : DECKNAME.DeckBShort;
  }

  const { logoutActionSuccess, logoutActionFailure } = logoutActions;
  try {
    yield call(callApi, {
      payload: {
        reqPath: `${API_ENDPOINTS.logout}/${deck}`,
        method: HTTP_METHODS.DELETE,
        successAction: logoutActionSuccess,
        failureAction: logoutActionFailure,
        showPopupSuccessMessage: showToast,
        showPopupFailureMessage: showToast,
        token: token,
      },
    });
  } catch (error) {
    console.log("error while logout: ", error);
    yield put(logoutFailure(error));
  }
}

export function* loginSaga() {
  yield takeEvery(loginActions.loginInitiated, login);
  yield takeEvery(logoutActions.logoutActionInitiated, logout);
}
