import React, { useEffect, useCallback } from "react";
import { useHistory } from "react-router";

import { Card, CardBody } from "core-components";
import { Icon, Text } from "shared-components";
import CommonFieldsComponent from "components/CalibrationExtraction/CommonFieldsComponent";
import FormikFieldsEditor from "components/FormikFieldsEditor";
import LidPidTuning from "./LidPidTuning";

import { HeadingTitle } from "components/CalibrationExtraction/HeadingTitle";
import TECOperations from "./TECOperations";
import DyeCalibration from "./DyeCalibration";
import ToleranceComponent from "./ToleranceComponent";

const CalibrationComponent = (props) => {
  let {
    handleSaveDetailsBtn,
    formik,
    isAdmin,
    formikRtpcrVars,
    handleRtpcrConfigSubmitButton,
    formikTECVars,
    handleTECConfigSubmitButton,
    lidPidStatus,
    handleLidPidButton,
    handleResetTEC,
    handleAutoTuneTEC,
    dyeOptions,
    formikDyeCalibration,
    handleDyeCalibrationButton,
    dyeCalibrationStatus,
    handleSaveToleranceBtn,
    toleranceData,
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

          {/* Tolerance Component */}
          <ToleranceComponent
            toleranceData={toleranceData}
            handleSaveToleranceBtn={handleSaveToleranceBtn}
          />

          {/**Rtpcr vars */}
          <FormikFieldsEditor
            formTitle={"Rtpcr Configuration"}
            formik={formikRtpcrVars}
            submitButtonLabel={"Save"}
            submitButtonHandler={handleRtpcrConfigSubmitButton}
          />

          {/** TEC vars */}
          <FormikFieldsEditor
            formTitle={"TEC Configuration"}
            formik={formikTECVars}
            submitButtonLabel={"Save"}
            submitButtonHandler={handleTECConfigSubmitButton}
          />

          {/** Lid PID Tuning */}
          <LidPidTuning
            lidPidStatus={lidPidStatus}
            handleButtonClick={handleLidPidButton}
          />

          {/** TEC Operations */}
          <TECOperations
            handleResetTEC={handleResetTEC}
            handleAutoTuneTEC={handleAutoTuneTEC}
          />

          {/**Calibration of dyes */}
          <DyeCalibration
            dyeOptions={dyeOptions}
            formikDyeCalibration={formikDyeCalibration}
            dyeCalibrationStatus={dyeCalibrationStatus}
            handleDyeCalibrationButton={handleDyeCalibrationButton}
          />
        </CardBody>
      </Card>
    </div>
  );
};

export default React.memo(CalibrationComponent);
