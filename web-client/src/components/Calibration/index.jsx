import React, { useEffect, useCallback } from "react";
import { useFormik } from "formik";

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

import {
  isValueValid,
  formikInitialState,
  validateAllFields,
  getRequestBody,
} from "./helper";

const CalibrationComponent = (props) => {
  let { configs, saveButtonClickHandler } = props;

  const formik = useFormik({
    initialValues: formikInitialState,
    enableReinitialize: true,
  });

  //store new data in local state
  useEffect(() => {
    if (configs) {
      Object.keys(formik.values).map((element) => {
        const { apiKey, name } = formik.values[element];

        const newValue = configs[apiKey] ? configs[apiKey] : "";
        const isValid = isValueValid(name, newValue);

        // set formik fields
        formik.setFieldValue(`${name}.isInvalid`, !isValid);
        formik.setFieldValue(`${name}.value`, newValue);
      });
    }
  }, [configs]);

  //validations and api call
  const onSubmit = (e) => {
    e.preventDefault();

    if (validateAllFields(formik.values) === true) {
      const requestBody = getRequestBody(formik.values);
      saveButtonClickHandler(requestBody);
    }
  };

  const handleBlurChange = useCallback((name, value) => {
    const isValid = isValueValid(name, value);
    formik.setFieldValue(`${name}.isInvalid`, !isValid);
  }, []);

  return (
    <div className="calibration-content px-5">
      <Card default className="my-5">
        <CardBody className="px-5 py-4">
          <Form onSubmit={onSubmit}>
            <Row>
              {Object.keys(formik.values).map((key, index) => {
                const element = formik.values[key];
                const {
                  type,
                  name,
                  label,
                  min,
                  max,
                  value,
                  isInvalid,
                  isInvalidMsg,
                } = element;

                return (
                  <Col md={6}>
                    <FormGroup>
                      <Label for="username">{label}</Label>
                      <Input
                        type={type}
                        name={name}
                        id={name}
                        placeholder={
                          type === "number" ? `${min} - ${max}` : "Type here"
                        }
                        value={value}
                        onChange={(event) => {
                          formik.setFieldValue(
                            `${name}.value`,
                            event.target.value
                          );
                        }}
                        onBlur={(event) =>
                          handleBlurChange(name, event.target.value)
                        }
                        onFocus={() =>
                          formik.setFieldValue(`${name}.isInvalid`, false)
                        }
                      />
                      {(isInvalid || value == null) && (
                        <div className="flex-70">
                          <Text Tag="p" size={14} className="text-danger">
                            {`${isInvalidMsg}`}
                          </Text>
                        </div>
                      )}
                    </FormGroup>
                  </Col>
                );
              })}
            </Row>
            <div className="text-right pt-4 pb-1 mb-3">
              <Button
                color="primary"
                disabled={!validateAllFields(formik.values)}
              >
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
