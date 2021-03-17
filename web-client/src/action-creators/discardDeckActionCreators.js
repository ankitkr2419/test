import { discardDeckActions } from "actions/discardDeckActions";

export const discardDeckInitiated = (params) => ({
  type: discardDeckActions.discardDeckInitiated,
  payload: {
    params
  }
});

export const discardDeckSuccess = () => ({
  type: discardDeckActions.discardDeckSuccess,
  payload: {}
});

export const discardDeckFailed = (serverErrors) => ({
  type: discardDeckActions.discardDeckFailed,
  payload: {
    serverErrors
  }
});
