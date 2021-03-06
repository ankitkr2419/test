import React, { useState } from "react";
import { useFormik } from "formik";

import {
  Button,
  Row,
  Card,
  CardBody,
  Col,
  Input,
  Label,
} from "core-components";
import { Center, Text } from "shared-components";
import {
  getToleranceInitialFormikState,
  isSaveToleranceBtnDisabled,
} from "./helper";
import { MAX_TOLERANCE_ALLOWED, MIN_TOLERANCE_ALLOWED } from "appConstants";

const ToleranceComponent = (props) => {
  const { toleranceData, handleSaveToleranceBtn } = props;

  const formik = useFormik({
    initialValues: getToleranceInitialFormikState(toleranceData),
    enableReinitialize: true,
  });

  const { tolerance } = formik.values;

  const handleBlur = (key, value) => {
    let isInvalid = false;
    const floatValue = parseFloat(value);
    if (
      !floatValue ||
      floatValue > MAX_TOLERANCE_ALLOWED ||
      floatValue < MIN_TOLERANCE_ALLOWED
    ) {
      isInvalid = true;
    }
    formik.setFieldValue(key, isInvalid);
  };

  const handleOnChange = (key, value) => {
    formik.setFieldValue(key, value);
  };

  const handleSaveBtn = () => {
    const requestBody = toleranceData.map((dataObj, index) => ({
      Name: dataObj.Name,
      Position: dataObj.Position,
      Tolerance: parseFloat(tolerance[index].value),
    }));
    handleSaveToleranceBtn(requestBody);
  };

  return (
    <Card default className="my-3">
      <CardBody>
        <Text
          Tag="h4"
          size={24}
          className="text-center text-gray text-bold mt-3 mb-4"
        >
          Tolerance Configuration
        </Text>
        {tolerance?.length > 0 ? (
          toleranceData.map((dyeObj, index) => (
            <Row className="my-3">
              <Col className="">
                <Label for="name" className="font-weight-bold">
                  Name
                </Label>
                <Input disabled value={dyeObj.Name} />
              </Col>

              <Col className="">
                <Label for="position" className="font-weight-bold">
                  Position
                </Label>
                <Input disabled value={dyeObj.Position} />
              </Col>

              <Col className="">
                <Label for="tolerance" className="font-weight-bold">
                  Tolerance
                </Label>
                <Input
                  name="tolerance"
                  id="tolerance"
                  placeholder={"Type here"}
                  value={tolerance[index].value}
                  onChange={(event) =>
                    handleOnChange(
                      `tolerance.${index}.value`,
                      event.target.value
                    )
                  }
                  onBlur={(e) =>
                    handleBlur(`tolerance.${index}.isInvalid`, e.target.value)
                  }
                  onFocus={() =>
                    formik.setFieldValue(`tolerance.${index}.isInvalid`, false)
                  }
                />
                {tolerance[index].isInvalid && (
                  <div className="flex-auto">
                    <Text Tag="p" size={14} className="text-danger">
                      {`Tolerance must be between ${MIN_TOLERANCE_ALLOWED} and ${MAX_TOLERANCE_ALLOWED}`}
                    </Text>
                  </div>
                )}
              </Col>
            </Row>
          ))
        ) : (
          <Text className="secondary m-5 p-5 font-weight-bold">
            Nothing to show
          </Text>
        )}

        {toleranceData && toleranceData.length > 0 && (
          <Row>
            <Col>
              <Center className="text-center">
                <Button
                  onClick={handleSaveBtn}
                  disabled={isSaveToleranceBtnDisabled(tolerance)}
                  color="primary"
                >
                  Save Tolerance
                </Button>
              </Center>
            </Col>
          </Row>
        )}
      </CardBody>
    </Card>
  );
};

export default React.memo(ToleranceComponent);
