import { restoreDeckActions } from "actions/restoreDeckActions";

export const restoreDeckInitiated = (params) => ({
  type: restoreDeckActions.restoreDeckInitiated,
  payload: {
    params,
  },
});

export const restoreDeckSuccess = () => ({
  type: restoreDeckActions.restoreDeckSuccess,
});

export const restoreDeckFailed = (serverErrors) => ({
  type: restoreDeckActions.restoreDeckFailed,
  payload: {
    serverErrors,
  },
});

export const restoreDeckReset = () => ({
  type: restoreDeckActions.restoreDeckReset,
});
