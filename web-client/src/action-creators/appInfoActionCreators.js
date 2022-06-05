import { appInfoAction, shutDownAction } from "actions/appInfoActions";

export const appInfoInitiated = () => ({
  type: appInfoAction.appInfoInitiated,
});

export const appInfoFailed = (error) => ({
  type: appInfoAction.appInfoFailure,
  payload: {
    error,
  },
});

export const shutdownInitiated = () => ({
  type: shutDownAction.shutdownInitiated,
});

export const shutdownSucess = (params) => ({
  type: shutDownAction.shutdownSuccess,
  payload: {
    ...params,
  },
});

