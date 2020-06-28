import { createSelector } from 'reselect';
// import { TARGET_CAPACITY } from '../../constants';
const PLATE_CAPACITY = 96;

const getWellListSelector = state => state.wellListReducer;

export const getWells = createSelector(
	getWellListSelector,
	wellListReducer => wellListReducer,
);

export const setSelectedToList = (state, { isSelected, index }) => state.setIn(['defaultList', index, 'isSelected'], isSelected);

export const getDefaultPlatesList = createSelector(
	() => {
		const arr = [];
		const initialPlateState = {
			isSelected : false,
			isRunning: false,
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
