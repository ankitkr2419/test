import { takeEvery, call } from "redux-saga/effects";
import { callApi } from "apis/apiHelper";
import { whiteLightActions } from "actions/whiteLightActions";
import {} from "action-creators/whiteLightActionCreators";
import { API_ENDPOINTS } from "appConstants";

export function* whiteLight(actions) {
  const { successAction, failureAction } = whiteLightActions;
  try {
    yield call(callApi, {
      payload: {
        reqPath: `${API_ENDPOINTS.whiteLight}`,
        successAction,
        failureAction,
        showPopupSuccessMessage: true,
        showPopupFailureMessage: true,
      },
    });
  } catch (error) {
    console.log("Error while turning on white-light: ", error);
  }
}

export function* whiteLightSaga() {
  yield takeEvery(whiteLightActions.initiateAction, whiteLight);
}
