import {
	createExperimentActions,
	listExperimentActions,
} from 'actions/experimentActions';

export const createExperiment = body => ({
	type: createExperimentActions.createAction,
	payload: {
		body,
	},
});

export const createExperimentFailed = errorResponse => ({
	type: createExperimentActions.failureAction,
	payload: {
		...errorResponse,
		error: true,
	},
});

export const createExperimentReset = () => ({
	type: createExperimentActions.createExperimentReset,
});

export const fetchExperiments = () => ({
	type: listExperimentActions.listAction,
});

export const fetchExperimentsFailed = errorResponse => ({
	type: listExperimentActions.failureAction,
	payload: {
		...errorResponse,
		error: true,
	},
});
