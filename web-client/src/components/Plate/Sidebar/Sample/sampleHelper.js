import { createSelector } from 'reselect';
import { fromJS } from 'immutable';

export const getSampleTargetIds = createSelector(
	sampleTargets => sampleTargets,
	sampleTargets => sampleTargets.map(ele => ele.target_id),
);

// get sample target list by adding the isSelected property as true or false respectively
// for selected and unselected targets in experiment targets list
export const getSampleTargetList = createSelector(
	getSampleTargetIds,
	(_, experimentTargets) => experimentTargets,
	(sampleTargetIds, experimentTargets) => {
		const sampleTargetList = experimentTargets.map(ele => {
			if (sampleTargetIds.includes(ele.get('target_id'))) {
				return ele.merge({ isSelected: true });
			}
			return ele.merge({ isSelected: false });
		});
		return fromJS(sampleTargetList);
	},
);

export const getInitSampleTargetList = createSelector(
	experimentTargetsList => experimentTargetsList,
	experimentTargetsList => experimentTargetsList.map(ele => ele.merge({ isSelected: true })),
);
