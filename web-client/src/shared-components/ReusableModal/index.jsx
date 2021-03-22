import React, { useState } from "react";
import PropTypes from "prop-types";
import { Button, Modal, ModalBody } from "core-components";
import { Center, Text, ButtonGroup, ButtonIcon } from "shared-components";
import { ModalHeader } from "reactstrap";

const ReusableModal = (props) => {
  const {
    textHead,
    textBody,
    isOpen,
    clickHandler,
    primaryBtn,
    secondaryBtn,
    toggleModal,
  } = props;

  const [modalOpen, setModalOpen] = useState(isOpen);

  return (
    <Modal isOpen={modalOpen} toggle={toggleModal} centered size="md">
        <ModalHeader>
            <Text tag="h4" className="text-center text-truncate font-weight-bold">
                {textHead}
            </Text>
        </ModalHeader>
      <ModalBody>
            <ButtonIcon
                position="absolute"
                placement="right"
                top={16}
                right={16}
                size={36}
                name="cross"
                onClick={() => {setModalOpen(!modalOpen)}}
                className="border-0"
            />
            <Text tag="h4" className="text-center text-truncate font-weight-bold">
            {textHead}
            </Text>

        <Center className="text-center p-0 m-0 pt-5">
          <ButtonGroup className="text-center mt-5">
            {secondaryBtn && (
              <Button
                onClick={() => {setModalOpen(!modalOpen)}}
                color="transparent"
                className="mr-4"
              >
                {secondaryBtn}
              </Button>
            )}

            {primaryBtn && (
              <Button
                onClick={() => {
                  clickHandler(true);
                }}
                color="primary"
                className="mr-4"
              >
                {primaryBtn}
              </Button>
            )}
          </ButtonGroup>
        </Center>
      </ModalBody>
    </Modal>
  );
};

ReusableModal.propTypes = {
  textBody: PropTypes.string,
  isOpen: PropTypes.bool,
  clickHandler: PropTypes.func,
};

ReusableModal.defaultProps = {
  textBody: "Are you sure you want to Exit?",
  isOpen: false,
};

export default ReusableModal;
