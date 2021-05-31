import { takeEvery, call } from "redux-saga/effects";
import { callApi } from "apis/apiHelper";
import {
  aspireDispenseAction,
  heatingAction,
  magnetAction,
  delayAction,
  piercingAction,
  shakingAction,
  tipPickupAction,
} from "actions/processesActions";
import {} from "action-creators/processesActionCreators";
import { API_ENDPOINTS, HTTP_METHODS } from "../appConstants";

export function* piercing(actions) {
  const { type, cartridgeWells, recipeID, token } = actions.payload;

  const { savePiercingSuccess, savePiercingFailed } = piercingAction;
  try {
    yield call(callApi, {
      payload: {
        body: {
          type: `cartridge_${type + 1}`,
          cartridge_wells: cartridgeWells,
        },
        reqPath: `${API_ENDPOINTS.piercing}/${recipeID}`,
        method: HTTP_METHODS.POST,
        successAction: savePiercingSuccess,
        failureAction: savePiercingFailed,
        showPopupSuccessMessage: true,
        showPopupFailureMessage: true,
        token: token,
      },
    });
  } catch (error) {
    console.log("error while login: ", error);
    savePiercingFailed(error);
  }
}

export function* tipPickUp(actions) {
  const { position, recipeID, token } = actions.payload;

  const { saveTipPickUpSuccess, saveTipPickUpFailed } = tipPickupAction;
  try {
    yield call(callApi, {
      payload: {
        body: {
          type: "pickup",
          position: parseInt(position),
        },
        reqPath: `${API_ENDPOINTS.tipOperation}/${recipeID}`,
        method: HTTP_METHODS.POST,
        successAction: saveTipPickUpSuccess,
        failureAction: saveTipPickUpFailed,
        showPopupSuccessMessage: true,
        showPopupFailureMessage: true,
        token: token,
      },
    });
  } catch (error) {
    console.log("error while login: ", error);
    saveTipPickUpFailed(error);
  }
}

export function* aspireDispense(actions) {
  const { body, recipeID, token } = actions.payload;

  const { saveAspireDispenseSuccess, saveAspireDispenseFailed } =
    aspireDispenseAction;
  try {
    yield call(callApi, {
      payload: {
        body: body,
        reqPath: `${API_ENDPOINTS.aspireDispense}/${recipeID}`,
        method: HTTP_METHODS.POST,
        successAction: saveAspireDispenseSuccess,
        failureAction: saveAspireDispenseFailed,
        showPopupSuccessMessage: true,
        showPopupFailureMessage: true,
        token: token,
      },
    });
  } catch (error) {
    console.log("error while login: ", error);
    saveAspireDispenseFailed(error);
  }
}

export function* shaking(actions) {
  const { body, recipeID, token } = actions.payload;

  const { saveShakingSuccess, saveShakingFailed } = shakingAction;

  try {
    yield call(callApi, {
      payload: {
        body: body,
        reqPath: `${API_ENDPOINTS.shaking}/${recipeID}`,
        method: HTTP_METHODS.POST,
        successAction: saveShakingSuccess,
        failureAction: saveShakingFailed,
        showPopupSuccessMessage: true,
        showPopupFailureMessage: true,
        token: token,
      },
    });
  } catch (error) {
    console.log("error while login: ", error);
    saveShakingFailed(error);
  }
}

export function* heating(actions) {
  const { body, recipeID, token } = actions.payload;

  const { saveHeatingSuccess, saveHeatingFailed } = heatingAction;
  try {
    yield call(callApi, {
      payload: {
        body: body,
        reqPath: `${API_ENDPOINTS.heating}/${recipeID}`,
        method: HTTP_METHODS.POST,
        successAction: saveHeatingSuccess,
        failureAction: saveHeatingFailed,
        showPopupSuccessMessage: true,
        showPopupFailureMessage: true,
        token: token,
      },
    });
  } catch (error) {
    console.log("error while login: ", error);
    saveHeatingFailed(error);
  }
}

export function* magnet(actions) {
  const { body, recipeID, token } = actions.payload;

  const { saveMagnetSuccess, saveMagnetFailed } = magnetAction;
  try {
    yield call(callApi, {
      payload: {
        body: body,
        reqPath: `${API_ENDPOINTS.magnet}/${recipeID}`,
        method: HTTP_METHODS.POST,
        successAction: saveMagnetSuccess,
        failureAction: saveMagnetFailed,
        showPopupSuccessMessage: true,
        showPopupFailureMessage: true,
        token: token,
      },
    });
  } catch (error) {
    console.log("error while login: ", error);
    saveMagnetFailed(error);
  }
}
export function* delay(actions) {
  const { body, recipeID, token } = actions.payload;

  const { saveDelaySuccess, saveDelayFailed } = delayAction;
  try {
    yield call(callApi, {
      payload: {
        body: body,
        reqPath: `${API_ENDPOINTS.delay}/${recipeID}`,
        method: HTTP_METHODS.POST,
        successAction: saveDelaySuccess,
        failureAction: saveDelayFailed,
        showPopupSuccessMessage: true,
        showPopupFailureMessage: true,
        token: token,
      },
    });
  } catch (error) {
    console.log("error while login: ", error);
    saveDelayFailed(error);
  }
}

export function* processesSaga() {
  yield takeEvery(piercingAction.savePiercingInitiated, piercing);
  yield takeEvery(tipPickupAction.saveTipPickUpInitiated, tipPickUp);
  yield takeEvery(
    aspireDispenseAction.saveAspireDispenseInitiated,
    aspireDispense
  );
  yield takeEvery(shakingAction.saveShakingInitiated, shaking);
  yield takeEvery(heatingAction.saveHeatingInitiated, heating);
  yield takeEvery(magnetAction.saveMagnetInitiated, magnet);
  yield takeEvery(delayAction.saveDelayInitiated, delay);
}
