import React, { useEffect, useCallback } from "react";
import { useHistory } from "react-router";

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
import { Icon, Text } from "shared-components";

import {
  isValueValid,
  formikInitialState,
  validateAllFields,
  getRequestBody,
} from "./helper";

import { HeadingTitle } from "components/CalibrationExtraction/HeadingTitle";
import CommonFieldsComponent from "components/CalibrationExtraction/CommonFieldsComponent";
import FormikFieldsEditor from "components/FormikFieldsEditor";

const CalibrationComponent = (props) => {
  let {
    handleSaveDetailsBtn,
    formik,
    isAdmin,
    formikRtpcrVars,
    handleRtpcrConfigSubmitButton,
  } = props;

  const history = useHistory();

  const handleBack = () => {
    history.goBack();
  };

  return (
    <div className="calibration-content px-5">
      <div className="d-flex align-items-center">
        {isAdmin && (
          <div style={{ cursor: "pointer" }} onClick={handleBack}>
            <Icon name="angle-left" size={32} className="text-white" />
            <HeadingTitle
              Tag="h5"
              className="text-white font-weight-bold ml-3 mb-0"
            >
              Go back to template page
            </HeadingTitle>
          </div>
        )}
      </div>
      <Card default className="my-5">
        <CardBody className="px-5 py-4">
          {/* Common fields - name, email, room temperature */}
          <CommonFieldsComponent
            formik={formik}
            handleSaveDetailsBtn={handleSaveDetailsBtn}
          />

          {/**Rtpcr vars */}
          <FormikFieldsEditor
            formik={formikRtpcrVars}
            submitButtonLabel={"Save"}
            submitButtonHandler={handleRtpcrConfigSubmitButton}
          />
        </CardBody>
      </Card>
    </div>
  );
};

export default React.memo(CalibrationComponent);
