import { takeEvery, put, call } from "redux-saga/effects";
import { callApi } from "apis/apiHelper";
import { API_ENDPOINTS, HTTP_METHODS } from "appConstants";
import { appInfoAction, shutDownAction } from "actions/appInfoActions";
import { appInfoFailed } from "action-creators/appInfoActionCreators";

export function* appInfo() {
  const { appInfoSuccess, appInfoFailure } = appInfoAction;

  try {
    yield call(callApi, {
      payload: {
        method: HTTP_METHODS.GET,
        body: null,
        reqPath: `${API_ENDPOINTS.appInfo}`,
        successAction: appInfoSuccess,
        failureAction: appInfoFailure,
        showPopupFailureMessage: true,
      },
    });
  } catch (error) {
    console.error("Error in fetching app info: ", error);
    yield put(appInfoFailed({ error }));
  }
}

export function* shutdown() {
  const { shutdownSuccess, shutdownFailure } = shutDownAction;
  try {
    yield call(callApi, {
      payload: {
        method: HTTP_METHODS.GET,
        body: null,
        reqPath: `${API_ENDPOINTS.shutdown}`,
        successAction: shutdownSuccess,
        failureAction: shutdownFailure,
      },
    });
  } catch (error) {
    console.error("Error in fetching app info: ", error);
  }
}

export function* appInfoSaga() {
  yield takeEvery(appInfoAction.appInfoInitiated, appInfo);
  yield takeEvery(shutDownAction.shutdownInitiated, shutdown);
}

