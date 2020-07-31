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
