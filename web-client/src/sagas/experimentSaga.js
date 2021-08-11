import { takeEvery, put, call } from 'redux-saga/effects';
import { callApi } from 'apis/apiHelper';
import {
	createExperimentActions,
	listExperimentActions,
} from 'actions/experimentActions';
import {
	createExperimentFailed,
	fetchExperimentsFailed,
} from 'action-creators/experimentActionCreators';
import { HTTP_METHODS } from 'appConstants';

export function* createExperiment(actions) {
	const {
		payload: { body, token },
	} = actions;

	const { successAction, failureAction } = createExperimentActions;

	try {
		yield call(callApi, {
			payload: {
				method: HTTP_METHODS.POST,
				body,
				reqPath: 'experiments',
				successAction,
				failureAction,
				token
			},
		});
	} catch (error) {
		console.error('error in adding experiment ', error);
		yield put(createExperimentFailed(error));
	}
}

export function* fetchExperiments(actions) {
	const { successAction, failureAction } = listExperimentActions;
	try {
		yield call(callApi, {
			payload: {
				body: null,
				reqPath: 'experiments',
				successAction,
				failureAction,
			},
		});
	} catch (error) {
		console.error('error in fetch stages ', error);
		yield put(fetchExperimentsFailed(error));
	}
}

export function* createExperimentSaga() {
	yield takeEvery(createExperimentActions.createAction, createExperiment);
}

export function* fetchExperimentsSaga() {
	yield takeEvery(listExperimentActions.listAction, fetchExperiments);
}
