import { createSelector } from 'reselect';

const runExperimentSelector = state => state.runExperimentReducer;

export const getRunExperimentReducer = createSelector(
	runExperimentSelector,
	runExperimentReducer => runExperimentReducer,
);
