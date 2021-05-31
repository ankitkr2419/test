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
    sequenceActions,
    deleteProcessActions,
} from "actions/processActions";
import {
    duplicateProcessFail,
    fetchProcessDataFail,
    sequenceFail,
    deleteProcessFail,
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

export function* changeSequence(actions) {
    const {
        payload: { recipeId, processList, token },
    } = actions;
    const { sequenceSuccess, sequenceFailure } = sequenceActions;

    try {
        yield call(callApi, {
            payload: {
                method: HTTP_METHODS.POST,
                body: processList,
                reqPath: `${API_ENDPOINTS.rearrangeProcesses}/${recipeId}`,
                successAction: sequenceSuccess,
                failureAction: sequenceFailure,
                showPopupSuccessMessage: true,
                showPopupFailureMessage: true,
                token,
            },
        });
    } catch (error) {
        console.error("Error in changing sequence", error);
        yield put(sequenceFail({ error }));
    }
}

export function* deleteProcess(actions) {
    const {
        payload: { processId, token },
    } = actions;
    const { deleteProcessSuccess, deleteProcessFailure } = deleteProcessActions;

    try {
        yield call(callApi, {
            payload: {
                method: HTTP_METHODS.DELETE,
                body: null,
                reqPath: `${API_ENDPOINTS.processes}/${processId}`,
                successAction: deleteProcessSuccess,
                failureAction: deleteProcessFailure,
                showPopupSuccessMessage: true,
                showPopupFailureMessage: true,
                token,
            },
        });
    } catch (error) {
        console.error("Error in delete process", error);
        yield put(deleteProcessFail({ error }));
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
    yield takeEvery(sequenceActions.sequenceInitiated, changeSequence);
    yield takeEvery(deleteProcessActions.deleteProcessInitiated, deleteProcess);
}
