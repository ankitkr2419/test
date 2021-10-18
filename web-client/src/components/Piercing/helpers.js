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
      height: 0,
    });
  }
  return wellsArray;
};

//this is called when data from edit-reducer is recieved
//i.e. during edit process.
export const getWellsArrayForEdit = (wellsArray, editReducerData = null) => {
  const selectedWells = editReducerData && editReducerData.cartridge_wells;
  const piercingHeights = editReducerData && editReducerData.piercing_heights;

  wellsArray.map((wellObj, i) => {
    let typeName = wellObj.type === 0 ? "cartridge_1" : "cartridge_2";
    if (
      selectedWells &&
      selectedWells.includes(i + 1) &&
      typeName === editReducerData.type
    ) {
      wellObj.isSelected = true;
      wellObj.height = piercingHeights[i];
    }
    return wellObj;
  });
  return wellsArray;
};

//this function updates and renders wellsArray after selecting or de-selecting wells
export const updateWellsArray = (
  wellsObjArray,
  currentWellObj,
  updatedHeight
) => {
  const updatedWellObjArray = wellsObjArray.map((wellObj) => {
    if (wellObj.id === currentWellObj.id) {
      return {
        ...wellObj,
        isSelected: !wellObj.isSelected, //toggle
        height: updatedHeight || wellObj.height, // set height
      };
    }
    return wellObj;
  });

  return updatedWellObjArray;
};
