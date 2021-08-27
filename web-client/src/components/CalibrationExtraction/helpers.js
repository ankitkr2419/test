export const formikInitialState = {
  name: { value: null, isInvalid: false },
  email: { value: null, isInvalid: false },
  roomTemperature: { value: null, isInvalid: false },
  motorNumber: { value: null, isInvalid: false },
  direction: { value: null, isInvalid: false },
  distance: { value: null, isInvalid: false },
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
