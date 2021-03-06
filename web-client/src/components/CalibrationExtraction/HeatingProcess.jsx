import React from "react";
import { useState } from "react";

import { Row, Col } from "core-components";
import { TemperatureInfo, TimeInfo } from "shared-components";

import { HeatingProcessBox } from "./Style";
import { MAX_TEMP_ALLOWED, MIN_TEMP_ALLOWED } from "appConstants";

const HeatingProcess = (props) => {
  const { formik } = props;

  const [invalidTemp, setInvalidTemp] = useState(false);

  const handleOnBlur = (tempValue) => {
    if (tempValue > MAX_TEMP_ALLOWED || tempValue < MIN_TEMP_ALLOWED) {
      setInvalidTemp(true);
    }
  };

  const handleOnFocus = () => {
    setInvalidTemp(false);
  };

  return (
    <>
      <HeatingProcessBox className="p-5">
        <div className="process-box mx-auto py-5">
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
              invalidTemperature={invalidTemp}
              handleOnFocus={handleOnFocus}
              handleOnBlur={handleOnBlur}
            />
          </div>
          <Row>
            <Col lg={8} className="py-4">
              {/* <TimeInfo /> */}
              <TimeInfo
                hours={formik.values.hours}
                handleHoursChange={(e) =>
                  formik.setFieldValue("hours", e.target.value)
                }
                mins={formik.values.mins}
                handleMinsChange={(e) =>
                  formik.setFieldValue("mins", e.target.value)
                }
                secs={formik.values.secs}
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
