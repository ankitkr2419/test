import { appInfoAction } from "actions/appInfoActions";

export const appInfoInitiated = () => ({
    type: appInfoAction.appInfoInitiated,
});

export const appInfoFailed = (error) => ({
  type: appInfoAction.appInfoFailure,
  payload: {
    error
  },
});
