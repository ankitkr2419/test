import { List, fromJS } from 'immutable';
import { createTemplateActions, listTemplateActions } from "actions/templateActions";

const listTemplateInitialState = {
  list: List(),
};

export const listTemplateReducer = (state = listTemplateInitialState, action) => {
  switch (action.type) {
    case listTemplateActions.listAction:
      return { ...state, isLoading: true };
      case listTemplateActions.successAction:
      return { ...state, list : fromJS(action.payload), isLoading: false };
      case listTemplateActions.failureAction:
      return { ...state, ...action.payload, isLoading: false };
    default:
      return state;
  }
};

const createTemplateInitialState = {
  data: {},
};

export const createTemplateReducer = (state = createTemplateInitialState, action) => {
  switch (action.type) {
    case createTemplateActions.createAction:
      return { ...state, ...action.payload, isLoading: true };
      case createTemplateActions.successAction:
      return { ...state, ...action.payload, isLoading: false };
      case createTemplateActions.failureAction:
      return { ...state, ...action.payload, isLoading: false };
    default:
      return state;
  }
};

