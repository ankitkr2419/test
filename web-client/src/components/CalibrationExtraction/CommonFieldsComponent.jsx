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
import { Center, Text } from "shared-components";
import {
  EMAIL_REGEX,
  EMAIL_REGEX_OR_EMPTY_STR,
  MAX_ROOM_TEMPERATURE,
  MIN_ROOM_TEMPERATURE,
} from "appConstants";
import { isSaveDetailsBtnDisabled } from "./helpers";

const CommonFieldsComponent = (props) => {
  const { handleSaveDetailsBtn, formik } = props;
  const { name, email, roomTemperature } = formik.values;

  const handleBlurName = (value) => {
    if (!value) {
      formik.setFieldValue("name.isInvalid", true);
    }
  };
  const handleBlurEmail = (value) => {
    if (!value.match(EMAIL_REGEX)) {
      formik.setFieldValue("email.isInvalid", true);
    }
  };
  const handleBlurRoomTemp = (value) => {
    if (
      !value ||
      value > MAX_ROOM_TEMPERATURE ||
      value <= MIN_ROOM_TEMPERATURE
    ) {
      formik.setFieldValue("roomTemperature.isInvalid", true);
    }
  };

  const handleOnChange = (key, value) => {
    formik.setFieldValue(key, value);
  };

  return (
    <Card default className="my-3">
      <CardBody>
        <Form>
          <Row form>
            <Col sm={4}>
              <FormGroup>
                <Label for="name" className="font-weight-bold">
                  Name
                </Label>
                <Input
                  name="name"
                  id="name"
                  placeholder={"Type here"}
                  value={name.value}
                  onChange={(event) =>
                    handleOnChange("name.value", event.target.value)
                  }
                  onBlur={(e) => handleBlurName(e.target.value)}
                  onFocus={() => handleOnChange("name.isInvalid", false)}
                />
                {name.isInvalid && (
                  <div className="flex-70">
                    <Text Tag="p" size={14} className="text-danger">
                      {"Enter valid name"}
                    </Text>
                  </div>
                )}
              </FormGroup>
            </Col>

            <Col sm={4}>
              <FormGroup>
                <Label for="email" className="font-weight-bold">
                  Email
                </Label>
                <Input
                  type="email"
                  name="email"
                  id="email"
                  placeholder={"Type here"}
                  value={email.value}
                  onChange={(event) =>
                    handleOnChange("email.value", event.target.value)
                  }
                  onBlur={(e) => handleBlurEmail(e.target.value)}
                  onFocus={() => handleOnChange("email.isInvalid", false)}
                />
                {email.isInvalid && (
                  <div className="flex-70">
                    <Text Tag="p" size={14} className="text-danger">
                      Email is invalid.
                    </Text>
                  </div>
                )}
              </FormGroup>
            </Col>

            <Col sm={4}>
              <FormGroup>
                <Label for="roomTemperature" className="font-weight-bold">
                  Room Temperature
                </Label>
                <Input
                  type="number"
                  name="roomTemperature"
                  id="roomTemperature"
                  placeholder={`${20} - ${30}`}
                  value={roomTemperature.value}
                  onChange={(event) =>
                    handleOnChange(
                      "roomTemperature.value",
                      parseInt(event.target.value)
                    )
                  }
                  onBlur={(e) => handleBlurRoomTemp(parseInt(e.target.value))}
                  onFocus={() =>
                    handleOnChange("roomTemperature.isInvalid", false)
                  }
                />
                {roomTemperature.isInvalid && (
                  <div className="flex-70">
                    <Text Tag="p" size={14} className="text-danger">
                      Room temperature should be between {MIN_ROOM_TEMPERATURE}{" "}
                      - {MAX_ROOM_TEMPERATURE}.
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
                  onClick={() => handleSaveDetailsBtn(formik.values)}
                  disabled={isSaveDetailsBtnDisabled(formik.values)}
                  color="primary"
                >
                  Save Details
                </Button>
              </Center>
            </Col>
          </Row>
        </Form>
      </CardBody>
    </Card>
  );
};

export default React.memo(CommonFieldsComponent);
