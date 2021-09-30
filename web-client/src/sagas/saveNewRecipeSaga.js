import { takeEvery, put, call } from "redux-saga/effects";
import saveNewRecipeActions, {
  saveNewRecipeAction,
  getTipsAndTubesAction,
  getCartridgeAction,
  getTipsAction,
  getTubesAction,
  getCartridge1Action,
  getCartridge2Action,
} from "actions/saveNewRecipeActions";
import {
  API_ENDPOINTS,
  CARTRIDGE_1,
  CARTRIDGE_2,
  HTTP_METHODS,
} from "appConstants";
import { callApi } from "apis/apiHelper";
import {
  getCartridge1ActionFailure,
  getCartridge2ActionFailure,
} from "action-creators/saveNewRecipeActionCreators";

export function* saveRecipe(actions) {
  //in dev
  //api call
}

export function* updateRecipe(actions) {
  const { updateRecipeSuccess, updateRecipeFailure } = saveNewRecipeAction;
  const requestBody = actions.payload.requestBody;
  const token = actions.payload.token;

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
        token: token,
      },
    });
  } catch (error) {
    console.error("error while starting", error);
    yield put(updateRecipeFailure(error));
  }
}

export function* getTipsAndTubes(actions) {
  //api call
  const { getTipsAndTubesFailure, getTipsAndTubesSuccess } =
    getTipsAndTubesAction;
  const token = actions.payload.token;

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
        token: token,
      },
    });
  } catch (error) {
    console.error("error while starting", error);
    yield put(getTipsAndTubesFailure(error));
  }
}

export function* getCartridge(actions) {
  //api call
  const { token, id, type } = actions.payload;
  const { getCartridgeSuccess, getCartridgeFailure } = getCartridgeAction;

  let successAction = getCartridgeSuccess;
  let failureAction = getCartridgeFailure;
  let reqPath = `${API_ENDPOINTS.cartridges}`;

  if (id) {
    if (type === CARTRIDGE_1) {
      const { getCartridge1Success, getCartridge1Failure } =
        getCartridge1Action;

      successAction = getCartridge1Success;
      failureAction = getCartridge1Failure;
    } else {
      const { getCartridge2Success, getCartridge2Failure } =
        getCartridge2Action;

      successAction = getCartridge2Success;
      failureAction = getCartridge2Failure;
    }
    reqPath = `${API_ENDPOINTS.cartridge}/${id}`;
  }

  try {
    yield call(callApi, {
      payload: {
        method: HTTP_METHODS.GET,
        body: null,
        reqPath: reqPath,
        successAction: successAction,
        failureAction: failureAction,
        // showPopupSuccessMessage: true,
        showPopupFailureMessage: true,
        token: token,
      },
    });
  } catch (error) {
    console.error("error while starting", error);
    if (type && type === CARTRIDGE_1) {
      yield put(getCartridge1ActionFailure(error));
    } else if (type && type === CARTRIDGE_2) {
      yield put(getCartridge2ActionFailure(error));
    } else {
      yield put(getCartridgeFailure(error));
    }
  }
}

// for tubes
export function* getTubes(actions) {
  const { getTubesSuccess, getTubesFailure } = getTubesAction;
  const token = actions.payload.token;

  try {
    yield call(callApi, {
      payload: {
        method: HTTP_METHODS.GET,
        body: null,
        reqPath: `${API_ENDPOINTS.tipsTubes}/${API_ENDPOINTS.tubes}`,
        successAction: getTubesSuccess,
        failureAction: getTubesFailure,
        // showPopupSuccessMessage: true,
        showPopupFailureMessage: true,
        token: token,
      },
    });
  } catch (error) {
    console.error("error while starting", error);
    yield put(getTubesFailure(error));
  }
}

//for tips
export function* getTips(actions) {
  const { getTipsSuccess, getTipsFailure } = getTipsAction;
  const token = actions.payload.token;

  try {
    yield call(callApi, {
      payload: {
        method: HTTP_METHODS.GET,
        body: null,
        reqPath: `${API_ENDPOINTS.tipsTubes}/${API_ENDPOINTS.tips}`,
        successAction: getTipsSuccess,
        failureAction: getTipsFailure,
        // showPopupSuccessMessage: true,
        showPopupFailureMessage: true,
        token: token,
      },
    });
  } catch (error) {
    console.error("error while starting", error);
    yield put(getTipsFailure(error));
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
  yield takeEvery(getCartridge1Action.getCartridge1Initiated, getCartridge);
  yield takeEvery(getCartridge2Action.getCartridge2Initiated, getCartridge);
  yield takeEvery(getTipsAction.getTipsInitiated, getTips);
  yield takeEvery(getTubesAction.getTubesInitiated, getTubes);
}
