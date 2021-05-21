export const getRequestBody = (recipeName, formikState) => {
  const requestBody = {
    name: recipeName,
    description: "",
    pos_1: formikState.tips.processDetails.tipPosition1.id,
    pos_2: formikState.tips.processDetails.tipPosition2.id,
    pos_3: formikState.tips.processDetails.tipPosition3.id,
    pos_4: formikState.tipPiercing.processDetails.position1.id,
    pos_5: formikState.tipPiercing.processDetails.position2.id,
    pos_6: formikState.deckPosition1.processDetails.tubeType.id,
    pos_7: formikState.deckPosition2.processDetails.tubeType.id,
    pos_cartridge_1: formikState.cartridge1.processDetails.cartridgeType.id,
    pos_9: formikState.deckPosition3.processDetails.tubeType.id,
    pos_cartridge_2: formikState.cartridge2.processDetails.cartridgeType.id,
    pos_11: formikState.deckPosition4.processDetails.tubeType.id,
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
    return optionsObj.filter((item) => item);
  } else {
    return null;
  }
};
