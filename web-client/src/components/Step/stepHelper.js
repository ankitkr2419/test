import { createSelector } from 'reselect';
import { MAX_RAMP_RATE, MIN_RAMP_RATE, MAX_TARGET_TEMPERATURE, MIN_TARGET_TEMPERATURE } from './stepConstants';

export const validateHoldTime = holdTime => holdTime.match(/^[1-9]\d{0,4}$/);

// Validate Ramp rate below Maximum and above Minimum ramp rate

export const validateRampRate = createSelector(
	rampRate => rampRate,
	(rampRate) => {
		const ramp_rate = parseFloat(rampRate);
		if (ramp_rate <= MAX_RAMP_RATE && ramp_rate >= MIN_RAMP_RATE) {
			return true;
		}
		return false;
	},
);

// Validate Target Temperature below Maximum and above Minimum Target temperature
export const validateTargetTemperature = createSelector(
	targetTemperature => targetTemperature,
	(targetTemperature) => {
		const target_temp = parseFloat(targetTemperature);
		if (target_temp <= MAX_TARGET_TEMPERATURE && target_temp >= MIN_TARGET_TEMPERATURE) {
			return true;
		}
		return false;
	},
);

// Validate create step form
export const validateStepForm = createSelector(
	state => state,
	({
		rampRate,
		targetTemperature,
		holdTime,
	}) => {
		if (rampRate !== '' && targetTemperature !== '' && holdTime !== '' && holdTime !== '0') {
			if (validateHoldTime(holdTime) &&
				validateRampRate(rampRate) &&
				validateTargetTemperature(targetTemperature)) {
				return true;
			}
			return false;
		}
		return false;
	},
);
