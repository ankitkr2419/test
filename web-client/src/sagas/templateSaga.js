import { takeEvery, put, call } from "redux-saga/effects";
import { callApi } from "apis/apiHelper";
import { createTemplateActions, listTemplateActions } from "actions/templateActions";
import { createTemplateFailed, fetchTemplatesFailed } from "actionCreators/templateActionCreators";
// WIP Mock testing
// import listJson from 'mockJson/listTemplates'

export function* createTemplate(actions) {
  const {
    payload: { body }
  } = actions;
  
  const { successAction, failureAction } = createTemplateActions;

  try {
    yield call(callApi, {
      payload: {
        method: "POST",
        body: body,
        reqPath: "template",
        successAction,
        failureAction,
      }
    });
  } catch (error) {
    console.error("error in create template ", error);
    yield put(createTemplateFailed(error));
  }
}

export function* fetchTemplates() {
  const { successAction, failureAction } = listTemplateActions;
  try {
    yield call(callApi, {
      payload: {
        body: null,
        reqPath: "templates",
        successAction,
        failureAction,
      }
    });
    // WIP Mock testing
    // yield put({
    //   type: successAction,
    //   payload: listJson,
    // })
  } catch (error) {
    console.error("error in create template ", error);
    yield put(fetchTemplatesFailed(error));
  }
}

export function* createTemplateSaga() {
  yield takeEvery(createTemplateActions.createAction, createTemplate);
}

export function* fetchTemplatesSaga() {
  yield takeEvery(listTemplateActions.listAction, fetchTemplates);
}
