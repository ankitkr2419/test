import { processAction } from "actions/processesActions";

//save process
export const saveProcessInitiated = (params) => ({
  type: processAction.saveProcessInitiated,
  payload: params,
});

export const saveProcessSuccess = (response) => ({
  type: processAction.saveProcessSuccess,
  payload: response,
});

export const saveProcessFailure = (error) => ({
  type: processAction.saveProcessFailed,
  payload: error,
});

export const saveProcessReset = () => ({
  type: processAction.saveProcessReset,
});
