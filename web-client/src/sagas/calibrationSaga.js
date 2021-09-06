import { takeEvery, put, call } from "redux-saga/effects";
import { callApi } from "apis/apiHelper";
import { API_ENDPOINTS, HTTP_METHODS } from "appConstants";
import {
  abortActions,
  calibrationActions,
  commonDetailsActions,
  fetchPidDetailsActions,
  heaterActions,
  motorActions,
  pidActions,
  shakerActions,
  updateCalibrationActions,
  updateCommonDetailsActions,
  updateMotorDetailsActions,
  updatePidDetailsActions,
  createTipsTubesActions,
  fetchRtpcrConfigsActions,
  updateRtpcrConfigsActions,
  fetchTECConfigsActions,
  updateTECConfigsActions,
  createCartridgesActions,
  deleteCartridgesActions,
  fetchToleranceActions,
  updateToleranceActions,
  fetchConsumableActions,
  updateConsumableActions,
  addConsumableActions,
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
  shakerFailed,
  heaterFailed,
  abortFailed,
  updateMotorDetailsFailed,
  createTipsOrTubesFailed,
  fetchRtpcrConfigsFailed,
  updateRtpcrConfigsFailed,
  fetchTECConfigsFailed,
  updateTECConfigsFailed,
  createCartridgesFailed,
  deleteCartridgesFailed,
  fetchToleranceFailed,
  updateToleranceFailed,
  fetchConsumableFailed,
  updateConsumableFailed,
  addConsumableFailed,
} from "action-creators/calibrationActionCreators";

export function* shaker(actions) {
  const {
    payload: { token, body, deckName },
  } = actions;
  const { shakerActionSuccess, shakerActionFailed } = shakerActions;

  try {
    yield call(callApi, {
      payload: {
        method: HTTP_METHODS.POST,
        body: body,
        reqPath: `${API_ENDPOINTS.startShaking}/${deckName}`,
        successAction: shakerActionSuccess,
        failureAction: shakerActionFailed,
        showPopupFailureMessage: true,
        token,
      },
    });
  } catch (error) {
    console.error("Error updating shaker configs", error);
    yield put(shakerFailed({ error }));
  }
}

export function* heater(actions) {
  const {
    payload: { token, body, deckName },
  } = actions;
  const { heaterActionSuccess, heaterActionFailed } = heaterActions;

  try {
    yield call(callApi, {
      payload: {
        method: HTTP_METHODS.POST,
        body: body,
        reqPath: `${API_ENDPOINTS.startHeating}/${deckName}`,
        successAction: heaterActionSuccess,
        failureAction: heaterActionFailed,
        showPopupFailureMessage: true,
        token,
      },
    });
  } catch (error) {
    console.error("Error updating heater configs", error);
    yield put(heaterFailed({ error }));
  }
}

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
        showPopupFailureMessage: true,
        token,
      },
    });
  } catch (error) {
    console.error("Error running pid", error);
    yield put(runPidFailed({ error }));
  }
}

