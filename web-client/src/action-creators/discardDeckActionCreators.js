import { discardDeckActions, discardTipActions } from "actions/discardDeckActions";

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

export const discardTipInitiated = (params) => ({
  type: discardTipActions.discardTipInitiated,
  payload: {
    params
  }
});

export const discardTipSuccess = () => ({
  type: discardTipActions.discardTipInitiated,
  payload: {}
});

export const discardTipFailed = (serverErrors) => ({
  type: discardTipActions.discardTipFailed,
  payload: {
    serverErrors
  }
});

export const discardTipInProgress = (discardTipInProgress) => ({
  type: discardTipActions.discardTipInProgress,
  payload: {
    discardTipInProgress
  }
});

export const discardTipInCompleted = (discardTipInCompleted) => ({
  type: discardTipActions.discardTipInCompleted,
  payload: {
    discardTipInCompleted
  }
});
