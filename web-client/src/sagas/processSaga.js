import { takeEvery, put, call } from "redux-saga/effects";
import { callApi } from "apis/apiHelper";
import { API_ENDPOINTS, HTTP_METHODS } from "appConstants";
import { processListActions } from "actions/processActions";

export function* fetchProcessList(actions) {
    const {
        payload: { recipeId },
    } = actions;
    const { processListSuccess, processListFailure } = processListActions;

    try {
        yield call(callApi, {
            payload: {
                method: HTTP_METHODS.GET,
                body: null,
                reqPath: `${API_ENDPOINTS.recipe}/${recipeId}/processes`,
                successAction: processListSuccess,
                failureAction: processListFailure,
                showPopupFailureMessage: true,
            },
        });
    } catch (error) {
        console.error("Error fetching processList", error);
    }
}

export function* processSaga() {
    yield takeEvery(processListActions.processListInitiated, fetchProcessList);
}
