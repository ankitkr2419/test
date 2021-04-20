import { takeEvery, put, call } from "redux-saga/effects";
import { callApi } from "apis/apiHelper";
import {
  runCleanUpAction,
  pauseCleanUpAction,
  resumeCleanUpAction,
  abortCleanUpAction,
} from "actions/cleanUpActions";
import { API_ENDPOINTS, HTTP_METHODS } from "appConstants";
import {
  runCleanUpActionFailed,
  pauseCleanUpActionFailed,
  resumeCleanUpActionFailed,
  abortCleanUpActionFailed,
} from "action-creators/cleanUpActionCreators";

export function* runUVCleaning(actions) {
  const {
    payload: {
      params: { time, deckName },
    },
  } = actions;
  const { runCleanUpSuccess, runCleanUpFailed } = runCleanUpAction;

  try {
    yield call(callApi, {
      payload: {
        method: HTTP_METHODS.GET,
        body: null,
        reqPath: `${API_ENDPOINTS.cleanUp}/${time}/${deckName}`,
        successAction: runCleanUpSuccess,
        failureAction: runCleanUpFailed,
      },
    });
  } catch (error) {
    console.error("error while starting", error);
    yield put(runCleanUpActionFailed(error));
  }
}

export function* pauseUVCleaning(actions) {
  const {
    payload: {
      params: { deckName },
    },
  } = actions;
  const { pauseCleanUpSuccess, pauseCleanUpFailed } = pauseCleanUpAction;

  try {
    yield call(callApi, {
      payload: {
        method: HTTP_METHODS.GET,
        body: null,
        reqPath: `${API_ENDPOINTS.pause}/${deckName}`,
        successAction: pauseCleanUpSuccess,
        failureAction: pauseCleanUpFailed,
      },
    });
  } catch (error) {
    console.error("error while pausing", error);
    yield put(pauseCleanUpActionFailed(error));
  }
}

export function* resumeUVCleaning(actions) {
  const {
    payload: {
      params: { deckName },
    },
  } = actions;
  const { resumeCleanUpSuccess, resumeCleanUpFailed } = resumeCleanUpAction;

  try {
    yield call(callApi, {
      payload: {
        method: HTTP_METHODS.GET,
        body: null,
        reqPath: `${API_ENDPOINTS.resume}/${deckName}`,
        successAction: resumeCleanUpSuccess,
        failureAction: resumeCleanUpFailed,
      },
    });
  } catch (error) {
    console.error("error while resuming", error);
    yield put(resumeCleanUpActionFailed(error));
  }
}

export function* abortUVCleaning(actions) {
  const {
    payload: {
      params: { deckName },
    },
  } = actions;
  const { abortCleanUpSuccess, abortCleanUpFailed } = abortCleanUpAction;

  try {
    yield call(callApi, {
      payload: {
        method: HTTP_METHODS.GET,
        body: null,
        reqPath: `${API_ENDPOINTS.abort}/${deckName}`,
        successAction: abortCleanUpSuccess,
        failureAction: abortCleanUpFailed,
      },
    });
  } catch (error) {
    console.error("error while aborting", error);
    yield put(abortCleanUpActionFailed(error));
  }
}

export function* cleanUpSaga() {
  yield takeEvery(runCleanUpAction.runCleanUpInitiated, runUVCleaning);
  yield takeEvery(pauseCleanUpAction.pauseCleanUpInitiated, pauseUVCleaning);
  yield takeEvery(resumeCleanUpAction.resumeCleanUpInitiated, resumeUVCleaning);
  yield takeEvery(abortCleanUpAction.abortCleanUpInitiated, abortUVCleaning);
}
