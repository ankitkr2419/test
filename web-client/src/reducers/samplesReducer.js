import { fromJS, List } from 'immutable';
import { listSampleActions } from 'actions/sampleActions';
import { updateSampleListSelector } from 'selectors/sampleSelectors';

const listSampleInitialState = fromJS({
	isLoading: false,
	list: List([]),
});

export const listSamplesReducer = (state = listSampleInitialState, action) => {
	switch (action.type) {
	case listSampleActions.listAction:
		return state.setIn(['isLoading'], true);
	case listSampleActions.successAction:
		return updateSampleListSelector(state, action);
	case listSampleActions.failureAction:
		return state.merge({
			error: fromJS(action.payload.error),
			isLoading: false,
		});
	case listSampleActions.addSample:
		return state.updateIn(['list'], myList => myList.push(fromJS(action.payload.sample)));
	case listSampleActions.resetSamples:
		return listSampleInitialState;
	default:
		return state;
	}
};
