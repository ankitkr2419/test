import React from "react";
import { useHistory } from "react-router";

import { Card, CardBody } from "core-components";
import { ButtonIcon, Icon, MlModal, Text } from "shared-components";
import { MODAL_BTN, MODAL_MESSAGE, ROUTES } from "appConstants";
import PidComponent from "./PidComponent";
import MotorComponent from "./MotorComponent";

import { HeadingTitle } from "./HeadingTitle";

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
    isAdmin,
  } = props;

  const history = useHistory();

  const handleBack = () => {
    history.push(ROUTES.recipeListing);
  };

  return (
    <div className="calibration-content px-5 pt-3">
      <div className="d-flex align-items-center pb-4">
        {isAdmin && (
          <div style={{ cursor: "pointer" }} onClick={handleBack}>
            <Icon name="angle-left" size={32} className="text-white" />
            <HeadingTitle
              Tag="h5"
              className="text-white font-weight-bold ml-3 mb-0"
            >
              Go back to recipe listing
            </HeadingTitle>
          </div>
        )}

        <ButtonIcon
          name="logout"
          size={28}
          className="ml-auto bg-white border-primary"
          onClick={toggleConfirmModal}
        />
      </div>

      <Card default className="my-3">
        <CardBody className="px-5 py-4">
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
