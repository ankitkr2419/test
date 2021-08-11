import {
	listSampleActions,
} from 'actions/sampleActions';

export const fetchSamples = (searchText, token) => ({
	type: listSampleActions.listAction,
	payload: {
		searchText,
		token,
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
