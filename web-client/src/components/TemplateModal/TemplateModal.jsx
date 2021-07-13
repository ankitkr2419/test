import React, { useEffect } from "react";
import PropTypes from "prop-types";
import {
  Button,
  Form,
  FormGroup,
  Row,
  Col,
  Input,
  Label,
  Modal,
  ModalBody,
} from "core-components";
import { Center, ButtonIcon, Text } from "shared-components";

const TemplateModal = (props) => {
  const {
    isCreateTemplateModalVisible,
    toggleTemplateModal,
    templateDescription,
    setTemplateDescription,
    templateName,
    setTemplateName,
    volume,
    setVolume,
    lidTemperature,
    setLidTemperature,
    addClickHandler,
    isFormValid,
    resetFormValues,
    isTemplateEdited,
    resetModalState,
  } = props;

  // disabled as we only need effect to be run while component is un-mounting
  // eslint-disable-next-line arrow-body-style
  useEffect(() => {
    return () => {
      resetFormValues();
    };
    // eslint-disable-next-line
  }, []);

  return (
    <>
      <Modal
        onClosed={() => resetModalState()}
        isOpen={isCreateTemplateModalVisible}
        toggle={toggleTemplateModal}
        centered
        size="lg"
      >
        <ModalBody>
          <Text
            tag="h4"
            className="modal-title text-center text-truncate font-weight-bold"
          >
            {isTemplateEdited ? "Edit Template" : "Create New Template"}
          </Text>
          <ButtonIcon
            position="absolute"
            placement="right"
            top={24}
            right={32}
            size={32}
            name="cross"
            onClick={toggleTemplateModal}
          />
          <Form>
            <Row form>
              <Col sm={6}>
                <FormGroup>
                  <Label for="template_name" className="font-weight-bold">
                    Template Name
                  </Label>
                  <Input
                    type="text"
                    name="template_name"
                    id="template_name"
                    placeholder="Type here"
                    maxLength={100}
                    value={templateName}
                    onChange={(event) => {
                      setTemplateName(event.target.value);
                    }}
                  />
                </FormGroup>
              </Col>
              <Col sm={6}>
                <FormGroup>
                  <Label
                    for="template_description"
                    className="font-weight-bold"
                  >
                    Description
                  </Label>
                  <Input
                    type="text"
                    name="template_description"
                    id="template_description"
                    placeholder="Type here"
                    value={templateDescription}
                    onChange={(event) => {
                      setTemplateDescription(event.target.value);
                    }}
                    maxLength={300}
                  />
                </FormGroup>
              </Col>
            </Row>

            <Row form className="mb-5 pb-5">
              <Col sm={6}>
                <FormGroup>
                  <Label for="volume_value" className="font-weight-bold">
                    Volume (µ units)
                  </Label>
                  <Input
                    type="number"
                    name="volume_name"
                    id="volume_name"
                    placeholder="Type here"
                    value={volume}
                    onChange={(event) => {
                      setVolume(parseInt(event.target.value));
                    }}
                  />
                </FormGroup>
              </Col>
              <Col sm={6}>
                <FormGroup>
                  <Label
                    for="lid_temperature_value"
                    className="font-weight-bold"
                  >
                    Lid Temperature (°C)
                  </Label>
                  <Input
                    type="number"
                    name="lid_temperature_name"
                    id="lid_temperature_name"
                    placeholder="Type here"
                    value={lidTemperature}
                    onChange={(event) => {
                      setLidTemperature(parseInt(event.target.value));
                    }}
                  />
                </FormGroup>
              </Col>
            </Row>
            <Center className="text-center p-0 m-0 pt-5">
              <Button
                onClick={addClickHandler}
                color="primary"
                disabled={isFormValid === false}
              >
                {isTemplateEdited ? "Save" : "Add"}
              </Button>
            </Center>
          </Form>
        </ModalBody>
      </Modal>
    </>
  );
};

TemplateModal.propTypes = {
  isCreateTemplateModalVisible: PropTypes.bool.isRequired,
  toggleTemplateModal: PropTypes.func.isRequired,
  templateDescription: PropTypes.string.isRequired,
  setTemplateDescription: PropTypes.func.isRequired,
  templateName: PropTypes.string.isRequired,
  setTemplateName: PropTypes.func.isRequired,
  addClickHandler: PropTypes.func.isRequired,
  isFormValid: PropTypes.bool.isRequired,
  resetModalState: PropTypes.func.isRequired,
  isTemplateEdited: PropTypes.bool.isRequired,
};

export default TemplateModal;
