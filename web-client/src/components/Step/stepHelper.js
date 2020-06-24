export const validateHoldTime = holdTime => holdTime.match(/^[0-9][0-9]:[0-9][0-9]$/);

// Validate create step form
export const validateStepForm = ({
	rampRate,
	targetTemperature,
	holdTime,
}) => {
	if (rampRate !== '' && targetTemperature !== '' && holdTime !== '') {
		if (validateHoldTime(holdTime)) {
			return true;
		}
		return false;
	}
	return false;
};
