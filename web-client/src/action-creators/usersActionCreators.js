import {
  createUserActions,
  deleteUserActions,
  updateUserActions,
} from "actions/usersActions";

export const createUserInitiated = (token, userData) => ({
  type: createUserActions.createUserInitiated,
  payload: {
    token,
    userData,
  },
});

export const createUserFailed = ({ error }) => ({
  type: createUserActions.createUserFailure,
  payload: { error },
});

export const deleteUserInitiated = (token, username) => ({
  type: deleteUserActions.deleteUserInitiated,
  payload: {
    token,
    username,
  },
});

export const deleteUserFailed = ({ error }) => ({
  type: deleteUserActions.deleteUserFailure,
  payload: { error },
});
