import {
  homingActions,
  deckHomingActions,
  discardTipAndHomingActions,
  homingInProgressActionsForDeckA,
  homingInProgressActionsForDeckB,
} from "actions/homingActions";

export const hideHomingModal = () => ({
  type: homingActions.hideHomingModaal,
});

export const showHomingModal = () => ({
  type: homingActions.showHomingModaal,
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

// homing progressing for deck A
export const homingActionInProgressDeckA = (payload) => ({
  type: homingInProgressActionsForDeckA.homingActionInProgressForDeckA,
  payload: payload,
});

export const homingActionInCompletedDeckA = (payload) => ({
  type: homingInProgressActionsForDeckA.homingActionInSuccessForDeckA,
  payload: payload,
});

export const homingActionResetDeckA = () => ({
  type: homingInProgressActionsForDeckA.resetHomingStateForDeckA,
});

// homing progressing for deck B
export const homingActionInProgressDeckB = (payload) => ({
  type: homingInProgressActionsForDeckB.homingActionInProgressForDeckB,
  payload: payload,
});

export const homingActionInCompletedDeckB = (payload) => ({
  type: homingInProgressActionsForDeckB.homingActionInSuccessForDeckB,
  payload: payload,
});

export const homingActionResetDeckB = () => ({
  type: homingInProgressActionsForDeckB.resetHomingStateForDeckB,
});

export const deckHomingActionInitiated = (params) => ({
  type: deckHomingActions.deckHomingActionInitiated,
  payload: {
    params,
  },
});

export const deckHomingActionSuccess = () => ({
  type: deckHomingActions.deckHomingActionSuccess,
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
});
