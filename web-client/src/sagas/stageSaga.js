import { takeEvery, put, call } from "redux-saga/effects";
import { callApi } from "apis/apiHelper";
import {
  addStageActions,
  listStageActions,
  updateStageActions,
  deleteStageActions,
} from "actions/stageActions";
import {
  addStageFailed,
  fetchStagesFailed,
  updateStageFailed,
  deleteStageFailed,
} from "action-creators/stageActionCreators";
import { HTTP_METHODS } from "../constants";

export function* addStage(actions) {
  const {
    payload: { body },
  } = actions;

  const { successAction, failureAction } = addStageActions;

  try {
    yield call(callApi, {
      payload: {
        method: HTTP_METHODS.POST,
        body: body,
        reqPath: `stage`,
        successAction,
        failureAction,
      },
    });
  } catch (error) {
    console.error("error in adding stage ", error);
    yield put(addStageFailed(error));
  }
}

export function* fetchStages(actions) {
  const {
    payload: { templateID },
  } = actions;
  const { successAction, failureAction } = listStageActions;
  try {
    yield call(callApi, {
      payload: {
        body: null,
        reqPath: `stages/${templateID}`,
        successAction,
        failureAction,
      },
    });
  } catch (error) {
    console.error("error in fetch stages ", error);
    yield put(fetchStagesFailed(error));
  }
}

export function* updateStage(actions) {
  const {
    payload: { stageID, body },
  } = actions;

  const { successAction, failureAction } = updateStageActions;

  try {
    yield call(callApi, {
      payload: {
        method: HTTP_METHODS.PUT,
        body: body,
        reqPath: `stage/${stageID}`,
        successAction,
        failureAction,
      },
    });
  } catch (error) {
    console.error("error while updating stage ", error);
    yield put(updateStageFailed(error));
  }
}

export function* deleteStage(actions) {
  const {
    payload: { stageID },
  } = actions;

  const { successAction, failureAction } = deleteStageActions;

  try {
    yield call(callApi, {
      payload: {
        method: HTTP_METHODS.DELETE,
        reqPath: `stage/${stageID}`,
        successAction,
        failureAction,
      },
    });
  } catch (error) {
    console.error("error while deleting stage ", error);
    yield put(deleteStageFailed(error));
  }
}

export function* addStageSaga() {
  yield takeEvery(addStageActions.addAction, addStage);
}

export function* fetchStagesSaga() {
  yield takeEvery(listStageActions.listAction, fetchStages);
}

export function* updateStageSaga() {
  yield takeEvery(updateStageActions.updateAction, updateStage);
}

export function* deleteStageSaga() {
  yield takeEvery(deleteStageActions.deleteAction, deleteStage);
}
