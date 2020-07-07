import {
	listSampleActions,
} from 'actions/sampleActions';

export const fetchSamples = searchText => ({
	type: listSampleActions.listAction,
	payload: {
		searchText,
	},
});

export const addSampleLocallyCreated = sample => ({
	type: listSampleActions.addSample,
	payload: {
		sample,
	},
});

export const fetchSamplesFailed = errorResponse => ({
	type: listSampleActions.failureAction,
	payload: {
		...errorResponse,
		error: true,
	},
});

export const resetSamples = () => ({ type: listSampleActions.resetSamples });
