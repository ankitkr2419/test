import { takeEvery, call } from 'redux-saga/effects';
// Purposefully commented code, will remvove once api is integrated
import { callApi } from 'apis/apiHelper';
import loginActions from 'actions/loginActions';
// import { loginFailed } from 'action-creators/loginActionCreators';
import { DECKNAME, HTTP_METHODS } from '../appConstants';

export function* login(actions) {
	const {
		payload: { body },
	} = actions;

	// const { successAction, failureAction } = loginActions;

	// try {
	// 	if (body.username === 'admin' && body.password === 'admin') {
	// 		yield put({
	// 			type: loginActions.successAction,
	// 		});
	// 	} else {
	// 		yield put(loginFailed(null));
	// 	}
	// } catch (error) {
	// 	console.error('error in login ', error);
	// 	yield put(loginFailed(error));
	// }

	let deckName = actions.payload.body.deckName === DECKNAME.DeckA ? "A" : "B"
	const { successAction, failureAction } = loginActions;
	try {
        yield call(callApi, {
            payload: {
                body: {
                    "username": body.email,
                    "password": body.password,
                    "role": body.role
                },
                // reqPath: 'users/admin/validate',
				reqPath: `login/${deckName}`,
                method: HTTP_METHODS.POST,
                successAction,
                failureAction,
				showPopupSuccessMessage: true,
				showPopupFailureMessage: true
            },
        });
    } catch (error) {
        // yield put(operatorLoginFailed(error));
		console.log('error while login: ', error);
		// yield put(loginFailed(error))
		// yield put(toast.error('failure!'))
    }
}

export function* loginSaga() {
	yield takeEvery(loginActions.loginInitiated, login);
}
