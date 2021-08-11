import { fromJS } from "immutable";
import { appInfoAction } from "actions/appInfoActions";

const appInfoInitialState = fromJS({
    isLoading: false,
    error: null,
    appInfo: null,
});

export const appInfoReducer = (state = appInfoInitialState, action) => {
    switch (action.type) {
        case appInfoAction.appInfoInitiated:
            return state.merge({
                isLoading: true,
                error: null,
                appInfo: null,
            });
        case appInfoAction.appInfoSuccess:
            return state.merge({
                isLoading: false,
                error: false,
                appInfo: action.payload.response,
            });
        case appInfoAction.appInfoFailure:
            return state.merge({
                isLoading: false,
                error: true,
                appInfo: null,
            });
        case appInfoAction.appInfoReset:
            return state.merge({
                isLoading: false,
                error: null,
                appInfo: null,
            });
        default:
            return state;
    }
};
