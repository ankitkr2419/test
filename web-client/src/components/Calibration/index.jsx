import React, { useEffect, useState } from "react";
import {
  Button,
  Form,
  FormGroup,
  Input,
  Label,
  Card,
  CardBody,
  Row,
  Col,
} from "core-components";
import { Text } from "shared-components";

import { getFormikInitialState } from "./helper";
import { useCallback } from "react";
import { useFormik } from "formik";

const CalibrationComponent = (props) => {
  let { configs, saveBtnClickHandler } = props;

  const formik = useFormik({
    initialValues: getFormikInitialState(),
    enableReinitialize: true,
  });

  //store new data in local state
  useEffect(() => {
    if (configs?.room_temperature) {
      formik.setFieldValue(`roomTemperature.value`, configs.room_temperature);
    }
    if (configs?.homing_time) {
      formik.setFieldValue(`homingTime.value`, configs.homing_time);
    }
    if (configs?.no_of_homing_cycles) {
      formik.setFieldValue(
        `noOfHomingCycles.value`,
        configs.no_of_homing_cycles
      );
    }
    if (configs?.cycle_time) {
      formik.setFieldValue(`cycleTime.value`, configs.cycle_time);
    }
    if (configs?.pid_temperature) {
      formik.setFieldValue(`pidTemperature.value`, configs.pid_temperature);
    }
    if (configs?.pid_minutes) {
      formik.setFieldValue(`pidMinutes.value`, configs.pid_minutes);
    }
  }, [configs]);

  //validations and api call
  const onSubmit = (e) => {
    e.preventDefault();

    const {
      roomTemperature,
      homingTime,
      noOfHomingCycles,
      cycleTime,
      pidTemperature,
      pidMinutes,
    } = formik.values;

    if (
      roomTemperature.value !== null &&
      homingTime.value !== null &&
      noOfHomingCycles.value !== null &&
      cycleTime.value !== null &&
      pidTemperature !== null &&
      pidMinutes !== null
    ) {
      saveBtnClickHandler({
        roomTemperature: parseInt(roomTemperature.value),
        homingTime: parseInt(homingTime.value),
        noOfHomingCycles: parseInt(noOfHomingCycles.value),
        cycleTime: parseInt(cycleTime.value),
        pidTemperature: parseInt(pidTemperature.value),
        pidMinutes: parseInt(pidMinutes.value),
      });
    }
  };

  const handleBlurChange = useCallback((name, value, min, max) => {
    if (!value || value < min || value > max) {
      formik.setFieldValue(`${name}.isInvalid`, true);
    }
  }, []);

  const isBtnDisabled = useCallback(() => {
    const {
      roomTemperature,
      homingTime,
      noOfHomingCycles,
      cycleTime,
      pidTemperature,
      pidMinutes,
    } = formik.values;

    return (
      !roomTemperature.value ||
      !homingTime.value ||
      !noOfHomingCycles.value ||
      !cycleTime.value ||
      !pidTemperature.value ||
      !pidMinutes.value ||
      roomTemperature.isInvalid ||
      homingTime.isInvalid ||
      noOfHomingCycles.isInvalid ||
      cycleTime.isInvalid ||
      pidTemperature.isInvalid ||
      pidMinutes.isInvalid
    );
  });

  return (
    <div className="calibration-content px-5">
      <Card default className="my-5">
        <CardBody className="px-5 py-4">
          <Form onSubmit={onSubmit}>
            <Row>
              {Object.keys(formik.values).map((key, index) => {
                const element = formik.values[key];
                const { type, name, label, min, max, value, isInvalid } =
                  element;
                return (
                  <Col md={6}>
                    <FormGroup>
                      <Label for="username">{label}</Label>
                      <Input
                        type={type}
                        name={name}
                        id={name}
                        placeholder={`${min} - ${max}`}
                        value={value}
                        onChange={(event) => {
                          formik.setFieldValue(
                            `${name}.value`,
                            event.target.value
                          );
                        }}
                        onBlur={(event) =>
                          handleBlurChange(
                            name,
                            parseInt(event.target.value),
                            min,
                            max
                          )
                        }
                        onFocus={() =>
                          formik.setFieldValue(`${name}.isInvalid`, false)
                        }
                      />
                      {(isInvalid || value == null) && (
                        <div className="flex-70">
                          <Text Tag="p" size={14} className="text-danger">
                            {`${label} should be between ${min} - ${max}`}
                          </Text>
                        </div>
                      )}
                    </FormGroup>
                  </Col>
                );
              })}
            </Row>
            <div className="text-right pt-4 pb-1 mb-3">
              <Button color="primary" disabled={isBtnDisabled()}>
                Save
              </Button>
            </div>
          </Form>
        </CardBody>
      </Card>
    </div>
  );
};

export default React.memo(CalibrationComponent);
