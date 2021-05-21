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

export const sideBarState = {
  aspire: {
    cartridge1: true,
    cartridge2: true,
    shaker: true,
    deckPosition: true,
  },
  dispense: {
    cartridge1: true,
    cartridge2: true,
    shaker: true,
    deckPosition: true,
  },
};

export const toggler = (isAspire, activeTab) => {
  const activeTabObj = isAspire ? sideBarState.aspire : sideBarState.dispense;
  if (isAspire) {
    // console.log(activeTabObj[Object.keys(activeTabObj)[activeTab - 1]]);
  } else {
    // console.log(activeTabObj[Object.keys(activeTabObj)[activeTab - 1]]);
  }
};
