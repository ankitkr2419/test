import { createSelector } from 'reselect';

const getListStepsReducer = state => state.listStepsReducer;

export const getStepList = createSelector(
	getListStepsReducer,
	listStepReducer => listStepReducer.updateIn(['list'], myList => myList.sortBy(ele => ele.get('created_at'))),
);
