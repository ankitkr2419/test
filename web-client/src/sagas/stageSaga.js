import { takeEvery, put, call } from 'redux-saga/effects';
import { callApi } from 'apis/apiHelper';
import {
	listStageActions,
	updateStageActions,
} from 'actions/stageActions';
import {
	fetchStagesFailed,
	updateStageFailed,
} from 'action-creators/stageActionCreators';
import { HTTP_METHODS } from 'appConstants';

export function* fetchStages(actions) {
	const {
		payload: { templateID, token },
	} = actions;
	const { successAction, failureAction } = listStageActions;
	try {
		yield call(callApi, {
			payload: {
				body: null,
				reqPath: `templates/${templateID}/stages`,
				successAction,
				failureAction,
				token
			},
		});
	} catch (error) {
		console.error('error in fetch stages ', error);
		yield put(fetchStagesFailed(error));
	}
}

export function* updateStage(actions) {
	const {
		payload: { stageId, body, token },
	} = actions;

	const { successAction, failureAction } = updateStageActions;

	try {
		yield call(callApi, {
			payload: {
				method: HTTP_METHODS.PUT,
				body,
				reqPath: `stages/${stageId}`,
				successAction,
				failureAction,
				token
			},
		});
	} catch (error) {
		console.error('error while updating stage ', error);
		yield put(updateStageFailed(error));
	}
}

export function* fetchStagesSaga() {
	yield takeEvery(listStageActions.listAction, fetchStages);
}

export function* updateStageSaga() {
	yield takeEvery(updateStageActions.updateAction, updateStage);
}
