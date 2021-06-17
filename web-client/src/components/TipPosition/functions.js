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
  tip_position_1: "Tip Position 1",
  tip_position_2: "Tip Position 2",
  tip_position_3: "Tip Position 3",
  piercing_tip_1: "Piercing Tip 1",
  piercing_tip_2: "Piercing Tip 2",
  sample_tube: "Sample Tube",
  shaker_tube: "Shaker Tube",
  cartridge_1: "Cartridge 1",
  extraction_tube: "Extraction Tube",
  cartridge_2: "Cartridge 2",
  pcr_tube: "PCR Tube",
};

export const typeName = {
  1: "cartridge1",
  2: "deck",
  3: "cartridge2",
};

export const tabApiNames = { cartridge_1: 1, deck: 2, cartridge_2: 3 };

export const getPosition = (wells) => {
  if (wells) {
    const selectedWell = wells.find((wellObj) => wellObj.isSelected);
    return selectedWell ? selectedWell.id : 0;
  }
  return 0;
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

export const getFormikInitialState = (editReducer = null) => {
  const cartridge1Wells =
    editReducer?.process_id && tabApiNames[editReducer.type] === 1
      ? getArray(13, 0, editReducer.position)
      : getArray(13, 0);

  const cartridge1TipHeight =
    editReducer?.process_id && tabApiNames[editReducer.type] === 1
      ? editReducer.height
      : null;

  const deckPosition =
    editReducer?.process_id && tabApiNames[editReducer.type] === 2
      ? apiNameToLabel[editReducer.deck_position]
      : null;

  const deckTipHeight =
    editReducer?.process_id && tabApiNames[editReducer.type] === 2
      ? editReducer.height
      : null;

  const cartridge2Wells =
    editReducer?.process_id && tabApiNames[editReducer.type] === 3
      ? getArray(13, 1, editReducer.position)
      : getArray(13, 1);

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
