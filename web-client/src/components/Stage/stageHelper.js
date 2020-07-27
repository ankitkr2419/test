import { createSelector } from 'reselect';
import { MAX_REPEAT_COUNT, MIN_REPEAT_COUNT } from './stageConstants';

export const validateRepeatCount = createSelector(
	repeatCount => repeatCount,
	(repeatCount) => {
		if (repeatCount >= MIN_REPEAT_COUNT && repeatCount <= MAX_REPEAT_COUNT) {
			return true;
		}
		return false;
	},
);

// Validate create stage form
export const validateStageForm = createSelector(
	stageFormStateJS => stageFormStateJS,
	({ stageType, stageRepeatCount, repeatCountError }) => {
		// if stage type is cycle check if repeatCountError is false for valid stage form
		if (
			stageType.value === 'cycle'
			&& stageRepeatCount !== ''
			&& repeatCountError === false
		) {
			return true;
		}
		// Repeat count is not applicable for hold stage so we don't validate
		// repeat count for hold stage
		if (stageType !== '' && stageType.value === 'hold') {
			return true;
		}
		return false;
	},
);
