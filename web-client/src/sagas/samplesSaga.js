import { takeEvery, put, call } from 'redux-saga/effects';
import { callApi } from 'apis/apiHelper';
import {
	listSampleActions,
} from 'actions/sampleActions';
import {
	fetchSamplesFailed,
} from 'action-creators/sampleActionCreators';

export function* fetchSamples(actions) {
	const {
		payload: { searchText },
	} = actions;

	const { successAction, failureAction } = listSampleActions;
	try {
		yield call(callApi, {
			payload: {
				body: null,
				reqPath: `samples/${searchText}`,
				successAction,
				failureAction,
			},
		});
		// yield put({
		// 	type: listSampleActions.successAction,
		// 	payload: {
		// 		response : [{
		// 			id: 'uuid',
		// 			name: 'test string 2',
		// 		},
		// 		{
		// 			id: 'uuid1',
		// 			name: 'test string 1',
		// 		}],
		// 	},
		// });
	} catch (error) {
		console.error('error in fetch stages ', error);
		yield put(fetchSamplesFailed(error));
	}
}

export function* fetchSamplesSaga() {
	yield takeEvery(listSampleActions.listAction, fetchSamples);
}
