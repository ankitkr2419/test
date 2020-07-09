export const addWellActions =  {
	addAction: 'ADD_WELL_INITIATED',
	successAction: 'ADD_WELL_SUCCEEDED',
	failureAction: 'ADD_WELL_FAILURE',
	addWellReset: 'ADD_WELL_RESET',
};

export const listWellActions =  {
	listAction: 'FETCH_WELLS_INITIATED',
	successAction: 'FETCH_WELLS_SUCCEEDED',
	failureAction: 'FETCH_WELLS_FAILURE',
	setSelectedWell: 'SET_SELECTED_WELL',
	resetSelectedWell: 'RESET_SELECTED_WELL',
	setMultiSelectedWell: 'SET_MULTI_SELECT_WELL',
	resetMultiSelectedWell: 'RESET_MULTI_SELECT_WELL',
	toggleMultiSelectOption: 'TOGGLE_MULTI_SELECT_OPTION',
	updateWellThroughSocket: 'UPDATE_WELL_THROUGH_SOCKET_DATA',
	resetWells: 'RESET_WELLS',
};
