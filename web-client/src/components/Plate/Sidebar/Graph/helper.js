import { EXPERIMENT_STATUS } from "appConstants";

export const formikInitialState = {
  xMin: { value: null, min: 0, max: 99, isInvalid: false },
  xMax: { value: null, min: 0, max: 99, isInvalid: false },
  yMin: { value: null, min: 0, max: 99, isInvalid: false },
  yMax: { value: null, min: 0, max: 99, isInvalid: false },
};

export const getRequestBody = (state) => ({
  xMax: parseInt(state.xMax.value),
  xMin: parseInt(state.xMin.value),
  yMax: parseFloat(state.yMax.value),
  yMin: parseFloat(state.yMin.value),
});

export const disbleApplyBtn = (state, status, isExpanded) => {
  const { xMax, xMin, yMax, yMin } = state;

  return (
    xMax.value === "" ||
    xMin.value === "" ||
    yMax.value === "" ||
    yMin.value === "" ||
    xMax.value === null ||
    xMin.value === null ||
    yMax.value === null ||
    yMin.value === null ||
    xMax.isInvalid ||
    xMin.isInvalid ||
    yMin.isInvalid ||
    yMax.isInvalid
  );
};
