import { takeEvery, put, call } from "redux-saga/effects";
import { callApi } from "apis/apiHelper";
import {
  addStepActions,
  listStepActions,
  updateStepActions,
  deleteStepActions,
} from "actions/stepActions";
import {
  addStepFailed,
  fetchStepsFailed,
  updateStepFailed,
  deleteStepFailed,
} from "action-creators/stepActionCreators";
import { HTTP_METHODS } from "../constants";

export function* addStep(actions) {
  const {
    payload: { body },
  } = actions;

  const { successAction, failureAction } = addStepActions;

  try {
    yield call(callApi, {
      payload: {
        method: HTTP_METHODS.POST,
        body: body,
        reqPath: 'step',
        successAction,
        failureAction,
      },
    });
  } catch (error) {
    console.error("error in adding step ", error);
    yield put(addStepFailed(error));
  }
}

export function* fetchSteps(actions) {
  const {
    payload: { stageID },
  } = actions;
  const { successAction, failureAction } = listStepActions;
  try {
    yield call(callApi, {
      payload: {
        body: null,
        reqPath: `steps/${stageID}`,
        successAction,
        failureAction,
      },
    });
  } catch (error) {
    console.error("error in fetch steps ", error);
    yield put(fetchStepsFailed(error));
  }
}

export function* updateStep(actions) {
  const {
    payload: { stepID, body },
  } = actions;

  const { successAction, failureAction } = updateStepActions;

  try {
    yield call(callApi, {
      payload: {
        method: HTTP_METHODS.PUT,
        body: body,
        reqPath: `step/${stepID}`,
        successAction,
        failureAction,
      },
    });
  } catch (error) {
    console.error("error while updating step ", error);
    yield put(updateStepFailed(error));
  }
}

export function* deleteStep(actions) {
  const {
    payload: { stepID },
  } = actions;

  const { successAction, failureAction } = deleteStepActions;

  try {
    yield call(callApi, {
      payload: {
        method: HTTP_METHODS.DELETE,
        reqPath: `step/${stepID}`,
        successAction,
        failureAction,
      },
    });
  } catch (error) {
    console.error("error while deleting step ", error);
    yield put(deleteStepFailed(error));
  }
}

export function* addStageSaga() {
  yield takeEvery(addStepActions.addAction, addStep);
}

export function* fetchStagesSaga() {
  yield takeEvery(listStepActions.listAction, fetchSteps);
}

export function* updateStageSaga() {
  yield takeEvery(updateStepActions.updateAction, updateStep);
}

export function* deleteStageSaga() {
  yield takeEvery(deleteStepActions.deleteAction, deleteStep);
}
