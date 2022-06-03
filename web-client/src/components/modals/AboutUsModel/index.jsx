import React, { useEffect, useState } from "react";

import PropTypes from "prop-types";
import { Modal, ModalBody } from "core-components";
import { Text, ButtonIcon } from "shared-components";
import { useDispatch, useSelector } from "react-redux";
import { commonDetailsInitiated } from "action-creators/calibrationActionCreators";

const AboutUsModel = (props) => {
  const { isOpen, toggleAboutUsModal } = props;

  const commonDetails = useSelector(
    (state) => state.commonDetailsReducer
  ).toJS();
  const {
    contact_number,
    machine_version,
    manufacturing_year,
    receiver_email,
    receiver_name,
    serial_number,
    software_version,
  } = commonDetails.details;

  return (
    <>
      <Modal isOpen={isOpen} toggle={toggleAboutUsModal} centered size="sm">
        <ModalBody className="p-0">
          <div className="d-flex justify-content-center align-items-center modal-heading mb-2">
            <Text className="mb-0 title font-weight-bold">ABOUT US</Text>
            <ButtonIcon
              position="absolute"
              placement="right"
              top={0}
              right={16}
              size={36}
              name="cross"
              onClick={toggleAboutUsModal}
              className="border-0"
            />
          </div>
          <div className="d-flex justify-content-center align-items-center">
            <Text className="font-weight-bold mr-1">Serial No - </Text>
            <Text>{serial_number}</Text>
            <Text className="font-weight-bold ml-2">Manufacturing Year - </Text>
            <Text>{manufacturing_year}</Text>
          </div>
          <div className="d-flex justify-content-center align-items-center">
            <Text className="font-weight-bold">Machine Version - </Text>
            <Text className="mr-4 ml-2">{machine_version}</Text>
            <Text className="font-weight-bold">Software Version - </Text>
            <Text className="mr-4 ml-2">{software_version}</Text>
          </div>
          <div className="d-flex justify-content-center align-items-center">
            <Text className="font-weight-bold">Input Power - </Text>
            <Text className="mr-4 ml-2">1.1 KW</Text>
            <Text className="font-weight-bold">Rated Voltage - </Text>
            <Text className="mr-4 ml-2">220-240 VAC</Text>
          </div>
          <div className="d-flex justify-content-center align-items-center">
            <Text className="font-weight-bold">Rated Current - </Text>
            <Text className="mr-4 ml-2">5.0 Amp AC</Text>
            <Text className="font-weight-bold">Rated Frequency - </Text>
            <Text className="mr-4 ml-2">50 Hz</Text>
          </div>
          <div className="d-flex justify-content-center align-items-center">
            <Text className="font-weight-bold">Contact Number - </Text>
            <Text className="ml-2">{contact_number}</Text>
          </div>
          <div className="d-flex justify-content-center align-items-center">
            <Text className="font-weight-bold">Email - </Text>
            <Text className="ml-2">{receiver_email}</Text>
          </div>
          <div className="d-flex justify-content-center align-items-center">
            <Text className="font-weight-bold ml-4">Address-</Text>
            <Text className="ml-4">
              Plot No.99-B, Lonavla Industrial Co-operative Estate Ltd.,
              Nangargaon, Lonavala, Maharashtra 410401, India.
            </Text>
          </div>
        </ModalBody>
      </Modal>
    </>
  );
};

AboutUsModel.propTypes = {
  confirmationText: PropTypes.string,
  isOpen: PropTypes.bool,
  // confirmationClickHandler: PropTypes.func,
};

AboutUsModel.defaultProps = {
  confirmationText: "Recipe Name",
  isOpen: false,
};

export default React.memo(AboutUsModel);

