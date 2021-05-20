import { takeEvery, call } from "redux-saga/effects";
import { callApi } from "apis/apiHelper";
import { piercingAction, tipPickupAction } from "actions/processesActions";
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

export function* processesSaga() {
  yield takeEvery(piercingAction.savePiercingInitiated, piercing);
  yield takeEvery(tipPickupAction.saveTipPickUpInitiated, tipPickUp);
}
