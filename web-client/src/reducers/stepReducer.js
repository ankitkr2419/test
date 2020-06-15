import { fromJS } from "immutable";
import {
  addStepActions,
  listStepActions,
  updateStepActions,
  deleteStepActions,
} from "actions/stepActions";

const listStepInitialState = fromJS({
  isLoading: true,
  list: [],
});

const createStepInitialState = {
  data: {},
};

const updateStepInitialState = {
  data: {},
};

const deleteStepInitialState = {
  data: {},
};

export const listStepsReducer = (
  state = listStepInitialState,
  action
) => {
  switch (action.type) {
    case listStepActions.listAction:
      return state.setIn(["isLoading"], true);
    case listStepActions.successAction:
      return state.merge({ list: fromJS(action.payload), isLoading: false });
    case listStepActions.failureAction:
      return state.merge({
        error: fromJS(action.payload.error),
        isLoading: false,
      });
    default:
      return state;
  }
};

export const createStepReducer = (
  state = createStepInitialState,
  action
) => {
  switch (action.type) {
    case addStepActions.addAction:
      return { ...state, isLoading: true };
    case addStepActions.successAction:
      return { ...state, ...action.payload, isLoading: false };
    case addStepActions.failureAction:
      return { ...state, ...action.payload, isLoading: false };
    default:
      return state;
  }
};

export const updateStepReducer = (
  state = updateStepInitialState,
  action
) => {
  switch (action.type) {
    case updateStepActions.updateAction:
      return { ...state, isLoading: true };
    case updateStepActions.successAction:
      return { ...state, ...action.payload, isLoading: false };
    case updateStepActions.failureAction:
      return { ...state, ...action.payload, isLoading: false };
    default:
      return state;
  }
};

export const deleteStepReducer = (
  state = deleteStepInitialState,
  action
) => {
  switch (action.type) {
    case deleteStepActions.deleteAction:
      return { ...state, isLoading: true };
    case deleteStepActions.successAction:
      return { ...state, ...action.payload, isLoading: false };
    case deleteStepActions.failureAction:
      return { ...state, ...action.payload, isLoading: false };
    default:
      return state;
  }
};
