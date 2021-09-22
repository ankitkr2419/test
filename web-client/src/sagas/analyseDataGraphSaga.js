import { takeEvery, put, call } from "redux-saga/effects";
import { callApi } from "apis/apiHelper";
import { API_ENDPOINTS, HTTP_METHODS } from "appConstants";

import {
  fetchAnalyseDataWithBaselineActions,
  fetchAnalyseDataWithThresholdActions,
} from "actions/analyseDataGraphActions";
import {
  fetchAnalyseDataThresholdFailed,
  fetchAnalyseDataBaselineFailed,
} from "action-creators/analyseDataGraphActionCreators";

export function* fetchAnalyseDataWithThresholdData(actions) {
  const {
    payload: { token, experimentId, target_id, auto_threshold, threshold },
  } = actions;
  const { successAction, failureAction } = fetchAnalyseDataWithThresholdActions;

  try {
    yield call(callApi, {
      payload: {
        method: HTTP_METHODS.POST,
        body: {
          target_id,
          auto_threshold,
          threshold,
        },
        reqPath: `${API_ENDPOINTS.setThreshold}/${experimentId}`,
        successAction: successAction,
        failureAction: failureAction,
        token,
      },
    });
  } catch (error) {
    console.error(
      "Error fetching analyseDataGraph data from api (threshold)",
      error
    );
    yield put(fetchAnalyseDataThresholdFailed({ error }));
  }
}

export function* fetchAnalyseDataWithBaselineData(actions) {
  const {
    payload: { token, experimentId, auto_baseline, start_cycle, end_cycle },
  } = actions;
  const { successAction, failureAction } = fetchAnalyseDataWithBaselineActions;

  try {
    yield call(callApi, {
      payload: {
        method: HTTP_METHODS.POST,
        body: {
          auto_baseline,
          start_cycle,
          end_cycle,
        },
        reqPath: `${API_ENDPOINTS.getBaseline}/${experimentId}`,
        successAction: successAction,
        failureAction: failureAction,
        token,
      },
    });
  } catch (error) {
    console.error(
      "Error fetching analyseDataGraph data from api (baseline)",
      error
    );
    yield put(fetchAnalyseDataBaselineFailed({ error }));
  }
}

export function* analyseDataGraphSaga() {
  yield takeEvery(
    fetchAnalyseDataWithThresholdActions.initiateAction,
    fetchAnalyseDataWithThresholdData
  );
  yield takeEvery(
    fetchAnalyseDataWithBaselineActions.initiateAction,
    fetchAnalyseDataWithBaselineData
  );
}
