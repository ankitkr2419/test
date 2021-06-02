import React from "react";

import { Row, Col } from "core-components";
import { TemperatureInfo, TimeInfo } from "shared-components";

import { HeatingProcessBox } from "./Style";

const HeatingProcess = (props) => {
  const { formik } = props;
  return (
    <>
      <HeatingProcessBox className="p-5">
        <div className="process-box mx-auto py-5">
          <div className="border-bottom-line">
            <TemperatureInfo
              temperatureHandler={(e) => {
                formik.setFieldValue("temperature", e.target.value);
              }}
              checkBoxHandler={(e) => {
                formik.setFieldValue("followTemperature", e.target.checked);
              }}
            />
          </div>
          <Row>
            <Col lg={8} className="py-4">
              {/* <TimeInfo /> */}
              <TimeInfo
                handleHoursChange={(e) =>
                  formik.setFieldValue("hours", e.target.value)
                }
                handleMinsChange={(e) =>
                  formik.setFieldValue("mins", e.target.value)
                }
                handleSecsChange={(e) =>
                  formik.setFieldValue("secs", e.target.value)
                }
              />
            </Col>
          </Row>
        </div>
      </HeatingProcessBox>
    </>
  );
};

HeatingProcess.propTypes = {};

export default HeatingProcess;
