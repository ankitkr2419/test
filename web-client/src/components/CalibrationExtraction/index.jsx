import React from "react";
import {
  Button,
  Form,
  FormGroup,
  FormError,
  Input,
  Label,
  Card,
  CardBody,
  Row,
  Col,
} from "core-components";
import { ButtonIcon, MlModal } from "shared-components";
import { MODAL_BTN, MODAL_MESSAGE } from "appConstants";
const CalibrationExtractionComponent = (props) => {
  const { deckName, handleLogout, showConfirmationModal, toggleConfirmModal } =
    props;
  return (
    <div className="calibration-content px-5">
      <ButtonIcon
        name="logout"
        size={28}
        className="ml-2 bg-white border-primary"
        onClick={toggleConfirmModal}
      />

      <Card default className="my-5">
        <CardBody className="px-5 py-4">
          <p>Extraction Flow: Calibration</p>
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
