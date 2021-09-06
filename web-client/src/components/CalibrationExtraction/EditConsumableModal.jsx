import React from "react";
import PropTypes from "prop-types";
import {
  Button,
  FormGroup,
  Row,
  Col,
  Input,
  InputGroupWithAddonText,
  Label,
  Modal,
  ModalBody,
  CheckBox,
} from "core-components";
import { ButtonGroup, ButtonIcon, Text } from "shared-components";
import {
  getPrevValue,
  validateHoldTime,
  validateRampRate,
  validateTargetTemperature,
} from "./stepHelper";
import {
  MIN_RAMP_RATE,
  MAX_RAMP_RATE,
  MAX_TARGET_TEMPERATURE,
  MIN_TARGET_TEMPERATURE,
  CYCLE_STAGE,
  HOLD_STAGE,
  COL_SIZE_HOLD_STAGE,
  COL_SIZE_CYCLE_STAGE,
  MIN_HOLD_TIME,
} from "./stepConstants";

const EditConsumableModal = (props) => {
  const { isUpdate, handleCrossBtn } = props;

  return (
    <>
      <Modal centered size="lg">
        <ModalBody>
          <Text
            tag="h4"
            size={24}
            className="modal-title text-center text-truncate text-capitalize font-weight-bold"
          >
            {isUpdate ? "Update Details" : "Add New Details"}
          </Text>
          <ButtonIcon
            position="absolute"
            placement="right"
            top={24}
            right={32}
            size={32}
            name="cross"
            onClick={handleCrossBtn}
          />
          <Row form className="mb-3 pb-3">
            <Col sm={colSize}>
              <FormGroup>
                <Label for="id" className="font-weight-bold">
                  ID
                </Label>
                <Input
                  type="number"
                  name="id"
                  id="id"
                  placeholder={"Type here"}
                  value={id}
                  // onChange={() => onChangeHandler(e.target)}
                  // onBlur={}
                  // onFocus={}
                  // invalid={}
                />
                <Text
                  Tag="p"
                  size={12}
                  className={`${id.isInvalid && "text-danger"} px-2 mb-0`}
                >
                  Invalid ID
                </Text>
              </FormGroup>
            </Col>

            <Col sm={colSize}>
              <FormGroup>
                <Label for="target_temperature" className="font-weight-bold">
                  Target Temperature
                </Label>
                <InputGroupWithAddonText addonText="unit °C">
                  <Input
                    type="number"
                    name="targetTemperature"
                    id="target_temperature"
                    placeholder={`${MIN_TARGET_TEMPERATURE} - ${MAX_TARGET_TEMPERATURE}`}
                    value={targetTemperature}
                    // onChange={onChangeHandler}
                    // onBlur={onTargetTemperatureBlurHandler}
                    // onFocus={onTargetTemperatureFocusHandler}
                    // invalid={targetTemperatureError}
                  />
                </InputGroupWithAddonText>
                <Text
                  Tag="p"
                  size={12}
                  className={`${
                    targetTemperatureError && "text-danger"
                  } px-2 mb-0`}
                >
                  Enter value between {MIN_TARGET_TEMPERATURE} to{" "}
                  {MAX_TARGET_TEMPERATURE}
                </Text>
              </FormGroup>
            </Col>
            <Col sm={colSize}>
              <FormGroup>
                <Label for="hold_time" className="font-weight-bold">
                  Hold Time
                </Label>
                <InputGroupWithAddonText addonText="unit sec">
                  <Input
                    type="number"
                    name="holdTime"
                    id="hold_time"
                    placeholder="seconds"
                    value={holdTime}
                    // onBlur={onHoldTimeBlurHandler}
                    // onFocus={onHoldTimeFocusHandler}
                    // onChange={onChangeHandler}
                    // invalid={holdTimeError}
                  />
                </InputGroupWithAddonText>
                {/* {holdTimeError && (
                  <Text Tag="p" size={12} className="text-danger px-2 mb-0">
                    Invalid Hold time
                  </Text>
                )} */}
                {dataCapture && (
                  <Text
                    Tag="p"
                    size={12}
                    className={`${holdTimeError && "text-danger"} px-2 mb-0`}
                  >
                    Enter value above {MIN_HOLD_TIME}
                  </Text>
                )}
              </FormGroup>
            </Col>
            {/* If the stage type is hold don't show datacapture checkbox */}
            {stageType !== HOLD_STAGE && (
              <Col sm={colSize}>
                <FormGroup>
                  <Label for="data_capture" className="font-weight-bold">
                    Data Capture
                  </Label>
                  <CheckBox
                    name="dataCapture"
                    id="dataCapture"
                    onChange={(event) => handleDataCapture(event)}
                    className="mr-2 ml-3 py-2"
                    checked={dataCapture}
                  />
                </FormGroup>
              </Col>
            )}
          </Row>
          <ButtonGroup className="text-center p-0 m-0 pt-5">
            {isUpdateForm === false && (
              <Button
                onClick={addClickHandler}
                color="primary"
                disabled={isFormValid === false}
              >
                Add
              </Button>
            )}
            {isUpdateForm === true && (
              <Button
                onClick={saveClickHandler}
                color="primary"
                disabled={isFormValid === false}
              >
                Save
              </Button>
            )}
          </ButtonGroup>
        </ModalBody>
      </Modal>
    </>
  );
};

export default EditConsumableModal;
