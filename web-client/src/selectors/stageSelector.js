import { createSelector } from 'reselect';

const getListStagesReducer = state => state.listStagesReducer;

export const getStageList = createSelector(
	getListStagesReducer,
	listStagesReducer => listStagesReducer.updateIn(['list'], myList => myList.sortBy(ele => ele.get('created_at'))),
);

// get cycle stage
export const getCycleStage = createSelector(
	getListStagesReducer,
	listStagesReducer => listStagesReducer.get('list').find(ele => ele.get('type') === 'cycle'),
);

// get hold stage Id
export const getHoldStageId = createSelector(
	getListStagesReducer,
	listStagesReducer => {
		const holdStage = listStagesReducer.get('list').find(ele => ele.get('type') === 'hold');
		// if hold stage is present return its ID or return null
		return holdStage !== undefined ? holdStage.get('id') : null;
	},
);

// get cycle stage Id
export const getCycleStageId = createSelector(
	getListStagesReducer,
	listStagesReducer => {
		const cycleStage = listStagesReducer.get('list').find(ele => ele.get('type') === 'cycle');
		// if cycle stage is present return its ID or return null
		return cycleStage !== undefined ? cycleStage.get('id') : null;
	},
);

// get cycle stage repeat count
export const getCycleRepeatCount = createSelector(
	getListStagesReducer,
	listStagesReducer => {
		const cycleStage = listStagesReducer.get('list').find(ele => ele.get('type') === 'cycle');
		// if cycle stage is present return its ID or return null
		return cycleStage !== undefined ? cycleStage.get('repeat_count') : null;
	},
);
