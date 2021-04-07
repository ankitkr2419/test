import React, { useState } from "react";
import { Icon, Text } from "shared-components";
import { Button, CheckBox, Modal, ModalBody } from "core-components";

const TipDiscardModal = (props) => {
  const { isOpen, handleSuccessBtn, deckName } = props;
  const [checked, setChecked] = useState(false);

  return (
    <Modal isOpen={isOpen} centered size="sm">
      <ModalBody className="p-0">
        <div className="d-flex justify-content-center align-items-center modal-heading mb-4">
          <Text className="mb-0 title font-weight-bold">{deckName}</Text>
        </div>
        <div className="d-flex justify-content-center align-items-center flex-column mb-3">
          <Text
            Tag="label"
            size="20"
            className="mb-4 title-heading font-weight-bold"
          >
            Homing Confirmation
          </Text>

          <div className="font-weight-light border border-light shadow-none bg-light large-btn mb-3">
            <div className="d-flex justify-content-center align-items-center flex-column">
              <Icon
                size={21}
                name="tip-pickup"
                className="text-primary mt-3 mb-3"
              />
              <CheckBox
                id="tip-discard"
                name="tip-discard"
                label="Tip Discard"
                className="mb-3"
                onClick={() => {
                  setChecked(!checked);
                }}
              />
            </div>
          </div>

          <Button
            color="primary"
            onClick={() => {
              handleSuccessBtn(checked);
            }}
          >
            Okay
          </Button>
        </div>
      </ModalBody>
    </Modal>
  );
};

export default TipDiscardModal;
