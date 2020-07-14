import { createSelector } from 'reselect';

const getListStepsReducer = state => state.listStepsReducer;
const getListStagesReducer = state => state.listStagesReducer;

export const getStepList = createSelector(
	getListStepsReducer,
	listStepReducer => listStepReducer.updateIn(['list'], myList => myList.sortBy(ele => ele.get('created_at'))),
);

export const getStageType = createSelector(
	getListStagesReducer,
	(_, stageId) => stageId,
	(listStagesReducer, stageId) => {
		const stage = listStagesReducer.get('list').find(ele => ele.get('id') === stageId).toJS();
		return stage.type;
	},
);
