import {
  aspireDispenseAction,
  heatingAction,
  magnetAction,
  delayAction,
  piercingAction,
  shakingAction,
  tipDiscardAction,
  tipPickupAction,
} from "actions/processesActions";

//peircing
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

//aspire-dispense
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

//tip-pickup
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

//shaking
export const saveShakingInitiated = (params) => ({
  type: shakingAction.saveShakingInitiated,
  payload: params,
});

export const saveShakingSuccess = (response) => ({
  type: shakingAction.saveShakingSuccess,
  payload: response,
});

export const saveShakingFailure = (error) => ({
  type: shakingAction.saveShakingFailed,
  payload: error,
});

//heating
export const saveHeatingInitiated = (params) => ({
  type: heatingAction.saveHeatingInitiated,
  payload: params,
});

export const saveHeatingSuccess = (response) => ({
  type: heatingAction.saveHeatingSuccess,
  payload: response,
});

export const saveHeatingFailure = (error) => ({
  type: heatingAction.saveHeatingFailed,
  payload: error,
});

//magnet
export const saveMagnetInitiated = (params) => ({
  type: magnetAction.saveMagnetInitiated,
  payload: params,
});

export const saveMagnetSuccess = (response) => ({
  type: magnetAction.saveMagnetSuccess,
  payload: response,
});

export const saveMagnetFailure = (error) => ({
  type: magnetAction.saveMagnetFailed,
  payload: error,
});

//delay
export const saveDelayInitiated = (params) => ({
  type: delayAction.saveDelayInitiated,
  payload: params,
});

export const saveDelaySuccess = (response) => ({
  type: delayAction.saveDelaySuccess,
  payload: response,
});

export const saveDelayFailure = (error) => ({
  type: delayAction.saveDelayFailed,
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
