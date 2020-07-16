import { fromJS } from 'immutable';
import {
	runExperimentActions,
	stopExperimentActions,
	experimentCompleteActions,
} from 'actions/runExperimentActions';
import loginActions from 'actions/loginActions';
import { EXPERIMENT_STATUS } from 'appConstants';
import { getTimeNow } from 'selectors/runExperimentSelector';

const runInitialState = fromJS({
	isLoading: false,
	experimentStatus: null,
	experimentStartedTime: null,
	/**
   * experiment completed details
   * e.g completion time, no of wells etc
   */
	data: {},
	/**
   * experiment failed details
   */
	failedData: null,
});

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

		// experiment failed
	case experimentCompleteActions.failed:
		return state.merge({
			experimentStatus:
					state.get('experimentStatus') === EXPERIMENT_STATUS.running
						? EXPERIMENT_STATUS.failed
						: null,
			failedData: fromJS(action.payload.data),
		});

	// stop experiment actions (abort)
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
