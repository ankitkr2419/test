import { takeEvery, put, call } from "redux-saga/effects";
import { callApi } from "apis/apiHelper";
import { API_ENDPOINTS, HTTP_METHODS } from "appConstants";
import { activityLogActions } from "actions/activityLogActions";
import { activityLogFailed } from "action-creators/activityLogActionCreators";

export function* fetchActivityLog(actions) {
  const {
    payload: { token },
  } = actions;
  const { activityLogSuccess, activityLogFailure } = activityLogActions;

  try {
    yield call(callApi, {
      payload: {
        method: HTTP_METHODS.GET,
        body: null,
        reqPath: `${API_ENDPOINTS.experiments}`,
        successAction: activityLogSuccess,
        failureAction: activityLogFailure,
        token,
      },
    });
  } catch (error) {
    console.error("Error fetching activity list", error);
    yield put(activityLogFailed({ error }));
  }
}

export function* activityLogSaga() {
  yield takeEvery(activityLogActions.activityLogInitiated, fetchActivityLog);
}
