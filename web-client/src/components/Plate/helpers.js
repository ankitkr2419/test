// const action types
export const rangeActions = {
  UPDATE_X_MAX: "UPDATE_X_MAX",
  UPDATE_X_MIN: "UPDATE_X_MIN",
  UPDATE_Y_MAX: "UPDATE_Y_MAX",
  UPDATE_Y_MIN: "UPDATE_Y_MIN",
  RESET_VALUES: "RESET_VALUES",
};

export const rangeInitialState = {
  xMaxValue: 0,
  xMinValue: 1,
  yMaxValue: 10,
  yMinValue: 0,
};

export const rangeReducer = (state, action) => {
  switch (action.type) {
    case rangeActions.UPDATE_X_MAX:
      return { ...state, xMaxValue: parseFloat(action.value) };

    case rangeActions.UPDATE_X_MIN:
      return { ...state, xMinValue: parseFloat(action.value) };

    case rangeActions.UPDATE_Y_MAX:
      return { ...state, yMaxValue: parseFloat(action.value) };

    case rangeActions.UPDATE_Y_MIN:
      return { ...state, yMinValue: parseFloat(action.value) };

    case rangeActions.RESET_VALUES:
      return rangeInitialState;

    default:
      throw new Error("Invalid action type");
  }
};
