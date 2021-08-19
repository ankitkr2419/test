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

export const resetGraphInitiated = ({ experimentId, token }) => ({
  type: resetGraphActions.resetGraphInitiated,
  payload: { experimentId, token },
});

export const resetGraphSuccess = ({ response }) => ({
  type: resetGraphActions.resetGraphSucceeded,
  payload: { response },
});

export const resetGraphFailed = ({ error }) => ({
  type: resetGraphActions.resetGraphFailure,
  payload: { error },
});
