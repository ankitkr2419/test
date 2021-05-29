import React, { useState } from "react";

import PropTypes from "prop-types";
import {
  Modal,
  ModalBody,
  Button,
  FormGroup,
  FormError,
  Label,
  Input,
} from "core-components";
import { Center, Text, ButtonIcon, Icon } from "shared-components";
import { Slider } from "antd";
import "antd/dist/antd.css";
import "./heightModal.scss";

const HeightModal = (props) => {
  const { isOpen, wellObj, handleSuccessBtn, handleCrossBtn } = props;

  const [InputValue, setInputValue] = useState(4);

  const onChange = (value) => setInputValue(value);

  const changeInputValue = (value, type) => {
    //type=0 means click and type=1 means edit
    let newValue = parseInt(value);
    if (type === 0) {
      newValue = InputValue + value;
    }
    if (newValue < 11 && newValue > 0) {
      setInputValue(newValue);
    }
  };

  return (
    <Modal isOpen={isOpen} centered className="set-height-modal">
      <ModalBody className="p-0">
        <div className="top-heading pt-3">
          <Text size="16" className="text-center font-weight-bold">
            Set Height for Well Number {wellObj.label}
          </Text>
          <ButtonIcon
            position="absolute"
            placement="right"
            top={16}
            right={16}
            size={36}
            name="cross"
            onClick={handleCrossBtn}
            className="border-0"
          />
        </div>
        <div className="d-flex justify-content-center align-items-center flex-column mt-4">
          <div className="slider-box d-flex align-items-center justify-content-center">
            <div className="slider-outter-box text-center bg-white">
              <div className="deck-level-box">
                {/* <div className="deck-level-inner"></div> */}
                <Slider
                  min={1}
                  max={10}
                  vertical
                  onChange={onChange}
                  value={typeof InputValue === "number" ? InputValue : 0}
                />
              </div>
            </div>
          </div>
          <FormGroup className="d-flex align-items-center justify-content-center mt-3 flex-column">
            <Label className="text-center w-100">
              <Text Tag="span" className="height-label">
                Enter Height
              </Text>
              <Text Tag="span" className="height-unit">
                {" "}
                (in cm)
              </Text>
            </Label>
            <Text className="d-flex justify-content-center align-items-center number mb-2">
              <Text
                Tag="span"
                className="minus"
                onClick={() => changeInputValue(-1, 0)}
              >
                <Icon name="minus-1" size="18" />
              </Text>
              <Input
                type="text"
                value={InputValue}
                onChange={(e) => changeInputValue(e.target.value, 1)}
              />
              <Text
                Tag="span"
                className="plus"
                onClick={() => changeInputValue(1, 0)}
              >
                <Icon name="plus-2" size="18" />
              </Text>
            </Text>
            {/* <InputNumber
                min={1}
                max={10}
                value={InputValue}
                onChange={onChange}
                placeholder="Enter Height"
              /> */}
            <Label for="height" className="text-center w-100 gray-text">
              * You can add values from x to y
            </Label>
            <FormError>Incorrect height</FormError>
          </FormGroup>
          <Center className="mb-4">
            <Button
              color="primary"
              className="mb-1"
              onClick={() => handleSuccessBtn(InputValue, wellObj.type)}
            >
              Done
            </Button>
          </Center>
        </div>
      </ModalBody>
    </Modal>
  );
};

HeightModal.propTypes = {
  confirmationText: PropTypes.string,
  isOpen: PropTypes.bool,
  confirmationClickHandler: PropTypes.func,
};

HeightModal.defaultProps = {
  confirmationText: "Are you sure you want to Exit?",
  isOpen: false,
};

export default HeightModal;
