export const getWellsInitialArray = (n, type, editReducerData = null) => {
  const wellsArray = [];

  for (let i = 0; i < n; i++) {
    wellsArray.push({
      id: i + 1,
      type: type,
      label: `${i + 1}`,
      footerText: "",
      isSelected: false,
      isDisabled: false,
    });
  }
  return wellsArray;
};

export const updatedWellsArray = (wellsArray, editReducerData = null) => {
  const selectedWells = editReducerData && editReducerData.cartridge_wells;

  wellsArray.map((wellObj, i) => {
    let typeName = wellObj.type === 0 ? "cartridge_1" : "cartridge_2";
    if (
      selectedWells &&
      selectedWells.includes(i + 1) &&
      typeName === editReducerData.type
    ) {
      wellObj.isSelected = true;
    }
    return wellObj;
  });
  return wellsArray;
};
