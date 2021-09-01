import {
  ASPIRE_WELLS,
  CATEGORY_NAME,
  DISPENSE_WELLS,
  NUMBER_OF_WELLS,
} from "appConstants";

// returns the position (index) of selected well
export const getPosition = (wells) => {
  if (wells) {
    const selectedWell = wells.find((wellObj) => wellObj.isSelected);
    return selectedWell ? selectedWell.id : 0;
  }
  return 0;
};

const catergories = {
  CATEGORY_1: "1",
  CATEGORY_2: "2",
  SHAKER: "3",
  DECK: "4",
};

// used to generated request body for API call.
export const getRequestBody = (activeTab, aspire, dispense) => {
  const aspireSelectedTabName = CATEGORY_NAME[aspire.selectedCategory];
  const dispenseSelectedTabName = CATEGORY_NAME[dispense.selectedCategory];

  const aspireWells = aspire[`cartridge${aspire.selectedCategory}Wells`];
  const dispenseWells = dispense[`cartridge${dispense.selectedCategory}Wells`];

  let cartridgeType = 1;
  if (
    aspire.selectedCategory === catergories.CATEGORY_1 ||
    aspire.selectedCategory === catergories.CATEGORY_2
  ) {
    cartridgeType = aspire.selectedCategory;
  }
  if (
    dispense.selectedCategory === catergories.CATEGORY_1 ||
    dispense.selectedCategory === catergories.CATEGORY_2
  ) {
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
export const getArray = (
  length,
  type,
  currentSelectedCategory,
  selectedPosition = null
) => {
  const array = [];

  for (let i = 0; i < length; i++) {
    let isSelected = false;
    if (
      selectedPosition &&
      i + 1 === selectedPosition &&
      currentSelectedCategory === true
    ) {
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

// function that checks and sets formik initial values
// based on update or create
const setFormikValue = (editReducer, name) => {
  if ((editReducer && editReducer[name]) || editReducer[name] === 0) {
    return `${editReducer[name]}`;
  }
  return "";
};

// this function generates the initial formik state according to the
// operation being performeed i.e. if NEW process is being created than
// empty data is loaded in formikState else for EDIT old values are loaded.
export const getFormikInitialState = (editReducer = null) => {
  let type = "category_1", // default
    category,
    category1,
    category2,
    aspireSelectedCategory = catergories.CATEGORY_1,
    dispenseSelectedCategory = catergories.CATEGORY_1;

  const CATEGORY_ID = {
    // type will change according to response from API
    well:
      type === "cartridge_1" ? catergories.CATEGORY_1 : catergories.CATEGORY_2, // default cartridge 1
    shaker: catergories.SHAKER,
    deck: catergories.DECK,
  };

  if (editReducer?.process_id) {
    type = editReducer.cartridge_type;

    // editReducer.category is a string in format like : well_to_well, well_to_shaker, etc..
    category = editReducer.category.split("_"); // [ 'well', 'to', 'deck' ] -> in case of well_to_deck
    category1 = category[0]; // 'well' -> in case of well_to_well
    category2 = category[2]; // 'deck' -> in case of well_to_well

    aspireSelectedCategory = CATEGORY_ID[category1];
    dispenseSelectedCategory = CATEGORY_ID[category2];
  }

  return {
    aspire: {
      cartridge1Wells: getArray(
        NUMBER_OF_WELLS,
        ASPIRE_WELLS,
        aspireSelectedCategory === catergories.CATEGORY_1,
        editReducer.source_position
      ),
      cartridge2Wells: getArray(
        NUMBER_OF_WELLS,
        ASPIRE_WELLS,
        aspireSelectedCategory === catergories.CATEGORY_2,
        editReducer.source_position
      ),
      deckPosition: "",
      aspireHeight: setFormikValue(editReducer, "aspire_height"),
      mixingVolume: setFormikValue(editReducer, "aspire_mixing_volume"),
      aspireVolume: setFormikValue(editReducer, "aspire_volume"),
      airVolume: setFormikValue(editReducer, "aspire_air_volume"),
      nCycles: setFormikValue(editReducer, "aspire_no_of_cycles"),

      selectedCategory: aspireSelectedCategory,
    },
    dispense: {
      cartridge1Wells: getArray(
        NUMBER_OF_WELLS,
        DISPENSE_WELLS,
        dispenseSelectedCategory === catergories.CATEGORY_1,
        editReducer.destination_position
      ),
      cartridge2Wells: getArray(
        NUMBER_OF_WELLS,
        DISPENSE_WELLS,
        dispenseSelectedCategory === catergories.CATEGORY_2,
        editReducer.destination_position
      ),
      deckPosition: "",
      dispenseHeight: setFormikValue(editReducer, "dispense_height"),
      mixingVolume: setFormikValue(editReducer, "dispense_mixing_volume"),
      nCycles: setFormikValue(editReducer, "dispense_no_of_cycles"),
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

  for (const key in formikData) {
    // check for wells
    if (key === "cartridge1Wells" || key === "cartridge2Wells") {
      isFilled = !formikData[`${key}`].every(
        (wellObj) => wellObj.isSelected === false
      );
    }
    // check for form fields other than wells
    else if (
      key !== currentKey &&
      key !== "selectedCategory" &&
      key !== "nCycles"
    ) {
      const value = formikData[key];
      isFilled = value ? true : false;
    }
    if (isFilled) break;
  }
  return isFilled;
};

export const tabNames = {
  1: "cartridge1",
  2: "cartridge2",
  3: "shaker",
  4: "deckPosition",
};

/** This function enables tab of aspire or dispense
 * if no parameter is sent it will enable tabs for both
 */
const enableAllTabs = (aspireOrDispense = null) => {
  Object.values(tabNames).forEach((key) => {
    // enable tabs for aspire AND dispense
    if (!aspireOrDispense) {
      disabledTabInitTab.aspire[key] = false;
      disabledTabInitTab.dispense[key] = false;
    }
    // enable tabs only for the param sent
    else {
      disabledTabInitTab[`${aspireOrDispense}`][key] = false;
    }
  });
};

//this function is used to change the disablility of
//different tabs according to different cases
export const toggler = (formik, isAspire) => {
  const disabledState = disabledTabInitTab;
  const formikData = formik.values[isAspire ? "aspire" : "dispense"];

  const currentTab = formikData.selectedCategory;
  const currentTabName = tabNames[currentTab];
  const aspireOrDispense = isAspire
    ? disabledState.aspire
    : disabledState.dispense;

  let isFilled = false;
  isFilled = checkIsFilled(formikData);

  //disable other tabs accordingly
  if (isFilled) {
    for (const key in aspireOrDispense) {
      const disableTab = key !== currentTabName;
      aspireOrDispense[`${key}`] = disableTab;
    }

    //also if cartridge1 is selected in Aspire, cart2 will be disabled in dispense
    if (currentTabName === "cartridge1" || currentTabName === "cartridge2") {
      const otherTab = !isAspire ? "aspire" : "dispense";
      const tabToDisableInNextPage =
        currentTabName === "cartridge1" ? "cartridge2" : "cartridge1";
      disabledState[otherTab][tabToDisableInNextPage] = true;
    }
  }
  // here we check conditions and enable tabs accordingly
  else {
    const formikAspire = formik.values.aspire;
    const formikDispense = formik.values.dispense;

    const aspireSelectedCategory = parseInt(formikAspire.selectedCategory);
    const dispenseSelectedCategory = parseInt(formikDispense.selectedCategory);
    const aspireFieldsAreFilled = checkIsFilled(formikAspire);
    const dispenseFieldsAreFilled = checkIsFilled(formikDispense);

    // if aspire and dispense BOTH have no fields, then enable all tabs
    if (!aspireFieldsAreFilled && !dispenseFieldsAreFilled) {
      enableAllTabs();
    } else {
      // in aspire tab
      if (isAspire) {
        /** if dispense selected category is cartridge1
         * then we enable cartridge2 and others in aspire AND vice-versa*/
        if (dispenseSelectedCategory <= 2) {
          disabledState.aspire[`cartridge${dispenseSelectedCategory}`] = false;
          disabledState.aspire["shaker"] = false;
          disabledState.aspire["deckPosition"] = false;
        } else {
          enableAllTabs("aspire");
        }
      }
      // in dispense tab
      else {
        /** if aspire selected category is cartridge1
         * then we enable cartridge2 and others in dispense AND vice-versa*/
        if (aspireSelectedCategory <= 2) {
          disabledState.dispense[`cartridge${aspireSelectedCategory}`] = false;
          disabledState.dispense["shaker"] = false;
          disabledState.dispense["deckPosition"] = false;
        } else {
          enableAllTabs("dispense");
        }
      }
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
