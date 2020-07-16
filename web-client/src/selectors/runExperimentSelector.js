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

export const getTimeNow = () => {
	const date = new Date();
	let hours = date.getHours();
	let minutes = date.getMinutes();
	const ampm = hours >= 12 ? 'pm' : 'am';
	hours %= 12;
	hours = hours || 12; // the hour '0' should be '12'
	minutes = minutes < 10 ? `0${minutes}` : minutes;
	const strTime = `${hours}:${minutes} ${ampm}`;
	return strTime;
};
