import { takeEvery, put, call } from "redux-saga/effects";
import { callApi } from "apis/apiHelper";
import { API_ENDPOINTS, HTTP_METHODS } from "appConstants";
import {
  calibrationActions,
  commonDetailsActions,
  motorActions,
  pidActions,
  updateCalibrationActions,
  updateCommonDetailsActions,
} from "actions/calibrationActions";
import {
  calibrationFailed,
  updateCalibrationFailed,
  runPidFailed,
  motorFailed,
  commonDetailsFailed,
  updateCommonDetailsFailed,
} from "action-creators/calibrationActionCreators";

export function* fetchCommonDetails(actions) {
  const {
    payload: { token },
  } = actions;
  const { commonDetailsSuccess, commonDetailsFailure } = commonDetailsActions;

  try {
    yield call(callApi, {
      payload: {
        method: HTTP_METHODS.GET,
        body: null,
        reqPath: `${API_ENDPOINTS.configs}/common`,
        successAction: commonDetailsSuccess,
        failureAction: commonDetailsFailure,
        showPopupFailureMessage: true,
        token,
      },
    });
  } catch (error) {
    console.error("Error fetching calibrations configs", error);
    yield put(commonDetailsFailed({ error }));
  }
}

export function* updateCommonDetails(actions) {
  const {
    payload: { token, data },
  } = actions;
  const { updateCommonDetaislSuccess, updateCommonDetaislFailure } =
    updateCommonDetailsActions;

  try {
    yield call(callApi, {
      payload: {
        method: HTTP_METHODS.PUT,
        body: { ...data },
        reqPath: `${API_ENDPOINTS.configs}/common`,
        successAction: updateCommonDetaislSuccess,
        failureAction: updateCommonDetaislFailure,
        showPopupSuccessMessage: true,
        showPopupFailureMessage: true,
        token,
      },
    });
  } catch (error) {
    console.error("Error updating calibrations configs", error);
    yield put(updateCommonDetailsFailed({ error }));
  }
}

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

export function* pidStart(actions) {
  const {
    payload: { token, deckName },
  } = actions;
  const { pidActionSuccess, pidActionFailure } = pidActions;

  try {
    yield call(callApi, {
      payload: {
        body: null,
        reqPath: `${API_ENDPOINTS.pidCalibration}/${deckName}`,
        successAction: pidActionSuccess,
        failureAction: pidActionFailure,
        token,
      },
    });
  } catch (error) {
    console.error("Error running pid", error);
    yield put(runPidFailed({ error }));
  }
}

export function* pidAbort(actions) {
  const {
    payload: { token, deckName },
  } = actions;
  const { pidAbortActionSuccess, pidAbortActionFailure } = pidActions;

  try {
    yield call(callApi, {
      payload: {
        body: null,
        reqPath: `${API_ENDPOINTS.abort}/${deckName}`,
        successAction: pidAbortActionSuccess,
        failureAction: pidAbortActionFailure,
        showPopupSuccessMessage: true,
        showPopupFailureMessage: true,
        token,
      },
    });
  } catch (error) {
    console.error("Error aborting pid", error);
    yield put(runPidFailed({ error }));
  }
}

export function* motor(actions) {
  const {
    payload: { token, body },
  } = actions;
  const { motorActionSuccess, motorActionFailure } = motorActions;

  try {
    yield call(callApi, {
      payload: {
        method: HTTP_METHODS.POST,
        body: body,
        reqPath: `${API_ENDPOINTS.manual}`,
        successAction: motorActionSuccess,
        failureAction: motorActionFailure,
        showPopupSuccessMessage: true,
        showPopupFailureMessage: true,
        token,
      },
    });
  } catch (error) {
    console.error("Error updating calibrations configs", error);
    yield put(motorFailed({ error }));
  }
}

export function* calibrationSaga() {
  yield takeEvery(
    commonDetailsActions.commonDetailsInitiated,
    fetchCommonDetails
  );
  yield takeEvery(
    updateCommonDetailsActions.updateCommonDetaislInitiated,
    updateCommonDetails
  );
  yield takeEvery(calibrationActions.calibrationInitiated, fetchCalibrations);
  yield takeEvery(
    updateCalibrationActions.updateCalibrationInitiated,
    updateCalibrations
  );
  yield takeEvery(pidActions.pidActionInitiated, pidStart);
  yield takeEvery(pidActions.pidAbortActionInitiated, pidAbort);
  yield takeEvery(motorActions.motorActionInitiated, motor);
}
