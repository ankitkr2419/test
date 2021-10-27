import React from "react";
import PropTypes from "prop-types";

import { Col, FormGroup, Label } from "core-components";

import Well from "components/Plate/Grid/Well";
import CoordinateItem from "components/Plate/Grid/CoordinateItem";
import Coordinate from "components/Plate/Grid/Coordinate";

export const WellComponent = (props) => {
  const { wellsObjArray, wellClickHandler } = props;

  return (
    <div className="mb-3 border-bottom-line">
      <FormGroup row>
        <Label for="select-well" md={12} className="mb-3	">
          Select well
        </Label>
        <Col md={12}>
          <Coordinate direction="horizontal" className="px-0 mx-0 well-no">
            {wellsObjArray &&
              wellsObjArray.map((wellObj, index) => {
                return (
                  <CoordinateItem
                    key={wellObj.id}
                    coordinateValue={`${wellObj.label}`}
                  />
                );
              })}
          </Coordinate>
          <div className="d-flex align-items-center well-box mt-2">
            {wellsObjArray &&
              wellsObjArray.map((wellObj) => {
                return (
                  <Well
                    key={wellObj.id}
                    id={wellObj.id}
                    height={wellObj.height}
                    isRunning={wellObj.isRunning}
                    isSelected={wellObj.isSelected}
                    isDisabled={wellObj.isDisabled}
                    className="well"
                    onClickHandler={() =>
                      wellClickHandler(wellObj.id, wellObj.type)
                    }
                  />
                );
              })}
          </div>
        </Col>
      </FormGroup>
    </div>
  );
};

WellComponent.propTypes = {
  isSelected: PropTypes.bool,
  isRunning: PropTypes.bool,
  isDisabled: PropTypes.bool,
};

WellComponent.defaultProps = {
  isSelected: false,
  isRunning: false,
  isDisabled: false,
};
