import { takeEvery, put, call } from 'redux-saga/effects';
import { callApi } from 'apis/apiHelper';
import runExperimentActions from 'actions/runExperimentActions';
import { runExperimentFailed } from 'action-creators/runExperimentActionCreators';

export function* runExperiments(actions) {
	const {
		payload: { experimentId },
	} = actions;

	const { successAction, failureAction } = runExperimentActions;
	try {
		yield call(callApi, {
			payload: {
				body: null,
				reqPath: `experiments/${experimentId}/run`,
				successAction,
				failureAction,
			},
		});
	} catch (error) {
		console.error('error while running experiment ', error);
		yield put(runExperimentFailed(error));
	}
}
export function* runExperimentsSaga() {
	yield takeEvery(runExperimentActions.runExperiment, runExperiments);
}
