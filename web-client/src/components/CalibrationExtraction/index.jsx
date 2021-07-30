import React from "react";
import { Card, CardBody } from "core-components";
import { ButtonIcon, Icon, MlModal, Text } from "shared-components";
import { MODAL_BTN, MODAL_MESSAGE } from "appConstants";
import PidComponent from "./PidComponent";
import MotorComponent from "./MotorComponent";

const CalibrationExtractionComponent = (props) => {
  const {
    deckName,
    progressData,
    handleBtnClick,
    handleLogout,
    handleMotorBtn,
    showConfirmationModal,
    toggleConfirmModal,
    formik,
  } = props;

  return (
    <div className="calibration-content px-5">
      <ButtonIcon
        name="logout"
        size={28}
        className="ml-auto bg-white border-primary"
        onClick={toggleConfirmModal}
      />

      <Card default className="my-3">
        <CardBody className="px-5 py-4">
          {/* <p>Extraction Flow: Calibration</p> */}
          <PidComponent
            progressData={progressData}
            handleBtnClick={handleBtnClick}
          />

          <MotorComponent formik={formik} handleMotorBtn={handleMotorBtn} />
        </CardBody>
      </Card>

      {showConfirmationModal && (
        <MlModal
          isOpen={showConfirmationModal}
          textHead={deckName}
          textBody={MODAL_MESSAGE.exitConfirmation}
          successBtn={MODAL_BTN.yes}
          failureBtn={MODAL_BTN.no}
          handleSuccessBtn={handleLogout}
          handleCrossBtn={toggleConfirmModal}
        />
      )}
    </div>
  );
};

export default React.memo(CalibrationExtractionComponent);
