export const changeInputValOnClick = (value) => {
  let newValue = InputValue + value;
  if (newValue < 11 && newValue > 0) {
    setInputValue(newValue);
  }
};

export const changeInputValOnEdit = (value) => {
  if (value < 11 && value > 0) {
    setInputValue(parseInt(value));
  }
};