export function* abort(actions) {
  const {
    payload: { token, deckName },
  } = actions;

  const { abortActionSuccess, abortActionFailed } = abortActions;

  try {
    yield call(callApi, {
      payload: {
        body: null,
        reqPath: `${API_ENDPOINTS.abort}/${deckName}`,
        successAction: abortActionSuccess,
        failureAction: abortActionFailed,
        showPopupFailureMessage: true,
        token,
      },
    });
  } catch (error) {
    console.error(`Error aborting`, error);
    yield put(abortFailed({ error }));
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

export function* updateMotorDetails(actions) {
  const {
    payload: { token, requestBody },
  } = actions;

  const { updateMotorDetaislSuccess, updateMotorDetaislFailure } =
    updateMotorDetailsActions;

  try {
    yield call(callApi, {
      payload: {
        method: HTTP_METHODS.PUT,
        body: requestBody,
        reqPath: `${API_ENDPOINTS.motor}`,
        successAction: updateMotorDetaislSuccess,
        failureAction: updateMotorDetaislFailure,
        showPopupFailureMessage: true,
        token,
      },
    });
  } catch (error) {
    console.error("Error updating motor configs", error);
    yield put(updateMotorDetailsFailed({ error }));
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

export function* createCartridges(actions) {
  const {
    payload: { token, body },
  } = actions;

  const { successAction, failureAction } = createCartridgesActions;

  try {
    yield call(callApi, {
      payload: {
        method: HTTP_METHODS.POST,
        body: body,
        reqPath: `${API_ENDPOINTS.cartridge}`,
        successAction: successAction,
        failureAction: failureAction,
        showPopupSuccessMessage: true,
        showPopupFailureMessage: true,
        token,
      },
    });
  } catch (error) {
    console.error("Error creating cartridges", error);
    yield put(createCartridgesFailed({ error }));
  }
}

export function* deleteCartridges(actions) {
  const {
    payload: { token, id },
  } = actions;

  const { successAction, failureAction } = deleteCartridgesActions;

  try {
    yield call(callApi, {
      payload: {
        method: HTTP_METHODS.DELETE,
        reqPath: `${API_ENDPOINTS.cartridge}/${id}`,
        successAction: successAction,
        failureAction: failureAction,
        showPopupSuccessMessage: true,
        showPopupFailureMessage: true,
        token,
      },
    });
  } catch (error) {
    console.error("Error deleting cartridges", error);
    yield put(deleteCartridgesFailed({ error }));
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

export function* fetchTolerance(actions) {
  const {
    payload: { token },
  } = actions;

  const { successAction, failureAction } = fetchToleranceActions;

  try {
    yield call(callApi, {
      payload: {
        method: HTTP_METHODS.GET,
        reqPath: `${API_ENDPOINTS.dyes}`,
        successAction: successAction,
        failureAction: failureAction,
        showPopupFailureMessage: true,
        token,
      },
    });
  } catch (error) {
    console.error("Error fetching tolerance configs", error);
    yield put(fetchToleranceFailed(error));
  }
}

export function* updateTolerance(actions) {
  const {
    payload: { token, requestBody },
  } = actions;

  const { successAction, failureAction } = updateToleranceActions;

  try {
    yield call(callApi, {
      payload: {
        method: HTTP_METHODS.PUT,
        body: requestBody,
        reqPath: `${API_ENDPOINTS.dyes}`,
        successAction: successAction,
        failureAction: failureAction,
        showPopupSuccessMessage: true,
        showPopupFailureMessage: true,
        token,
      },
    });
  } catch (error) {
    console.error("Error updating tolerance configs", error);
    yield put(updateToleranceFailed({ error }));
  }
}

export function* fetchConsumable(actions) {
  const {
    payload: { token },
  } = actions;

  const { successAction, failureAction } = fetchConsumableActions;

  try {
    yield call(callApi, {
      payload: {
        method: HTTP_METHODS.GET,
        reqPath: `${API_ENDPOINTS.consumable}`,
        successAction: successAction,
        failureAction: failureAction,
        showPopupFailureMessage: true,
        token,
      },
    });
  } catch (error) {
    console.error("Error fetching consumable distance configs", error);
    yield put(fetchConsumableFailed(error));
  }
}

export function* updateConsumable(actions) {
  const {
    payload: { token, requestBody },
  } = actions;

  const { successAction, failureAction } = updateConsumableActions;

  try {
    yield call(callApi, {
      payload: {
        method: HTTP_METHODS.PUT,
        body: requestBody,
        reqPath: `${API_ENDPOINTS.consumable}`,
        successAction: successAction,
        failureAction: failureAction,
        showPopupSuccessMessage: true,
        showPopupFailureMessage: true,
        token,
      },
    });
  } catch (error) {
    console.error("Error updating consumable distance configs", error);
    yield put(updateConsumableFailed({ error }));
  }
}

export function* addConsumable(actions) {
  const {
    payload: { token, requestBody },
  } = actions;

  const { successAction, failureAction } = addConsumableActions;

  try {
    yield call(callApi, {
      payload: {
        method: HTTP_METHODS.POST,
        body: requestBody,
        reqPath: `${API_ENDPOINTS.consumable}`,
        successAction: successAction,
        failureAction: failureAction,
        showPopupSuccessMessage: true,
        showPopupFailureMessage: true,
        token,
      },
    });
  } catch (error) {
    console.error("Error adding consumable distance configs", error);
    yield put(addConsumableFailed({ error }));
  }
}

export function* calibrationSaga() {
  yield takeEvery(shakerActions.shakerActionInitiated, shaker);
  yield takeEvery(heaterActions.heaterActionInitiated, heater);
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
  yield takeEvery(abortActions.abortActionInitiated, abort);
  yield takeEvery(motorActions.motorActionInitiated, motor);
  yield takeEvery(
    fetchPidDetailsActions.fetchPidActionInitiated,
    fetchPidDetails
  );
  yield takeEvery(
    updatePidDetailsActions.updatePidActionInitiated,
    updatePidDetails
  );
  yield takeEvery(
    updateMotorDetailsActions.updateMotorDetaislInitiated,
    updateMotorDetails
  );
  yield takeEvery(createTipsTubesActions.initiateAction, createTipsOrTubes);
  yield takeEvery(createCartridgesActions.initiateAction, createCartridges);
  yield takeEvery(deleteCartridgesActions.initiateAction, deleteCartridges);
  yield takeEvery(fetchRtpcrConfigsActions.initiateAction, fetchRtpcrConfigs);
  yield takeEvery(updateRtpcrConfigsActions.initiateAction, updateRtpcrConfigs);
  yield takeEvery(fetchTECConfigsActions.initiateAction, fetchTECConfigs);
  yield takeEvery(updateTECConfigsActions.initiateAction, updateTECConfigs);
  yield takeEvery(fetchToleranceActions.initiateAction, fetchTolerance);
  yield takeEvery(updateToleranceActions.initiateAction, updateTolerance);
  yield takeEvery(fetchConsumableActions.initiateAction, fetchConsumable);
  yield takeEvery(updateConsumableActions.initiateAction, updateConsumable);
  yield takeEvery(addConsumableActions.initiateAction, addConsumable);
}
