import { whiteLightActions } from "actions/whiteLightActions";

export const whiteLightInitiated = (payload) => ({
  type: whiteLightActions.initiateAction,
  payload,
});

export const whiteLightSuccess = (response) => ({
  type: whiteLightActions.successAction,
  payload: response,
});

export const whiteLightFailed = (error) => ({
  type: whiteLightActions.failureAction,
  payload: error,
});

export const whiteLightReset = () => ({
  type: whiteLightActions.resetAction,
});
