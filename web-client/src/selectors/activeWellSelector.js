import { createSelector } from 'reselect';

const getActiveWellReducerReducer = state => state.activeWellReducer;

export const getActiveLoadedWellFlag = createSelector(
	getActiveWellReducerReducer,
	activeWellReducer => activeWellReducer.get('isDataLoaded'),
);

export const getActiveLoadedWells = createSelector(
	getActiveWellReducerReducer,
	activeWellReducer => activeWellReducer.get('list'),
);
