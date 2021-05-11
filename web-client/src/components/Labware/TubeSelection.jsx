import React from "react";

import { FormGroup, Label, FormError, Select } from "core-components";

import styled from "styled-components";

const CommmonTubeFields = styled.div`
  .label-name {
    width: 9.125rem;
  }
  .input-field {
    width: 14.125rem;
    height: 2.25rem;
    .height-icon-box {
      position: absolute;
      top: 3px;
      right: 0.75rem;
    }
  }
`;

const TubeSelection = (props) => {
  const { handleOptionChange, options } = props;
  return (
    <>
      <CommmonTubeFields>
        <FormGroup>
          <Label className="mb-3 font-weight-bold px-0">Select Tube</Label>
        </FormGroup>
        <FormGroup className="d-flex align-items-center mb-2">
          <Label className="px-0 label-name">Tube Type</Label>
          <div className="d-flex flex-column input-field position-relative">
            <Select
              placeholder="Select Option"
              className=""
              size="sm"
              options={options}
              onChange={handleOptionChange}
            />
            <FormError>Incorrect Tube Type</FormError>
          </div>
        </FormGroup>
      </CommmonTubeFields>
    </>
  );
};

TubeSelection.propTypes = {};

export default React.memo(TubeSelection);
