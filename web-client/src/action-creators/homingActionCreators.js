import {
  homingActions
} from "actions/homingActions";

export const homingActionInitiated = () => ({
  type: homingActions.homingActionInitiated,
  payload: {}
});

export const homingActionSuccess = (homingData) => ({
  type: homingActions.homingActionSuccess,
  payload: {
    homingData
  }
});

export const homingActionFailed = (serverErrors) => ({
  type: homingActions.homingActionFailed,
  payload: {
    serverErrors
  }
})
