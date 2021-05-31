import { fromJS } from "immutable";
import {
    processListActions,
    duplicateProcessActions,
} from "actions/processActions";
import { resetIsOpenInProcessList } from "components/ProcessListing/helper";
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
            const list = resetIsOpenInProcessList(action.payload?.response);
            return state.merge({
                isLoading: false,
                processList: list,
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

        //if we want to set new processList
        case processListActions.setProcessList:
            return state.merge({
                processList: action.payload.processList,
            });
        case duplicateProcessActions.duplicateProcessInitiated:
            return state.merge({
                isLoading: true,
                error: null,
            });

        case duplicateProcessActions.duplicateProcessSuccess:
            const newProcessList = [...state.toJS().processList, action.payload.response];
            
            //add isOpen to new process and reset isOpen for old processes
            const processListAfterIsOpenReset = resetIsOpenInProcessList(newProcessList);
            
            return state.merge({
                isLoading: false,
                error: null,
                processList: processListAfterIsOpenReset,
            });
            
        case duplicateProcessActions.duplicateProcessFailure:
            return state.merge({
                isLoading: false,
                error: true,
            });

        case duplicateProcessActions.duplicateProcessReset:
            return state.merge({
                error: null,
            });
        default:
            return state;
    }
};
