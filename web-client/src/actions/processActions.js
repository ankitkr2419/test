export const processListActions = {
    processListInitiated: "PROCESS_LIST_INITIATED",
    processListSuccess: "PROCESS_LIST_SUCCESS",
    processListFailure: "PROCESS_LIST_FAILURE",
    processListReset: "PROCESS_LIST_RESET",
    setProcessList: "SET_PROCESS_LIST"//set updated list
};

export const duplicateProcessActions = {
    duplicateProcessInitiated: "DUPLICATE_PROCESS_INITIATED",
    duplicateProcessSuccess: "DUPLICATE_PROCESS_SUCCESS",
    duplicateProcessFailure: "DUPLICATE_PROCESS_FAILURE",
    duplicateProcessReset: "DUPLICATE_PROCESS_RESET",
}

export const fetchProcessDataActions = {
    fetchProcessDataInitiated: "FETCH_PROCESS_DATA_INITIATED",
    fetchProcessDataSuccess: "FETCH_PROCESS_DATA_SUCCESS",
    fetchProcessDataFailure: "FETCH_PROCESS_DATA_FAILURE",
    fetchProcessDataReset: "FETCH_PROCESS_DATA_RESET" 
}

export const sequenceActions = {
    sequenceInitiated: "SEQUENCE_INITIATED",
    sequenceSuccess: "SEQUENCE_SUCCESS",
    sequenceFailure: "SEQUENCE_FAILURE",
    sequenceReset: "SEQUENCE_RESET",
}

export const deleteProcessActions = {
    deleteProcessInitiated: "DELETE_PROCESS_INITIATED",
    deleteProcessSuccess: "DELETE_PROCESS_SUCCESS",
    deleteProcessFailure: "DELETE_PROCESS_FAILURE",
}