import { createSelector } from 'reselect';

const getExperimentsTarget = state => state.listExperimentTargetsReducer;

export const getExperimentTargets = createSelector(
	getExperimentsTarget,
	experimentTargetsReducer => experimentTargetsReducer,
);
