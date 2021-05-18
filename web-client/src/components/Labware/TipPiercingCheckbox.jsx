import React from "react";

import { ImageIcon } from "shared-components";
import { FormGroup, Label, CheckBox } from "core-components";
import { ProcessSetting } from "./Styles";

import labwarePiercing from "assets/images/labware-plate-piercing.png";

const TipPiercingCheckbox = (props) => {
  const { formik } = props;

  const position1 = formik.values.tipPiercing.processDetails.position1.id;
  const position2 = formik.values.tipPiercing.processDetails.position2.id;

  const tipPositions = formik.values.tipPiercing.processDetails;

  const getTipPiercingCheckbox = () => {
    return Object.keys(tipPositions).map((position, index) => {
      let isChecked = tipPositions[position].id;
      return (
        <CheckBox
          id={`${position}`}
          name={`${position}`}
          label={`Position ${index + 1}`}
          className={index > 0 ? "ml-4" : ""}
          checked={isChecked ? true : false}
          onChange={(e) => {
            formik.setFieldValue(
              `tipPiercing.processDetails.position${index + 1}.id`,
              e.target.checked ? 3 : null //backend dependency!
            );
          }}
        />
      );
    });
  };

  return (
    <>
      {/* heading label */}
      <div className="mb-3">
        <FormGroup row>
          <Label
            for="select-tip-piercing"
            md={12}
            className="mb-3 font-weight-bold"
          >
            Select Tip Piercing
          </Label>
        </FormGroup>
      </div>

      <div className="d-flex align-items-center">
        {/* checkboxes */}
        {getTipPiercingCheckbox()}

        {/* Side Images */}
        <ProcessSetting>
          <div className="piercing-info">
            <ul className="list-unstyled piercing-position active">
              {position1 && (
                <li className="highlighted piercing-position-1"></li>
              )}
              {position2 && (
                <li className="highlighted piercing-position-2 active"></li>
              )}
            </ul>
            <ImageIcon
              src={labwarePiercing}
              alt="Piercing Process"
              className=""
            />
          </div>
        </ProcessSetting>
      </div>
    </>
  );
};

export default React.memo(TipPiercingCheckbox);
