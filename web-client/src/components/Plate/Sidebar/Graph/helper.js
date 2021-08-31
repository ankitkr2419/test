import { EXPERIMENT_STATUS } from "appConstants";

const initialObj = { value: null, min: 0, max: 99, isInvalid: false };

export const formikInitialState = {
  xMin: initialObj,
  xMax: initialObj,
  yMin: initialObj,
  yMax: initialObj,
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
    status !== EXPERIMENT_STATUS.progressComplete
  );
};
