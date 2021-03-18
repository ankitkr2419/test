import {
  homingActions,
  deckHomingActions
} from "actions/homingActions";

export const homingActionInitiated = () => ({
  type: homingActions.homingActionInitiated,
  payload: {}
});

export const homingActionSuccess = (homingData) => ({
  type: homingActions.homingActionSuccess,
  payload: {
    homingData
  }
});

export const homingActionFailed = (serverErrors) => ({
  type: homingActions.homingActionFailed,
  payload: {
    serverErrors
  }
});

export const deckHomingActionInitiated = (params) => ({
  type: deckHomingActions.deckHomingActionInitiated,
  payload: {
    params
  }
});

export const deckHomingActionSuccess = () => ({
  type: deckHomingActions.deckHomingActionSuccess,
  payload: {}
});

export const deckHomingActionFailed = (serverErrors) => ({
  type: deckHomingActions.deckHomingActionFailed,
  payload: {
    serverErrors
  }
});