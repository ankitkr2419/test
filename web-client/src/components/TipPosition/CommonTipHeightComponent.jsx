import React from "react";

import { FormGroup, Label, Input, FormError } from "core-components";
import { Icon, Text } from "shared-components";
import { typeName } from "./functions";

export const CommonTipHeightComponent = (props) => {
  const { formik, activeTab } = props;

  const handleHeightValueChange = (e) => {
    let disableOtherTabs = true;
    let otherFieldIsNotSelected = false;

    const currentTab = typeName[activeTab];
    const tipHeightValue = parseFloat(e.target.value);

    if (currentTab !== "deck") {
      otherFieldIsNotSelected = formik.values[currentTab].wellsArray
        .map((wellObj) => wellObj.isSelected)
        .every((value) => value === false);
    } else {
      otherFieldIsNotSelected = formik.values[currentTab].deckPosition === null;
    }

    // if tipHeight is empty than enable all tabs
    if (!tipHeightValue && otherFieldIsNotSelected) {
      disableOtherTabs = false;
    }
    for (const key in formik.values) {
      if (key !== currentTab) {
        formik.setFieldValue(`${key}.isDisabled`, disableOtherTabs);
      }
    }
    //set new height
    formik.setFieldValue(`${typeName[activeTab]}.tipHeight`, tipHeightValue);
  };

  return (
    <div className="pt-4">
      <FormGroup className="d-flex align-items-center mb-3">
        <Label for="tip-selection" className="px-0 mr-1">
          Tip Height
        </Label>
        <div className="tip-height-input-box ml-4">
          <div className="d-flex flex-column position-relative">
            <Input
              type="number"
              name="tip-height"
              id="tip-height"
              placeholder="Type here"
              className="tip-height-input"
              value={formik.values[`${typeName[activeTab]}`].tipHeight}
              onChange={(e) => handleHeightValueChange(e)}
            />
            {/* <Icon name="height" size={16} className="height-icon-btn" /> */}
            <Text Tag="span" className="font-weight-bold height-icon-btn">
              mm
            </Text>
          </div>
          <FormError>Incorrect Tip Height</FormError>
        </div>
      </FormGroup>
    </div>
  );
};
