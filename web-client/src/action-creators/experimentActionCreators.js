import {
  createExperimentActions,
  listExperimentActions,
} from "actions/experimentActions";

export const createExperiment = (body, token) => ({
  type: createExperimentActions.createAction,
  payload: {
    body,
    token,
  },
});

export const createExperimentSucceeded = (body) => ({
  type: createExperimentActions.successAction,
  payload: { response: body },
});

export const createExperimentFailed = (errorResponse) => ({
  type: createExperimentActions.failureAction,
  payload: {
    ...errorResponse,
    error: true,
  },
});

export const createExperimentReset = () => ({
  type: createExperimentActions.createExperimentReset,
});

export const fetchExperiments = (token) => ({
  type: listExperimentActions.listAction,
  payload: token,
});

export const fetchExperimentsFailed = (errorResponse) => ({
  type: listExperimentActions.failureAction,
  payload: {
    ...errorResponse,
    error: true,
  },
});
