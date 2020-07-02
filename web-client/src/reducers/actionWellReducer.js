import { fromJS, List } from 'immutable';
import activeWellActions from 'actions/activeWellActions';

const runInitialState = fromJS({
	isLoading: true,
	isDataLoaded: false,
	list: List([]),
});

export const activeWellReducer = (state = runInitialState, action) => {
	switch (action.type) {
	case activeWellActions.listAction:
		return state.merge({ isLoading: true, isDataLoaded: false });
	case activeWellActions.successAction:
		return state.merge({
			list: fromJS(action.payload.response || []),
			isLoading: false,
			isDataLoaded: true,
		});
	case activeWellActions.failureAction:
		return state.merge({ isError: true, isLoading: false });
	default:
		return state;
	}
};
