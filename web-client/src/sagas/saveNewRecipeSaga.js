import { takeEvery, put, call } from "redux-saga/effects";
import saveNewRecipeActions, {
  saveNewRecipeAction,
  getTipsAndTubesAction,
  getCartridgeAction,
} from "actions/saveNewRecipeActions";
import { API_ENDPOINTS, HTTP_METHODS } from "appConstants";
import { callApi } from "apis/apiHelper";

export function* saveRecipe(actions) {
  //in dev
  //api call
}

export function* updateRecipe(actions) {
  const { updateRecipeSuccess, updateRecipeFailure } = saveNewRecipeAction;
  const requestBody = actions.payload.params;

  try {
    yield call(callApi, {
      payload: {
        method: HTTP_METHODS.POST,
        body: requestBody,
        reqPath: `${API_ENDPOINTS.saveAndUpdateRecipes}`,
        successAction: updateRecipeSuccess,
        failureAction: updateRecipeFailure,
        // showPopupSuccessMessage: true,
        showPopupFailureMessage: true,
        token: "",
      },
    });
  } catch (error) {
    console.error("error while starting", error);
    yield put(updateRecipeFailure(error));
  }
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
  yield takeEvery(saveNewRecipeAction.updateRecipeInitiated, updateRecipe);
  yield takeEvery(
    getTipsAndTubesAction.getTipsAndTubesInitiated,
    getTipsAndTubes
  );
  yield takeEvery(getCartridgeAction.getCartridgeInitiated, getCartridge);
}
