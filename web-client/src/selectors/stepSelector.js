import { createSelector } from 'reselect';

const getHoldListStepsReducer = state => state.listHoldStepsReducer;
const getCycleListStepsReducer = state => state.listCycleStepsReducer;

// const getListStagesReducer = state => state.listStagesReducer;

export const getHoldStepList = createSelector(
	getHoldListStepsReducer,
	listHoldStepReducer => listHoldStepReducer.updateIn(['list'], myList => myList.sortBy(ele => ele.get('created_at'))),
);

export const getCycleStepList = createSelector(
	getCycleListStepsReducer,
	listCycleStepReducer => listCycleStepReducer.updateIn(['list'], myList => myList.sortBy(ele => ele.get('created_at'))),
);
