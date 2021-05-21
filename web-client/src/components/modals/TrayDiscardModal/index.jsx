import React, { useState } from "react";

import PropTypes from "prop-types";
import { Modal, ModalBody, Button, CheckBox } from "core-components";
import { Center, Text, ButtonIcon, ImageIcon } from "shared-components";
import alertIcon from "assets/images/alertIcon.svg";
import collectAndEmptyTrayImage from "assets/images/collect-and-empty-tray.jpg";
import doneThumbsUpImage from "assets/images/done-thumbs-up-image.svg";
import {TrayDiscardSection, DiscardTrayBox} from './Style'

const TrayDiscardModal = (props) => {
  const {
    trayDiscardModal,
    toggleTrayDiscardModal,
    deckName,
    handleSuccessBtn,
    nextModal,
  } = props;

  const [enableContent, setEnableContent] = useState(false);

  const toggleContents = () => {
    setEnableContent(!enableContent);
  };

  return (
    <>
      <DiscardTrayBox>
        <Modal isOpen={trayDiscardModal} centered size="md">
          <ModalBody className="p-0">
            <div className="d-flex justify-content-center align-items-center modal-heading">
              <Text className="mb-0 title font-weight-bold">{deckName}</Text>
              <ButtonIcon
                position="absolute"
                placement="right"
                top={0}
                right={16}
                size={36}
                name="cross"
                onClick={toggleTrayDiscardModal}
                className="border-0"
              />
            </div>

            <TrayDiscardSection className="gray-scale-box d-flex justify-content-center align-items-center">
              <Center className="mt-4">
                {nextModal ? (
                  <>
                    <ImageIcon
                      src={alertIcon}
                      alt="alert icon not available"
                      className="mb-4"
                    />
                    <Text Tag="h5" size={18} className="text-center mx-5 mb-0">
                      Tray will be discarded from Machine!
                    </Text>
                    <Text Tag="h5" size={18} className="text-center mx-5 mb-4">
                      You can collect and empty the tray.
                    </Text>
                    <ImageIcon
                      src={collectAndEmptyTrayImage}
                      alt=""
                      className="mb-4 mx-auto d-block"
                    />

                    <Button
                      className="mb-4 "
                      color="primary"
                      size="sm"
                      onClick={handleSuccessBtn}
                    >
                      Continue to Discard
                    </Button>
                  </>
                ) : (
                  <>
                    <div
                      className={
                        enableContent
                          ? "status-box my-5"
                          : "gray-scale-box status-box my-5"
                      }
                    >
                      <ImageIcon
                        className="mb-4"
                        src={doneThumbsUpImage}
                        alt=""
                      />
                      <CheckBox
                        id="done"
                        name="done"
                        label="Successfully emptied the discarded tray & inserted back."
                        className="mb-5"
                        onClick={toggleContents}
                      />
                      <Button
                        color="primary"
                        disabled={!enableContent}
                        onClick={handleSuccessBtn}
                      >
                        Yes
                      </Button>
                    </div>
                  </>
                )}
              </Center>
            </TrayDiscardSection>
          </ModalBody>
        </Modal>
      </DiscardTrayBox>
    </>
  );
};

TrayDiscardModal.propTypes = {
  //confirmationText: PropTypes.string,
  isOpen: PropTypes.bool,
  nextModal: PropTypes.bool,
  confirmationClickHandler: PropTypes.func,
};

TrayDiscardModal.defaultProps = {
  //confirmationText: 'Are you sure you want to Exit?',
  nextModal: true,
  isOpen: false,
};

export default React.memo(TrayDiscardModal);
