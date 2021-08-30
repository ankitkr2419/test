import { MAX_PID_TEMP, MIN_PID_TEMP } from "appConstants";

export const tipTubeTypeOptions = [
  { label: "Tip", value: "tip" },
  { label: "Tube", value: "tube" },
];

export const formikInitialState = {
  name: { value: null, isInvalid: false },
  email: { value: null, isInvalid: false },
  roomTemperature: { value: null, isInvalid: false },
  motorNumber: { value: null, isInvalid: false },
  direction: { value: null, isInvalid: false },
  distance: { value: null, isInvalid: false },
  pidTemperature: { value: null, isInvalid: false },

  //TipsTubes fields
  tipTubeName: { value: null, isInvalid: false },
  tipTubeType: { value: tipTubeTypeOptions[0], isInvalid: false },
  allowedPositions: {
    value: {
      1: true,
      2: true,
      3: true,
      4: true,
      5: true,
    },
    isInvalid: false,
  },
  volume: { value: null, isInvalid: false },
  height: { value: null, isInvalid: false },
  ttBase: { value: null, isInvalid: false },
};

/**
 * This method used to convert formik values of allowed positions to api ready response
 * Formik: {1: true, 2: true, 3: false})
 * Output: [1, 2])
 */
export const formikToArray = (allowedPositions) => {
  let keys = Object.keys(allowedPositions.value);
  //take key with true values i.e. selected positions
  let allowedPositionsArr = keys.filter((key) => {
    return allowedPositions.value[key];
  });
  //convert to int array
  return allowedPositionsArr.map((position) => parseInt(position));
};

export const isSaveDetailsBtnDisabled = (state) => {
  const { name, email, roomTemperature } = state;
  if (
    !name.value ||
    !email.value ||
    !roomTemperature.value ||
    name.isInvalid ||
    email.isInvalid ||
    roomTemperature.isInvalid
  ) {
    return true;
  }
  return false;
};

export const isPidUpdateBtnDisabled = (state) => {
  const { pidTemperature } = state;
  const { value, isInvalid } = pidTemperature;
  if (!value || isInvalid || value > MAX_PID_TEMP || value < MIN_PID_TEMP) {
    return true;
  }
  return false;
};

export const isBtnDisabled = (state) => {
  const { motorNumber, direction, distance } = state;
  if (
    !motorNumber.value ||
    !distance.value ||
    motorNumber.isInvalid ||
    distance.isInvalid ||
    direction.isInvalid ||
    direction.value === "" ||
    direction.value === null ||
    direction.value === undefined
  ) {
    return true;
  }
  return false;
};

export const isTipsTubesButtonDisabled = (state) => {
  const { tipTubeName, tipTubeType, allowedPositions, volume, height, ttBase } =
    state;

  let arrayOfAllowedPositions = formikToArray(allowedPositions);

  if (
    !tipTubeName.value ||
    !tipTubeType.value ||
    !volume.value ||
    !height.value ||
    !ttBase.value ||
    tipTubeName.isInvalid ||
    tipTubeType.isInvalid ||
    volume.isInvalid ||
    height.isInvalid ||
    ttBase.isInvalid ||
    arrayOfAllowedPositions.length === 0
  ) {
    return true;
  }
  return false;
};
