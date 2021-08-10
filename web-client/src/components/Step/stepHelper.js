import { createSelector } from "reselect";
import {
  MAX_RAMP_RATE,
  MIN_RAMP_RATE,
  MAX_TARGET_TEMPERATURE,
  MIN_TARGET_TEMPERATURE,
  MIN_HOLD_TIME,
} from "./stepConstants";

export const getPrevValue = (steps, stepId, key) => {
  const holdStepsArr = steps.toJS();

  if (holdStepsArr?.length > 0) {
    const stepObj = holdStepsArr.find((step) => step.id === stepId);
    if (stepObj) {
      const prevValue = stepObj[key];
      return prevValue;
    }
  }
  return null;
};

export const validateHoldTime = (holdTime, dataCapture) => {
  if (dataCapture === true) {
    if (parseInt(holdTime) >= MIN_HOLD_TIME) {
      return true;
    }
    return false;
  } else {
    if (holdTime !== "") {
      return true;
    }
  }
  return false;
};

// Validate Ramp rate. Should be below Maximum and above Minimum ramp rate
export const validateRampRate = createSelector(
  (rampRate) => rampRate,
  (rampRate) => {
    const ramp_rate = parseFloat(rampRate);
    if (ramp_rate <= MAX_RAMP_RATE && ramp_rate >= MIN_RAMP_RATE) {
      return true;
    }
    return false;
  }
);

// Validate Target Temperature. Should be below Maximum and above Minimum Target temperature
export const validateTargetTemperature = createSelector(
  (targetTemperature) => targetTemperature,
  (targetTemperature) => {
    const target_temp = parseFloat(targetTemperature);
    if (
      target_temp <= MAX_TARGET_TEMPERATURE &&
      target_temp >= MIN_TARGET_TEMPERATURE
    ) {
      return true;
    }
    return false;
  }
);

// Validate create step form
export const validateStepForm = createSelector(
  (state) => state,
  ({ rampRate, targetTemperature, holdTime, dataCapture }) => {
    if (
      rampRate !== "" &&
      targetTemperature !== "" &&
      holdTime !== "" &&
      holdTime !== "0"
    ) {
      if (
        validateHoldTime(holdTime, dataCapture) &&
        validateRampRate(rampRate) &&
        validateTargetTemperature(targetTemperature)
      ) {
        return true;
      }
      return false;
    }
    return false;
  }
);
