import wellGraphActions, {
  resetGraphActions,
  updateGraphActions,
} from "actions/wellGraphActions";

export const wellGraphSucceeded = (data) => ({
  type: wellGraphActions.successAction,
  payload: {
    data,
  },
});

export const updateGraphInitiated = (payload) => ({
  type: updateGraphActions.updateGraphInitiated,
  payload,
});

export const updateGraphSuccess = (response) => ({
  type: updateGraphActions.updateGraphSucceeded,
  payload: response,
});

export const updateGraphFailed = (error) => ({
  type: updateGraphActions.updateGraphFailure,
  payload: error,
});

export const resetGraphInitiated = (payload) => ({
  type: resetGraphActions.resetGraphInitiated,
  payload,
});

export const resetGraphSuccess = (response) => ({
  type: resetGraphActions.resetGraphSucceeded,
  payload: response,
});

export const resetGraphFailed = (error) => ({
  type: resetGraphActions.resetGraphFailure,
  payload: error,
});
