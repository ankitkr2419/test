import { takeEvery, put, call } from "redux-saga/effects";
import { callApi } from "apis/apiHelper";
import {
  runRecipeAction,
  pauseRecipeAction,
  resumeRecipeAction,
  abortRecipeAction,
  recipeListingAction,
  stepRunRecipeAction,
  publishRecipeAction,
  deleteRecipeAction,
  duplicateRecipeActions,
} from "actions/recipeActions";
import { API_ENDPOINTS, HTTP_METHODS, DECKNAME } from "appConstants";
import {
  runRecipeFailed as runrecipeFailure,
  resumeRecipeFailed as resumeRecipeFailure,
  pauseRecipeFailed as pauseRecipeFailure,
  abortRecipeFailed as abortRecipeFailure,
  recipeListingFailed as recipeListingFailure,
  duplicateRecipeFail,
} from "action-creators/recipeActionCreators";

import { toast } from "react-toastify";

export function* runRecipe(actions) {
  const {
    payload: {
      params: { recipeId, deckName, token },
    },
  } = actions;
  const { runRecipeSuccess, runRecipeFailed } = runRecipeAction;

  try {
    yield call(callApi, {
      payload: {
        method: HTTP_METHODS.GET,
        body: null,
        reqPath: `${API_ENDPOINTS.run}/${recipeId}/${
          deckName === DECKNAME.DeckA ? "A" : "B"
        }`,
        successAction: runRecipeSuccess,
        failureAction: runRecipeFailed,
        showPopupFailureMessage: true,
        token,
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
      params: { deckName, token },
    },
  } = actions;
  const { resumeRecipeSuccess, resumeRecipeFailed } = resumeRecipeAction;

  try {
    yield call(callApi, {
      payload: {
        method: HTTP_METHODS.GET,
        body: null,
        reqPath: `${API_ENDPOINTS.resume}/${
          deckName === DECKNAME.DeckA ? "A" : "B"
        }`,
        successAction: resumeRecipeSuccess,
        failureAction: resumeRecipeFailed,
        showPopupFailureMessage: true,
        token,
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
      params: { deckName, token },
    },
  } = actions;

  const { abortRecipeSuccess, abortRecipeFailed } = abortRecipeAction;

  try {
    yield call(callApi, {
      payload: {
        method: HTTP_METHODS.GET,
        body: null,
        reqPath: `${API_ENDPOINTS.abort}/${
          deckName === DECKNAME.DeckA ? "A" : "B"
        }`,
        successAction: abortRecipeSuccess,
        failureAction: abortRecipeFailed,
        showPopupFailureMessage: true,
        token,
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
      params: { deckName, token },
    },
  } = actions;
  const { pauseRecipeSuccess, pauseRecipeFailed } = pauseRecipeAction;

  try {
    yield call(callApi, {
      payload: {
        method: HTTP_METHODS.GET,
        body: null,
        reqPath: `${API_ENDPOINTS.pause}/${
          deckName === DECKNAME.DeckA ? "A" : "B"
        }`,
        successAction: pauseRecipeSuccess,
        failureAction: pauseRecipeFailed,
        showPopupFailureMessage: true,
        token,
      },
    });
  } catch (error) {
    console.error("Error in pausing a recipe", error);
    yield put(pauseRecipeFailure(error));
  }
}

export function* recipeListing(actions) {
  const { recipeListingSuccess, recipeListingFailed } = recipeListingAction;

  const token = actions.payload.token;

  try {
    yield call(callApi, {
      payload: {
        method: HTTP_METHODS.GET,
        body: null,
        reqPath: API_ENDPOINTS.recipeListing,
        successAction: recipeListingSuccess,
        failureAction: recipeListingFailed,
        showPopupFailureMessage: true,
        token,
      },
    });
  } catch (error) {
    yield put(toast(error));
    console.error("Error in fetching the recipes", error);
    yield put(recipeListingFailure(error));
  }
}

export function* stepRunRecipe(actions) {
  const {
    payload: {
      params: { recipeId, deckName, token },
    },
  } = actions;
  const { runRecipeSuccess, runRecipeFailed } = runRecipeAction;

  try {
    yield call(callApi, {
      payload: {
        method: HTTP_METHODS.GET,
        body: null,
        reqPath: `${API_ENDPOINTS.stepRun}/${recipeId}/${
          deckName === DECKNAME.DeckA ? "A" : "B"
        }`,
        successAction: runRecipeSuccess,
        failureAction: runRecipeFailed,
        showPopupFailureMessage: true,
        token,
      },
    });
  } catch (error) {
    console.error("Error in running a recipe", error);
    yield put(runrecipeFailure(error));
  }
}

export function* nextStepRunRecipe(actions) {
  const {
    payload: {
      params: { deckName, token },
    },
  } = actions;
  const { runRecipeSuccess, runRecipeFailed } = runRecipeAction;

  try {
    yield call(callApi, {
      payload: {
        method: HTTP_METHODS.GET,
        body: null,
        reqPath: `${API_ENDPOINTS.runNextStep}/${
          deckName === DECKNAME.DeckA ? "A" : "B"
        }`,
        successAction: runRecipeSuccess,
        failureAction: runRecipeFailed,
        showPopupFailureMessage: true,
        token,
      },
    });
  } catch (error) {
    console.error("Error in running a recipe", error);
    yield put(runrecipeFailure(error));
  }
}

export function* publishRecipe(actions) {
  const {
    payload: {
      params: { recipeId, token, isPublished },
    },
  } = actions;
  const { publishRecipeSuccess, publishRecipeFailed } = publishRecipeAction;

  /**isPublished: true means we should call unpublish api
   * isPublished: false means we should call publish api
   */

  try {
    yield call(callApi, {
      payload: {
        method: HTTP_METHODS.POST,
        body: null,
        reqPath: `${API_ENDPOINTS.recipeListing}/${recipeId}/${
          isPublished ? "unpublish" : "publish"
        }`,
        successAction: publishRecipeSuccess,
        failureAction: publishRecipeFailed,
        showPopupSuccessMessage: true,
        showPopupFailureMessage: true,
        token,
      },
    });
  } catch (error) {
    console.error("Error in publish recipe", error);
  }
}

export function* deleteRecipe(actions) {
  const {
    payload: {
      params: { recipeId, token },
    },
  } = actions;
  const { deleteRecipeSuccess, deleteRecipeFailure } = deleteRecipeAction;

  try {
    yield call(callApi, {
      payload: {
        method: HTTP_METHODS.DELETE,
        body: null,
        reqPath: `${API_ENDPOINTS.recipeListing}/${recipeId}`,
        successAction: deleteRecipeSuccess,
        failureAction: deleteRecipeFailure,
        showPopupSuccessMessage: true,
        showPopupFailureMessage: true,
        token,
      },
    });
  } catch (error) {
    console.error("Error in publish recipe", error);
  }
}

export function* duplicateRecipe(actions) {
  const {
    payload: { recipeId, token, recipeName },
  } = actions;
  const { duplicateRecipeSuccess, duplicateRecipeFailure } =
    duplicateRecipeActions;

  try {
    yield call(callApi, {
      payload: {
        method: HTTP_METHODS.POST,
        body: { recipeName },
        reqPath: `${API_ENDPOINTS.duplicateRecipe}/${recipeId}`,
        successAction: duplicateRecipeSuccess,
        failureAction: duplicateRecipeFailure,
        showPopupSuccessMessage: true,
        showPopupFailureMessage: true,
        token,
      },
    });
  } catch (error) {
    console.error("Error in creating duplicate recipe", error);
    yield put(duplicateRecipeFail({ error }));
  }
}

export function* recipeActionSaga() {
  yield takeEvery(runRecipeAction.runRecipeInitiated, runRecipe);
  yield takeEvery(abortRecipeAction.abortRecipeInitiated, abortRecipe);
  yield takeEvery(pauseRecipeAction.pauseRecipeInitiated, pauseRecipe);
  yield takeEvery(resumeRecipeAction.resumeRecipeInitiated, resumeRecipe);
  yield takeEvery(recipeListingAction.recipeListingInitiated, recipeListing);
  yield takeEvery(stepRunRecipeAction.stepRunRecipeInitiated, stepRunRecipe);
  yield takeEvery(
    stepRunRecipeAction.nextStepRunRecipeInitiated,
    nextStepRunRecipe
  );
  yield takeEvery(publishRecipeAction.publishRecipeInitiated, publishRecipe);
  yield takeEvery(deleteRecipeAction.deleteRecipeInitiated, deleteRecipe);
  yield takeEvery(
    duplicateRecipeActions.duplicateRecipeInitiated,
    duplicateRecipe
  );
}
