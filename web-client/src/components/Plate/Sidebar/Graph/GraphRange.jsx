import React, { useCallback } from "react";
import PropTypes from "prop-types";
import { Text } from "shared-components";
import { FormGroup, Label, Input, Button } from "core-components";

const GraphRange = ({ className }) => {
  return (
    <div className={`graph-range d-flex ${className}`}>
      <Text Tag="h4" size={19} className="flex-10 title mb-0 pr-3">
        Range
      </Text>
      <div className="d-flex align-items-center flex-wrap flex-90">
        <FormGroup className="d-flex align-items-center flex-40 px-2">
          <Label className="flex-20 text-right mb-0 p-1">X Axis</Label>
          <Input
            type="number"
            className="px-2 py-1 ml-2"
            placeholder="Min value"
          />
          <Input
            type="number"
            className="px-2 py-1 ml-2"
            placeholder="Max value"
          />
        </FormGroup>
        <FormGroup className="d-flex align-items-center flex-40 px-2">
          <Label className="flex-20 text-right mb-0 p-1">Y Axis</Label>
          <Input
            type="number"
            className="px-2 py-1 ml-2"
            placeholder="Min value"
          />
          <Input
            type="number"
            className="px-2 py-1 ml-2"
            placeholder="Max value"
          />
        </FormGroup>
        <Button color="primary" size="sm" outline className="mb-3">
          Apply
        </Button>
      </div>
    </div>
  );
};

GraphRange.propTypes = {
  className: PropTypes.string
};

GraphRange.defaultProps = {
  className: ""
};

export default GraphRange;
