import {
  aspireDispenseAction,
  piercingAction,
  tipPickupAction,
} from "actions/processesActions";

export const saveAspireDispenseInitiated = (params) => ({
  type: aspireDispenseAction.saveAspireDispenseInitiated,
  payload: params,
});

export const saveAspireDispenseSuccess = (response) => ({
  type: aspireDispenseAction.saveAspireDispenseSuccess,
  payload: response,
});

export const saveAspireDispenseFailure = (response) => ({
  type: aspireDispenseAction.saveAspireDispenseFailed,
  payload: response,
});

export const saveTipPickupInitiated = (params) => ({
  type: tipPickupAction.saveTipPickUpInitiated,
  payload: params,
});

export const saveTipPickupSuccess = (response) => ({
  type: tipPickupAction.saveTipPickUpSuccess,
  payload: response,
});

export const saveTipPickupFailure = (response) => ({
  type: tipPickupAction.saveTipPickUpFailed,
  payload: response,
});
