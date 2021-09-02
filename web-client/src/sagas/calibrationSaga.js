import { takeEvery, put, call } from "redux-saga/effects";
import { callApi } from "apis/apiHelper";
import { API_ENDPOINTS, HTTP_METHODS } from "appConstants";
import {
  calibrationActions,
  commonDetailsActions,
  fetchPidDetailsActions,
  motorActions,
  pidActions,
  updateCalibrationActions,
  updateCommonDetailsActions,
  updatePidDetailsActions,
  createTipsTubesActions,
  fetchRtpcrConfigsActions,
  updateRtpcrConfigsActions,
  fetchTECConfigsActions,
  updateTECConfigsActions,
} from "actions/calibrationActions";
import {
  calibrationFailed,
  updateCalibrationFailed,
  runPidFailed,
  motorFailed,
  commonDetailsFailed,
  updateCommonDetailsFailed,
  fetchPidFailed,
  updatePidFailed,
  createTipsOrTubesFailed,
  fetchRtpcrConfigsFailed,
  updateRtpcrConfigsFailed,
  fetchTECConfigsFailed,
  updateTECConfigsFailed,
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

export function* fetchPidDetails(actions) {
  const {
    payload: { token },
  } = actions;
  const { fetchPidActionSuccess, fetchPidActionFailed } =
    fetchPidDetailsActions;

  try {
    yield call(callApi, {
      payload: {
        method: HTTP_METHODS.GET,
        body: null,
        reqPath: `${API_ENDPOINTS.pidUpdate}`,
        successAction: fetchPidActionSuccess,
        failureAction: fetchPidActionFailed,
        showPopupFailureMessage: true,
        token,
      },
    });
  } catch (error) {
    console.error("Error fetching pid configs", error);
    yield put(fetchPidFailed({ error }));
  }
}

export function* updatePidDetails(actions) {
  const {
    payload: { token, body },
  } = actions;

  const { updatePidActionSuccess, updatePidActionFailed } =
    updatePidDetailsActions;

  try {
    yield call(callApi, {
      payload: {
        method: HTTP_METHODS.PUT,
        body: body,
        reqPath: `${API_ENDPOINTS.pidUpdate}`,
        successAction: updatePidActionSuccess,
        failureAction: updatePidActionFailed,
        showPopupSuccessMessage: true,
        showPopupFailureMessage: true,
        token,
      },
    });
  } catch (error) {
    console.error("Error updating pid configs", error);
    yield put(updatePidFailed({ error }));
  }
}

export function* createTipsOrTubes(actions) {
  const {
    payload: { token, body },
  } = actions;

  const { successAction, failureAction } = createTipsTubesActions;

  try {
    yield call(callApi, {
      payload: {
        method: HTTP_METHODS.POST,
        body: body,
        reqPath: `${API_ENDPOINTS.tipTube}`,
        successAction: successAction,
        failureAction: failureAction,
        showPopupSuccessMessage: true,
        showPopupFailureMessage: true,
        token,
      },
    });
  } catch (error) {
    console.error("Error creating tips or tubes", error);
    yield put(createTipsOrTubesFailed({ error }));
  }
}

export function* fetchRtpcrConfigs(actions) {
  const {
    payload: { token },
  } = actions;

  const { successAction, failureAction } = fetchRtpcrConfigsActions;

  try {
    yield call(callApi, {
      payload: {
        method: HTTP_METHODS.GET,
        body: null,
        reqPath: `${API_ENDPOINTS.rtpcrConfigs}`,
        successAction: successAction,
        failureAction: failureAction,
        showPopupFailureMessage: true,
        token,
      },
    });
  } catch (error) {
    console.error("Error fetching rtpcr configs", error);
    yield put(fetchRtpcrConfigsFailed({ error }));
  }
}

export function* updateRtpcrConfigs(actions) {
  const {
    payload: { token, requestBody },
  } = actions;

  const { successAction, failureAction } = updateRtpcrConfigsActions;

  try {
    yield call(callApi, {
      payload: {
        method: HTTP_METHODS.PUT,
        body: requestBody,
        reqPath: `${API_ENDPOINTS.rtpcrConfigs}`,
        successAction: successAction,
        failureAction: failureAction,
        showPopupSuccessMessage: true,
        showPopupFailureMessage: true,
        token,
      },
    });
  } catch (error) {
    console.error("Error updating rtpcr configs", error);
    yield put(updateRtpcrConfigsFailed({ error }));
  }
}

export function* fetchTECConfigs(actions) {
  const {
    payload: { token },
  } = actions;

  const { successAction, failureAction } = fetchTECConfigsActions;

  try {
    yield call(callApi, {
      payload: {
        method: HTTP_METHODS.GET,
        body: null,
        reqPath: `${API_ENDPOINTS.tecConfigs}`,
        successAction: successAction,
        failureAction: failureAction,
        showPopupFailureMessage: true,
        token,
      },
    });
  } catch (error) {
    console.error("Error fetching TEC configs", error);
    yield put(fetchTECConfigsFailed({ error }));
  }
}

export function* updateTECConfigs(actions) {
  const {
    payload: { token, requestBody },
  } = actions;

  const { successAction, failureAction } = updateTECConfigsActions;

  try {
    yield call(callApi, {
      payload: {
        method: HTTP_METHODS.PUT,
        body: requestBody,
        reqPath: `${API_ENDPOINTS.tecConfigs}`,
        successAction: successAction,
        failureAction: failureAction,
        showPopupSuccessMessage: true,
        showPopupFailureMessage: true,
        token,
      },
    });
  } catch (error) {
    console.error("Error updating TEC configs", error);
    yield put(updateTECConfigsFailed({ error }));
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
  yield takeEvery(
    fetchPidDetailsActions.fetchPidActionInitiated,
    fetchPidDetails
  );
  yield takeEvery(
    updatePidDetailsActions.updatePidActionInitiated,
    updatePidDetails
  );
  yield takeEvery(createTipsTubesActions.initiateAction, createTipsOrTubes);
  yield takeEvery(fetchRtpcrConfigsActions.initiateAction, fetchRtpcrConfigs);
  yield takeEvery(updateRtpcrConfigsActions.initiateAction, updateRtpcrConfigs);
  yield takeEvery(fetchTECConfigsActions.initiateAction, fetchTECConfigs);
  yield takeEvery(updateTECConfigsActions.initiateAction, updateTECConfigs);
}
