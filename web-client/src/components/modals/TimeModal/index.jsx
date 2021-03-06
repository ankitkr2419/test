import React from "react";

import PropTypes from "prop-types";
import {
  Modal,
  ModalBody,
  Button,
  Form,
  FormGroup,
  FormError,
  Input,
  Label,
  Row,
  Col,
} from "core-components";
import { Center, Text, ButtonIcon } from "shared-components";

import { EnterTimeForm } from './EnterTimeForm';

const TimeModal = (props) => {
  const {
    timeModal,
    toggleTimeModal,
    handleChangeTime,
    submitTime,
    deckName,
  } = props;

  return (
    <>
      {/* Operator Login Modal */}
      <Modal isOpen={timeModal} toggle={toggleTimeModal} centered size="sm">
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
              onClick={toggleTimeModal}
              className="border-0"
            />
          </div>
          <div className="d-flex justify-content-center align-items-center flex-column h-100 py-4">
            <Text
              Tag="h5"
              size={20}
              className="text-center font-weight-bold mt-3 mb-4"
            >
              <Text Tag="span" className="mb-1">
                Enter Time Here
              </Text>
            </Text>

            <Form>
              <EnterTimeForm className="col-11 mx-auto">
                <Row>
                  <Col md={7} className="mx-auto">
                    <FormGroup
                      row
                      className="d-flex align-items-center justify-content-center row-small-gutter"
                    >
                      <Col sm={4}>
                        <Input
                          type="number"
                          name="hours"
                          id="hours"
                          placeholder=""
                          min={0}
                          onChange={handleChangeTime}
                        />
                        <Label for="hours" className="font-weight-bold">
                          Hours
                        </Label>
                        <FormError>Incorrect Hours</FormError>
                      </Col>
                      <Col sm={4}>
                        <Input
                          type="number"
                          name="minutes"
                          id="minutes"
                          placeholder=""
                          min={0}
                          onChange={handleChangeTime}
                        />
                        <Label for="minutes" className="font-weight-bold px-2">
                          Minutes
                        </Label>
                        <FormError>Incorrect Minutes</FormError>
                      </Col>
                      <Col sm={4}>
                        <Input
                          type="number"
                          name="seconds"
                          id="seconds"
                          placeholder=""
                          min={0}
                          onChange={handleChangeTime}
                        />
                        <Label for="seconds" className="font-weight-bold">
                          Seconds
                        </Label>
                        <FormError>Incorrect Seconds</FormError>
                      </Col>
                    </FormGroup>
                  </Col>
                </Row>
                <Center className="mt-3 mb-4">
                  <Button color="primary" onClick={submitTime}>
                    Next
                  </Button>
                </Center>
              </EnterTimeForm>
            </Form>
          </div>
        </ModalBody>
      </Modal>
    </>
  );
};

TimeModal.propTypes = {
  confirmationText: PropTypes.string,
  isOpen: PropTypes.bool,
  confirmationClickHandler: PropTypes.func,
};

TimeModal.defaultProps = {
  confirmationText: "Are you sure you want to Exit?",
  isOpen: false,
};

export default React.memo(TimeModal);
