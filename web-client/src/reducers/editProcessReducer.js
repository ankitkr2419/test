import { fromJS } from "immutable";
import { fetchProcessDataActions } from "actions/processActions";

const editProcessInitialState = fromJS({
    isLoading: false,
    error: null,
});

export const editProcessReducer = (state = editProcessInitialState, action) => {
    switch (action.type) {
        case fetchProcessDataActions.fetchProcessDataInitiated:
            return state.merge({
                isLoading: true,
                error: null,
            });

        case fetchProcessDataActions.fetchProcessDataSuccess:
            return state.merge({
                isLoading: false,
                error: false,
                ...action.payload.response
            });
        case fetchProcessDataActions.fetchProcessDataFailure:
            return state.merge({
                isLoading: false,
                error: true,
            });
        case fetchProcessDataActions.fetchProcessDataReset:
            return editProcessInitialState;
        default:
            return state;
    }
};
