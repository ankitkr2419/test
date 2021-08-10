export const formikInitialState = {
  motorNumber: { value: null, isInvalid: false },
  direction: { value: null, isInvalid: false },
  distance: { value: null, isInvalid: false },
};

export const isBtnDisabled = (state) => {
  const { motorNumber, direction, distance } = state;
  if (
    motorNumber.isInvalid ||
    direction.isInvalid ||
    distance.isInvalid ||
    !motorNumber.value ||
    !direction.value ||
    !distance.value
  ) {
    return true;
  }
  return false;
};
