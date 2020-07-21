import temperatureGraphActions from 'actions/temperatureGraphActions';

export const temperatureDataSucceeded = (data) => ({
	type: temperatureGraphActions.successAction,
	payload: {
		data,
	},
});
