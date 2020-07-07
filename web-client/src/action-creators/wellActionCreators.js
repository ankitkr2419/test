import {
	addWellActions,
	listWellActions,
} from 'actions/wellActions';

export const addWell = (experimentId, body) => ({
	type: addWellActions.addAction,
	payload: {
		body,
		experimentId,
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

export const fetchWells = experimentId => ({
	type: listWellActions.listAction,
	payload: {
		experimentId,
	},
});

export const fetchWellsFailed = errorResponse => ({
	type: listWellActions.failureAction,
	payload: {
		...errorResponse,
		error: true,
	},
});

export const setSelectedWell = (index, isSelected) => ({
	type: listWellActions.setSelectedWell,
	payload: {
		isSelected,
		index,
	},
});

export const resetSelectedWells = () => ({
	type: listWellActions.resetSelectedWell,
});

export const setMultiSelectedWell = (index, isMultiSelected) => ({
	type: listWellActions.setMultiSelectedWell,
	payload: {
		isMultiSelected,
		index,
	},
});

export const resetMultiSelectedWells = () => ({
	type: listWellActions.resetMultiSelectedWell,
});

export const toggleMultiSelectOption = () => ({
	type: listWellActions.toggleMultiSelectOption,
});
