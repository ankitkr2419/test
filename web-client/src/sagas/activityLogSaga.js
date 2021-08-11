import { takeEvery, put, call } from "redux-saga/effects";
import { callApi } from "apis/apiHelper";
import { API_ENDPOINTS, HTTP_METHODS } from "appConstants";
import {
  activityLogActions,
  expandLogActions,
  mailReportActions,
} from "actions/activityLogActions";
import {
  expandLogFailed,
  activityLogFailed,
  mailReportFailed,
} from "action-creators/activityLogActionCreators";

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

export function* sendMail(actions) {
  const {
    payload: { token, experimentId, showToast },
  } = actions;
  const { mailReportSuccess, mailReportFailure } = mailReportActions;

  try {
    yield call(callApi, {
      payload: {
        method: HTTP_METHODS.GET,
        reqPath: `${API_ENDPOINTS.emailReport}/${experimentId}`,
        successAction: mailReportSuccess,
        failureAction: mailReportFailure,
        showPopupSuccessMessage: showToast,
        showPopupFailureMessage: showToast,
        token,
      },
    });
  } catch (error) {
    console.error("Error sending mail", error);
    yield put(mailReportFailed({ error }));
  }
}

export function* expand(actions) {
  const {
    payload: {
      params: { x_axis_min, x_axis_max, y_axis_min, y_axis_max },
      token,
      experimentId,
    },
  } = actions;

  const { expandLogSuccess, expandLogFailure } = expandLogActions;
  const queryStr = `x_axis_min=${x_axis_min}&x_axis_max=${x_axis_max}&y_axis_min=${y_axis_min}&y_axis_max=${y_axis_max}`;

  try {
    yield call(callApi, {
      payload: {
        method: HTTP_METHODS.GET,
        // body: body,
        reqPath: `${API_ENDPOINTS.graphUpdate}/${experimentId}?${queryStr}`,
        successAction: expandLogSuccess,
        failureAction: expandLogFailure,
        token,
      },
    });
  } catch (error) {
    console.error("Error fetching activity list", error);
    yield put(expandLogFailed({ error }));
  }
}

export function* activityLogSaga() {
  yield takeEvery(activityLogActions.activityLogInitiated, fetchActivityLog);
  yield takeEvery(mailReportActions.mailReportInitiated, sendMail);
  yield takeEvery(expandLogActions.expandLogInitiated, expand);
}
