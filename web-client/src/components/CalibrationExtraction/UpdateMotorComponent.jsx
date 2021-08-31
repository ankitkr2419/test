import React from "react";
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
  checkIsFieldInvalid,
  isMotorUpdateBtnDisabled,
  updateMotorInitialFormikState,
} from "./helpers";

const UpdateMotorComponent = (props) => {
  const { handleUpdateMotorDetailsBtn } = props;

  const formik = useFormik({
    initialValues: updateMotorInitialFormikState,
    enableReinitialize: true,
  });

  const handleBlur = ({ name, value }) => {
    const isInvalid = checkIsFieldInvalid(formik.values[name], value);
    formik.setFieldValue(`${name}.isInvalid`, isInvalid);
  };

  const handleOnChange = (key, value) => {
    formik.setFieldValue(key, value);
  };

  return (
    <Card default className="my-3">
      <CardBody>
        <Row>
          {Object.values(formik.values).map((e) => (
            <Col className="mb-4" key={e.id} md={3}>
              <Label for="id" className="font-weight-bold">
                {e.label}
              </Label>
              <Input
                name={e.name}
                id={e.id}
                placeholder="Type here"
                value={e.value}
                onChange={(event) =>
                  handleOnChange(`${e.name}.value`, event.target.value)
                }
                onBlur={(event) => handleBlur(event.target)}
                onFocus={() => handleOnChange(`${e.name}.isInvalid`, false)}
              />
              {e.isInvalid && (
                <div className="flex-70">
                  <Text Tag="p" size={14} className="text-danger">
                    {e.label} is invalid.
                  </Text>
                </div>
              )}
            </Col>
          ))}
        </Row>

        <Row>
          <Col>
            <Center className="text-center pt-3">
              <Button
                className="w-auto"
                disabled={isMotorUpdateBtnDisabled(formik.values)}
                color="primary"
                onClick={() => handleUpdateMotorDetailsBtn(formik.values)}
              >
                Update Motor Details
              </Button>
            </Center>
          </Col>
        </Row>
      </CardBody>
    </Card>
  );
};

export default React.memo(UpdateMotorComponent);
