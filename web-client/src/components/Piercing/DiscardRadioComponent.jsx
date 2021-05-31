import React from "react";

import { Col, FormGroup, Label, Select, Radio } from "core-components";

export const DiscardRadioComponent = () => {
  return (
    <div className="mb-3 border-bottom-line">
      <FormGroup row className="align-items-center">
        <Label for="discard" md={12}>
          Discard
        </Label>
        <Col md={12}>
          <div className="d-flex mt-3">
            <Radio
              id="pickup-passing"
              name="discard-option"
              label="At pickup passing"
              className="mb-3 mr-4"
            />
            <Radio
              id="discard-box"
              name="discard-option"
              label="At Discard Box"
              className="mb-3 mr-4"
            />
          </div>
        </Col>
      </FormGroup>
    </div>
  );
};
