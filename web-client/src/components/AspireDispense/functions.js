export const getRequestBody = (activeTab, aspire, dispense) => {
  const aspireSelectedTabName = getCategoryName(aspire.selectedCategory);
  const dispenseSelectedTabName = getCategoryName(activeTab);

  /** Aspire category is maintained using formik.
   *  Dispense category is directly maintained using 'activeTab' state.
   */

  const aspireWells = aspire[`cartridge${aspire.selectedCategory}Wells`];
  const dispenseWells = dispense[`cartridge${activeTab}Wells`];

  return {
    category: `${aspireSelectedTabName}_to_${dispenseSelectedTabName}`,
    cartridge_type: `cartridge_${activeTab}`,
    source_position: getPosition(aspireWells),
    aspire_height: parseFloat(aspire.aspireHeight ? aspire.aspireHeight : 0),
    aspire_mixing_volume: parseFloat(
      aspire.mixingVolume ? aspire.mixingVolume : 0
    ),
    aspire_no_of_cycles: parseFloat(aspire.nCycles ? aspire.nCycles : 0),
    aspire_volume: parseFloat(aspire.aspireVolume ? aspire.aspireVolume : 0),
    aspire_air_volume: parseFloat(aspire.airVolume ? aspire.airVolume : 0),
    dispense_height: parseFloat(
      dispense.dispenseHeight ? dispense.dispenseHeight : 0
    ),
    dispense_mixing_volume: parseFloat(
      dispense.mixingVolume ? dispense.mixingVolume : 0
    ),
    dispense_no_of_cycles: parseFloat(dispense.nCycles ? dispense.nCycles : 0),
    destination_position: getPosition(dispenseWells),
  };
};

export const getCategoryLabel = (tabID) => {
  switch (tabID) {
    case "1":
      return "Category 1";
    case "2":
      return "Category 2";
    case "3":
      return "Shaker";
    case "4":
      return "Deck Position";
    default:
      return;
  }
};

export const getCategoryName = (tabID) => {
  switch (tabID) {
    case "1":
      return "well";
    case "2":
      return "well";
    case "3":
      return "shaker";
    case "4":
      return "deck";
    default:
      return;
  }
};

export const getPosition = (wells) => {
  const selectedWell = wells.find((wellObj) => wellObj.isSelected);
  return selectedWell.id;
};

// footerText can be: "aspire-from" or "selected"
export const getArray = (length, type, selectedPosition = null) => {
  const array = [];

  for (let i = 0; i < length; i++) {
    let isSelected = false;
    if (selectedPosition && i + 1 === selectedPosition) {
      isSelected = true;
    }
    array.push({
      id: i + 1,
      type: type,
      label: `${i + 1}`,
      footerText: "",
      isSelected: isSelected,
      isDisabled: false,
    });
  }
  return array;
};

export const getFormikInitialState = (editReducer = null) => {
  let type;
  if (editReducer?.process_id) {
    type = editReducer.cartridge_type;
  }

  return {
    aspire: {
      cartridge1Wells:
        type === "cartridge_1"
          ? getArray(8, 0, editReducer.source_position)
          : getArray(8, 0),
      cartridge2Wells:
        type === "cartridge_2"
          ? getArray(8, 0, editReducer.source_position)
          : getArray(8, 0),
      deckPosition: "",
      aspireHeight: editReducer?.aspire_height ? editReducer.aspire_height : "",
      mixingVolume: editReducer?.aspire_mixing_volume
        ? editReducer.aspire_mixing_volume
        : "",
      aspireVolume: editReducer?.aspire_volume ? editReducer.aspire_volume : "",
      airVolume: editReducer?.aspire_air_volume
        ? editReducer.aspire_air_volume
        : "",
      nCycles: editReducer?.aspire_no_of_cycles
        ? editReducer.aspire_no_of_cycles
        : "",
      selectedCategory: "",
    },
    dispense: {
      cartridge1Wells: getArray(8, 0, editReducer.destination_position),
      cartridge2Wells: getArray(8, 0, editReducer.destination_position),
      deckPosition: "",
      dispenseHeight: editReducer?.dispense_height
        ? editReducer.dispense_height
        : "",
      mixingVolume: editReducer?.dispense_mixing_volume
        ? editReducer.dispense_mixing_volume
        : "",
      nCycles: editReducer?.dispense_no_of_cycles
        ? editReducer.dispense_no_of_cycles
        : "",
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
    default:
      return;
  }
};

//return true if any value is filled or updated, else returns false
//this will check for all keys except currentKey
const checkIsFilled = (formikData, isAspire, currentKey) => {
  let isFilled = false;

  if (currentKey === "cartridge1Wells" || currentKey === "cartridge2Wells") {
    isFilled = true;

    //if cartridge1 is selected in Aspire, cart2 will be discarded in dispense
    const aspireOrDispense = !isAspire ? "aspire" : "dispense";
    const disabledTabInNextPage =
      currentKey === "cartridge1Wells" ? "cartridge2" : "cartridge1";

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
