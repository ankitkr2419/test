import { takeEvery, put, call } from 'redux-saga/effects';
import { callApi } from 'apis/apiHelper';
import {
	createExperimentTargetActions,
	listExperimentTargetActions,
} from 'actions/experimentTargetActions';
import {
	createExperimentTargetFailed,
	fetchExperimentTargetsFailed,
} from 'action-creators/experimentTargetActionCreators';
import { HTTP_METHODS } from 'appConstants';

export function* createExperimentTarget(actions) {
	const {
		payload: { body, experimentId, token },
	} = actions;

	const { successAction, failureAction } = createExperimentTargetActions;

	try {
		yield call(callApi, {
			payload: {
				method: HTTP_METHODS.POST,
				body,
				reqPath: `experiments/${experimentId}/targets`,
				successAction,
				failureAction,
				token
			},
		});
	} catch (error) {
		console.error('error in adding experiment ', error);
		yield put(createExperimentTargetFailed(error));
	}
}

export function* fetchExperimentTargets(actions) {
	const {
		payload: { experimentId, token },
	} = actions;

	const { successAction, failureAction } = listExperimentTargetActions;
	try {
		yield call(callApi, {
			payload: {
				body: null,
				reqPath: `experiments/${experimentId}/targets`,
				successAction,
				failureAction,
				token
			},
		});
	} catch (error) {
		console.error('error in fetch stages ', error);
		yield put(fetchExperimentTargetsFailed(error));
	}
}

export function* createExperimentTargetSaga() {
	yield takeEvery(createExperimentTargetActions.createAction, createExperimentTarget);
}

export function* fetchExperimentTargetsSaga() {
	yield takeEvery(listExperimentTargetActions.listAction, fetchExperimentTargets);
}
