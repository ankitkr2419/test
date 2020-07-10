export const validateHoldTime = holdTime => holdTime.match(/^[1-9]\d{0,4}$/);

// Validate create step form
export const validateStepForm = ({
	rampRate,
	targetTemperature,
	holdTime,
}) => {
	if (rampRate !== '' && targetTemperature !== '' && holdTime !== '' && holdTime !== '0') {
		if (validateHoldTime(holdTime)) {
			return true;
		}
		return false;
	}
	return false;
};
