import React from "react";

import { Row, Col, FormGroup, Label, Input, FormError } from "core-components";
import { Text } from "shared-components";

import styled from "styled-components";

const TimeInfoBox = styled.div`
  .time-info-box {
    .time-box {
      label {
        font-size: 0.75rem;
        line-height: 0.875rem;
      }
    }
  }
`;

const TimeInfo = (props) => {
  const {
    hours,
    mins,
    secs,
    handleHoursChange,
    handleMinsChange,
    handleSecsChange,
  } = props;
  return (
    <>
      <TimeInfoBox>
        <div className="time-info-box">
          <FormGroup row className="d-flex time-box">
            <Col sm={2} className="mt-2">
              <Text>Time</Text>
            </Col>
            <Col sm={8}>
              <Row>
                <Col sm={3}>
                  <Input
                    type="number"
                    name="hours"
                    id="hours"
                    placeholder=""
                    value={hours}
                    onChange={handleHoursChange}
                  />
                  <Label for="hours" className="font-weight-bold">
                    Hours
                  </Label>
                  <FormError>Incorrect Hours</FormError>
                </Col>
                <Col sm={3}>
                  <Input
                    type="number"
                    name="minutes"
                    id="minutes"
                    placeholder=""
                    value={mins}
                    onChange={handleMinsChange}
                  />
                  <Label for="minutes" className="font-weight-bold px-1">
                    Minutes
                  </Label>
                  <FormError>Incorrect Minutes</FormError>
                </Col>
                <Col sm={3}>
                  <Input
                    type="number"
                    name="seconds"
                    id="seconds"
                    placeholder=""
                    value={secs}
                    onChange={handleSecsChange}
                  />
                  <Label for="minutes" className="font-weight-bold">
                    Seconds
                  </Label>
                  <FormError>Incorrect Seconds</FormError>
                </Col>
              </Row>
            </Col>
          </FormGroup>
        </div>
      </TimeInfoBox>
    </>
  );
};

TimeInfo.propTypes = {};

export default TimeInfo;
