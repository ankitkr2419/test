import { EMAIL_REGEX, NAME_REGEX } from "appConstants";

export const formikInitialState = {
  name: {
    type: "text",
    name: "name",
    apiKey: "name",
    label: "Name",
    value: null,
    isInvalid: false,
    isInvalidMsg: "Please enter a valid name",
  },
  email: {
    type: "email",
    name: "email",
    apiKey: "email",
    label: "Email",
    value: null,
    isInvalid: false,
    isInvalidMsg: "Email is invalid",
  },
  roomTemperature: {
    type: "number",
    name: "roomTemperature",
    apiKey: "room_temperature",
    label: "Room Temperature",
    value: null,
    min: 20,
    max: 30,
    isInvalid: false,
    isInvalidMsg: "Room temperature should be between 20 - 30",
  },
  homingTime: {
    type: "number",
    name: "homingTime",
    apiKey: "homing_time",
    label: "Homing Time",
    min: 16,
    max: 30,
    value: null,
    isInvalid: false,
    isInvalidMsg: "Homing time should be between 16 - 30",
  },
  noOfHomingCycles: {
    type: "number",
    name: "noOfHomingCycles",
    apiKey: "no_of_homing_cycles",
    label: "No. Of Homing Cycles",
    min: 0,
    max: 100,
    value: null,
    isInvalid: false,
    isInvalidMsg: "No. of homing cycles should be between 0 - 100",
  },
  cycleTime: {
    type: "number",
    name: "cycleTime",
    apiKey: "cycle_time",
    label: "Cycle Time",
    min: 2,
    max: 30,
    value: null,
    isInvalid: false,
    isInvalidMsg: "Cycle time should be between 2 - 30",
  },
  pidTemperature: {
    type: "number",
    name: "pidTemperature",
    apiKey: "pid_temperature",
    label: "PID Temperature",
    min: 50,
    max: 75,
    value: null,
    isInvalid: false,
    isInvalidMsg: "PID temperature should be between 50 - 75",
  },
  pidMinutes: {
    type: "number",
    name: "pidMinutes",
    apiKey: "pid_minutes",
    label: "PID Minutes",
    min: 20,
    max: 40,
    value: null,
    isInvalid: false,
    isInvalidMsg: "PID minutes should be between 20 - 40",
  },
};

export const validateAllFields = (state) => {
  for (const key in state) {
    const { name, value, isInvalid } = state[key];
    if (
      value === null ||
      value === "" ||
      isInvalid === true ||
      isValueValid(name, value) === false
    ) {
      return false;
    }
  }
  return true;
};

export const isValueValid = (name, value) => {
  const element = formikInitialState[name];
  const { type, min, max } = element;

  if (type === "number" && (value < min || value > max)) {
    return false;
  } else if (type === "email" && !value && value.match(EMAIL_REGEX) === null) {
    return false;
  } else if (type === "text" && value === "") {
    //&& value.match(NAME_REGEX) === null) {
    return false;
  }
  return true;
};

export const getRequestBody = (state) => {
  const body = {};
  for (const key in state) {
    const element = state[key];
    const { type, name, value } = element;
    body[name] = type === "number" ? parseInt(value) : value;
  }
  return body;
};
