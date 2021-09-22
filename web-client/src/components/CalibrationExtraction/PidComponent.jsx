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
} from "core-components";
import { Center, ButtonIcon, Text } from "shared-components";
import {
  MAX_PID_MIN,
  MAX_PID_TEMP,
  MIN_PID_MIN,
  MIN_PID_TEMP,
} from "appConstants";
import { isBtnDisabled, isPidUpdateBtnDisabled } from "./helpers";

const PidComponent = (props) => {
  const { handlePidBtn, formik } = props;
  const { pidTemperature, pidMinutes } = formik.values;

  const handleBlurPidTemp = (value) => {
    if (!value || value > MAX_PID_TEMP || value < MIN_PID_TEMP) {
      formik.setFieldValue("pidTemperature.isInvalid", true);
    }
  };
  const handleBlurPidMinutes = (value) => {
    if (!value || value > MAX_PID_MIN || value < MIN_PID_MIN) {
      formik.setFieldValue("pidMinutes.isInvalid", true);
    }
  };

  const handleOnChange = (key, value) => {
    formik.setFieldValue(key, value);
  };

  return (
    <Card default className="my-3 w-100">
      <CardBody>
        <Form>
          <Row form>
            <Col sm={4}>
              <FormGroup>
                <Label for="pid_temperature" className="font-weight-bold ml-3">
                  PID Temperature
                </Label>
                <Input
                  className="ml-3"
                  type="number"
                  name="pid_temperature"
                  id="pid_temperature"
                  placeholder={`${MIN_PID_TEMP} - ${MAX_PID_TEMP}`}
                  value={pidTemperature.value}
                  onChange={(event) =>
                    handleOnChange(
                      "pidTemperature.value",
                      parseInt(event.target.value)
                    )
                  }
                  onBlur={(e) => handleBlurPidTemp(parseInt(e.target.value))}
                  onFocus={() =>
                    handleOnChange("pidTemperature.isInvalid", false)
                  }
                />
                {pidTemperature.isInvalid && (
                  <div className="flex-70">
                    <Text Tag="p" size={14} className="text-danger">
                      PID temperature should be between {MIN_PID_TEMP} -{" "}
                      {MAX_PID_TEMP}.
                    </Text>
                  </div>
                )}
              </FormGroup>
            </Col>
            <Col>
              <Center className="text-center pt-4">
                <Button
                  onClick={() => handlePidBtn(formik.values)}
                  disabled={isPidUpdateBtnDisabled(formik.values)}
                  color="primary"
                >
                  Update PID
                </Button>
              </Center>
            </Col>
          </Row>
        </Form>
      </CardBody>
    </Card>
  );
};

export default React.memo(PidComponent);
