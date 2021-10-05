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
  CheckBox,
} from "core-components";
import { Center, Text } from "shared-components";
import { calibrationStatusMessage } from "./helper";
import { PID_STATUS } from "appConstants";

const DyeCalibration = (props) => {
  let {
    dyeCalibrationStatus,
    dyeOptions,
    formikDyeCalibration,
    handleDyeCalibrationButton,
  } = props;
  let { selectedDye, kitID } = formikDyeCalibration.values;

  const handleBlurKitID = (value) => {
    if (!value || value.length !== kitID.allowedLength) {
      formikDyeCalibration.setFieldValue("kitID.isInvalid", true);
    }
  };
  const handleOnChange = (key, value) => {
    formikDyeCalibration.setFieldValue(key, value);
  };

  const handleSubmit = (e) => {
    e.preventDefault();
    if (selectedDye.isInvalid === false && kitID.isInvalid === false) {
      let selectedDyeID = selectedDye.value.value;
      let selectedKitID = kitID.value;

      //create request body
      handleDyeCalibrationButton({
        dye_id: selectedDyeID,
        kit_id: `${selectedKitID}`, //string format as per api
      });
    }
  };

  let isDyeCalibrationDisabled =
    selectedDye.value === null ||
    kitID.value === null ||
    selectedDye.isInvalid ||
    kitID.isInvalid ||
    dyeCalibrationStatus === PID_STATUS.running ||
    dyeCalibrationStatus === PID_STATUS.progressing;

  let message = calibrationStatusMessage(dyeCalibrationStatus);

  return (
    <Card default className="my-3">
      <CardBody>
        <Text
          Tag="h4"
          size={24}
          className="text-center text-gray text-bold mt-3 mb-4"
        >
          Dye Calibration
        </Text>

        <Form onSubmit={handleSubmit}>
          <Row form>
            <Col sm={4}>
              <FormGroup>
                <Label for="dye" className="font-weight-bold">
                  Dye
                </Label>
                <div>
                  <Select
                    placeholder="Select Type"
                    options={dyeOptions}
                    value={selectedDye.value}
                    onChange={(value) =>
                      handleOnChange("selectedDye.value", value)
                    }
                    isSearchable={false}
                  />
                </div>
              </FormGroup>
            </Col>

            <Col sm={4}>
              <FormGroup>
                <Label for="kitId" className="font-weight-bold">
                  Kit ID
                </Label>
                <Input
                  type="text"
                  name="kitId"
                  id="kitId"
                  maxLength={kitID.allowedLength}
                  placeholder="Type Kit Id"
                  value={kitID.value}
                  onChange={(e) =>
                    handleOnChange("kitID.value", e.target.value)
                  }
                  onBlur={(e) => handleBlurKitID(e.target.value)}
                  onFocus={() => handleOnChange("kitID.isInvalid", false)}
                />
                {(kitID.isInvalid || kitID.value == null) && (
                  <div className="flex-70">
                    <Text Tag="p" size={14} className="text-danger">
                      8 characters required
                    </Text>
                  </div>
                )}
              </FormGroup>
            </Col>
            {message && (
              <Col sm={4} className="d-flex align-items-center">
                <Text Tag="h4" size={16} className="text-gray">
                  {message}
                </Text>
              </Col>
            )}
          </Row>
          <Center className="text-center pt-3">
            <Button
              /*disabled={isDyeCalibrationDisabled}  (NOTE:may need in future) */ color="primary"
            >
              Start
            </Button>
          </Center>
        </Form>
      </CardBody>
    </Card>
  );
};

export default React.memo(DyeCalibration);
