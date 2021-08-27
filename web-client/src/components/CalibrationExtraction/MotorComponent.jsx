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
  Modal,
  ModalBody,
} from "core-components";
import { Center, ButtonIcon, Text } from "shared-components";
import {
  MAX_MOTOR_DIRECTION,
  MAX_MOTOR_DISTANCE,
  MAX_MOTOR_NUMBER,
  MIN_MOTOR_DIRECTION,
  MIN_MOTOR_DISTANCE,
  MIN_MOTOR_NUMBER,
} from "appConstants";
import { isBtnDisabled } from "./helpers";

const MotorComponent = (props) => {
  const { handleMotorBtn, formik } = props;
  const { motorNumber, direction, distance } = formik.values;

  const handleBlurMotorNumber = (value) => {
    if (!value || value > MAX_MOTOR_NUMBER || value < MIN_MOTOR_NUMBER) {
      formik.setFieldValue("motorNumber.isInvalid", true);
    }
  };
  const handleBlurDirection = (value) => {
    if (
      value === null ||
      value === "" ||
      (value !== MIN_MOTOR_DIRECTION && value !== MAX_MOTOR_DIRECTION)
    ) {
      formik.setFieldValue("direction.isInvalid", true);
    }
  };
  const handleBlurDistance = (value) => {
    if (!value || value > MAX_MOTOR_DISTANCE || value <= MIN_MOTOR_DISTANCE) {
      formik.setFieldValue("distance.isInvalid", true);
    }
  };

  const handleOnChange = (key, value) => {
    formik.setFieldValue(key, value);
  };

  return (
    <Card default className="my-3">
      <CardBody>
        <Form onSubmit={handleMotorBtn}>
          <Row form>
            <Col sm={4}>
              <FormGroup>
                <Label for="motor_number" className="font-weight-bold">
                  Motor Number
                </Label>
                <Input
                  type="number"
                  name="motor_number"
                  id="motor_number"
                  placeholder={`${MIN_MOTOR_NUMBER} - ${MAX_MOTOR_NUMBER}`}
                  value={motorNumber.value}
                  onChange={(event) =>
                    handleOnChange(
                      "motorNumber.value",
                      parseInt(event.target.value)
                    )
                  }
                  onBlur={(e) =>
                    handleBlurMotorNumber(parseInt(e.target.value))
                  }
                  onFocus={() => handleOnChange("motorNumber.isInvalid", false)}
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
                <Input
                  type="number"
                  name="direction"
                  id="direction"
                  placeholder={`${MIN_MOTOR_DIRECTION} / ${MAX_MOTOR_DIRECTION}`}
                  value={direction.value}
                  max={MAX_MOTOR_DIRECTION}
                  min={MIN_MOTOR_DIRECTION}
                  onChange={(event) =>
                    handleOnChange(
                      "direction.value",
                      parseInt(event.target.value)
                    )
                  }
                  onBlur={(e) => handleBlurDirection(parseInt(e.target.value))}
                  onFocus={() => handleOnChange("direction.isInvalid", false)}
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
                      parseInt(event.target.value)
                    )
                  }
                  onBlur={(e) => handleBlurDistance(parseInt(e.target.value))}
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
                <Button disabled={isBtnDisabled(formik.values)} color="primary">
                  Start Motor
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
