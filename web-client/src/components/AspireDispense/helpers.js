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

  // initially source and destination will deckPosition
  let sourcePosition = aspire.deckPosition;
  let destinationPosition = dispense.deckPosition;

  let cartridgeType = 1;

  // if aspire is well then source will be position of wells.
  if (
    aspire.selectedCategory === catergories.CATEGORY_1 ||
    aspire.selectedCategory === catergories.CATEGORY_2
  ) {
    cartridgeType = aspire.selectedCategory;
    sourcePosition = getPosition(aspireWells);
  }

  // if dispense is well then destination will be position of wells.
  if (
    dispense.selectedCategory === catergories.CATEGORY_1 ||
    dispense.selectedCategory === catergories.CATEGORY_2
  ) {
    cartridgeType = dispense.selectedCategory;
    destinationPosition = getPosition(dispenseWells);
  }

  return {
    category: `${aspireSelectedTabName}_to_${dispenseSelectedTabName}`,
    cartridge_type: `cartridge_${cartridgeType}`,
    source_position: parseInt(sourcePosition),
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
    destination_position: parseInt(destinationPosition),
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

const getCategoryId = (type, cartridgeType = null) => {
  const { CATEGORY_1, CATEGORY_2, DECK, SHAKER } = catergories;

  // type will change according to response from API
  switch (type) {
    case "well":
      return cartridgeType === "cartridge_1" ? CATEGORY_1 : CATEGORY_2;
    case "shaker":
      return SHAKER;
    case "deck":
      return DECK;
  }
};

// this function generates the initial formik state according to the
// operation being performeed i.e. if NEW process is being created than
// empty data is loaded in formikState else for EDIT old values are loaded.
export const getFormikInitialState = (editReducer = null) => {
  let type = "cartridge_1", // default
    category,
    category1,
    category2,
    aspireSelectedCategory = catergories.CATEGORY_1, // default
    dispenseSelectedCategory = catergories.CATEGORY_1; // default

  // for source
  let sourcePosForCartridge = null;
  let sourcePosForDeck = "";

  // for destination
  let destPosForCartridge = null;
  let destPosForDeck = "";

  if (editReducer?.process_id) {
    type = editReducer.cartridge_type;

    // editReducer.category is a string in format like : well_to_well, well_to_shaker, etc..
    category = editReducer.category.split("_"); // [ 'well', 'to', 'deck' ] -> in case of well_to_deck
    category1 = category[0]; // 'well' -> in case of well_to_well
    category2 = category[2]; // 'deck' -> in case of well_to_well

    aspireSelectedCategory = getCategoryId(category1, type);
    dispenseSelectedCategory = getCategoryId(category2, type);

    // source position
    if (category1 === "well") {
      sourcePosForCartridge = editReducer.source_position;
    } else if (category1 === "deck") {
      sourcePosForDeck = editReducer.source_position;
    }

    // destination position
    if (category2 === "well") {
      destPosForCartridge = editReducer.destination_position;
    } else if (category2 === "deck") {
      destPosForDeck = editReducer.destination_position;
    }
  }

  return {
    aspire: {
      cartridge1Wells: getArray(
        NUMBER_OF_WELLS,
        ASPIRE_WELLS,
        aspireSelectedCategory === catergories.CATEGORY_1,
        sourcePosForCartridge
      ),
      cartridge2Wells: getArray(
        NUMBER_OF_WELLS,
        ASPIRE_WELLS,
        aspireSelectedCategory === catergories.CATEGORY_2,
        sourcePosForCartridge
      ),
      deckPosition: sourcePosForDeck,
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
        destPosForCartridge
      ),
      cartridge2Wells: getArray(
        NUMBER_OF_WELLS,
        DISPENSE_WELLS,
        dispenseSelectedCategory === catergories.CATEGORY_2,
        destPosForCartridge
      ),
      deckPosition: destPosForDeck,
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
const enableAllTabs = (aspireOrDispense = null, except = null) => {
  Object.values(tabNames).forEach((key) => {
    // enable tabs for aspire AND dispense
    if (!aspireOrDispense) {
      disabledTabInitTab.aspire[key] = false;
      disabledTabInitTab.dispense[key] = false;
    }
    // enable tabs only for the param sent
    else {
      if (except === null) {
        // enable all
        disabledTabInitTab[`${aspireOrDispense}`][key] = false;
      } else {
        // enable all except one
        disabledTabInitTab[`${aspireOrDispense}`][key] = key === except;
      }
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

    // opposite of currentTab, if currentTab === aspire, then otherTab === dispense and vice versa
    const otherTab = !isAspire ? "aspire" : "dispense";

    //also if cartridge1 is selected in Aspire, cart2 will be disabled in dispense
    if (currentTabName === "cartridge1" || currentTabName === "cartridge2") {
      const tabToDisableInNextPage =
        currentTabName === "cartridge1" ? "cartridge2" : "cartridge1";
      disabledState[otherTab][tabToDisableInNextPage] = true;
    }

    // if shaker tab is selected in aspire, then shaker tab will get disabled in dispense
    if (currentTabName === "shaker") {
      disabledState[otherTab].shaker = true;
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
      //enable all tabs in aspire and dispense
      enableAllTabs();
    } else {
      // in aspire tab
      if (isAspire) {
        /** if dispense selected category is cartridge1
         * then we enable cartridge2 and others in aspire AND vice-versa*/

        // either cartridge1 or cartridge2
        if (dispenseSelectedCategory === 1 || dispenseSelectedCategory === 2) {
          const exceptValue =
            dispenseSelectedCategory === 1 ? tabNames[2] : tabNames[1];
          enableAllTabs("aspire", exceptValue);
        } else if (dispenseSelectedCategory === 3) {
          /** if shaker is selected in dispense, then in aspire
           * we enable all tabs except shaker */
          // enable all aspire tabs except shaker
          enableAllTabs("aspire", "shaker");
        } else {
          // enable all aspire tabs
          enableAllTabs("aspire");
        }
      }
      // in dispense tab
      else {
        /** if aspire selected category is cartridge1
         * then we enable cartridge2 and others in dispense AND vice-versa*/

        // either cartridge1 or cartridge2
        if (aspireSelectedCategory === 1 || aspireSelectedCategory === 2) {
          const exceptValue =
            aspireSelectedCategory === 1 ? tabNames[2] : tabNames[1];
          enableAllTabs("dispense", exceptValue);
        } else if (aspireSelectedCategory === 3) {
          /** if shaker is selected in aspire, then in dispense
           * we enable all tabs except shaker */
          // enable all dispense tabs except shaker
          enableAllTabs("dispense", "shaker");
        } else {
          // enable all dispense tabs except shaker
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
