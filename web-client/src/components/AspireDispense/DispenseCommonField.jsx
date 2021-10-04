import React from "react";

import { FormGroup, Label, Input, FormError } from "core-components";
import { Text, Icon } from "shared-components";
import { CommmonFields } from "./Style";
import { setFormikField } from "./helpers";

const DispenseCommonField = (props) => {
  const { formik, currentTab } = props;
  const mixingVolume = formik.values.dispense.mixingVolume;
  const nCyclesDisabled = mixingVolume === null || mixingVolume === "";

  // if value of mixing volume is cleared then value of nCycles must also be cleared
  const handleBlur = (value) => {
    if (value === "") {
      setFormikField(formik, false, currentTab, "nCycles", "");
    }
  };

  return (
    <CommmonFields>
      <FormGroup className="d-flex align-items-center mb-2">
        <Label for="dispense-height" className="px-0 label-name">
          Dispense Height
        </Label>
        <div className="d-flex flex-column input-field position-relative">
          <Input
            type="text"
            name="dispenseHeight"
            id="dispense-height"
            placeholder="Type here"
            className="dispense-input-field"
            value={formik.values.dispense.dispenseHeight}
            onChange={(e) =>
              setFormikField(
                formik,
                false,
                currentTab,
                e.target.name,
                e.target.value
              )
            }
          />
          <Text Tag="span" className="font-weight-bold height-icon-box">
            {/* <Icon name="height" size={20} /> */}
            mm
          </Text>
          <FormError>Incorrect Dispense Height</FormError>
        </div>
      </FormGroup>

      <FormGroup className="d-flex align-items-center mb-2">
        <Label for="mixing-volume" className="px-0 label-name">
          Mixing Volume
        </Label>
        <div className="d-flex flex-column input-field position-relative">
          <Input
            type="text"
            name="mixingVolume"
            id="mixing-volume"
            placeholder="Type here"
            className="dispense-input-field"
            value={formik.values.dispense.mixingVolume}
            onChange={(e) =>
              setFormikField(
                formik,
                false,
                currentTab,
                e.target.name,
                e.target.value
              )
            }
            onBlur={(e) => handleBlur(e.target.value)}
          />
          <Text Tag="span" className="font-weight-bold height-icon-box">
            μl
          </Text>
          <FormError>Incorrect Mixing Volume</FormError>
        </div>

        <Text
          Tag="span"
          className={`d-flex align-items-center ${
            nCyclesDisabled ? "disabled" : ""
          } ml-4`}
        >
          <Label for="no-of-cycles" className="px-0 label-name">
            No. Of Cycles
          </Label>
          <Text Tag="span" className="d-flex flex-column cycle-input">
            <Input
              type="text"
              name="nCycles"
              id="no-of-cycles"
              value={formik.values.dispense.nCycles}
              placeholder=""
              onChange={(e) =>
                setFormikField(
                  formik,
                  false,
                  currentTab,
                  e.target.name,
                  e.target.value
                )
              }
            />
            <FormError>Incorrect No. Of Cycles</FormError>
          </Text>
        </Text>
      </FormGroup>
    </CommmonFields>
  );
};

DispenseCommonField.propTypes = {};

export default React.memo(DispenseCommonField);
