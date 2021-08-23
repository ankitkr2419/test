import { takeEvery, put, call } from "redux-saga/effects";
import { callApi } from "apis/apiHelper";
import { API_ENDPOINTS, HTTP_METHODS } from "appConstants";
import { saveReportActions } from "actions/reportActions";
import { saveReportFailed } from "action-creators/reportActionCreators";

export function* saveReport(actions) {
  const {
    payload: { token, experimentId, data },
  } = actions;
  const { saveReportSuccess, saveReportFailure } = saveReportActions;

  try {
    yield call(callApi, {
      payload: {
        method: HTTP_METHODS.POST,
        body: data,
        reqPath: `${API_ENDPOINTS.uploadReport}/${experimentId}`,
        successAction: saveReportSuccess,
        failureAction: saveReportFailure,
        showPopupSuccessMessage: true,
        showPopupFailureMessage: true,
        token,
        isMultipartFormData: true
      },
    });
  } catch (error) {
    console.error("Error in saving the report: ", error);
    yield put(saveReportFailed({ error }));
  }
}

export function* reportSaga() {
  yield takeEvery(saveReportActions.saveReportInitiated, saveReport);
}
