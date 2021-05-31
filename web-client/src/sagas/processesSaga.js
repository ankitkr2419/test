import { takeEvery, call } from "redux-saga/effects";
import { callApi } from "apis/apiHelper";
import {
  aspireDispenseAction,
  heatingAction,
  magnetAction,
  delayAction,
  piercingAction,
  shakingAction,
  tipDiscardAction,
  tipPickupAction,
  processAction,
} from "actions/processesActions";
import {} from "action-creators/processesActionCreators";
import { API_ENDPOINTS, HTTP_METHODS } from "../appConstants";

// //piercing
// export function* piercing(actions) {
//   const { body, recipeID, token } = actions.payload;

//   const { savePiercingSuccess, savePiercingFailed } = piercingAction;
//   try {
//     yield call(callApi, {
//       payload: {
//         body: body,
//         reqPath: `${API_ENDPOINTS.piercing}/${recipeID}`,
//         method: HTTP_METHODS.POST,
//         successAction: savePiercingSuccess,
//         failureAction: savePiercingFailed,
//         showPopupSuccessMessage: true,
//         showPopupFailureMessage: true,
//         token: token,
//       },
//     });
//   } catch (error) {
//     console.log("error while login: ", error);
//     savePiercingFailed(error);
//   }
// }

// //tip-pickup
// export function* tipPickUp(actions) {
//   const { position, recipeID, token } = actions.payload;

//   const { saveTipPickUpSuccess, saveTipPickUpFailed } = tipPickupAction;
//   try {
//     yield call(callApi, {
//       payload: {
//         body: {
//           type: "pickup",
//           position: parseInt(position),
//         },
//         reqPath: `${API_ENDPOINTS.tipOperation}/${recipeID}`,
//         method: HTTP_METHODS.POST,
//         successAction: saveTipPickUpSuccess,
//         failureAction: saveTipPickUpFailed,
//         showPopupSuccessMessage: true,
//         showPopupFailureMessage: true,
//         token: token,
//       },
//     });
//   } catch (error) {
//     console.log("error while login: ", error);
//     saveTipPickUpFailed(error);
//   }
// }

// //aspire-dispense
// export function* aspireDispense(actions) {
//   const { body, recipeID, token } = actions.payload;

//   const { saveAspireDispenseSuccess, saveAspireDispenseFailed } =
//     aspireDispenseAction;
//   try {
//     yield call(callApi, {
//       payload: {
//         body: body,
//         reqPath: `${API_ENDPOINTS.aspireDispense}/${recipeID}`,
//         method: HTTP_METHODS.POST,
//         successAction: saveAspireDispenseSuccess,
//         failureAction: saveAspireDispenseFailed,
//         showPopupSuccessMessage: true,
//         showPopupFailureMessage: true,
//         token: token,
//       },
//     });
//   } catch (error) {
//     console.log("error while login: ", error);
//     saveAspireDispenseFailed(error);
//   }
// }

// //heating
// export function* heating(actions) {
//   const { body, recipeID, token } = actions.payload;

//   const { saveHeatingSuccess, saveHeatingFailed } = heatingAction;
//   try {
//     yield call(callApi, {
//       payload: {
//         body: body,
//         reqPath: `${API_ENDPOINTS.heating}/${recipeID}`,
//         method: HTTP_METHODS.POST,
//         successAction: saveHeatingSuccess,
//         failureAction: saveHeatingFailed,
//         showPopupSuccessMessage: true,
//         showPopupFailureMessage: true,
//         token: token,
//       },
//     });
//   } catch (error) {
//     console.log("error while login: ", error);
//     saveHeatingFailed(error);
//   }
// }

// //tip-discard
// export function* tipDiscard(actions) {
//   const { body, recipeID, token } = actions.payload;

//   const { saveTipDiscardSuccess, saveTipDiscardFailed } = tipDiscardAction;
//   try {
//     yield call(callApi, {
//       payload: {
//         body: body,
//         reqPath: `${API_ENDPOINTS.tipDiscard}/${recipeID}`,
//         method: HTTP_METHODS.POST,
//         successAction: saveTipDiscardSuccess,
//         failureAction: saveTipDiscardFailed,
//         showPopupSuccessMessage: true,
//         showPopupFailureMessage: true,
//         token: token,
//       },
//     });
//   } catch (error) {
//     console.log("error while login: ", error);
//     saveTipDiscardFailed(error);
//   }
// }

