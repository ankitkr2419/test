import React from "react";
import PropTypes from "prop-types";

import { Col } from "core-components";
import { Text } from "shared-components";

import Well from "components/Plate/Grid/Well";
import CoordinateItem from "components/Plate/Grid/CoordinateItem";
import Coordinate from "components/Plate/Grid/Coordinate";

export const WellComponent = (props) => {
  const { wellsObjArray, wellClickHandler } = props;
  return (
    <div className="mb-3 border-bottom-line">
      <div className="row">
        <Col md={12}>
          <Text Tag="h6" md={12} className="mb-1">
            Select well
          </Text>
        </Col>

        <Col className="mb-4" md={12}>
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
              wellsObjArray.map((wellObj, index) => {
                return (
                  <Well
                    key={wellObj.id}
                    id={wellObj.id}
                    isRunning={wellObj.isRunning}
                    isSelected={wellObj.isSelected}
                    // isDisabled={wellObj.isDisabled}
                    className={`well ${wellObj.footerText}`}
                    onClickHandler={() =>
                      wellClickHandler(wellObj.id, wellObj.type)
                    }
                  />
                );
              })}
          </div>
        </Col>
      </div>
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
