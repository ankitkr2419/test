import { takeEvery, call, put } from "redux-saga/effects";
// Purposefully commented code, will remvove once api is integrated
import { callApi } from "apis/apiHelper";
import loginActions, { logoutActions } from "actions/loginActions";
// import { loginFailed } from 'action-creators/loginActionCreators';
import { API_ENDPOINTS, DECKNAME, HTTP_METHODS } from "../appConstants";
import { logoutFailure } from "action-creators/loginActionCreators";

export function* login(actions) {
  const {
    payload: { body },
  } = actions;

  let deckName =
    actions.payload.body.deckName === DECKNAME.DeckA
      ? DECKNAME.DeckAShort
      : DECKNAME.DeckBShort;
  const { successAction, failureAction } = loginActions;
  try {
    yield call(callApi, {
      payload: {
        body: {
          username: body.email,
          password: body.password,
          role: body.role,
        },
        reqPath: `${API_ENDPOINTS.login}/${deckName}`,
        method: HTTP_METHODS.POST,
        successAction,
        failureAction,
        showPopupSuccessMessage: true,
        showPopupFailureMessage: true,
      },
    });
  } catch (error) {
    // yield put(operatorLoginFailed(error));
    console.log("error while login: ", error);
    // yield put(loginFailed(error))
    // yield put(toast.error('failure!'))
  }
}

export function* logout(actions) {
  const {
    payload: { deckName, token },
  } = actions;

  let deck =
    deckName === DECKNAME.DeckA ? DECKNAME.DeckAShort : DECKNAME.DeckBShort;

  const { logoutActionSuccess, logoutActionFailure } = logoutActions;
  try {
    yield call(callApi, {
      payload: {
        reqPath: `${API_ENDPOINTS.logout}/${deck}`,
        method: HTTP_METHODS.DELETE,
        successAction: logoutActionSuccess,
        failureAction: logoutActionFailure,
        showPopupSuccessMessage: true,
        showPopupFailureMessage: true,
        token: token,
      },
    });
  } catch (error) {
    // yield put(operatorLoginFailed(error));
    console.log("error while logout: ", error);
    // yield put(loginFailed(error))
    // yield put(toast.error('failure!'))
    yield put(logoutFailure(error));
  }
}

export function* loginSaga() {
  yield takeEvery(loginActions.loginInitiated, login);
  yield takeEvery(logoutActions.logoutActionInitiated, logout);
}
