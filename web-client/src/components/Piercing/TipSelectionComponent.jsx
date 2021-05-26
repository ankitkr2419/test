import React from "react";
import PropTypes from "prop-types";

import { Col, FormGroup, Label, Select } from "core-components";

export const TipSelectionComponent = (props) => {
  const {heading} = props;
  return (
    <div className="process-box mx-auto">
      <div className="mb-3 border-bottom-line">
        <FormGroup row className="align-items-center">
          <Label for="tip-selection" md={2}>
            {heading}
          </Label>
          <Col md={4}>
            <Select
              placeholder="Select Tip"
              className=""
              size="sm"
            //   options={taskOptions}
            //   value={task}
            //   onChange={handleTaskChange}
            />
          </Col>
        </FormGroup>
      </div>
    </div>
  );
};

TipSelectionComponent.propTypes = {
    heading: PropTypes.string
};

TipSelectionComponent.defaultProps = {
  heading: "Select Tips"
};
