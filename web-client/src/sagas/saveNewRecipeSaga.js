import { takeEvery, put, call } from "redux-saga/effects";
import saveNewRecipeActions, {
  getRecipeDetailsAction,
} from "actions/saveNewRecipeActions";
import { API_ENDPOINTS, HTTP_METHODS } from "appConstants";
import { callApi } from 'apis/apiHelper';

export function* saveRecipe(actions) {
  //in dev
  //api call
}

export function* getRecipeDetails(actions) {
  //api call
  const {
    getRecipeDetailsSuccess,
    getRecipeDetailsFailure,
  } = getRecipeDetailsAction;
 
  try {
    yield call(callApi, {
      payload: {
        method: HTTP_METHODS.GET,
        body: null,
        reqPath: `${API_ENDPOINTS.tipsTubes}/`,
        successAction: getRecipeDetailsSuccess,
        failureAction: getRecipeDetailsFailure,
        // showPopupSuccessMessage: true,
        showPopupFailureMessage: true,
      },
    });
  } catch (error) {
    console.error("error while starting", error);
    yield put(getRecipeDetailsFailure(error));
  }
}

export function* saveNewRecipeSaga() {
  yield takeEvery(saveNewRecipeActions.saveRecipe, saveRecipe);
  yield takeEvery(
    getRecipeDetailsAction.getRecipeDetailsInitiated,
    getRecipeDetails
  );
}
