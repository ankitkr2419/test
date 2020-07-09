import { fromJS } from 'immutable';
import {
	runExperimentActions,
	stopExperimentActions,
	experimentCompleteActions,
} from 'actions/runExperimentActions';
import loginActions from 'actions/loginActions';
import { EXPERIMENT_STATUS } from 'appConstants';

const runInitialState = fromJS({
	isLoading: false,
	experimentStatus: null,
	experimentStartedTime: null,
	/**
   * experiment completed details
   * e.g completion time, no of wells etc
   */
	data: {},
});

const getTimeNow = () => {
	const date = new Date();
	let hours = date.getHours();
	let minutes = date.getMinutes();
	const ampm = hours >= 12 ? 'pm' : 'am';
	hours %= 12;
	hours = hours || 12; // the hour '0' should be '12'
	minutes = minutes < 10 ? `0${minutes}` : minutes;
	const strTime = `${hours}:${minutes} ${ampm}`;
	return strTime;
};

export const runExperimentReducer = (state = runInitialState, action) => {
	switch (action.type) {
	case runExperimentActions.runExperiment:
		return runInitialState;
	case runExperimentActions.successAction:
		return state.merge({
			isLoading: false,
			experimentStatus: EXPERIMENT_STATUS.running,
			experimentStartedTime: getTimeNow(),
		});
	case runExperimentActions.failureAction:
		return state.merge({
			isLoading: false,
			experimentStatus: EXPERIMENT_STATUS.runFailed,
			experimentStartedTime: getTimeNow(),
		});

		// experiment completed
	case experimentCompleteActions.success:
		return state.merge({
			experimentStatus:
          state.get('experimentStatus') === EXPERIMENT_STATUS.running
          	? EXPERIMENT_STATUS.success
          	: null,
			data: fromJS(action.payload.data),
		});

		// stop experiment actions
	case stopExperimentActions.stopExperiment:
		return state;
	case stopExperimentActions.successAction:
		return state.merge({
			isLoading: false,
			experimentStatus: EXPERIMENT_STATUS.stopped,
			experimentStoppedTime: getTimeNow(),
		});
	case stopExperimentActions.failureAction:
		return runInitialState;
	case loginActions.loginReset:
		return runInitialState;
	default:
		return state;
	}
};
