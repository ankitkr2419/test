import { createSelector } from 'reselect';

const getActiveWellReducer = state => state.activeWellReducer;

export const getActiveLoadedWellFlag = createSelector(
	getActiveWellReducer,
	activeWellReducer => activeWellReducer.get('isDataLoaded'),
);

export const getActiveLoadedWells = createSelector(
	getActiveWellReducer,
	activeWellReducer => activeWellReducer.get('list'),
);
