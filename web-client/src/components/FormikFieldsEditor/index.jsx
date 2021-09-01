import React, { useEffect, useCallback } from "react";
import PropTypes from "prop-types";

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
import { Center, Text } from "shared-components";
import { isValueValid, validateAllFields, getRequestBody } from "./helper";

const FormikFieldsEditor = (props) => {
  let { formik, submitButtonLabel, submitButtonHandler } = props;

  const handleBlurChange = useCallback((name, value) => {
    const isValid = isValueValid(formik.values, name, value);
    formik.setFieldValue(`${name}.isInvalid`, !isValid);
  }, []);

  const handleOnChange = (event, name) => {
    formik.setFieldValue(`${name}.value`, event.target.value);
  };

  const handleOnFocus = (name) => {
    formik.setFieldValue(`${name}.isInvalid`, false);
  };

  //Validations and api call
  const onSubmit = (e) => {
    e.preventDefault();

    if (validateAllFields(formik.values) === true) {
      const requestBody = getRequestBody(formik.values);
      submitButtonHandler(requestBody);
    }
  };

  return (
    <Card default className="my-3">
      <CardBody>
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
                <Col md={4} key={name}>
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
                      onChange={(event) => handleOnChange(event, name)}
                      onBlur={(event) =>
                        handleBlurChange(name, event.target.value)
                      }
                      onFocus={() => handleOnFocus(name)}
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
          <Center className="text-center pt-3">
            <Button
              color="primary"
              disabled={!validateAllFields(formik.values)}
            >
              {submitButtonLabel || ""}
            </Button>
          </Center>
        </Form>
      </CardBody>
    </Card>
  );
};

FormikFieldsEditor.propTypes = {
  formik: PropTypes.object.isRequired,
  submitButtonLabel: PropTypes.string,
  submitButtonHandler: PropTypes.func.isRequired,
};

FormikFieldsEditor.defaultProps = {
  submitButtonLabel: "Save",
};

export default React.memo(FormikFieldsEditor);
