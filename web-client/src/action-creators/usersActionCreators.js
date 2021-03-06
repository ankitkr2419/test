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

export const updateUserInitiated = (token, oldUsername, updatedUserData) => ({
  type: updateUserActions.updateUserInitiated,
  payload: {
    token,
    oldUsername,
    updatedUserData,
  },
});

export const updateUserFailed = ({ error }) => ({
  type: updateUserActions.updateUserFailure,
  payload: { error },
});
