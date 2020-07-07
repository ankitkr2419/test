import { createSelector } from 'reselect';

const runExperimentSelector = state => state.runExperimentReducer;

export const getRunExperimentReducer = createSelector(
	runExperimentSelector,
	runExperimentReducer => runExperimentReducer,
);

export const getExperimentStatus = createSelector(
	runExperimentSelector,
	runExperimentReducer => runExperimentReducer.get('experimentStatus'),
);
