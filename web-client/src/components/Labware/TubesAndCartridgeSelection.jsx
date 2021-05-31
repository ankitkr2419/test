import React from "react";
import { getOptions } from "./functions";
import labwareDeckPosition1 from "assets/images/labware-plate-deck-position-1.png";
import labwareDeckPosition2 from "assets/images/labware-plate-deck-position-2.png";
import labwareDeckPosition3 from "assets/images/labware-plate-deck-position-3.png";
import labwareDeckPosition4 from "assets/images/labware-plate-deck-position-4.png";
import labwareCartridePosition1 from "assets/images/labware-plate-cartridge-1.png";
import labwareCartridePosition2 from "assets/images/labware-plate-cartridge-2.png";
import LabelAndDropDown from "./LabelAndDropDown";

const deckImages = [
  labwareDeckPosition1,
  labwareDeckPosition2,
  labwareDeckPosition3,
  labwareDeckPosition4,
];
const cartridgeImages = [labwareCartridePosition1, labwareCartridePosition2];

const TubeAndCartridgeSelection = (props) => {
  const { formik, position, allOptions, allowedPositions, isDeck } = props;

  const recipeData = formik.values;
  const images = isDeck ? deckImages : cartridgeImages;
  const key = isDeck ? "deckPosition" : "cartridge";

  const type = isDeck ? "tubeType" : "cartridgeType";

  // gets options array is desired format i.e. {value: "abc", label: "xyz"}
  const options = getOptions(allOptions, allowedPositions);
  const selectedOptionID =
    recipeData[`${key}${position}`].processDetails[type].id;
  const index =
    options && options.map((item) => item.value).indexOf(selectedOptionID);

  //sets values to formik
  const handleOptionChange = (event) => {
    formik.setFieldValue(
      `${key}${position}.processDetails.${type}.id`,
      event.value
    );
    formik.setFieldValue(
      `${key}${position}.processDetails.${type}.label`,
      event.label
    );
  };

  return (
    options && (
      <LabelAndDropDown
        isDeck={isDeck}
        handleOptionChange={(event) => handleOptionChange(event)}
        value={options[index]}
        options={options}
        position={position}
        images={images}
        typeValue={selectedOptionID}
      />
    )
  );
};

export default React.memo(TubeAndCartridgeSelection);
