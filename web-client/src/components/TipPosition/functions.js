import {
  CARTRIDGE_1_WELLS,
  CARTRIDGE_2_WELLS,
  NUMBER_OF_WELLS,
} from "appConstants";

export const deckPositionNames = [
  "Tip Position 1",
  "Tip Position 2",
  "Tip Position 3",
  "Piercing Tip 1",
  "Piercing Tip 2",
  "Sample Tube",
  "Shaker Tube",
  "Cartridge 1",
  "Extraction Tube",
  "Cartridge 2",
  "PCR Tube",
];

export const apiNameToLabel = {
  1: "Tip Position 1",
  2: "Tip Position 2",
  3: "Tip Position 3",
  4: "Piercing Tip 1",
  5: "Piercing Tip 2",
  6: "Sample Tube",
  7: "Shaker Tube",
  8: "Cartridge 1",
  9: "Extraction Tube",
  10: "Cartridge 2",
  11: "PCR Tube",
};

export const typeName = {
  1: "cartridge1",
  2: "deck",
  3: "cartridge2",
};

export const typeNameAPI = {
  1: "cartridge_1",
  2: "deck",
  3: "cartridge_2",
};

export const tabApiNames = { cartridge_1: 1, deck: 2, cartridge_2: 3 };

export const getPosition = (wells) => {
  const selectedWell = wells.find((wellObj) => wellObj.isSelected);
  return selectedWell ? selectedWell.id : 0;
};

//this function updates and renders wellsArray after selecting or de-selecting wells
export const updateWellsArray = (wellsObjArray, currentWellObj) => {
  const updatedWellObjArray = wellsObjArray.map((wellObj) => {
    if (wellObj.id === currentWellObj.id) {
      return {
        ...wellObj,
        isSelected: !wellObj.isSelected, //toggle
      };
    }
    return wellObj;
  });

  return updatedWellObjArray;
};

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
      height: null,
      isSelected: isSelected,
      isDisabled: false,
    });
  }
  return array;
};

export const getFormikInitialState = (
  cartridge1Details,
  cartridge2Details,
  editReducer = null
) => {
  const cartridge1Wells =
    editReducer?.process_id && tabApiNames[editReducer.type] === 1
      ? getArray(
          cartridge1Details?.wells_count || NUMBER_OF_WELLS,
          CARTRIDGE_1_WELLS,
          editReducer.position
        )
      : getArray(
          cartridge1Details?.wells_count || NUMBER_OF_WELLS,
          CARTRIDGE_1_WELLS
        );

  const cartridge1TipHeight =
    editReducer?.process_id && tabApiNames[editReducer.type] === 1
      ? editReducer.height
      : null;

  const deckPosition =
    editReducer?.process_id && tabApiNames[editReducer.type] === 2
      ? parseInt(editReducer.position)
      : null;

  const deckTipHeight =
    editReducer?.process_id && tabApiNames[editReducer.type] === 2
      ? editReducer.height
      : null;

  const cartridge2Wells =
    editReducer?.process_id && tabApiNames[editReducer.type] === 3
      ? getArray(
          cartridge2Details?.wells_count || NUMBER_OF_WELLS,
          CARTRIDGE_2_WELLS,
          editReducer.position
        )
      : getArray(
          cartridge2Details?.wells_count || NUMBER_OF_WELLS,
          CARTRIDGE_2_WELLS
        );

  const cartridge2TipHeight =
    editReducer?.process_id && tabApiNames[editReducer.type] === 3
      ? editReducer.height
      : null;

  return {
    cartridge1: {
      wellsArray: cartridge1Wells,
      tipHeight: cartridge1TipHeight,
      isDisabled: false,
    },
    cartridge2: {
      wellsArray: cartridge2Wells,
      tipHeight: cartridge2TipHeight,
      isDisabled: false,
    },
    deck: {
      deckPosition: deckPosition,
      tipHeight: deckTipHeight,
      isDisabled: false,
    },
  };
};
