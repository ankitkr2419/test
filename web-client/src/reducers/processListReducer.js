import { fromJS } from "immutable";
import { processListActions } from "actions/processActions";

const processListInitialState = fromJS({
    isLoading: false,
    error: null,
    processList: [],
});

export const processListReducer = (state = processListInitialState, action) => {
    switch (action.type) {
        case processListActions.processListInitiated:
            return state.merge({
                isLoading: true,
                error: null,
            });
        case processListActions.processListSuccess:
            return state.merge({
                isLoading: false,
                processList: action.payload?.response,
                error: null,
            });
        case processListActions.processListFailure:
            return state.merge({
                isLoading: false,
                processList: [],
                error: true,
            });
        case processListActions.processListReset:
            return state.merge({
                processList: [],
                error: null,
            });
        default:
            return state;
    }
};
