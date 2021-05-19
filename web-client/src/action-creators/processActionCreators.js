import { processListActions } from "actions/processActions";

export const processListInitiated = (params) => ({
    type: processListActions.processListInitiated,
    payload: {
        ...params,
    }
});


