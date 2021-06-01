import React from "react";

import { Row, Col, FormGroup, Label, Input, FormError } from "core-components";
import { TemperatureInfo, TimeInfo } from "shared-components";

import { ShakingProcessBox } from "./Style";
import { isDisabled, setRpmFormikField } from "./functions";

const ShakingProcess = (props) => {
  const { formik, activeTab, temperature } = props;

  return (
    <>
      <ShakingProcessBox>
        <div className="process-box mx-auto">
          {temperature && (
            <div className="border-bottom-line">
              <TemperatureInfo
                temperature={formik.values.temperature}
                followTemp={formik.values.followTemperature}
                temperatureHandler={(e) => {
                  formik.setFieldValue("temperature", e.target.value);
                }}
                checkBoxHandler={(e) => {
                  formik.setFieldValue("followTemperature", e.target.checked);
                }}
              />
            </div>
          )}

          {/* RPM 1 */}
          <Row>
            <Col lg={3} className="py-4">
              <FormGroup row className="d-flex align-items-center">
                <Label for="tip-selection" className="mb-0">
                  RPM
                </Label>
                <div className="ml-3 rpm-input">
                  <Input
                    type="text"
                    name="rpm1"
                    id="rpm"
                    placeholder="Type here"
                    value={formik.values.rpm1}
                    onChange={(e) =>
                      setRpmFormikField(
                        formik,
                        activeTab,
                        "rpm1",
                        e.target.value
                      )
                    }
                  />
                  <FormError>Incorrect RPM</FormError>
                </div>
              </FormGroup>
            </Col>
            <Col lg={9} className="border-left-line py-4">
              <Row>
                <Col sm={11} className="ml-4 mr-auto">
                  <TimeInfo
                    hours={formik.values.hours1}
                    handleHoursChange={(e) =>
                      formik.setFieldValue("hours1", e.target.value)
                    }
                    mins={formik.values.mins1}
                    handleMinsChange={(e) =>
                      formik.setFieldValue("mins1", e.target.value)
                    }
                    secs={formik.values.secs1}
                    handleSecsChange={(e) =>
                      formik.setFieldValue("secs1", e.target.value)
                    }
                  />
                </Col>
              </Row>
            </Col>
          </Row>

          {/* RPM 2 */}
          <Row className={isDisabled.rpm2 && "disabled"}>
            <Col lg={3}>
              <FormGroup row className="d-flex align-items-center">
                <Label for="tip-selection" className="mb-0">
                  RPM
                </Label>
                <div className="ml-3 rpm-input">
                  <Input
                    type="text"
                    name="rpm2"
                    id="rpm"
                    placeholder="Type here"
                    value={formik.values.rpm2}
                    onChange={(e) =>
                      setRpmFormikField(
                        formik,
                        activeTab,
                        "rpm2",
                        e.target.value
                      )
                    }
                  />
                  <FormError>Incorrect RPM</FormError>
                </div>
              </FormGroup>
            </Col>
            <Col lg={9} className="border-left-line">
              <Row>
                <Col sm={11} className="ml-4 mr-auto">
                  <TimeInfo
                    hours={formik.values.hours2}
                    handleHoursChange={(e) =>
                      formik.setFieldValue("hours2", e.target.value)
                    }
                    secs={formik.values.secs2}
                    mins={formik.values.mins2}
                    handleMinsChange={(e) =>
                      formik.setFieldValue("mins2", e.target.value)
                    }
                    handleSecsChange={(e) =>
                      formik.setFieldValue("secs2", e.target.value)
                    }
                  />
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
