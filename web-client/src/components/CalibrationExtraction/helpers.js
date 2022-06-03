import {
  MAX_CARTRIDGE_ID,
  MIN_CARTRIDGE_ID,
  MAX_WELLS_COUNT,
  MIN_WELLS_COUNT,
  MAX_PID_TEMP,
  MAX_TEMP_ALLOWED,
  MAX_TIME_ALLOWED,
  MIN_PID_TEMP,
  MIN_TEMP_ALLOWED,
  CARTRIDGE_WELLS,
  timeConstants,
} from "appConstants";

export const tipTubeTypeOptions = [
  { label: "Tip", value: "tip" },
  { label: "Tube", value: "tube" },
];

export const formikInitialStateForTipsTubes = {
  tipTubeId: { value: null, isInvalid: false },
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

export const formikInitialState = {
  name: { value: null, isInvalid: false },
  email: { value: null, isInvalid: false },
  roomTemperature: { value: null, isInvalid: false },
  motorNumber: { value: null, isInvalid: false },
  direction: { value: null, isInvalid: false },
  distance: { value: 1, isInvalid: false }, // default value should be 1mm
  pidTemperature: { value: null, isInvalid: false },
  serialNo: { value: null, isInvalid: false },
  manufacturingYear: { value: null, isInvalid: false },
  machineVersion: { value: null, isInvalid: false },
  softwareVersion: { value: null, isInvalid: false },
  contactNumber: { value: null, isInvalid: false },
  ...formikInitialStateForTipsTubes, //TipsTubes fields
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

//formik state for shaker
export const shakerInitialFormikState = {
  withHeating: null,
  temperature: null,
  followTemperature: false,
  rpm1: { value: 0, isInvalid: false },
  rpm2: { value: 0, isInvalid: false },
  hours1: 0,
  mins1: 0,
  secs1: 0,
  hours2: 0,
  mins2: 0,
  secs2: 0,
  rpm2IsDisabled: false,
};

const { SEC_IN_ONE_HOUR, SEC_IN_ONE_MIN, MIN_IN_ONE_HOUR } = timeConstants;

/** This function checks for validity of input data and
 *  returns the request body.
 */
export const getShakerRequestBody = (formik, activeTab) => {
  const formikValues = formik.values;

  const time1 =
    parseInt(formikValues.hours1) * MIN_IN_ONE_HOUR * SEC_IN_ONE_MIN +
    parseInt(formikValues.mins1) * SEC_IN_ONE_MIN +
    parseInt(formikValues.secs1);

  const time2 =
    parseInt(formikValues.hours2) * MIN_IN_ONE_HOUR * SEC_IN_ONE_MIN +
    parseInt(formikValues.mins2) * SEC_IN_ONE_MIN +
    parseInt(formikValues.secs2);

  if (time1 !== 0) {
    if (time1 > MAX_TIME_ALLOWED) {
      return false;
    }
  }

  if (time2 !== 0) {
    if (time2 > MAX_TIME_ALLOWED) {
      return false;
    }
  }

  const rpm1 = parseInt(formikValues.rpm1.value);
  if (!rpm1 || rpm1 === 0) {
    return false;
  }

  const temperature = parseInt(formikValues.temperature);
  if (temperature !== 0) {
    if (temperature < MIN_TEMP_ALLOWED || temperature > MAX_TEMP_ALLOWED) {
      return false;
    }
  }

  const body = {
    with_temp: activeTab === "2",
    temperature: temperature ? temperature : 0,
    follow_temp: formikValues.followTemperature,
    rpm_1: parseInt(formikValues.rpm1.value)
      ? parseInt(formikValues.rpm1.value)
      : 0,
    rpm_2: parseInt(formikValues.rpm2.value)
      ? parseInt(formikValues.rpm2.value)
      : 0,
    time_1: time1 ? time1 : 0,
    time_2: time2 ? time2 : 0,
  };

  return body;
};

/**
 * Helpers for Heater.
 */

//formik state for shaker
export const heaterInitialFormikState = {
  temperature: 0,
  followTemperature: false,
  hours: 0,
  mins: 0,
  secs: 0,
};

export const getHeaterRequestBody = (formik) => {
  const formikValues = formik.values;

  const time =
    parseInt(formikValues.hours) * 60 * 60 +
    parseInt(formikValues.mins) * 60 +
    parseInt(formikValues.secs);

  if (time !== 0) {
    if (time < 10 || time > 3660) {
      return false;
    }
  }

  const temperature = parseInt(formikValues.temperature);
  if (temperature !== 0) {
    if (temperature < 20 || temperature > 120) {
      return false;
    }
  }

  const body = {
    temperature: temperature,
    follow_temp: formikValues.followTemperature,
    duration: time,
  };

  return body;
};

// initial formik state for Update Motor Component
export const updateMotorInitialFormikState = {
  id: {
    label: "ID",
    name: "id",
    type: "number",
    min: 0,
    max: 20,
    value: null,
    isInvalid: false,
  },
  deck: {
    label: "Deck",
    name: "deck",
    type: "text",
    min: null,
    max: null,
    value: null,
    isInvalid: false,
  },
  number: {
    label: "Number",
    name: "number",
    type: "number",
    min: 0,
    max: 10,
    value: null,
    isInvalid: false,
  },
  name: {
    label: "Name",
    name: "name",
    type: "text",
    min: null,
    max: null,
    value: null,
    isInvalid: false,
  },
  ramp: {
    label: "Ramp",
    name: "ramp",
    type: "number",
    min: 1,
    max: 3000,
    value: null,
    isInvalid: false,
  },
  steps: {
    label: "Steps",
    name: "steps",
    type: "number",
    min: 1,
    max: 3000,
    value: null,
    isInvalid: false,
  },
  slow: {
    label: "Slow",
    name: "slow",
    type: "number",
    min: 100,
    max: 9000,
    value: null,
    isInvalid: false,
  },
  fast: {
    label: "Fast",
    name: "fast",
    type: "number",
    min: 100,
    max: 16000,
    value: null,
    isInvalid: false,
  },
};

export const checkIsFieldInvalid = (fieldObj, value) => {
  const { name, max, min, type } = fieldObj;

  // if type is number => check for max and min
  if (type === "number") {
    if (!parseInt(value) || parseInt(value) > max || parseInt(value) < min) {
      return true;
    }
  }
  // else check accordingly for string type => deck and name
  else {
    // for deck value should only be either A or B
    if (name === "deck") {
      if (
        value === "" ||
        (value !== "A" && value !== "B" && value !== "a" && value !== "b")
      ) {
        return true;
      }
    }
    // for name and id => value should not be empty string
    else {
      if (value === "") {
        return true;
      }
    }
  }

  return false;
};

export const isMotorUpdateBtnDisabled = (state) => {
  let isInvalid = false;
  Object.keys(state).forEach((key) => {
    const element = state[key];
    isInvalid = checkIsFieldInvalid(element, element.value);
    if (isInvalid === true) {
      return;
    }
  });
  return isInvalid;
};

export const isTipsTubesButtonDisabled = (state) => {
  const {
    tipTubeId,
    tipTubeName,
    tipTubeType,
    allowedPositions,
    volume,
    height,
    ttBase,
  } = state;

  let arrayOfAllowedPositions = formikToArray(allowedPositions);

  if (
    !tipTubeId.value ||
    !tipTubeName.value ||
    !tipTubeType.value ||
    !volume.value ||
    !height.value ||
    !ttBase.value ||
    tipTubeId.isInvalid ||
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

// helpers for cartridges

// formik initial state
export const cartridgeFormikInitialState = {
  id: { value: null, isInvalid: false },
  type: { value: "Cartridge 1", isInvalid: false },
  description: { value: null, isInvalid: false },
  wellsCount: { value: null, isInvalid: false },
  distance: [],
  height: [],
  volume: [],
};

export const checkIsCartridgeFieldInvalid = (key, value) => {
  const filteredKey = key.split(".")[0];

  const {
    MAX_DISTANCE,
    MIN_DISTANCE,
    MAX_HEIGHT,
    MIN_HEIGHT,
    MAX_VOLUME,
    MIN_VOLUME,
  } = CARTRIDGE_WELLS;

  switch (filteredKey) {
    case "id":
      if (!Number.isInteger(parseFloat(value))) {
        return true;
      }
      return false;

    case "type":
      if (value === "" || value === undefined || value === null) {
        return true;
      }
      return false;

    case "wellsCount":
      const wellsCount = parseInt(value);
      if (!wellsCount || wellsCount > 13 || wellsCount < 0) {
        return true;
      }
      return false;

    case "distance":
      const distance = parseFloat(value);
      if (!distance || distance > MAX_DISTANCE || distance < MIN_DISTANCE) {
        return true;
      }
      return false;

    case "volume":
      const volume = parseFloat(value);
      if (
        !volume ||
        !Number.isInteger(volume) ||
        volume > MAX_VOLUME ||
        volume < MIN_VOLUME
      ) {
        return true;
      }
      return false;

    case "height":
      const height = parseFloat(value);
      if (!height || height > MAX_HEIGHT || height < MIN_HEIGHT) {
        return true;
      }
      return false;
  }
};

export const isAddWellsBtnDisabled = (state) => {
  const { id, type, wellsCount } = state;

  if (
    type.value === "" ||
    !parseInt(id.value) ||
    parseInt(id.value) > MAX_CARTRIDGE_ID ||
    parseInt(id.value) < MIN_CARTRIDGE_ID ||
    !parseInt(wellsCount.value) ||
    parseInt(wellsCount.value) > MAX_WELLS_COUNT ||
    parseInt(wellsCount.value) < MIN_WELLS_COUNT ||
    id.isInvalid ||
    wellsCount.isInvalid ||
    type.isInvalid
  ) {
    return true;
  }
  return false;
};

export const isCreateCartridgesBtnDisabled = (state) => {
  const { distance, volume, height } = state;

  const {
    MAX_DISTANCE,
    MIN_DISTANCE,
    MAX_HEIGHT,
    MIN_HEIGHT,
    MAX_VOLUME,
    MIN_VOLUME,
  } = CARTRIDGE_WELLS;

  // get array of all distance, height and volume values
  const allDistanceValues = distance.map((disObj) => disObj.value);
  const allHeightValues = height.map((htObj) => htObj.value);
  const allVolumeValues = volume.map((volObj) => volObj.value);

  // check if any value in allDistanceValues array is out-of-range
  const isAnyDistanceOutOfRange = allDistanceValues.some(
    (distanceValue) =>
      !distanceValue ||
      distanceValue > MAX_DISTANCE ||
      distanceValue < MIN_DISTANCE
  );
  // check if any value in allVolumeValues array is out-of-range
  const isAnyVolumeOutOfRange = allVolumeValues.some(
    (volumeValue) =>
      !volumeValue || volumeValue > MAX_VOLUME || volumeValue < MIN_VOLUME
  );
  // check if any value in allHeightValues array is out-of-range
  const isAnyHeightOutOfRange = allHeightValues.some(
    (heightValue) =>
      !heightValue || heightValue > MAX_HEIGHT || heightValue < MIN_HEIGHT
  );

  // check if any value is invalid
  const allDistanceInvalidValues = distance.map((disObj) => disObj.isInvalid);
  const allHeightInvalidValues = height.map((htObj) => htObj.isInvalid);
  const allVolumeInvalidValues = volume.map((volObj) => volObj.isInvalid);

  if (
    isAnyDistanceOutOfRange ||
    isAnyHeightOutOfRange ||
    isAnyVolumeOutOfRange ||
    allDistanceInvalidValues.some((distance) => distance === true) ||
    allHeightInvalidValues.some((height) => height === true) ||
    allVolumeInvalidValues.some((volume) => volume === true)
  ) {
    return true;
  }
  return false;
};

const getCartridgeWellsBody = (state) => {
  const { id, distance, height, volume, wellsCount } = state;

  const cartridgeWells = [];

  for (let index = 0; index < parseInt(wellsCount.value); index++) {
    cartridgeWells[index] = {
      id: parseInt(id.value),
      well_num: parseInt(index + 1),
      distance: parseFloat(distance[index].value),
      height: parseInt(height[index].value),
      volume: parseInt(volume[index].value),
    };
  }
  return cartridgeWells;
};

export const getRequestBody = (state) => {
  const { id, type, description, wellsCount } = state;

  const requestBody = {
    cartridges: [
      {
        id: parseInt(id.value),
        type: type.value === "Cartridge 1" ? "cartridge_1" : "cartridge_2",
        description: description.value,
        wells_count: parseInt(wellsCount.value),
      },
    ],
    cartridge_wells: getCartridgeWellsBody(state),
  };

  return requestBody;
};

// consumable formik initial state
export const consumableFormikInitialState = {
  id: { value: null, isInvalid: false },
  name: { value: null, isInvalid: false },
  description: { value: null, isInvalid: false },
  distance: { value: null, isInvalid: false },
};

export const checkConsumableFieldIsInvalid = (name, value) => {
  switch (name) {
    case "id":
      if (value === "" || !Number.isInteger(parseFloat(value))) {
        return true;
      }
      return false;
    case "name":
      if (value === "") {
        return true;
      }
      return false;
    case "description":
      if (value === "") {
        return true;
      }
      return false;
    case "distance":
      if (value === "" || Number.isNaN(parseFloat(value))) {
        return true;
      }
      return false;
    default:
      break;
  }
};

export const isConsumableModalBtnDisabled = (state) => {
  const { id, name, description, distance } = state;

  if (
    id.isInvalid ||
    name.isInvalid ||
    description.isInvalid ||
    distance.isInvalid ||
    !id.value ||
    !name.value ||
    !description.value ||
    !distance.value
  ) {
    return true;
  }
  return false;
};
