export const validateHoldTime = holdTime => holdTime.match(/^[1-9]\d{0,4}$/);

// Validate Ramp rate below 6 and above 0.5
export const validateRampRate = rampRate => {
	const ramp = parseFloat(rampRate);
	if (ramp <= 6 && ramp >= 0.5) {
		return true;
	}
	return false;
};

// Validate Target Temperature below 120 and above 22
export const validateTargetTemperature = targetTemperature => {
	const target = parseFloat(targetTemperature);
	if (target <= 120 && target >= 22) {
		return true;
	}
	return false;
};

// Validate create step form
export const validateStepForm = ({
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
};
