import { takeEvery, put, call } from "redux-saga/effects";
import { callApi } from "apis/apiHelper";
import { API_ENDPOINTS, HTTP_METHODS } from "appConstants";
import { temperatureGraphActions } from "actions/temperatureGraphActions";
import { temperatureApiGraphFailed } from "action-creators/temperatureGraphActionCreators";

export function* updateTemperatureGraph(actions) {
  const {
    payload: { token, experimentId },
  } = actions;

  const { temperatureGraphSuccess, temperatureGraphFailed } =
    temperatureGraphActions;

  try {
    yield call(callApi, {
      payload: {
        method: HTTP_METHODS.GET,
        reqPath: `${API_ENDPOINTS.experiments}/${experimentId}/${API_ENDPOINTS.temperature}`,
        successAction: temperatureGraphSuccess,
        failureAction: temperatureGraphFailed,
        showPopupSuccessMessage: true,
        showPopupFailureMessage: true,
        token,
      },
    });
  } catch (error) {
    console.error("Error updating graph", error);
    yield put(temperatureApiGraphFailed({ error }));
  }
}

export function* temperatureGraphSaga() {
  yield takeEvery(
    temperatureGraphActions.temperatureGraphInitated,
    updateTemperatureGraph
  );
}
