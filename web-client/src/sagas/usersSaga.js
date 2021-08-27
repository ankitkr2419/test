import { takeEvery, put, call } from "redux-saga/effects";
import { callApi } from "apis/apiHelper";
import { API_ENDPOINTS, HTTP_METHODS } from "appConstants";
import {
  createUserActions,
  deleteUserActions,
  updateUserActions,
} from "actions/usersActions";
import {
  createUserFailed,
  deleteUserFailed,
  updateUserFailed,
} from "action-creators/usersActionCreators";

export function* createUser(actions) {
  const {
    payload: { token, userData },
  } = actions;
  const { createUserSuccess, createUserFailure } = createUserActions;

  try {
    yield call(callApi, {
      payload: {
        method: HTTP_METHODS.POST,
        body: userData,
        reqPath: `${API_ENDPOINTS.users}`,
        successAction: createUserSuccess,
        failureAction: createUserFailure,
        showPopupSuccessMessage: true,
        showPopupFailureMessage: true,
        token,
      },
    });
  } catch (error) {
    console.error("Error create user", error);
    yield put(createUserFailed({ error }));
  }
}

export function* deleteUser(actions) {
  const {
    payload: { token, username },
  } = actions;
  const { deleteUserSuccess, deleteUserFailure } = deleteUserActions;

  try {
    yield call(callApi, {
      payload: {
        method: HTTP_METHODS.DELETE,
        body: null,
        reqPath: `${API_ENDPOINTS.users}/${username}`,
        successAction: deleteUserSuccess,
        failureAction: deleteUserFailure,
        showPopupSuccessMessage: true,
        showPopupFailureMessage: true,
        token,
      },
    });
  } catch (error) {
    console.error("Error deleting user", error);
    yield put(deleteUserFailed({ error }));
  }
}

export function* updateUser(actions) {
  const {
    payload: { token, oldUsername, updatedUserData },
  } = actions;
  const { updateUserSuccess, updateUserFailure } = updateUserActions;

  try {
    yield call(callApi, {
      payload: {
        method: HTTP_METHODS.PUT,
        body: updatedUserData,
        reqPath: `${API_ENDPOINTS.users}/${oldUsername}`,
        successAction: updateUserSuccess,
        failureAction: updateUserFailure,
        showPopupSuccessMessage: true,
        showPopupFailureMessage: true,
        token,
      },
    });
  } catch (error) {
    console.error("Error updating user", error);
    yield put(updateUserFailed({ error }));
  }
}

export function* usersSaga() {
  yield takeEvery(createUserActions.createUserInitiated, createUser);
  yield takeEvery(deleteUserActions.deleteUserInitiated, deleteUser);
  yield takeEvery(updateUserActions.updateUserInitiated, updateUser);
}
