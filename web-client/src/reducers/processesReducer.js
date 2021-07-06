import { processAction } from "actions/processesActions";

const initialState = { error: null };

export const processesReducer = (state = initialState, actions) => {
  switch (actions.type) {
    case processAction.saveProcessSuccess:
      return {
        ...state,
        error: false,
      };

    case processAction.saveProcessFailed:
      return {
        ...state,
        error: true,
      };

    case processAction.saveProcessReset:
      return {
        ...state,
        error: null,
      };

    default:
      return state;
  }
};
