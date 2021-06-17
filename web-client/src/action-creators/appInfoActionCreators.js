import { appInfoAction } from "actions/appInfoActions";

export const appInfoInitiated = () => ({
    type: appInfoAction.appInfoInitiated,
    payload: {},
});

export const appInfoFailed = (error) => ({
  type: appInfoAction.appInfoFailure,
  payload: {
    error
  },
});
