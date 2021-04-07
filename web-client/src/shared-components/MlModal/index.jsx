import React from "react";
import PropTypes from "prop-types";
import { Button, Modal, ModalBody } from "core-components";
import { Center, Text, ButtonGroup, ButtonIcon } from "shared-components";
import { Progress } from "reactstrap";
import styled from "styled-components";

const CustomProgressBar = styled.div`
  .custom-progress-bar {
    border-radius: 7px;
    background-color: #b2dad131;
    border: 1px solid #b2dad131;
    .progress-bar {
      height: 0.875rem;
      background-color: #92c4bc;
      border-radius: 7px 0px 0px 7px;
    }
  }
`;
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
    progressPercentage,
    isProgressBarVisible,
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
        <div className="d-flex justify-content-center align-items-center flex-column h-100 py-5">
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
          <Text Tag="h5" size={20} className="text-center my-3">
            <Text Tag="span" className="font-weight-normal" size={20}>
              {textBody}
            </Text>
          </Text>
        </div>

        {/* Conditional rendering for progress bar */}
        {isProgressBarVisible && (
          <CustomProgressBar className="mx-5">
            <Progress
              value={progressPercentage}
              className="custom-progress-bar"
            />
          </CustomProgressBar>
        )}

        <Center className="text-center p-0 m-0 pt-5">
          <ButtonGroup className="text-center my-5">
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
  progressPercentage: PropTypes.number,
  isProgressBarVisible: PropTypes.bool,
};

MlModal.defaultProps = {
  isOpen: false,
  showCrossBtn: true,
  isProgressBarVisible: false,
};

export default MlModal;
