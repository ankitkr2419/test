import { takeEvery, put } from 'redux-saga/effects';
// Purposefully commented code, will remvove once api is integrated
// import { callApi } from 'apis/apiHelper';
import loginActions from 'actions/loginActions';
import { loginFailed } from 'action-creators/loginActionCreators';
// import { HTTP_METHODS } from '../constants';

export function* login(actions) {
	const {
		payload: { body },
	} = actions;

	// const { successAction, failureAction } = loginActions;

	try {
		if (body.username === 'admin' && body.password === 'admin') {
			yield put({
				type: loginActions.successAction,
			});
		} else {
			yield put(loginFailed(null));
		}

		// TODO enable once backend api is available
		// yield call(callApi, {
		// 	payload: {
		// 		method: HTTP_METHODS.POST,
		// 		body,
		// 		reqPath: 'login',
		// 		successAction,
		// 		failureAction,
		// 	},
		// });
	} catch (error) {
		console.error('error in login ', error);
		yield put(loginFailed(error));
	}
}

export function* loginSaga() {
	yield takeEvery(loginActions.loginInitiated, login);
}
