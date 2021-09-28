import React from "react";

import { FormGroup, Label, Input, FormError } from "core-components";
import { Text, Icon } from "shared-components";
import { CommmonFields } from "./Style";
import { setFormikField } from "./helpers";

const AspireCommonField = (props) => {
  const { formik, currentTab } = props;
  const mixingVolume = formik.values.aspire.mixingVolume;
  const nCyclesDisabled = mixingVolume === null || mixingVolume === "";

  // if value of mixing volume is cleared then value of nCycles must also be cleared
  const handleBlur = (value) => {
    if (value === "") {
      setFormikField(formik, true, currentTab, "nCycles", "");
    }
  };

  return (
    <CommmonFields>
      <FormGroup className="d-flex align-items-center mb-2">
        <Label for="aspire-height" className="px-0 label-name">
          Aspire Height
        </Label>
        <div className="d-flex flex-column input-field position-relative">
          <Input
            type="text"
            name="aspireHeight"
            id="aspire-height"
            placeholder="Type here"
            className="aspire-input-field"
            value={formik.values.aspire.aspireHeight}
            onChange={(e) =>
              setFormikField(
                formik,
                true,
                currentTab,
                e.target.name,
                e.target.value
              )
            }
          />
          <Text Tag="span" className="font-weight-bold height-icon-box">
            {/* <Icon name="height" alt={"mm"} size={20} /> */}
            mm
          </Text>
          <FormError>Incorrect Aspire Height</FormError>
        </div>
      </FormGroup>

      <FormGroup className="d-flex align-items-center mb-2">
        <Label for="mixing-volume" className="px-0 label-name">
          Mixing Volume
        </Label>
        <div className="d-flex flex-column input-field position-relative">
          <Input
            className="aspire-input-field"
            type="text"
            name="mixingVolume"
            id="mixing-volume"
            placeholder="Type here"
            value={formik.values.aspire.mixingVolume}
            onChange={(e) =>
              setFormikField(
                formik,
                true,
                currentTab,
                e.target.name,
                e.target.value
              )
            }
            onBlur={(e) => handleBlur(e.target.value)}
          />
          <Text Tag="span" className="font-weight-bold height-icon-box">
            ml
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
              placeholder=""
              value={formik.values.aspire.nCycles}
              onChange={(e) =>
                setFormikField(
                  formik,
                  true,
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

      <FormGroup className="d-flex align-items-center mb-2">
        <Label for="aspire-volume" className="px-0 label-name">
          Aspire Volume
        </Label>
        <div className="d-flex flex-column input-field position-relative">
          <Input
            className="aspire-input-field"
            type="text"
            name="aspireVolume"
            id="aspire-volume"
            placeholder="Type here"
            value={formik.values.aspire.aspireVolume}
            onChange={(e) =>
              setFormikField(
                formik,
                true,
                currentTab,
                e.target.name,
                e.target.value
              )
            }
          />
          <Text Tag="span" className="font-weight-bold height-icon-box">
            ml
          </Text>
          <FormError>Incorrect Aspire Volume</FormError>
        </div>
      </FormGroup>

      <FormGroup className="d-flex align-items-center mb-2">
        <Label for="air-volume" className="px-0 label-name">
          Air Volume
        </Label>
        <div className="d-flex flex-column input-field position-relative">
          <Input
            className="aspire-input-field"
            type="text"
            name="airVolume"
            id="air-volume"
            placeholder="Type here"
            value={formik.values.aspire.airVolume}
            onChange={(e) =>
              setFormikField(
                formik,
                true,
                currentTab,
                e.target.name,
                e.target.value
              )
            }
          />
          <Text Tag="span" className="font-weight-bold height-icon-box">
            ml
          </Text>
          <FormError>Incorrect Air Volume</FormError>
        </div>
      </FormGroup>
    </CommmonFields>
  );
};

AspireCommonField.propTypes = {};

export default React.memo(AspireCommonField);
