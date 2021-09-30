import React from "react";

import {
  Button,
  Form,
  FormGroup,
  Row,
  Card,
  CardBody,
  Col,
  Input,
  Label,
  Select,
} from "core-components";
import { Center, ButtonIcon, Text } from "shared-components";
import {
  MAX_MOTOR_DIRECTION,
  MAX_MOTOR_DISTANCE,
  MAX_MOTOR_NUMBER,
  MIN_MOTOR_DIRECTION,
  MIN_MOTOR_DISTANCE,
  MIN_MOTOR_NUMBER,
  MOTOR_NUMBER_OPTIONS,
  MOTOR_DIRECTION_OPTIONS,
} from "appConstants";
import { isBtnDisabled } from "./helpers";

const MotorComponent = (props) => {
  const { handleMotorBtn, handleSenseAndHitBtn, formik } = props;
  const { motorNumber, direction, distance } = formik.values;

  const handleBlurDistance = (value) => {
    if (!value || value > MAX_MOTOR_DISTANCE || value <= MIN_MOTOR_DISTANCE) {
      formik.setFieldValue("distance.isInvalid", true);
    }
  };

  const handleOnChange = (key, value) => {
    formik.setFieldValue(key, value);
  };

  let optionValue;
  //sets values to formik
  const handleOptionChange = (event, isMotorNumber = false) => {
    let value = null;
    let label = null;

    value = event.value;
    label = event.label;
    optionValue = { value: value, label: label };

    if (isMotorNumber === true) {
      formik.setFieldValue("motorNumber.value", value);
    } else {
      formik.setFieldValue("direction.value", value);
    }
  };

  return (
    <Card default className="my-3">
      <CardBody>
        <Text
          Tag="h4"
          size={24}
          className="text-center text-gray text-bold mt-3 mb-4"
        >
          {"Start Motor"}
        </Text>
        <Form>
          <Row form>
            <Col sm={4}>
              <FormGroup>
                <Label for="motor_number" className="font-weight-bold">
                  Motor Number
                </Label>
                <Select
                  placeholder="Select Motor Number"
                  className=""
                  size="sm"
                  value={optionValue}
                  options={MOTOR_NUMBER_OPTIONS}
                  onChange={(e) => handleOptionChange(e, true)}
                />
                {motorNumber.isInvalid && (
                  <div className="flex-70">
                    <Text Tag="p" size={14} className="text-danger">
                      Motor number should be between {MIN_MOTOR_NUMBER} -{" "}
                      {MAX_MOTOR_NUMBER}.
                    </Text>
                  </div>
                )}
              </FormGroup>
            </Col>

            <Col sm={4}>
              <FormGroup>
                <Label for="direction" className="font-weight-bold">
                  Direction
                </Label>
                <Select
                  placeholder="Select Motor Direction"
                  className=""
                  size="sm"
                  value={optionValue}
                  options={MOTOR_DIRECTION_OPTIONS}
                  onChange={handleOptionChange}
                />
                {direction.isInvalid && (
                  <div className="flex-70">
                    <Text Tag="p" size={14} className="text-danger">
                      Motor direction can only be either {MIN_MOTOR_DIRECTION}{" "}
                      or {MAX_MOTOR_DIRECTION}.
                    </Text>
                  </div>
                )}
              </FormGroup>
            </Col>

            <Col sm={4}>
              <FormGroup>
                <Label for="distance" className="font-weight-bold">
                  Distance
                </Label>
                <Input
                  type="number"
                  name="distance"
                  id="distance"
                  placeholder={`${MIN_MOTOR_DISTANCE} - ${MAX_MOTOR_DISTANCE}`}
                  value={distance.value}
                  onChange={(event) =>
                    handleOnChange(
                      "distance.value",
                      parseFloat(event.target.value)
                    )
                  }
                  onBlur={(e) => handleBlurDistance(parseFloat(e.target.value))}
                  onFocus={() => handleOnChange("distance.isInvalid", false)}
                />
                {distance.isInvalid && (
                  <div className="flex-70">
                    <Text Tag="p" size={14} className="text-danger">
                      Motor distance should be between {MIN_MOTOR_DISTANCE} -{" "}
                      {MAX_MOTOR_DISTANCE}.
                    </Text>
                  </div>
                )}
              </FormGroup>
            </Col>
          </Row>

          <Row>
            <Col>
              <Center className="text-center pt-3">
                <Button
                  className="mx-3"
                  disabled={isBtnDisabled(formik.values)}
                  color="primary"
                  onClick={handleMotorBtn}
                >
                  Start Motor
                </Button>

                <Button
                  className="mx-3"
                  color="primary"
                  disabled={motorNumber.value === null}
                  onClick={handleSenseAndHitBtn}
                >
                  {"Sense & Hit"}
                </Button>
              </Center>
            </Col>
          </Row>
        </Form>
      </CardBody>
    </Card>
  );
};

export default React.memo(MotorComponent);
