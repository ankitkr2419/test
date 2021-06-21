import {
  homingActions,
  deckHomingActions,
  discardTipAndHomingActions,
} from "actions/homingActions";

export const hideHomingModal = () => ({
  type: homingActions.hideHomingModaal,
  payload: {},
});

export const showHomingModal = () => ({
  type: homingActions.showHomingModaal,
  payload: {},
});

export const homingActionInitiated = (params) => ({
  type: homingActions.homingActionInitiated,
  payload: { params },
});

export const homingActionSuccess = (homingData) => ({
  type: homingActions.homingActionSuccess,
  payload: {
    homingData,
  },
});

export const homingActionFailed = (serverErrors) => ({
  type: homingActions.homingActionFailed,
  payload: {
    serverErrors,
  },
});

export const homingActionInProgress = (homingActionInProgress) => ({
  type: homingActions.homingActionInProgress,
  payload: homingActionInProgress,
});

export const homingActionInCompleted = (homingActionInCompleted) => ({
  type: homingActions.homingActionInCompleted,
  payload: homingActionInCompleted,
});

export const deckHomingActionInitiated = (params) => ({
  type: deckHomingActions.deckHomingActionInitiated,
  payload: {
    params,
  },
});

export const deckHomingActionSuccess = () => ({
  type: deckHomingActions.deckHomingActionSuccess,
  payload: {},
});

export const deckHomingActionFailed = (serverErrors) => ({
  type: deckHomingActions.deckHomingActionFailed,
  payload: {
    serverErrors,
  },
});

export const discardTipAndHomingActionInitiated = (params) => ({
  type: discardTipAndHomingActions.discardTipAndHomingActionInitiated,
  payload: {
    params,
  },
});

export const discardTipAndHomingActionSuccess = (homingData) => ({
  type: discardTipAndHomingActions.discardTipAndHomingActionSuccess,
  payload: {
    homingData,
  },
});

export const discardTipAndHomingActionFailed = (serverErrors) => ({
  type: discardTipAndHomingActions.discardTipAndHomingActionFailed,
  payload: {
    serverErrors,
  },
});

export const discardTipAndHomingActionReset = () => ({
  type: discardTipAndHomingActions.discardTipAndHomingActionReset,
  payload: {},
});
