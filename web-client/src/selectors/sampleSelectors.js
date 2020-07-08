import { createSelector } from 'reselect';
import { List, fromJS } from 'immutable';

const getListSamplesReducer = state => state.listSamplesReducer;

export const getSamples = createSelector(
	getListSamplesReducer,
	listSamplesReducer => listSamplesReducer,
);

const getSelectElement = (label, value) => fromJS({ label, value });

export const updateSampleListSelector = createSelector(
	state => state,
	(state, action) => action,
	(state, action) => {
		// if no sample returned
		if (action.payload.response !== null) {
			// iterate with list
			const tempState = state.updateIn(['list'], (myList) => {
				// if list is empty add all elements present in response
				if (myList.size === 0) {
					return List(action.payload.response.map(ele => getSelectElement(ele.name, ele.id)));
				}
				// if list is not empty concat response data with list
				return myList.concat(List(action.payload.response.map((ele) => {
					// duplicate element check
					const index = myList.findIndex(
						myListEle => myListEle.get('value') === ele.id,
					);
					if (index === -1) {
						return getSelectElement(ele.name, ele.id);
					}
					return null;
				})).filter(ele => ele !== null));
			});
			return tempState.setIn(['isLoading'], false);
		}
		return state.setIn(['isLoading'], false);
	},
);
