import { takeEvery, put, call } from 'redux-saga/effects';
import { callApi } from 'apis/apiHelper';
import { homingActions, deckHomingActions } from "actions/homingActions";
import { API_ENDPOINTS, HTTP_METHODS } from "appConstants";
import { homingActionFailed as homingActionFailure, deckHomingActionFailed as deckHomingActionFailure } from "action-creators/homingActionCreators";

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
		yield put(homingActionFailure(error));
	}
}

export function* deckHoming(actions) {
	const { payload: { params : { deckName } }} = actions;
	const { deckHomingActionSuccess, deckHomingActionFailed } = deckHomingActions;

	try {
		yield call(callApi, {
			payload: {
				method: HTTP_METHODS.GET,
				body: null,
				reqPath: `${API_ENDPOINTS.homing}/${deckName}`,
				deckHomingActionSuccess,
				deckHomingActionFailed,
			},
		});
	} catch(error) {
		console.error("error while deck homing confirmation", error);
		yield put(deckHomingActionFailure(error));
	}
}


export function* homingActionSaga() {
	yield takeEvery(homingActions.homingActionInitiated, homingAction);
	yield takeEvery(deckHomingActions.deckHomingActionInitiated, deckHoming);
}
