import { CATEGORY_NAME } from "appConstants";

// returns the position (index) of selected well
export const getPosition = (wells) => {
  if (wells) {
    const selectedWell = wells.find((wellObj) => wellObj.isSelected);
    return selectedWell ? selectedWell.id : 0;
  }
  return 0;
};

// used to generated request body for API call.
export const getRequestBody = (activeTab, aspire, dispense) => {
  const aspireSelectedTabName = CATEGORY_NAME[aspire.selectedCategory];
  const dispenseSelectedTabName = CATEGORY_NAME[dispense.selectedCategory];

  const aspireWells = aspire[`cartridge${aspire.selectedCategory}Wells`];
  const dispenseWells = dispense[`cartridge${dispense.selectedCategory}Wells`];

  let cartridgeType = 1;
  if (aspire.selectedCategory === "1" || aspire.selectedCategory === "2") {
    cartridgeType = aspire.selectedCategory;
  }
  if (dispense.selectedCategory === "1" || dispense.selectedCategory === "2") {
    cartridgeType = dispense.selectedCategory;
  }

  return {
    category: `${aspireSelectedTabName}_to_${dispenseSelectedTabName}`,
    cartridge_type: `cartridge_${cartridgeType}`,
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

// footerText can be: "aspire-from" or "selected"
// generates array of objects for wells.
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

// this function generates the initial formik state according to the
// operation being performeed i.e. if NEW process is being created than
// empty data is loaded in formikState else for EDIT old values are loaded.
export const getFormikInitialState = (editReducer = null) => {
  let type = "category_1";
  let category, category1, category2;
  let aspireSelectedCategory = "1";
  let dispenseSelectedCategory = "1";

  const CATEGORY_ID = {
    well: type === "category_1" ? "1" : "2",
    shaker: "3",
    deck: "4",
  };

  if (editReducer?.process_id) {
    type = editReducer.cartridge_type;

    category = editReducer.category.split("_");
    category1 = category[0];
    category2 = category[2];

    aspireSelectedCategory = CATEGORY_ID[category1];
    dispenseSelectedCategory = CATEGORY_ID[category2];
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
      selectedCategory: aspireSelectedCategory,
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
      selectedCategory: dispenseSelectedCategory,
    },
  };
};

// initial state of all tabs.
export const disabledTabInitTab = {
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

//return true if any value is filled or updated, else returns false
//this will check for all keys except currentKey
const checkIsFilled = (formikData, currentKey = null) => {
  let isFilled = false;

  // if key is wells array
  if (currentKey === "cartridge1Wells" || currentKey === "cartridge2Wells") {
    isFilled = !formikData[`${currentKey}`].every(
      (wellObj) => wellObj.isSelected === false
    );
  }
  // other than wells array.
  else {
    for (const key in formikData) {
      if (
        key !== currentKey &&
        key !== "cartridge1Wells" &&
        key !== "cartridge2Wells" &&
        key !== "selectedCategory" &&
        key !== "nCycles"
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
export const toggler = (formik, isAspire) => {
  const tabNames = {
    1: "cartridge1",
    2: "cartridge2",
    3: "shaker",
    4: "deckPosition",
  };

  const disabledState = disabledTabInitTab;
  const formikData = formik.values[isAspire ? "aspire" : "dispense"];

  const currentTab = formikData.selectedCategory;
  const currentTabName = tabNames[currentTab];
  const aspireOrDispense = isAspire
    ? disabledState.aspire
    : disabledState.dispense;

  let isFilled = false;

  //check for cartridge
  if (currentTabName === "cartridge1" || currentTabName === "cartridge2") {
    const cartWells =
      currentTabName === "cartridge1" ? "cartridge1Wells" : "cartridge2Wells";
    isFilled = checkIsFilled(formikData, cartWells);
  }
  //check for others only if isFilled is not true
  if (!isFilled) {
    isFilled = checkIsFilled(formikData);
  }

  //disable other tabs accordingly
  if (isFilled) {
    for (const key in aspireOrDispense) {
      if (key !== currentTabName) {
        aspireOrDispense[key] = true;
      }
    }

    //also if cartridge1 is selected in Aspire, cart2 will be discarded in dispense
    if (currentTabName === "cartridge1" || currentTabName === "cartridge2") {
      const otherTab = !isAspire ? "aspire" : "dispense";
      const tabToDisableInNextPage =
        currentTabName === "cartridge1" ? "cartridge2" : "cartridge1";
      disabledState[otherTab][tabToDisableInNextPage] = true;
    }
  }
  // enable all tabs
  else {
    for (const key in disabledState.aspire) {
      disabledState.aspire[key] = false;
    }
    for (const key in disabledState.dispense) {
      disabledState.dispense[key] = false;
    }
  }

  return disabledState;
};

// sets formik field
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
};
