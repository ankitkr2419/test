import { processListActions, duplicateProcessActions, fetchProcessDataActions } from "actions/processActions";

export const processListInitiated = (params) => ({
    type: processListActions.processListInitiated,
    payload: {
        ...params,
    }
});

export const setProcessList = (params) => ({
    type: processListActions.setProcessList,
    payload: {
        ...params,
    }
});

export const duplicateProcessInitiated = (params) => ({
    type: duplicateProcessActions.duplicateProcessInitiated,
    payload: {
        ...params,
    }
});

export const duplicateProcessFail = (params) => ({
    type: duplicateProcessActions.duplicateProcessFailure,
    payload: {
        ...params,
    }
});

export const duplicateProcessReset = () => ({
    type: duplicateProcessActions.duplicateProcessReset,
    payload: {}
});

export const fetchProcessDataInitiated = (params) => ({
    type: fetchProcessDataActions.fetchProcessDataInitiated,
    payload: {
        ...params
    }
});

export const fetchProcessDataFail = (params) => ({
    type: fetchProcessDataActions.fetchProcessDataFailure,
    payload: {
        ...params
    }
});

export const fetchProcessDataReset = () => ({
    type: fetchProcessDataActions.fetchProcessDataReset,
    payload: {}
})