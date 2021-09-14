import { EXPERIMENT_STATUS } from "appConstants";

export const getFormikInitialState = (maxSteps) => {
  return {
    xMin: { value: 0, min: 0, max: 99, isInvalid: false },
    xMax: { value: maxSteps || 99, min: 0, max: 99, isInvalid: false },
    yMin: { value: 0, min: 0, max: 99, isInvalid: false },
    yMax: { value: 10, min: 0, max: 99, isInvalid: false },
  };
};

export const getRequestBody = (state) => ({
  xMax: parseInt(state.xMax.value),
  xMin: parseInt(state.xMin.value),
  yMax: parseFloat(state.yMax.value),
  yMin: parseFloat(state.yMin.value),
});

export const disbleApplyBtn = (state, status, isExpanded) => {
  const { xMax, xMin, yMax, yMin } = state;

  /*
  Disable btn if -> 
    1. Any value is invalid (or) 
    2. All the values are empty (or) 
    3. Experiment is progressing 
  */
  return (
    (isExpanded === false &&
      (status === null || status === EXPERIMENT_STATUS.progressing)) ||
    ((xMax.value === "" || xMax.value === null) &&
      (xMin.value === "" || xMin.value === null) &&
      (yMax.value === "" || yMax.value === null) &&
      (yMin.value === "" || yMin.value === null)) ||
    xMax.isInvalid ||
    xMin.isInvalid ||
    yMin.isInvalid ||
    yMax.isInvalid
  );
};

export const disbleResetBtn = (status, isExpanded) => {
  if (
    isExpanded === false &&
    (status === EXPERIMENT_STATUS.progressing || status === null)
  ) {
    return true;
  }
  return false;
};
