import { fromJS } from "immutable";
import {
    processListActions,
    duplicateProcessActions,
} from "actions/processActions";

const processListInitialState = fromJS({
    isLoading: false,
    error: null,
    processList: [],
    tempDuplicateProcess: null,
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

        case duplicateProcessActions.duplicateProcessInitiated:
            return state.merge({
                isLoading: true,
                error: null,
                tempDuplicateProcess: null,
            });

        //TODO: remove
        // case duplicateProcessActions.duplicateProcessSuccess:
        //     const newProcessObj = action.payload.response;
        //     return state.merge({
        //         isLoading: false,
        //         error: null,
        //         processList: [...state.toJS().processList, newProcessObj],
        //     });
        case duplicateProcessActions.duplicateProcessSuccess:
            const newProcessObj = action.payload.response;
            return state.merge({
                isLoading: false,
                error: null,
                tempDuplicateProcess: newProcessObj
            });

        case duplicateProcessActions.duplicateProcessFailure:
            return state.merge({
                isLoading: false,
                error: true,
                tempDuplicateProcess: null,
            });

        case duplicateProcessActions.duplicateProcessReset:
            return state.merge({
                tempDuplicateProcess: null,
            })
        default:
            return state;
    }
};
