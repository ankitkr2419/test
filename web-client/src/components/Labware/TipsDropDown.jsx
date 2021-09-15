import React from "react";

import labwareTips from "assets/images/labware-plate-tips.png";
import { ImageIcon } from "shared-components";
import { FormGroup, Label, FormError, Select } from "core-components";
import { ProcessSetting } from "./Styles";
import { getOptions } from "./helpers";

const TipsDropdown = (props) => {
  const { formik, tipsOptions } = props;

  const tips = formik.values.tips;
  const tipPosition1Value = tips.processDetails.tipPosition1.id;
  const tipPosition2Value = tips.processDetails.tipPosition2.id;
  const tipPosition3Value = tips.processDetails.tipPosition3.id;

  const getDropDowns = () => {
    const tipPositions = formik.values.tips.processDetails;

    return Object.keys(tipPositions).map((tipPosition, i) => {
      let options = getOptions(tipsOptions, i + 1, "tips");
      let id = tipPositions[tipPosition].id;
      //match the id of option with current tipPosition ID
      let index = options && options.map((item) => item.value).indexOf(id);

      return (
        options && (
          <FormGroup key={i} className="d-flex align-items-center mb-4">
            <Label for={`tip-position-${i + 1}`} className="px-0 label-name">
              Tip Position {i + 1}
            </Label>
            <div className="d-flex flex-column input-field position-relative">
              <Select
                isClearable
                placeholder="Select Option"
                className=""
                size="sm"
                value={options[index]}
                options={options}
                onChange={(e) => {
                  formik.setFieldValue(
                    `tips.processDetails.tipPosition${i + 1}.id`,
                    e?.value || null
                  );
                  formik.setFieldValue(
                    `tips.processDetails.tipPosition${i + 1}.label`,
                    e?.label || null
                  );
                }}
              />
              <FormError>Incorrect Tip Position {index + 1}</FormError>
            </div>
          </FormGroup>
        )
      );
    });
  };

  return (
    <>
      <div className="">
        <div className="mb-3">
          <FormGroup row>
            <Label
              for="select-tip-position"
              md={12}
              className="mb-3 font-weight-bold"
            >
              Select Tip Position
            </Label>
          </FormGroup>
        </div>
        <div className="">{getDropDowns()}</div>
      </div>
      <ProcessSetting>
        <div className="tips-info">
          <ul className="list-unstyled tip-position active">
            {tipPosition1Value && (
              <li className="highlighted tip-position-1"></li>
            )}
            {tipPosition2Value && (
              <li className="highlighted tip-position-2 active"></li>
            )}
            {tipPosition3Value && (
              <li className="highlighted tip-position-3 active"></li>
            )}
          </ul>
          <ImageIcon src={labwareTips} alt="Tip Pickup Process" className="" />
        </div>
      </ProcessSetting>
    </>
  );
};

export default React.memo(TipsDropdown);
