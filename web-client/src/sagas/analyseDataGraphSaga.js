import { takeEvery, put, call } from "redux-saga/effects";
import { callApi } from "apis/apiHelper";
import { API_ENDPOINTS, HTTP_METHODS } from "appConstants";

import {
  fetchAnalyseDataWithBaselineActions,
  fetchAnalyseDataWithThresholdActions,
} from "actions/analyseDataGraphActions";
import { fetchAnalyseDataThresholdFailed } from "action-creators/analyseDataGraphActionCreators";

export function* fetchAnalyseDataWithThresholdData(actions) {
  const {
    payload: { token, experimentId, thresholdDataForApi },
  } = actions;
  const { successAction, failureAction } = fetchAnalyseDataWithThresholdActions;

  try {
    yield call(callApi, {
      payload: {
        method: HTTP_METHODS.POST,
        body: {
          ...thresholdDataForApi,
        },
        reqPath: `${API_ENDPOINTS.setThreshold}/${experimentId}`,
        successAction: successAction,
        failureAction: failureAction,
        token,
      },
    });
  } catch (error) {
    console.error("Error fetching analyseDataGraph data from api", error);
    yield put(fetchAnalyseDataThresholdFailed({ error }));
  }
}

export function* analyseDataGraphSaga() {
  yield takeEvery(
    fetchAnalyseDataWithThresholdActions.initiateAction,
    fetchAnalyseDataWithThresholdData
  );
}
