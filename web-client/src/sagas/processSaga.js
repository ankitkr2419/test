import { takeEvery, put, call } from "redux-saga/effects";
import { callApi } from "apis/apiHelper";
import {
    API_ENDPOINTS,
    HTTP_METHODS,
    SELECT_PROCESS_PROPS,
} from "appConstants";
import {
    processListActions,
    duplicateProcessActions,
    fetchProcessDataActions,
} from "actions/processActions";
import {
    duplicateProcessFail,
    fetchProcessDataFail,
} from "action-creators/processActionCreators";
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

/**
 * Generalized method to get process data | apiEndpoint is dynamic depend upon processType
 */
export function* fetchProcessData(actions) {
    const {
        payload: { processId, type, token },
    } = actions;
    const { fetchProcessDataSuccess, fetchProcessDataFailure } =
        fetchProcessDataActions;

    //get apiEndPoint depend on processType
    const obj = SELECT_PROCESS_PROPS.find((obj) => obj.processType === type);
    const apiEndPoint = obj?.apiEndpoint;
    if (!apiEndPoint) {
        console.error("api endpoint not found!");
        return;
    }

    try {
        yield call(callApi, {
            payload: {
                method: HTTP_METHODS.GET,
                body: null,
                reqPath: `${apiEndPoint}/${processId}`,
                successAction: fetchProcessDataSuccess,
                failureAction: fetchProcessDataFailure,
                showPopupFailureMessage: true,
                token,
            },
        });
    } catch (error) {
        console.error("Error in fetching process data", error);
        yield put(fetchProcessDataFail({ error }));
    }
}

export function* processSaga() {
    yield takeEvery(processListActions.processListInitiated, fetchProcessList);
    yield takeEvery(
        duplicateProcessActions.duplicateProcessInitiated,
        duplicateProcess
    );
    yield takeEvery(
        fetchProcessDataActions.fetchProcessDataInitiated,
        fetchProcessData
    );
}
