import React from "react";

import { FormGroup, Label, Input, FormError } from "core-components";
import { Text, Icon } from "shared-components";
import { CommmonFields } from "./Style";

const DispenseCommonField = (props) => {
  return (
    <CommmonFields>
      <FormGroup className="d-flex align-items-center mb-2">
        <Label for="dispense-height" className="px-0 label-name">
          Dispense Height
        </Label>
        <div className="d-flex flex-column input-field position-relative">
          <Input
            type="text"
            name="dispense-height"
            id="dispense-height"
            placeholder="Type here"
            className="dispense-input-field"
          />
          <Text Tag="span" className="height-icon-box">
            <Icon name="height" size={20} />
          </Text>
          <FormError>Incorrect Dispense Height</FormError>
        </div>
      </FormGroup>

      <FormGroup className="d-flex align-items-center mb-2">
        <Label for="mixing-volume" className="px-0 label-name">
          Mixing Volume
        </Label>
        <div className="d-flex flex-column input-field">
          <Input
            type="text"
            name="mixing-volume"
            id="mixing-volume"
            placeholder="Type here"
          />
          <FormError>Incorrect Mixing Volume</FormError>
        </div>

        <Text Tag="span" className="d-flex align-items-center disabled ml-4">
          <Label for="no-of-cycles" className="px-0 label-name">
            No. Of Cycles
          </Label>
          <Text Tag="span" className="d-flex flex-column cycle-input">
            <Input
              type="text"
              name="no-of-cycles"
              id="no-of-cycles"
              placeholder=""
            />
            <FormError>Incorrect No. Of Cycles</FormError>
          </Text>
        </Text>
      </FormGroup>

      {/* <FormGroup className="d-flex align-items-center mb-2">
        <Label for="dispense-volume" className="px-0 label-name">
          Dispense Volume
        </Label>
        <div className="d-flex flex-column input-field">
          <Input
            type="text"
            name="dispense-volume"
            id="dispense-volume"
            placeholder="Type here"
          />
          <FormError>Incorrect Dispense Volume</FormError>
        </div>
      </FormGroup> */}

      {/* <FormGroup className="d-flex align-items-center mb-2">
        <Label for="dispense-blow" className="px-0 label-name">
          Dispense Blow
        </Label>
        <div className="d-flex flex-column input-field">
          <Input
            type="text"
            name="dispense-blow"
            id="dispense-blow"
            placeholder="Type here"
          />
          <FormError>Incorrect Dispense Blow</FormError>
        </div>
      </FormGroup> */}
    </CommmonFields>
  );
};

DispenseCommonField.propTypes = {};

export default React.memo(DispenseCommonField);
