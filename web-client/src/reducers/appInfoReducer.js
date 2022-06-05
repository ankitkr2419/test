import { fromJS } from "immutable";
import { appInfoAction, shutDownAction } from "actions/appInfoActions";

const appInfoInitialState = fromJS({
  isShutdownLoading: false,
  isLoading: false,
  error: null,
  appInfo: null,
});

export const appInfoReducer = (state = appInfoInitialState, action) => {
  switch (action.type) {
    case appInfoAction.appInfoInitiated:
      return state.merge({
        isShutdownLoading: false,
        isLoading: true,
        error: null,
        appInfo: null,
      });
    case appInfoAction.appInfoSuccess:
      return state.merge({
        isShutdownLoading: false,
        isLoading: false,
        error: false,
        appInfo: action.payload.response,
      });
    case appInfoAction.appInfoFailure:
      return state.merge({
        isShutdownLoading: false,
        isLoading: false,
        error: true,
        appInfo: null,
      });
    case appInfoAction.appInfoReset:
      return state.merge({
        isShutdownLoading: false,
        isLoading: false,
        error: null,
        appInfo: null,
      });
    case shutDownAction.shutdownInitiated:
      return state.merge({
        isShutdownLoading: true,
        isLoading: false,
        error: false,
        appInfo: state.appInfo,
      });
    case shutDownAction.shutdownFailure:
      return state;
    default:
      return state;
  }
};
