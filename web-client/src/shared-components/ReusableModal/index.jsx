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
      <ModalBody className="p-0">
        {(textHead) &&
            <div className="d-flex justify-content-center align-items-center modal-heading">
                    <Text className="mb-0 title font-weight-bold">{textHead}</Text>
                
                    <ButtonIcon
                        position="absolute"
                        placement="right"
                        top={0}
                        right={16}
                        size={36}
                        name="cross"
                        onClick={() => {setModalOpen(!modalOpen)}}
                        className="border-0"
                    />
            </div>
        }
            <div className="d-flex justify-content-center align-items-center flex-column h-100 py-4">
                {(!textHead) &&
                    <ButtonIcon
                        position="absolute"
                        placement="right"
                        top={5}
                        right={16}
                        size={36}
                        name="cross"
                        onClick={() => {setModalOpen(!modalOpen)}}
                        className="border-0"
                    />
                }
                <Text Tag="h5" size="20" className="text-center font-weight-bold mt-3 mb-4">
                    <Text Tag="span" className="mb-1">{textBody}</Text>
                </Text>
            </div>

        <Center className="text-center p-0 m-0 pt-5">
          <ButtonGroup className="text-center mt-5">
            {secondaryBtn && (
              <Button
                onClick={() => {setModalOpen(!modalOpen)}}
                color="transparent"
                className="mr-4 border-primary"
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
  isOpen: false,
};

export default ReusableModal;
