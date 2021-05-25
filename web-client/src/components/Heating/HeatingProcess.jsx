import React from "react";

import { Row, Col } from "core-components";
import { TemperatureInfo, TimeInfo } from "shared-components";

import { HeatingProcessBox } from "./Style";

const HeatingProcess = (props) => {
  return (
    <>
      <HeatingProcessBox className="p-5">
        <div className="process-box mx-auto py-5">
          <div className="border-bottom-line">
            <TemperatureInfo />
          </div>
          <Row>
            <Col lg={8} className="py-4">
              <TimeInfo />
            </Col>
          </Row>
        </div>
      </HeatingProcessBox>
    </>
  );
};

HeatingProcess.propTypes = {};

export default HeatingProcess;
