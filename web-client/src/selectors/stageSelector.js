import { createSelector } from 'reselect';

const getListStagesReducer = state => state.listStagesReducer;

export const getStageList = createSelector(
	getListStagesReducer,
	listStagesReducer => listStagesReducer.updateIn(['list'], myList => myList.sortBy(ele => ele.get('created_at'))),
);
