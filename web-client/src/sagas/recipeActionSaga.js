import { takeEvery, put, call } from "redux-saga/effects";
import { callApi } from "apis/apiHelper";
import {
  runRecipeAction,
  pauseRecipeAction,
  resumeRecipeAction,
  abortRecipeAction,
  recipeListingAction,
} from "actions/recipeActions";
import { API_ENDPOINTS, HTTP_METHODS } from "appConstants";
import {
  runRecipeFailed as runrecipeFailure,
  resumeRecipeFailed as resumeRecipeFailure,
  pauseRecipeFailed as pauseRecipeFailure,
  abortRecipeFailed as abortRecipeFailure,
  recipeListingFailed as recipeListingFailure,
} from "action-creators/recipeActionCreators";

import { toast } from "react-toastify";

export function* runRecipe(actions) {
  const {
    payload: {
      params: { recipeId, deckName },
    },
  } = actions;
  const { runRecipeSuccess, runRecipeFailed } = runRecipeAction;

  try {
    yield call(callApi, {
      payload: {
        method: HTTP_METHODS.GET,
        body: null,
        reqPath: `${API_ENDPOINTS.run}/${recipeId}/${deckName}`,
        successAction: runRecipeSuccess,
        failureAction: runRecipeFailed,
      },
    });
  } catch (error) {
    console.error("Error in running a recipe", error);
    yield put(runrecipeFailure(error));
  }
}

export function* resumeRecipe(actions) {
  const {
    payload: {
      params: { deckName },
    },
  } = actions;
  const { resumeRecipeSuccess, resumeRecipeFailed } = resumeRecipeAction;

  try {
    yield call(callApi, {
      payload: {
        method: HTTP_METHODS.GET,
        body: null,
        reqPath: `${API_ENDPOINTS.resume}/${deckName}`,
        successAction: resumeRecipeSuccess,
        failureAction: resumeRecipeFailed,
      },
    });
  } catch (error) {
    console.error(" Error in resuming a recipe", error);
    yield put(resumeRecipeFailure(error));
  }
}

export function* abortRecipe(actions) {
  const {
    payload: {
      params: { deckName },
    },
  } = actions;
  const { abortRecipeSuccess, abortRecipeFailed } = abortRecipeAction;

  try {
    yield call(callApi, {
      payload: {
        method: HTTP_METHODS.GET,
        body: null,
        reqPath: `${API_ENDPOINTS.abort}/${deckName}`,
        successAction: abortRecipeSuccess,
        failureAction: abortRecipeFailed,
      },
    });
  } catch (error) {
    console.error(" Error in aborting a recipe", error);
    yield put(abortRecipeFailure(error));
  }
}

export function* pauseRecipe(actions) {
  const {
    payload: {
      params: { deckName },
    },
  } = actions;
  const { pauseRecipeSuccess, pauseRecipeFailed } = pauseRecipeAction;

  try {
    yield call(callApi, {
      payload: {
        method: HTTP_METHODS.GET,
        body: null,
        reqPath: `${API_ENDPOINTS.pause}/${deckName}`,
        successAction: pauseRecipeSuccess,
        failureAction: pauseRecipeFailed,
      },
    });
  } catch (error) {
    console.error("Error in pausing a recipe", error);
    yield put(pauseRecipeFailure(error));
  }
}

export function* recipeListing() {
  const { recipeListingSuccess, recipeListingFailed } = recipeListingAction;

  try {
    yield call(callApi, {
      payload: {
        method: HTTP_METHODS.GET,
        body: null,
        reqPath: API_ENDPOINTS.recipeListing,
        successAction: recipeListingSuccess,
        failureAction: recipeListingFailed,
      },
    });
  } catch (error) {
    yield put(toast(error));
    // console.error("Error in fetching the recipes", error);
    yield put(recipeListingFailure(error));
  }
}

export function* recipeActionSaga() {
  yield takeEvery(runRecipeAction.runRecipeInitiated, runRecipe);
  yield takeEvery(abortRecipeAction.abortRecipeInitiated, abortRecipe);
  yield takeEvery(pauseRecipeAction.pauseRecipeInitiated, pauseRecipe);
  yield takeEvery(resumeRecipeAction.resumeRecipeInitiated, resumeRecipe);
  yield takeEvery(recipeListingAction.recipeListingInitiated, recipeListing);
}
