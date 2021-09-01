import {
  temperatureDataSuccess,
  temperatureGraphActions,
} from "actions/temperatureGraphActions";

export const temperatureDataSucceeded = (data) => {
  return {
    type: temperatureDataSuccess.successAction,
    payload: {
      data,
    },
  };
};

export const temperatureApiGraphInitiated = ({ experimentId, token }) => {
  return {
    type: temperatureGraphActions.temperatureGraphInitated,
    payload: {
      experimentId,
      token,
    },
  };
};

export const temperatureApiGraphSucceeded = ({ response }) => {
  return {
    type: temperatureGraphActions.temperatureGraphSuccess,
    payload: { response },
  };
};

export const temperatureApiGraphFailed = ({ error }) => {
  return {
    type: temperatureGraphActions.temperatureGraphFailed,
    payload: error,
  };
};
