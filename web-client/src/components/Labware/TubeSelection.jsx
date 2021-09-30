import React from "react";
import { FormGroup, Label, FormError, Select } from "core-components";
import { CommmonTubeFields } from "./CommmonTubeFields";

const TubeSelection = (props) => {
  const { handleOptionChange, options, value } = props;
  return (
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
            value={value}
            options={options}
            onChange={handleOptionChange}
            isSearchable={false}
          />
          <FormError>Incorrect Tube Type</FormError>
        </div>
      </FormGroup>
    </CommmonTubeFields>
  );
};

TubeSelection.propTypes = {};

export default React.memo(TubeSelection);
