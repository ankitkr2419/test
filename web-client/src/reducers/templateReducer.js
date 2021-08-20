import { fromJS } from "immutable";
import {
  createTemplateActions,
  listTemplateActions,
  updateTemplateActions,
  deleteTemplateActions,
  finishCreateTemplateActions,
} from "actions/templateActions";
import { getTemplateDetails } from "components/Template/templateHelper";

const listTemplateInitialState = fromJS({
  isLoading: true,
  list: [],
});

const createTemplateInitialState = {
  data: {},
  errorCreatingTemplate: null,
  isLoading: true,
  isTemplateCreated: false,
};

const updateTemplateInitialState = {
  data: {},
  errorUpdatingTemplate: null,
  isTemplateUpdated: false,
};

const deleteTemplateInitialState = {
  data: {},
  errorDeletingTemplate: null,
  isTemplateDeleted: false,
};

const finishCreateTemplateInitialState = {
  isLoading: false,
  errorFinishCreateTemplate: null,
};

// eslint-disable-next-line arrow-body-style
export const getTemplateById = (state, templateId) => {
  const result = state
    .get("list")
    .filter((ele) => ele.get("id") === templateId);
  if (result !== null && result.size !== 0) {
    return result.get(0);
  }
  return null;
};

export const listTemplatesReducer = (
  state = listTemplateInitialState,
  action
) => {
  switch (action.type) {
    case listTemplateActions.listAction:
      return state.setIn(["isLoading"], true);
    case listTemplateActions.successAction:
      return state.merge({
        list: fromJS(action.payload.response || []),
        isLoading: false,
      });
    case listTemplateActions.failureAction:
      return state.merge({
        error: fromJS(action.payload.error),
        isLoading: false,
      });
    // appending the template list with created template
    // getTemplateDetails function is used for storing only the template details in template reducer.
    // stages data is filtered out
    case createTemplateActions.successAction:
      return state.updateIn(["list"], (list) =>
        list.push(getTemplateDetails(action.payload.response))
      );
    default:
      return state;
  }
};

export const createTemplateReducer = (
  state = createTemplateInitialState,
  action
) => {
  switch (action.type) {
    case createTemplateActions.createAction:
      return { ...state, isLoading: true, isTemplateCreated: false };
    case createTemplateActions.successAction:
      return {
        ...state,
        ...action.payload,
        isLoading: false,
        isTemplateCreated: true,
        errorCreatingTemplate: false,
      };
    case createTemplateActions.failureAction:
      return {
        ...state,
        ...action.payload,
        isLoading: false,
        errorCreatingTemplate: true,
      };
    case createTemplateActions.createTemplateReset:
      return createTemplateInitialState;
    default:
      return state;
  }
};

export const updateTemplateReducer = (
  state = updateTemplateInitialState,
  action
) => {
  switch (action.type) {
    case updateTemplateActions.updateAction:
      return { ...state, isLoading: true };
    case updateTemplateActions.successAction:
      return {
        ...state,
        ...action.payload,
        isLoading: false,
        isTemplateUpdated: true,
        errorUpdatingTemplate: false,
      };
    case updateTemplateActions.failureAction:
      return {
        ...state,
        ...action.payload,
        isLoading: false,
        isTemplateUpdated: true,
        errorUpdatingTemplate: true,
      };
    case updateTemplateActions.updateTemplateReset:
      return updateTemplateInitialState;
    default:
      return state;
  }
};

export const deleteTemplateReducer = (
  state = deleteTemplateInitialState,
  action
) => {
  switch (action.type) {
    case deleteTemplateActions.deleteAction:
      return { ...state, isLoading: true, isTemplateDeleted: false };
    case deleteTemplateActions.successAction:
      return {
        ...state,
        ...action.payload,
        isLoading: false,
        isTemplateDeleted: true,
        errorDeletingTemplate: false,
      };
    case deleteTemplateActions.failureAction:
      return {
        ...state,
        ...action.payload,
        isLoading: false,
        isTemplateDeleted: true,
        errorDeletingTemplate: true,
      };
    case deleteTemplateActions.deleteTemplateReset:
      return deleteTemplateInitialState;
    default:
      return state;
  }
};

export const finishCreateTemplateReducer = (
  state = finishCreateTemplateInitialState,
  action
) => {
  switch (action.type) {
    case finishCreateTemplateActions.createAction:
      return {
        ...state,
        isLoading: true,
        errorFinishCreateTemplate: null,
      };
    case finishCreateTemplateActions.successAction:
      return {
        ...state,
        isLoading: false,
        errorFinishCreateTemplate: false,
      };
    case finishCreateTemplateActions.failureAction:
      return {
        ...state,
        isLoading: false,
        errorFinishCreateTemplate: true,
      };
    case finishCreateTemplateActions.createTemplateReset:
      return finishCreateTemplateInitialState;
    default:
      return state;
  }
};
