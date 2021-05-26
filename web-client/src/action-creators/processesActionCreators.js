import {
  aspireDispenseAction,
  piercingAction,
  tipDiscardAction,
  tipPickupAction,
} from "actions/processesActions";

export const savePiercingInitiated = (params) => ({
  type: piercingAction.savePiercingInitiated,
  payload: params,
});

export const savePiercingSuccess = (response) => ({
  type: piercingAction.savePiercingSuccess,
  payload: response,
});

export const savePiercingFailure = (error) => ({
  type: piercingAction.savePiercingFailed,
  payload: error,
});

export const saveAspireDispenseInitiated = (params) => ({
  type: aspireDispenseAction.saveAspireDispenseInitiated,
  payload: params,
});

export const saveAspireDispenseSuccess = (response) => ({
  type: aspireDispenseAction.saveAspireDispenseSuccess,
  payload: response,
});

export const saveAspireDispenseFailure = (error) => ({
  type: aspireDispenseAction.saveAspireDispenseFailed,
  payload: error,
});

export const saveTipPickupInitiated = (params) => ({
  type: tipPickupAction.saveTipPickUpInitiated,
  payload: params,
});

export const saveTipPickupSuccess = (response) => ({
  type: tipPickupAction.saveTipPickUpSuccess,
  payload: response,
});

export const saveTipPickupFailure = (error) => ({
  type: tipPickupAction.saveTipPickUpFailed,
  payload: error,
});

//tip-discard
export const saveTipDiscardInitiated = (params) => ({
  type: tipDiscardAction.saveTipDiscardInitiated,
  payload: params,
});

export const saveTipDiscardSuccess = (response) => ({
  type: tipDiscardAction.saveTipDiscardSuccess,
  payload: response,
});

export const saveTipDiscardFailure = (error) => ({
  type: tipDiscardAction.saveTipDiscardFailed,
  payload: error,
});
