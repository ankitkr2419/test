import { createSelector } from 'reselect';
import { TARGET_LINE_COLORS } from 'appConstants';

const getExperimentsTarget = state => state.listExperimentTargetsReducer;

export const getExperimentTargets = createSelector(
	getExperimentsTarget,
	experimentTargetsReducer => experimentTargetsReducer,
);

// separate target list maintained for filtering chart
export const getExperimentGraphTargets = createSelector(
	getExperimentsTarget,
	experimentTargetsReducer => experimentTargetsReducer.get('graphTargets'),
);

export const addIsActiveFlag = createSelector(
	data => data,
	(data) => {
		if (data === null || data.length === 0) {
			return [];
		}

		return data.map((ele, index) => ({
			...ele,
			isActive: true,
			lineColor: TARGET_LINE_COLORS[index],
		}));
	},
);