import { createSelector } from 'reselect';

const PLATE_CAPACITY = 96;

const getWellListSelector = state => state.wellListReducer;
const getAddWellsSelector = state => state.addWellsReducer;


export const getWells = createSelector(
	getWellListSelector,
	wellListReducer => wellListReducer,
);

/**
 * return array of selected position
 */
export const getWellsPosition = createSelector(
	wellListReducer => wellListReducer,
	wellListReducer => wellListReducer.get('defaultList').map((ele, index) => {
		if (ele.get('isSelected') === true || ele.get('isMultiSelected') === true) {
			return index;
		}
		return null;
	}).filter(ele => ele !== null),
);

export const setSelectedToList = (state, { isSelected, index }) => state.setIn(['defaultList', index, 'isSelected'], isSelected);

export const setMultiSelectedToList = (state, { isMultiSelected, index }) => state.setIn(['defaultList', index, 'isMultiSelected'], isMultiSelected);

export const resetWellDefaultList = state => state.updateIn(['defaultList'], myDefaultList => myDefaultList.map(ele => ele.setIn(['isSelected'], false)));
export const resetMultiWellDefaultList = state => state.updateIn(['defaultList'], myDefaultList => myDefaultList.map(ele => ele.setIn(['isMultiSelected'], false)));

/**
 * getDefaultPlatesList return wells default data w.r.t PLATE_CAPACITY.
 */
export const getDefaultWellsList = createSelector(
	() => {
		const arr = [];
		const initialPlateState = {
			isSelected: false,
			isWellFilled: false,
			isRunning: false,
			isMultiSelected: false,
			status: '', // red, green, orange
			initial: '',
			id: null,
		};

		for (let i = 0; i !== PLATE_CAPACITY; i += 1) {
			arr.push(initialPlateState);
		}
		return arr;
	},
);

/**
 * Return isWellSaved flag from addWellReducer
 */
export const getWellsSavedStatus = createSelector(
	getAddWellsSelector,
	addWellReducer => addWellReducer.get('isWellSaved'),
);
/**
 * Get well data with w.r.t position
 * @param {*} wells
 * @param {*} position
 */
const getSelectedWell = (wells, position) => wells.filter(ele => ele.position === position)[0];

/**
 * updateWellListSelector accepts current state and action
 * It will update default list with updated action response
 * It will fill the popover data and well config data
 */
export const updateWellListSelector = createSelector(
	state => state,
	(state, action) => action,
	(state, action) => {
		// if no wells selected
		if (action.payload.response !== null) {
			const {
				payload: { response },
			} = action;
			// get position of selected wells from response
			const positions = response.map(ele => ele.position);
			// iteration over default map
			const tempState = state.updateIn(['defaultList'], myDefaultList => myDefaultList.map((ele, index) => {
				// find the index present in response data
				if (positions.includes(index)) {
					// get selected well data by index
					const selectedWell = getSelectedWell(response, index);
					// merge selected well data and modify local fields.
					return ele.merge({
						isWellFilled: true,
						...selectedWell,
						initial: selectedWell.task[0],
						status: selectedWell.color_code || 'green',
						sample: selectedWell.sample_name,
					});
				}
				return ele;
			}));
			return tempState.setIn(['isLoading'], false);
		}
		return state.setIn(['isLoading'], false);
	},
);
