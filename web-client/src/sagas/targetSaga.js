import { takeEvery, put, call } from 'redux-saga/effects';
import { callApi } from 'apis/apiHelper';
import {
	saveTargetActions,
	listTargetActions,
	listTargetByTemplateIDActions,
} from 'actions/targetActions';

import {
	saveTargetFailed,
	fetchMasterTargetsFailed,
	fetchTargetsByTemplateIDFailed,
} from 'action-creators/targetActionCreators';

export function* fetchMasterTargets(actions) {
	const {
		payload: { token }
	} = actions;
	const { successAction, failureAction } = listTargetActions;
	try {
		yield call(callApi, {
			payload: {
				body: null,
				reqPath: 'targets',
				successAction,
				failureAction,
				token
			},
		});
	} catch (error) {
		console.error('error in fetch targets ', error);
		yield put(fetchMasterTargetsFailed(error));
	}
}

export function* fetchTargetsByTemplateID(actions) {
	const {
		payload: { templateID, token },
	} = actions;
	const { successAction, failureAction } = listTargetByTemplateIDActions;
	try {
		yield call(callApi, {
			payload: {
				body: null,
				reqPath: `templates/${templateID}/targets`,
				successAction,
				failureAction,
				token
			},
		});
	} catch (error) {
		console.error('error in fetch targets by template ID', error);
		yield put(fetchTargetsByTemplateIDFailed(error));
	}
}

export function* saveTarget(actions) {
	const {
		payload: { templateID, body, token },
	} = actions;

	const { successAction, failureAction } = saveTargetActions;

	try {
		yield call(callApi, {
			payload: {
				method: 'POST',
				body,
				reqPath: `templates/${templateID}/targets`,
				successAction,
				failureAction,
				token
			},
		});
	} catch (error) {
		console.error('error in save target ', error);
		yield put(saveTargetFailed(error));
	}
}

export function* fetchMasterTargetsSaga() {
	yield takeEvery(listTargetActions.listAction, fetchMasterTargets);
}

export function* fetchTargetsByTemplateIDSaga() {
	yield takeEvery(listTargetByTemplateIDActions.listAction, fetchTargetsByTemplateID);
}

export function* saveTargetSaga() {
	yield takeEvery(saveTargetActions.saveAction, saveTarget);
}
