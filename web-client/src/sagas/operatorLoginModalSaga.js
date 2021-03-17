import { takeEvery, put, call } from 'redux-saga/effects';
import { callApi } from 'apis/apiHelper'

import {HTTP_METHODS} from 'appConstants';
import operatorLoginModalActions from 'actions/operatorLoginModalActions';
import { operatorLoginFailed } from 'action-creators/operatorLoginModalActionCreators';

export function* operatorLogin(actions){

    // const { email, password, role } = actions.payload;
    const { successAction, failureAction } = operatorLoginModalActions;
    try {
        yield call(callApi, {
            payload: {
                //will be removed in future
                body: {
                    "username": "admin",
                    "password": "admin",
                    "role": "admin"
                },
                // body: {
                //     "username": email,
                //     "password": password,
                //     "role": role
                // },
                reqPath: 'users/admin/validate',
                method: HTTP_METHODS.POST,
                successAction,
                failureAction,
            },
        });
    } catch (error) {
        yield put(operatorLoginFailed(error));
    }
}

export function* operatorLoginModalSaga(){
    yield takeEvery(operatorLoginModalActions.loginInitiated, operatorLogin)
}
