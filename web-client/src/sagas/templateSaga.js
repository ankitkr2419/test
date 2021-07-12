import { takeEvery, put, call } from "redux-saga/effects";
import { callApi } from "apis/apiHelper";
import {
  createTemplateActions,
  listTemplateActions,
  updateTemplateActions,
  deleteTemplateActions,
} from "actions/templateActions";
import {
  createTemplateFailed,
  fetchTemplatesFailed,
  fetchTemplates as fetchTemplatesActions,
  updateTemplateFailed,
  deleteTemplateFailed,
} from "action-creators/templateActionCreators";
import { HTTP_METHODS } from "appConstants";

export function* createTemplate(actions) {
  const {
    payload: { body, token },
  } = actions;

  const { successAction, failureAction } = createTemplateActions;

  try {
    yield call(callApi, {
      payload: {
        method: HTTP_METHODS.POST,
        body,
        reqPath: "templates",
        successAction,
        failureAction,
        token,
      },
    });
  } catch (error) {
    console.error("error in create template ", error);
    yield put(createTemplateFailed(error));
  }
}

// function* createTemplateSuccess() {
//   //   yield put(fetchTemplatesActions());
// }

export function* fetchTemplates(actions) {
  const {
    payload: { token },
  } = actions;

  const { successAction, failureAction } = listTemplateActions;
  try {
    yield call(callApi, {
      payload: {
        body: null,
        reqPath: "templates",
        successAction,
        failureAction,
        token,
      },
    });
  } catch (error) {
    console.error("error in fetch template ", error);
    yield put(fetchTemplatesFailed(error));
  }
}

export function* updateTemplate(actions) {
  const {
    payload: { templateID, body, token },
  } = actions;

  const { successAction, failureAction } = updateTemplateActions;

  try {
    yield call(callApi, {
      payload: {
        method: HTTP_METHODS.PUT,
        body,
        reqPath: `templates/${templateID}`,
        successAction,
        failureAction,
        token,
      },
    });
  } catch (error) {
    console.error("error while updating template ", error);
    yield put(updateTemplateFailed(error));
  }
}

export function* deleteTemplate(actions) {
  const {
    payload: { templateID, token },
  } = actions;

  const { successAction, failureAction } = deleteTemplateActions;

  try {
    yield call(callApi, {
      payload: {
        method: HTTP_METHODS.DELETE,
        reqPath: `templates/${templateID}`,
        successAction,
        failureAction,
        token,
        showPopupSuccessMessage: true,
        showPopupFailureMessage: true,
      },
    });
  } catch (error) {
    console.error("error while deleting template ", error);
    yield put(deleteTemplateFailed(error));
  }
}

export function* createTemplateSaga() {
  yield takeEvery(createTemplateActions.createAction, createTemplate);
}

export function* createTemplateSuccessSaga() {
  // yield takeEvery(createTemplateActions.successAction, createTemplateSuccess);
}

export function* fetchTemplatesSaga() {
  yield takeEvery(listTemplateActions.listAction, fetchTemplates);
}

export function* updateTemplateSaga() {
  yield takeEvery(updateTemplateActions.updateAction, updateTemplate);
}

export function* deleteTemplateSaga() {
  yield takeEvery(deleteTemplateActions.deleteAction, deleteTemplate);
}
