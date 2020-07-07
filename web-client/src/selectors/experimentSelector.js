import { createSelector } from 'reselect';

const getListExperimentsReducer = state => state.listExperimentsReducer;
const getCreateExperimentReducer = state => state.createExperimentReducer;

export const getExperiments = createSelector(
	getListExperimentsReducer,
	experimentReducer => experimentReducer,
);

export const getExperimentId = createSelector(
	getCreateExperimentReducer,
	createExperimentReducer => createExperimentReducer.get('id'),
);

export const getExperimentTemplate = createSelector(
	getCreateExperimentReducer,
	createExperimentReducer => ({
		templateName : createExperimentReducer.get('template_name'),
		templateId: createExperimentReducer.get('template_id'),
	}),
);

export const getIsExperimentSaved = createSelector(
	getCreateExperimentReducer,
	createExperimentReducer => createExperimentReducer.get('isExperimentSaved'),
);
