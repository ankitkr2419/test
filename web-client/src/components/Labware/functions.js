//this will be changed to null values instead of 0 in future
//after backend gets fixed.
export const getRequestBody = (recipeName, formikState) => {
  const requestBody = {
    name: recipeName,
    description: "",
    pos_1: formikState.tips.processDetails.tipPosition1.id
      ? formikState.tips.processDetails.tipPosition1.id
      : 0,
    pos_2: formikState.tips.processDetails.tipPosition2.id
      ? formikState.tips.processDetails.tipPosition2.id
      : 0,
    pos_3: formikState.tips.processDetails.tipPosition3.id
      ? formikState.tips.processDetails.tipPosition3.id
      : 0,
    pos_4: formikState.tipPiercing.processDetails.position1.id
      ? formikState.tipPiercing.processDetails.position1.id
      : 0,
    pos_5: formikState.tipPiercing.processDetails.position2.id
      ? formikState.tipPiercing.processDetails.position2.id
      : 0,
    pos_6: formikState.deckPosition1.processDetails.tubeType.id
      ? formikState.deckPosition1.processDetails.tubeType.id
      : 0,
    pos_7: formikState.deckPosition2.processDetails.tubeType.id
      ? formikState.deckPosition2.processDetails.tubeType.id
      : 0,
    pos_cartridge_1: formikState.cartridge1.processDetails.cartridgeType.id
      ? formikState.cartridge1.processDetails.cartridgeType.id
      : 0,
    pos_9: formikState.deckPosition3.processDetails.tubeType.id
      ? formikState.deckPosition3.processDetails.tubeType.id
      : 0,
    pos_cartridge_2: formikState.cartridge2.processDetails.cartridgeType.id
      ? formikState.cartridge2.processDetails.cartridgeType.id
      : 0,
    pos_11: formikState.deckPosition4.processDetails.tubeType.id
      ? formikState.deckPosition4.processDetails.tubeType.id
      : 0,
  };

  return requestBody;
};

export const getOptions = (lowerLimit, higherLimit, options) => {
  const optionsObj = [];
  if (options) {
    options.forEach((optionObj) => {
      if (optionObj.id >= lowerLimit && optionObj.id <= higherLimit) {
        optionsObj.push({
          value: optionObj.id,
          label: optionObj.name ? optionObj.name : optionObj.description,
        });
      }
    });
  }
  return optionsObj;
};

export const getOptionsForTubesAndCartridges = (options, position) => {
  let optionsObj;
  if (options) {
    optionsObj = options.map((optionObject) => {
      if (optionObject.id === position) {
        return {
          value: optionObject.id,
          label: optionObject.name
            ? optionObject.name
            : optionObject.description,
        };
      }
    });
  }
  return optionsObj.filter((item) => item);
};
