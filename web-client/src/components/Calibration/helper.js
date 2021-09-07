import { EMAIL_REGEX_OR_EMPTY_STR, NAME_REGEX } from "appConstants";
import { PID_STATUS } from "appConstants";

export const formikInitialState = {
  name: {
    type: "text",
    name: "name",
    apiKey: "receiver_name",
    label: "Name",
    value: null,
    isInvalid: false,
    isInvalidMsg: "Please enter a valid name",
  },
  email: {
    type: "email",
    name: "email",
    apiKey: "receiver_email",
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
};

export const formikInitialStateRtpcrVars = {
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
  pidLidTemperature: {
    type: "number",
    name: "pidLidTemperature",
    apiKey: "pid_lid_temp",
    label: "PID Lid Temperature",
    min: 50,
    max: 75,
    value: null,
    isInvalid: false,
    isInvalidMsg: "PID Lid temperature should be between 50 - 75",
  },
  scanSpeed: {
    type: "number",
    name: "scanSpeed",
    apiKey: "scan_speed",
    label: "Scan Speed",
    min: 0,
    max: 9999,
    value: null,
    isInvalid: false,
    isInvalidMsg: "Scan Speed should be between 0 - 9999",
  },
  scanTime: {
    type: "number",
    name: "scanTime",
    apiKey: "scan_time",
    label: "Scan Time",
    min: 0,
    max: 9999,
    value: null,
    isInvalid: false,
    isInvalidMsg: "Scan Time should be between 0 - 9999",
  },
  startCycle: {
    type: "number",
    name: "startCycle",
    apiKey: "start_cycle",
    label: "Start Cycle",
    min: 0,
    max: 9999,
    value: null,
    isInvalid: false,
    isInvalidMsg: "Start Cycle should be between 0 - 9999",
  },
  endCycle: {
    type: "number",
    name: "endCycle",
    apiKey: "end_cycle",
    label: "End Cycle",
    min: 0,
    max: 9999,
    value: null,
    isInvalid: false,
    isInvalidMsg: "End Cycle should be between 0 - 9999",
  },
  opticalCalibrationCycleCount: {
    type: "number",
    name: "opticalCalibrationCycleCount",
    apiKey: "optical_calibration_cycle_count",
    label: "Optical Calibration Cycle Count",
    min: 0,
    max: 9999,
    value: null,
    isInvalid: false,
    isInvalidMsg: "Optical Calibration Cycle Count should be between 0 - 9999",
  },
};

//TEC variables
export const formikInitialStateTECVars = {
  currentLimitation: {
    type: "number",
    name: "currentLimitation",
    apiKey: "current_limitation",
    label: "Current Limitation",
    min: 0,
    max: 9999,
    value: null,
    isInvalid: false,
    isInvalidMsg: "Current Limitation should be between 0 - 9999",
  },
  voltageLimitation: {
    type: "number",
    name: "voltageLimitation",
    apiKey: "voltage_limitation",
    label: "Voltage Limitation",
    min: 0,
    max: 9999,
    value: null,
    isInvalid: false,
    isInvalidMsg: "Voltage Limitation should be between 0 - 9999",
  },
  currentErrorThreshold: {
    type: "number",
    name: "currentErrorThreshold",
    apiKey: "current_error_threshold",
    label: "Current Error Threshold",
    min: 0,
    max: 9999,
    value: null,
    isInvalid: false,
    isInvalidMsg: "Current Error Threshold should be between 0 - 9999",
  },
  voltageErrorThreshold: {
    type: "number",
    name: "voltageErrorThreshold",
    apiKey: "voltage_error_threshold",
    label: "Voltage Error Threshold",
    min: 0,
    max: 9999,
    value: null,
    isInvalid: false,
    isInvalidMsg: "Voltage Error Threshold should be between 0 - 9999",
  },
  peltierMaxCurrent: {
    type: "number",
    name: "peltierMaxCurrent",
    apiKey: "peltier_max_current",
    label: "Peltier Max Current",
    min: 0,
    max: 9999,
    value: null,
    isInvalid: false,
    isInvalidMsg: "Peltier Max Current should be between 0 - 9999",
  },
  peltierDeltaTemperature: {
    type: "number",
    name: "peltierDeltaTemperature",
    apiKey: "peltier_delta_temperature",
    label: "Peltier Delta Temperature",
    min: 0,
    max: 9999,
    value: null,
    isInvalid: false,
    isInvalidMsg: "Peltier Delta Temperature should be between 0 - 9999",
  },
};

export const formikInitialStateDyeCalibration = {
  selectedDye: { value: null, isInvalid: false },
  kitID: { value: null, min: 0, isInvalid: false },
};

export const getToleranceInitialFormikState = (data) => {
  const toleranceArr = data?.map((dataObj) => ({
    value: dataObj.Tolerance,
    isInvalid: false,
  }));
  return {
    tolerance: toleranceArr,
  };
};

export const isSaveToleranceBtnDisabled = (tolerance) => {
  const allTolValues = tolerance.map((tolObj) => tolObj.value);
  const allTolInvalidValues = tolerance.map((tolObj) => tolObj.isInvalid);

  if (
    allTolValues.some((v) => !v || v > 100 || v < 0) ||
    allTolInvalidValues.some((v) => v === true)
  ) {
    return true;
  }
  return false;
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
  } else if (
    type === "email" &&
    value.match(EMAIL_REGEX_OR_EMPTY_STR) === null
  ) {
    return false;
  }
  //TODO : check if its required anymore else delete
  // else if (type === "text") {
  //   //&& value.match(NAME_REGEX) === null) {
  //   return false;
  // }
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

/**create dye list from tolerance data to use it in dropdown */
export const createDyeOptions = (toleranceReducerData) => {
  let dyeList = [];
  if (toleranceReducerData?.length) {
    dyeList = toleranceReducerData?.map((obj) => ({
      label: obj.Name,
      value: obj.ID,
    }));
  }
  return dyeList;
};

export const calibrationStatusMessage = (dyeCalibrationStatus) => {
  switch (dyeCalibrationStatus) {
    case PID_STATUS.running:
      return "Calibration Running";
    case PID_STATUS.runFailed:
      return "Calibration Failed";
    case PID_STATUS.progressing:
      return "Calibration In Progress";
    case PID_STATUS.progressComplete:
      return "Calibration Completed";
    default:
      return "";
  }
};
