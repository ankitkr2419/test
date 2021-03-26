import React from "react";
import PropTypes from "prop-types";
import { Button, Modal, ModalBody } from "core-components";
import { Center, Text, ButtonGroup, ButtonIcon } from "shared-components";

const MlModal = (props) => {
  const {
    textHead,
    textBody,
    isOpen,
    showCrossBtn,
    successBtn,
    failureBtn,
    handleCrossBtn,
    handleSuccessBtn,
  } = props;

  return (
    <Modal isOpen={isOpen} centered size="md">
      <ModalBody className="p-0">
        {textHead && (
          <div className="d-flex justify-content-center align-items-center modal-heading">
            <Text className="mb-0 title font-weight-bold">{textHead}</Text>

            {showCrossBtn && (
              <ButtonIcon
                position="absolute"
                placement="right"
                top={0}
                right={16}
                size={36}
                name="cross"
                onClick={handleCrossBtn}
                className="border-0"
              />
            )}
          </div>
        )}
        <div className="d-flex justify-content-center align-items-center flex-column h-100 py-4">
          {!textHead && showCrossBtn && (
            <ButtonIcon
              position="absolute"
              placement="right"
              top={5}
              right={16}
              size={36}
              name="cross"
              onClick={handleCrossBtn}
              className="border-0"
            />
          )}
          <Text
            Tag="h5"
            size="20"
            className="text-center font-weight-bold mt-3 mb-4"
          >
            <Text Tag="span" className="mb-1">
              {textBody}
            </Text>
          </Text>
        </div>

        <Center className="text-center p-0 m-0 pt-5">
          <ButtonGroup className="text-center mt-5">
            {failureBtn && (
              <Button
                onClick={handleCrossBtn}
                color="transparent"
                className="mr-4 border-primary"
              >
                {failureBtn}
              </Button>
            )}

            {successBtn && (
              <Button
                onClick={handleSuccessBtn}
                color="primary"
                className="mr-4"
              >
                {successBtn}
              </Button>
            )}
          </ButtonGroup>
        </Center>
      </ModalBody>
    </Modal>
  );
};

MlModal.propTypes = {
  textHead: PropTypes.string,
  textBody: PropTypes.string,
  isOpen: PropTypes.bool,
  clickHandler: PropTypes.func,
  showCrossBtn: PropTypes.bool,
  successBtn: PropTypes.string,
  failureBtn: PropTypes.string,
  toggleModal: PropTypes.func,
};

MlModal.defaultProps = {
  isOpen: false,
  showCrossBtn: true,
};

export default MlModal;
