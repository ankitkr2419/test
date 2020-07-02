import activeWellActions from 'actions/activeWellActions';

export const fetchActiveWells = experimentId => ({
	type: activeWellActions.listAction,
	payload: {
		experimentId,
	},
});

export const fetchActiveWellsFailed = errorResponse => ({
	type: activeWellActions.failureAction,
	payload: {
		...errorResponse,
		error: true,
	},
});
