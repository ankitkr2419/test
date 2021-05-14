import React from "react";

import { Col, FormGroup, Label } from "core-components";

import Well from "components/Plate/Grid/Well";
import CoordinateItem from "components/Plate/Grid/CoordinateItem";
import Coordinate from "components/Plate/Grid/Coordinate";

export const WellComponent = (props) => {
  const { labelArray } = props;
  return (
    <div className="mb-3 border-bottom-line">
      <FormGroup row>
        <Label for="select-well" md={12} className="mb-3	">
          Select well
        </Label>
        <Col md={12}>
          <Coordinate direction="horizontal" className="px-0 mx-0 well-no">
            {labelArray.map((label, index) => {
              return (
                <CoordinateItem key={index} coordinateValue={`${label}`} />
              );
            })}
          </Coordinate>
          <div className="d-flex align-items-center well-box mt-2">
            {labelArray.map((label, index) => {
              return <Well key={index} className="well" />;
            })}
          </div>
        </Col>
      </FormGroup>
    </div>
  );
};
