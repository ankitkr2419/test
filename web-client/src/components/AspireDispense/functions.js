// footerText can be: "aspire-from" or "selected"

export const getArray = (length, type) => {
  const array = [];
  for (let i = 0; i < length; i++) {
    array.push({
      id: i + 1,
      type: type,
      label: `${i + 1}`,
      footerText: "",
      isSelected: false,
      isDisabled: false,
    });
  }
  return array;
};

export const getFormikInitialState = () => {
  return {
    aspire: {
      cartridge1Wells: getArray(8, 0),
      cartridge2Wells: getArray(8, 0),
      deckPosition: null,
      aspireHeight: null,
      mixingVolume: null,
      aspireVolume: null,
      airVolume: null,
      nCycles: null,
    },
    dispense: {
      cartridge1Wells: getArray(8, 0),
      cartridge2Wells: getArray(8, 0),
      deckPosition: null,
      dispenseHeight: null,
      mixingVolume: null,
      nCycles: null,
    },
  };
};

export const disabledTab = {
  aspire: {
    cartridge1: false,
    cartridge2: false,
    shaker: false,
    deckPosition: false,
  },
  dispense: {
    cartridge1: false,
    cartridge2: false,
    shaker: false,
    deckPosition: false,
  },
};

const getCurrentTabName = (currentTab) => {
  switch (currentTab) {
    case "1":
      return "cartridge1";
    case "2":
      return "cartridge2";
    case "3":
      return "shaker";
    case "4":
      return "deckPosition";
  }
};

//return true if any value is filled or updated, else returns false
//this will check for all keys but currentKey
const checkIsFilled = (formikData, isAspire, currentKey) => {
  let isFilled = false;

  if (currentKey === "cartridge1Wells" || currentKey === "cartridge2Wells") {
    isFilled = true;

    //if cartridge1 is selected in Aspire, cart2 will be discarded in dispense
    const aspireOrDispense = !isAspire ? "aspire" : "dispense";
    const disabledTabInNextPage =
      currentKey === "cartridge1Wells" ? "cartridge1" : "cartridge2";

    disabledTab[aspireOrDispense][disabledTabInNextPage] = true;
  } else {
    for (const key in formikData) {
      if (
        key !== currentKey &&
        key !== "cartridge1Wells" &&
        key !== "cartridge2Wells"
      ) {
        const value = formikData[key];
        isFilled = value ? true : false;
        if (isFilled) break;
      }
    }
  }
  return isFilled;
};

//this function is used to change the disablility of
//different tabs according to different cases
export const toggler = (
  formik,
  isAspire,
  currentTab,
  fieldName,
  fieldValue
) => {
  const currentTabName = getCurrentTabName(currentTab);

  const aspireOrDispense = isAspire ? disabledTab.aspire : disabledTab.dispense;

  if (fieldValue) {
    for (const key in aspireOrDispense) {
      if (key !== currentTabName) {
        aspireOrDispense[key] = true;
      }
    }
  } else {
    const formikData = formik.values[isAspire ? "aspire" : "dispense"];
    const isFilled = checkIsFilled(formikData, isAspire, fieldName);

    if (!isFilled) {
      for (const key in aspireOrDispense) {
        aspireOrDispense[key] = false;
      }
    }
  }
};

export const setFormikField = (
  formik,
  isAspire,
  currentTab,
  fieldName,
  fieldValue
) => {
  const isAspireName = isAspire ? "aspire" : "dispense";

  //set formik field
  const formikFieldKey = `${isAspireName}.${fieldName}`;
  formik.setFieldValue(formikFieldKey, fieldValue);

  const convertedFieldName = fieldName.split(".")[0];
  toggler(formik, isAspire, currentTab, convertedFieldName, fieldValue);
};