// //magnet
// export function* magnet(actions) {
//   const { body, recipeID, token } = actions.payload;

//   const { saveMagnetSuccess, saveMagnetFailed } = magnetAction;
//   try {
//     yield call(callApi, {
//       payload: {
//         body: body,
//         reqPath: `${API_ENDPOINTS.magnet}/${recipeID}`,
//         method: HTTP_METHODS.POST,
//         successAction: saveMagnetSuccess,
//         failureAction: saveMagnetFailed,
//         showPopupSuccessMessage: true,
//         showPopupFailureMessage: true,
//         token: token,
//       },
//     });
//   } catch (error) {
//     console.log("error while login: ", error);
//     saveMagnetFailed(error);
//   }
// }

// //delay
// export function* delay(actions) {
//   const { body, recipeID, token } = actions.payload;

//   // const { saveDelaySuccess, saveDelayFailed } = delayAction;
//   const { saveProcessSuccess, saveProcessFailed } = processAction;
//   try {
//     yield call(callApi, {
//       payload: {
//         body: body,
//         reqPath: `${API_ENDPOINTS.delay}/${recipeID}`,
//         method: HTTP_METHODS.POST,
//         // successAction: saveDelaySuccess,
//         // failureAction: saveDelayFailed,
//         successAction: saveProcessSuccess,
//         failureAction: saveProcessFailed,
//         showPopupSuccessMessage: true,
//         showPopupFailureMessage: true,
//         token: token,
//       },
//     });
//   } catch (error) {
//     console.log("error while login: ", error);
//     // saveDelayFailed(error);
//     saveProcessFailed(error);
//   }
// }

// //shaking
// export function* shaking(actions) {
//   const { body, recipeID, token } = actions.payload;

//   const { saveShakingSuccess, saveShakingFailed } = shakingAction;
//   try {
//     yield call(callApi, {
//       payload: {
//         body: body,
//         reqPath: `${API_ENDPOINTS.shaking}/${recipeID}`,
//         method: HTTP_METHODS.POST,
//         successAction: saveShakingSuccess,
//         failureAction: saveShakingFailed,
//         showPopupSuccessMessage: true,
//         showPopupFailureMessage: true,
//         token: token,
//       },
//     });
//   } catch (error) {
//     console.log("error while login: ", error);
//     saveShakingFailed(error);
//   }
// }

export function* saveProcess(actions) {
  const { body, id, token, api, method } = actions.payload;

  const { saveProcessSuccess, saveProcessFailed } = processAction;
  try {
    yield call(callApi, {
      payload: {
        body: body,
        reqPath: `${api}/${id}`,
        method: method ? method : HTTP_METHODS.POST,
        successAction: saveProcessSuccess,
        failureAction: saveProcessFailed,
        showPopupSuccessMessage: true,
        showPopupFailureMessage: true,
        token: token,
      },
    });
  } catch (error) {
    console.log("error while login: ", error);
    saveProcessFailed(error);
  }
}

export function* processesSaga() {
  // yield takeEvery(piercingAction.savePiercingInitiated, piercing);
  // yield takeEvery(tipPickupAction.saveTipPickUpInitiated, tipPickUp);
  // yield takeEvery(
  //   aspireDispenseAction.saveAspireDispenseInitiated,
  //   aspireDispense
  // );
  // yield takeEvery(shakingAction.saveShakingInitiated, shaking);
  // yield takeEvery(heatingAction.saveHeatingInitiated, heating);
  // yield takeEvery(magnetAction.saveMagnetInitiated, magnet);
  // yield takeEvery(delayAction.saveDelayInitiated, delay);
  // yield takeEvery(tipDiscardAction.saveTipDiscardInitiated, tipDiscard);

  yield takeEvery(processAction.saveProcessInitiated, saveProcess);
}
