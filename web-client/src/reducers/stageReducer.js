import { fromJS } from "immutable";
import {
  addStageActions,
  listStageActions,
  updateStageActions,
  deleteStageActions,
} from "actions/stageActions";

const listStageInitialState = fromJS({
  isLoading: true,
  list: [],
});

const createStageInitialState = {
  data: {},
};

const updateStageInitialState = {
  data: {},
};

const deleteStageInitialState = {
  data: {},
};

export const listStagesReducer = (
  state = listStageInitialState,
  action
) => {
  switch (action.type) {
    case listStageActions.listAction:
      return state.setIn(["isLoading"], true);
    case listStageActions.successAction:
      return state.merge({ list: fromJS(action.payload), isLoading: false });
    case listStageActions.failureAction:
      return state.merge({
        error: fromJS(action.payload.error),
        isLoading: false,
      });
    default:
      return state;
  }
};

export const createStageReducer = (
  state = createStageInitialState,
  action
) => {
  switch (action.type) {
    case addStageActions.addAction:
      return { ...state, isLoading: true };
    case addStageActions.successAction:
      return { ...state, ...action.payload, isLoading: false };
    case addStageActions.failureAction:
      return { ...state, ...action.payload, isLoading: false };
    default:
      return state;
  }
};

export const updateStageReducer = (
  state = updateStageInitialState,
  action
) => {
  switch (action.type) {
    case updateStageActions.updateAction:
      return { ...state, isLoading: true };
    case updateStageActions.successAction:
      return { ...state, ...action.payload, isLoading: false };
    case updateStageActions.failureAction:
      return { ...state, ...action.payload, isLoading: false };
    default:
      return state;
  }
};

export const deleteStageReducer = (
  state = deleteStageInitialState,
  action
) => {
  switch (action.type) {
    case deleteStageActions.deleteAction:
      return { ...state, isLoading: true };
    case deleteStageActions.successAction:
      return { ...state, ...action.payload, isLoading: false };
    case deleteStageActions.failureAction:
      return { ...state, ...action.payload, isLoading: false };
    default:
      return state;
  }
};
