import wellGraphActions, { updateGraphActions } from "actions/wellGraphActions";

export const wellGraphSucceeded = (data) => ({
  type: wellGraphActions.successAction,
  payload: {
    data,
  },
});

export const updateGraphInitiated = ({ query, experimentId, token }) => ({
  type: updateGraphActions.updateGraphInitiated,
  payload: { query, experimentId, token },
});

export const updateGraphSuccess = ({ response }) => ({
  type: updateGraphActions.updateGraphSucceeded,
  payload: { response },
});

export const updateGraphFailed = ({ error }) => ({
  type: updateGraphActions.updateGraphFailure,
  payload: { error },
});
