import { takeEvery, put, call } from 'redux-saga/effects';
import { callApi } from 'apis/apiHelper';
import { homingActions } from "actions/homingActions";
import { API_ENDPOINTS, HTTP_METHODS } from "appConstants";
import { homingActionFailed } from 'action-creators/homingActionCreators';

export function* homingAction() {
	const { homingActionSuccess, homingActionFailed } = homingActions;
	try {
		yield call(callApi, {
			payload: {
        method: HTTP_METHODS.GET,
				body: null,
				reqPath: API_ENDPOINTS.homing,
				homingActionSuccess,
				homingActionFailed,
			},
		});
	} catch (error) {
		console.error('error while homing confirmation ', error);
		yield put(homingActionFailed(error));
	}
}
export function* homingActionSaga() {
	yield takeEvery(homingActions.homingActionInitiated, homingAction);
}
