import React from "react";

import { Row, Col, FormGroup, Label, Input, FormError } from "core-components";
import { TemperatureInfo, TimeInfo } from "shared-components";

import { ShakingProcessBox } from "./Style";

const ShakingProcess = (props) => {
  const { temperature } = props;
  return (
    <>
      <ShakingProcessBox>
        <div className="process-box mx-auto">
          {temperature && (
            <div className="border-bottom-line">
              <TemperatureInfo />
            </div>
          )}

          <Row>
            <Col lg={3} className="py-4">
              <FormGroup row className="d-flex align-items-center">
                <Label for="tip-selection" className="mb-0">
                  RPM
                </Label>
                <div className="ml-3 rpm-input">
                  <Input
                    type="text"
                    name="rpm"
                    id="rpm"
                    placeholder="Type here"
                  />
                  <FormError>Incorrect RPM</FormError>
                </div>
              </FormGroup>
            </Col>
            <Col lg={9} className="border-left-line py-4">
              <Row>
                <Col sm={11} className="ml-4 mr-auto">
                  <TimeInfo />
                </Col>
              </Row>
            </Col>
          </Row>
          <Row className="disabled">
            <Col lg={3}>
              <FormGroup row className="d-flex align-items-center">
                <Label for="tip-selection" className="mb-0">
                  RPM
                </Label>
                <div className="ml-3 rpm-input">
                  <Input
                    type="text"
                    name="rpm"
                    id="rpm"
                    placeholder="Type here"
                  />
                  <FormError>Incorrect RPM</FormError>
                </div>
              </FormGroup>
            </Col>
            <Col lg={9} className="border-left-line">
              <Row>
                <Col sm={11} className="ml-4 mr-auto">
                  <TimeInfo />
                </Col>
              </Row>
            </Col>
          </Row>
        </div>
      </ShakingProcessBox>
    </>
  );
};

ShakingProcess.propTypes = {};

export default ShakingProcess;
