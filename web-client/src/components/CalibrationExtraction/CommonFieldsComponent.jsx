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
  const {
    name,
    email,
    roomTemperature,
    serialNo,
    manufacturingYear,
    machineVersion,
    softwareVersion,
    contactNumber,
  } = formik.values;

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
        <Text
          Tag="h4"
          size={24}
          className="text-center text-gray text-bold mt-3 mb-4"
        >
          {"Common Details"}
        </Text>

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

            <Col sm={4}>
              <FormGroup>
                <Label for="serialNo" className="font-weight-bold">
                  serial Number
                </Label>
                <Input
                  name="serialNo"
                  id="serialNo"
                  placeholder={"Type here"}
                  value={serialNo.value}
                  onChange={(event) => {
                    handleOnChange("serialNo.value", event.target.value);
                  }}
                  onFocus={() => handleOnChange("serialNo.isInvalid", false)}
                />
                {serialNo.isInvalid && (
                  <div className="flex-70">
                    <Text Tag="p" size={14} className="text-danger">
                      "Enter serial number"
                    </Text>
                  </div>
                )}
              </FormGroup>
            </Col>

            <Col sm={4}>
              <FormGroup>
                <Label for="manufacturingYear" className="font-weight-bold">
                  Manufacturung Year
                </Label>
                <Input
                  name="manufacturungYear"
                  id="manufacturungYear"
                  placeholder={"Type here"}
                  value={manufacturingYear.value}
                  onChange={(event) => {
                    handleOnChange(
                      "manufacturingYear.value",
                      event.target.value
                    );
                  }}
                  onFocus={() =>
                    handleOnChange("manufacturingYear.isInvalid", false)
                  }
                />
                {manufacturingYear.isInvalid && (
                  <div className="flex-70">
                    <Text Tag="p" size={14} className="text-danger">
                      "Enter manufacturing Year"
                    </Text>
                  </div>
                )}
              </FormGroup>
            </Col>

            <Col sm={4}>
              <FormGroup>
                <Label for="machineVersion" className="font-weight-bold">
                  Machine Version
                </Label>
                <Input
                  name="machineVersion"
                  id="machineVersion"
                  placeholder={"Type here"}
                  value={machineVersion.value}
                  onChange={(event) => {
                    handleOnChange("machineVersion.value", event.target.value);
                  }}
                  onFocus={() =>
                    handleOnChange("machineVersion.isInvalid", false)
                  }
                />
                {machineVersion.isInvalid && (
                  <div className="flex-70">
                    <Text Tag="p" size={14} className="text-danger">
                      "Enter Machine Version"
                    </Text>
                  </div>
                )}
              </FormGroup>
            </Col>

            <Col sm={4}>
              <FormGroup>
                <Label for="softwareVersion" className="font-weight-bold">
                  Software Version
                </Label>
                <Input
                  name="softwareVersion"
                  id="softwareVersion"
                  placeholder={"Type here"}
                  value={softwareVersion.value}
                  onChange={(event) => {
                    handleOnChange("softwareVersion.value", event.target.value);
                  }}
                  onFocus={() =>
                    handleOnChange("softwareVersion.isInvalid", false)
                  }
                />
                {machineVersion.isInvalid && (
                  <div className="flex-70">
                    <Text Tag="p" size={14} className="text-danger">
                      "Enter Software Version"
                    </Text>
                  </div>
                )}
              </FormGroup>
            </Col>

            <Col sm={4}>
              <FormGroup>
                <Label for="contactNumber" className="font-weight-bold">
                  Contact Number
                </Label>
                <Input
                  name="contactNumber"
                  id="contactNumber"
                  placeholder={"Type here"}
                  value={contactNumber.value}
                  onChange={(event) => {
                    handleOnChange("contactNumber.value", event.target.value);
                  }}
                  onFocus={() =>
                    handleOnChange("contactNumber.isInvalid", false)
                  }
                />
                {contactNumber.isInvalid && (
                  <div className="flex-70">
                    <Text Tag="p" size={14} className="text-danger">
                      "Enter Contact Number"
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
