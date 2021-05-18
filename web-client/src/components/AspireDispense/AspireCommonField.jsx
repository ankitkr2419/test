import React from "react";

import { FormGroup, Label, Input, FormError } from "core-components";
import { Text, Icon } from "shared-components";
import { CommmonFields } from "./Style";

const AspireCommonField = (props) => {
  return (
    <>
      <CommmonFields>
        <FormGroup className="d-flex align-items-center mb-2">
          <Label for="aspire-height" className="px-0 label-name">
            Aspire Height
          </Label>
          <div className="d-flex flex-column input-field position-relative">
            <Input
              type="text"
              name="aspire-height"
              id="aspire-height"
              placeholder="Type here"
              className="aspire-input-field"
            />
            <Text Tag="span" className="height-icon-box">
              <Icon name="height" size={20} />
            </Text>
            <FormError>Incorrect Aspire Height</FormError>
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

        <FormGroup className="d-flex align-items-center mb-2">
          <Label for="aspire-volume" className="px-0 label-name">
            Aspire Volume
          </Label>
          <div className="d-flex flex-column input-field">
            <Input
              type="text"
              name="aspire-volume"
              id="aspire-volume"
              placeholder="Type here"
            />
            <FormError>Incorrect Aspire Volume</FormError>
          </div>
        </FormGroup>

        <FormGroup className="d-flex align-items-center mb-2">
          <Label for="air-volume" className="px-0 label-name">
            Air Volume
          </Label>
          <div className="d-flex flex-column input-field">
            <Input
              type="text"
              name="air-volume"
              id="air-volume"
              placeholder="Type here"
            />
            <FormError>Incorrect Air Volume</FormError>
          </div>
        </FormGroup>
      </CommmonFields>
    </>
  );
};

AspireCommonField.propTypes = {};

export default React.memo(AspireCommonField);
