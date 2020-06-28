import {
	addWellActions,
	listWellActions,
} from 'actions/wellActions';

export const addWell = body => ({
	type: addWellActions.addAction,
	payload: {
		body,
	},
});

export const addWellFailed = errorResponse => ({
	type: addWellActions.failureAction,
	payload: {
		...errorResponse,
		error: true,
	},
});

export const addWellReset = () => ({
	type: addWellActions.addWellReset,
});

export const fetchWells = templateID => ({
	type: listWellActions.listAction,
	payload: {
		templateID,
	},
});

export const fetchWellsFailed = errorResponse => ({
	type: listWellActions.failureAction,
	payload: {
		...errorResponse,
		error: true,
	},
});

export const setWellSelected = (index, isSelected) => ({
	type: listWellActions.setWellSelected,
	payload: {
		isSelected,
		index,
	},
});
