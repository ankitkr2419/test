export const formikInitialState = {
  roomTemperature: {
    type: "number",
    name: "roomTemperature",
    apiKey: "room_temperature",
    label: "Room Temperature",
    value: null,
    min: 20,
    max: 30,
    isInvalid: false,
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
  },
};

export const validateAllFields = (state) => {
  for (const key in state) {
    const { name, value, isInvalid } = state[key];
    if (
      value === null ||
      value === "" ||
      isInvalid === true ||
      isValueWithinRange(name, value) === false
    ) {
      return false;
    }
  }
  return true;
};

export const isValueWithinRange = (name, value) => {
  const element = formikInitialState[name];
  const { min, max } = element;
  if (value < min || value > max) {
    return false;
  }
  return true;
};

export const getRequestBody = (state) => {
  const body = {};
  for (const key in state) {
    const element = state[key];
    const { name, value } = element;
    body[name] = parseInt(value);
  }
  return body;
};
