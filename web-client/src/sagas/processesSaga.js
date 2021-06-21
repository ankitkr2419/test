import { takeEvery, call } from "redux-saga/effects";
import { callApi } from "apis/apiHelper";
import { processAction } from "actions/processesActions";
import { HTTP_METHODS } from "../appConstants";

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
  yield takeEvery(processAction.saveProcessInitiated, saveProcess);
}
