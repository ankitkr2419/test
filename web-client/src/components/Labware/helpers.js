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

export const getOptions = (options, position, type) => {
  let optionsObj;

  if (options) {
    optionsObj = options.map((optionObject) => {
      if (type === "cartridge_1" || type === "cartridge_2") {
        const { id, description } = optionObject;
        if (type === optionObject.type) {
          return {
            value: id,
            label: description,
          };
        }
      } else {
        const { id, name, description, allowed_positions } = optionObject;
        if (allowed_positions?.includes(position)) {
          return {
            value: id,
            label: name ? name : description,
          };
        }
      }
    });
    return optionsObj.filter((item) => item);
  } else {
    return null;
  }
};
