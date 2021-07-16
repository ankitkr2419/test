import React from "react";

import { FormGroup, Label, Input, FormError, CheckBox } from "core-components";
import {} from "shared-components";

import styled from "styled-components";

const TemperatureInfoBox = styled.div`
  .temperature-info {
    .custom-checkbox {
      > label {
        font-size: 0.875rem;
        line-height: 1rem;
        color: #666666;
      }
      .custom-control-label::after {
        left: -2.25rem;
      }
    }
    label {
      font-size: 1rem;
      line-height: 1.125rem;
      color: #666666;
    }
    .temperature-box {
      margin-left: 2.375rem;
    }
  }
`;

const TemperatureInfo = (props) => {
  const { temperature, followTemp, temperatureHandler, checkBoxHandler } =
    props;
  return (
    <TemperatureInfoBox>
      <FormGroup className="d-flex temperature-info">
        <Label for="discard" className="px-0 mt-2">
          Temperature
        </Label>
        <div className="temperature-box">
          <Input
            type="text"
            name="temperature"
            id="temperature"
            placeholder="Type here"
            value={temperature}
            onChange={temperatureHandler}
          />
          <FormError>Incorrect Temperature</FormError>

          <div className="d-flex mt-3">
            <CheckBox
              id="follow-temperature"
              name="follow-temperature"
              label="Follow Temperature"
              className="mb-3 mr-4"
              checked={followTemp}
              onChange={checkBoxHandler}
            />
          </div>
        </div>
      </FormGroup>
    </TemperatureInfoBox>
  );
};

TemperatureInfo.propTypes = {};

export default TemperatureInfo;
