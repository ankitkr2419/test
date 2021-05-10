import { takeEvery, put, call } from "redux-saga/effects";
import saveNewRecipeActions, {
  getTipsAndTubesAction,
  getCartridgeAction,
} from "actions/saveNewRecipeActions";
import { API_ENDPOINTS, HTTP_METHODS } from "appConstants";
import { callApi } from "apis/apiHelper";

export function* saveRecipe(actions) {
  //in dev
  //api call
}

export function* getTipsAndTubes(actions) {
  //api call
  const {
    getTipsAndTubesFailure,
    getTipsAndTubesSuccess,
  } = getTipsAndTubesAction;

  try {
    yield call(callApi, {
      payload: {
        method: HTTP_METHODS.GET,
        body: null,
        reqPath: `${API_ENDPOINTS.tipsTubes}/`,
        successAction: getTipsAndTubesSuccess,
        failureAction: getTipsAndTubesFailure,
        // showPopupSuccessMessage: true,
        showPopupFailureMessage: true,
      },
    });
  } catch (error) {
    console.error("error while starting", error);
    yield put(getTipsAndTubesFailure(error));
  }
}

export function* getCartridge(actions) {
  //api call
  const { getCartridgeSuccess, getCartridgeFailure } = getCartridgeAction;

  try {
    yield call(callApi, {
      payload: {
        method: HTTP_METHODS.GET,
        body: null,
        reqPath: `${API_ENDPOINTS.cartridge}`,
        successAction: getCartridgeSuccess,
        failureAction: getCartridgeFailure,
        // showPopupSuccessMessage: true,
        showPopupFailureMessage: true,
      },
    });
  } catch (error) {
    console.error("error while starting", error);
    yield put(getCartridgeFailure(error));
  }
}

export function* saveNewRecipeSaga() {
  yield takeEvery(saveNewRecipeActions.saveRecipe, saveRecipe);
  yield takeEvery(
    getTipsAndTubesAction.getTipsAndTubesInitiated,
    getTipsAndTubes
  );
  yield takeEvery(getCartridgeAction.getCartridgeInitiated, getCartridge);
}
