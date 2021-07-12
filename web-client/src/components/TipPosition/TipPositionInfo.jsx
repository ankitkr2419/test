import React from "react";

import { Col, FormGroup } from "core-components";
import { Text } from "shared-components";

import Well from "components/Plate/Grid/Well";
import CoordinateItem from "components/Plate/Grid/CoordinateItem";
import Coordinate from "components/Plate/Grid/Coordinate";
import { TipPositionInfoBox } from "./Style";
import { typeName } from "./functions";
import { CommonTipHeightComponent } from "./CommonTipHeightComponent";

const TipPositionInfo = (props) => {
  const { formik, activeTab, wellClickHandler } = props;
  const wellsObjArray = formik.values[`${typeName[activeTab]}`].wellsArray;

  return (
    <>
      <TipPositionInfoBox>
        <div className="process-box tip-position-box mx-auto">
          <div className="mb-3 border-bottom-line">
            <FormGroup row>
              <Text
                tag="h6"
                for="select-well"
                md={12}
                className="mb-3	font-weight-bold"
              >
                Select well
              </Text>
              <Col md={12}>
                <Coordinate
                  direction="horizontal"
                  className="px-0 mx-0 well-no"
                >
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
                        <>
                          <Well
                            key={wellObj.id}
                            id={wellObj.id}
                            isRunning={wellObj.isRunning}
                            isSelected={wellObj.isSelected}
                            // isDisabled={wellObj.isDisabled}
                            className={`well mb-3`}
                            onClickHandler={() =>
                              wellClickHandler(wellObj.id, wellObj.type)
                            }
                          />
                        </>
                      );
                    })}
                </div>

                <Coordinate
                  direction="horizontal"
                  className="px-0 mx-0 well-no"
                >
                  {wellsObjArray &&
                    wellsObjArray.map((wellObj, index) => {
                      return (
                        <CoordinateItem
                          key={wellObj.id}
                          coordinateValue={
                            wellObj.height &&
                            wellObj.isSelected &&
                            `${wellObj.height} mm`
                          }
                        />
                      );
                    })}
                </Coordinate>
              </Col>
            </FormGroup>
          </div>
          <CommonTipHeightComponent formik={formik} activeTab={activeTab} />
        </div>
      </TipPositionInfoBox>
    </>
  );
};

TipPositionInfo.propTypes = {};

export default TipPositionInfo;
