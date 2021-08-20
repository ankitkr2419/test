import { takeEvery, put, call } from "redux-saga/effects";
import { callApi } from "apis/apiHelper";
import { API_ENDPOINTS, HTTP_METHODS } from "appConstants";
import {
  calibrationActions,
  updateCalibrationActions,
} from "actions/calibrationActions";
import {
  calibrationFailed,
  updateCalibrationFailed,
} from "action-creators/calibrationActionCreators";

export function* fetchCalibrations(actions) {
  const {
    payload: { token },
  } = actions;
  const { calibrationSuccess, calibrationFailure } = calibrationActions;

  try {
    yield call(callApi, {
      payload: {
        method: HTTP_METHODS.GET,
        body: null,
        reqPath: `${API_ENDPOINTS.configs}`,
        successAction: calibrationSuccess,
        failureAction: calibrationFailure,
        showPopupFailureMessage: true,
        token,
      },
    });
  } catch (error) {
    console.error("Error fetching calibrations configs", error);
    yield put(calibrationFailed({ error }));
  }
}

export function* updateCalibrations(actions) {
  const {
    payload: { token, data },
  } = actions;
  const { updateCalibrationSuccess, updateCalibrationFailure } =
    updateCalibrationActions;

  try {
    yield call(callApi, {
      payload: {
        method: HTTP_METHODS.PUT,
        body: { ...data },
        reqPath: `${API_ENDPOINTS.configs}`,
        successAction: updateCalibrationSuccess,
        failureAction: updateCalibrationFailure,
        showPopupSuccessMessage: true,
        showPopupFailureMessage: true,
        token,
      },
    });
  } catch (error) {
    console.error("Error updating calibrations configs", error);
    yield put(updateCalibrationFailed({ error }));
  }
}

export function* calibrationSaga() {
  yield takeEvery(calibrationActions.calibrationInitiated, fetchCalibrations);
  yield takeEvery(
    updateCalibrationActions.updateCalibrationInitiated,
    updateCalibrations
  );
}
