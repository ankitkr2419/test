import { takeEvery, put, call } from 'redux-saga/effects';
import { callApi } from 'apis/apiHelper';
import { runExperimentActions, stopExperimentActions } from 'actions/runExperimentActions';
import { runExperimentFailed } from 'action-creators/runExperimentActionCreators';

export function* runExperiment(actions) {
	const {
		payload: { experimentId, token },
	} = actions;

	const { successAction, failureAction } = runExperimentActions;
	try {
		yield call(callApi, {
			payload: {
				body: null,
				reqPath: `experiments/${experimentId}/run`,
				successAction,
				failureAction,
				token,
			},
		});
	} catch (error) {
		console.error('error while running experiment ', error);
		yield put(runExperimentFailed(error));
	}
}

export function* stopExperiment(actions) {
	const {
		payload: { experimentId, token },
	} = actions;

	const { successAction, failureAction } = stopExperimentActions;
	try {
		yield call(callApi, {
			payload: {
				body: null,
				reqPath: `experiments/${experimentId}/stop`,
				successAction,
				failureAction,
				token,
			},
		});
	} catch (error) {
		console.error('error while stopping experiment ', error);
		yield put(runExperimentFailed(error));
	}
}

export function* runExperimentSaga() {
	yield takeEvery(runExperimentActions.runExperiment, runExperiment);
}

export function* stopExperimentSaga() {
	yield takeEvery(stopExperimentActions.stopExperiment, stopExperiment);
}
