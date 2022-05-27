import { whiteLightDeckActions } from "actions/whiteLightActions";

export const whiteLightDeckInitiated = (params) => ({
  type: whiteLightDeckActions.initiateAction,
  payload: { params },
});

export const whiteLightDeckSuccess = (response) => ({
  type: whiteLightDeckActions.successAction,
  payload: response,
});

export const whiteLightDeckFailed = (error) => ({
  type: whiteLightDeckActions.failureAction,
  payload: error,
});

export const whiteLightDeckReset = () => ({
  type: whiteLightDeckActions.resetAction,
});
