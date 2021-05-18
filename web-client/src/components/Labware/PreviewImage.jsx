import React from "react";

import { ImageIcon } from "shared-components";

import { ProcessSetting } from "./Styles";
import labwarePlate from "assets/images/labware-plate.png";

const PreviewImage = (props) => {
  const { formik } = props;
  return (
    <div className="img-box">
      <ProcessSetting>
        <div className="tips-info">
          <ul className="list-unstyled tip-position active">
            {formik.values.tips.processDetails.tipPosition1.id && (
              <li className="highlighted tip-position-1"></li>
            )}
            {formik.values.tips.processDetails.tipPosition2.id && (
              <li className="highlighted tip-position-2"></li>
            )}
            {formik.values.tips.processDetails.tipPosition3.id && (
              <li className="highlighted tip-position-3"></li>
            )}
          </ul>
        </div>

        <div className="piercing-info">
          <ul className="list-unstyled piercing-position active">
            {formik.values.tipPiercing.processDetails.position1.id && (
              <li className="highlighted piercing-position-1"></li>
            )}
            {formik.values.tipPiercing.processDetails.position2.id && (
              <li className="highlighted piercing-position-2"></li>
            )}
          </ul>
        </div>

        <div className="deck-position-info">
          <ul className="list-unstyled deck-position active">
            {formik.values.deckPosition1.processDetails.tubeType.id && (
              <li className="highlighted deck-position-1 active" />
            )}
            {formik.values.deckPosition2.processDetails.tubeType.id && (
              <li className="highlighted deck-position-2 active" />
            )}
            {formik.values.deckPosition3.processDetails.tubeType.id && (
              <li className="highlighted deck-position-3 active" />
            )}
            {formik.values.deckPosition4.processDetails.tubeType.id && (
              <li className="highlighted deck-position-4 active" />
            )}
          </ul>
        </div>

        <div className="cartridge-position-info">
          <ul className="list-unstyled cartridge-position active">
            {formik.values.cartridge1.processDetails.cartridgeType.id && (
              <li className="highlighted cartridge-position-1 active" />
            )}
            {formik.values.cartridge2.processDetails.cartridgeType.id && (
              <li className="highlighted cartridge-position-2 active" />
            )}
          </ul>
        </div>

        <ImageIcon src={labwarePlate} alt="Labware Plate" className="" />
      </ProcessSetting>
    </div>
  );
};

export default React.memo(PreviewImage);
