import { takeEvery, put, call } from 'redux-saga/effects';
import { callApi } from 'apis/apiHelper';
import activeWellActions from 'actions/activeWellActions';
import { fetchActiveWellsFailed } from 'action-creators/activeWellActionCreators';

export function* fetchActiveWells() {
	const { successAction, failureAction } = activeWellActions;
	try {
		yield call(callApi, {
			payload: {
				body: null,
				reqPath: 'activewells',
				successAction,
				failureAction,
			},
		});
	} catch (error) {
		console.error('error while fetching active wells ', error);
		yield put(fetchActiveWellsFailed(error));
	}
}
export function* fetchActiveWellsSaga() {
	yield takeEvery(activeWellActions.listAction, fetchActiveWells);
}
