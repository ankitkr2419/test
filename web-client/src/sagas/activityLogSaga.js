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
      body,
      token,
      experimentId,
    },
  } = actions;

  const { expandLogSuccess, expandLogFailure } = expandLogActions;

  try {
    yield call(callApi, {
      payload: {
        method: HTTP_METHODS.GET,
        // body: body,
        reqPath: `${API_ENDPOINTS.graphUpdate}/${experimentId}`,
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
}
