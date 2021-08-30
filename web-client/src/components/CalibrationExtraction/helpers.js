import {
  MAX_PID_TEMP,
  MAX_TEMP_ALLOWED,
  MAX_TIME_ALLOWED,
  MIN_PID_TEMP,
  MIN_TEMP_ALLOWED,
  timeConstants,
} from "appConstants";

export const formikInitialState = {
  name: { value: null, isInvalid: false },
  email: { value: null, isInvalid: false },
  roomTemperature: { value: null, isInvalid: false },
  motorNumber: { value: null, isInvalid: false },
  direction: { value: null, isInvalid: false },
  distance: { value: null, isInvalid: false },
  pidTemperature: { value: null, isInvalid: false },
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
