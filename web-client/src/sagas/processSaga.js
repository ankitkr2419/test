import { takeEvery, put, call } from "redux-saga/effects";
import { callApi } from "apis/apiHelper";
import { API_ENDPOINTS, HTTP_METHODS } from "appConstants";
import {
    processListActions,
    duplicateProcessActions,
} from "actions/processActions";
import { duplicateProcessFail } from "action-creators/processActionCreators";
export function* fetchProcessList(actions) {
    const {
        payload: { recipeId, token },
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
                token,
            },
        });
    } catch (error) {
        console.error("Error fetching processList", error);
    }
}

export function* duplicateProcess(actions) {
    const {
        payload: { processId, token },
    } = actions;
    const { duplicateProcessSuccess, duplicateProcessFailure } =
        duplicateProcessActions;

    try {
        yield call(callApi, {
            payload: {
                method: HTTP_METHODS.GET,
                body: null,
                reqPath: `${API_ENDPOINTS.duplicateProcess}/${processId}`,
                successAction: duplicateProcessSuccess,
                failureAction: duplicateProcessFailure,
                showPopupSuccessMessage: true,
                showPopupFailureMessage: true,
                token,
            },
        });
    } catch (error) {
        console.error("Error in creating duplicate process", error);
        yield put(duplicateProcessFail({ error }));
    }
}

export function* processSaga() {
    yield takeEvery(processListActions.processListInitiated, fetchProcessList);
    yield takeEvery(
        duplicateProcessActions.duplicateProcessInitiated,
        duplicateProcess
    );
}
