import { fromJS } from "immutable";
import {
  runExperimentInProgressActions,
  runExperimentActions,
  stopExperimentActions,
  experimentCompleteActions,
} from "actions/runExperimentActions";
import loginActions from "actions/loginActions";
import { EXPERIMENT_STATUS } from "appConstants";
import { getTimeNow } from "selectors/runExperimentSelector";

const runProgressInitialState = fromJS({
  progressStatus: null,
  progress: 0,
  remainingTime: null,
  totalTime: null,
});

export const runExpProgressReducer = (
  state = runProgressInitialState,
  action
) => {
  switch (action.type) {
    case runExperimentInProgressActions.runExperimentProgressAction:
      const { progressDetails } = action.payload;

      return state.merge({
        progressStatus: EXPERIMENT_STATUS.progressing,
        progress: progressDetails.progress,
        remainingTime: progressDetails.remaining_time,
        totalTime: progressDetails.total_time,
      });

    case runExperimentInProgressActions.runExperimentProgressSuccessAction:
      const { progressSucceeded } = action.payload;

      return state.merge({
        progressStatus: EXPERIMENT_STATUS.progressComplete,
        progress: progressSucceeded.progress,
        remainingTime: progressSucceeded.remaining_time,
        totalTime: progressSucceeded.total_time,
      });

    case loginActions.loginReset:
      return runProgressInitialState;

    default:
      return state;
  }
};

const runInitialState = fromJS({
  isLoading: false,
  experimentStatus: null,
  experimentStartedTime: null,
  experimentStoppedTime: null,
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
          state.get("experimentStatus") === EXPERIMENT_STATUS.running
            ? EXPERIMENT_STATUS.success
            : null,
        data: fromJS(action.payload.data),
      });

    // experiment failed
    case experimentCompleteActions.failed:
      return state.merge({
        experimentStatus:
          state.get("experimentStatus") === EXPERIMENT_STATUS.running
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
