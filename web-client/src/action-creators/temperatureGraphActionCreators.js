import temperatureGraphActions from "actions/temperatureGraphActions";

export const temperatureDataSucceeded = (data) => {
  return {
    type: temperatureGraphActions.successAction,
    payload: {
      data,
    },
  };
};
