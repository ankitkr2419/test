export const constants = {
  ROOM_TEMPERATURE: {
    min: 20,
    max: 30,
  },
  HOMING_TIME: {
    min: 16,
    max: 30,
  },
  NO_OF_HOMING_CYCLE: {
    min: 0,
    max: 100,
  },
};

export const isValidRoomTemp = (value) => {
  return (
    value !== null &&
    value >= constants.ROOM_TEMPERATURE.min &&
    value <= constants.ROOM_TEMPERATURE.max
  );
};

export const isValidHomingTime = (value) => {
  return (
    value !== null &&
    value >= constants.HOMING_TIME.min &&
    value <= constants.HOMING_TIME.max
  );
};

export const isValidNoOfHomingCycle = (value) => {
  return (
    value !== null &&
    value >= constants.NO_OF_HOMING_CYCLE.min &&
    value <= constants.NO_OF_HOMING_CYCLE.max
  );
};
