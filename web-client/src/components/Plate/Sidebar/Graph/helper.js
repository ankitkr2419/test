import { EXPERIMENT_STATUS } from "appConstants";

export const formikInitialState = {
  xMin: { value: null, min: 0, max: 99, isInvalid: false },
  xMax: { value: null, min: 0, max: 99, isInvalid: false },
  yMin: { value: null, min: 0, max: 99, isInvalid: false },
  yMax: { value: null, min: 0, max: 99, isInvalid: false },
};

export const getRequestBody = (state) => ({
  xMax: state.xMax.value,
  xMin: state.xMin.value,
  yMax: state.yMax.value,
  yMin: state.yMin.value,
});

export const disbleApplyBtn = (state, status) => {
  const { xMax, xMin, yMax, yMin } = state;
  return (
    xMax.isInvalid ||
    xMin.isInvalid ||
    yMin.isInvalid ||
    yMax.isInvalid ||
    (status !== EXPERIMENT_STATUS.progressing &&
      status !== EXPERIMENT_STATUS.progressComplete)
  );
};
