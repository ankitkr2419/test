import { EXPERIMENT_STATUS } from "appConstants";

const initialObj = { value: null, min: 0, max: 99, isInvalid: false };

export const formikInitialState = {
  xMin: initialObj,
  xMax: initialObj,
  yMin: initialObj,
  yMax: initialObj,
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
    (isExpanded === false &&
      (status === null || status === EXPERIMENT_STATUS.progressing)) ||
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

export const disbleResetBtn = (status, isExpanded) => {
  if (
    isExpanded === false &&
    (status === EXPERIMENT_STATUS.progressing || status === null)
  ) {
    return true;
  }
  return false;
};
