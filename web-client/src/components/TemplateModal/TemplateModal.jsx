import React, { useCallback, useEffect, useState } from "react";
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
import { MAX_LID_TEMP, MAX_VOLUME, MIN_LID_TEMP, MIN_VOLUME } from "appConstants";

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
    lid_temp,
    setLidTemperature,
    addClickHandler,
    isFormValid,
    resetFormValues,
    isTemplateEdited,
    resetModalState,
  } = props;

  const [volumeInvalid, setVolumeInvalid] = useState(false);
  const [lidTempInvalid, setLidTempInvalid] = useState(false);

  // disabled as we only need effect to be run while component is un-mounting
  // eslint-disable-next-line arrow-body-style
  useEffect(() => {
    return () => {
      resetFormValues();
    };
    // eslint-disable-next-line
  }, []);

  const onVolumeBlurHandler = useCallback(
    (volume) => {
      if (volume < MIN_VOLUME || volume > MAX_VOLUME) {
        setVolumeInvalid(true);
      }
    },
    [setVolumeInvalid]
  );

  const onLidTempBlurHandler = useCallback(
    (temp) => {
      if (temp < MIN_LID_TEMP || temp > MAX_LID_TEMP) {
        setLidTempInvalid(true);
      }
    },
    [setLidTempInvalid]
  );

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
                    placeholder="10-250 (µ units)"
                    value={volume}
                    onChange={(event) => {
                      setVolume(parseInt(event.target.value));
                    }}
                    onBlur={(event) =>
                      onVolumeBlurHandler(parseInt(event.target.value))
                    }
                    onFocus={() => setVolumeInvalid(false)}
                  />
                  {volumeInvalid && (
                    <div className="flex-70">
                      <Text Tag="p" size={14} className="text-danger">
                        Volume should be between {MIN_VOLUME} -{MAX_VOLUME}.
                      </Text>
                    </div>
                  )}
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
                    placeholder="30-100 (°C)"
                    value={lid_temp}
                    onChange={(event) => {
                      setLidTemperature(parseInt(event.target.value));
                    }}
                    onBlur={(event) =>
                      onLidTempBlurHandler(parseInt(event.target.value))
                    }
                    onFocus={() => setLidTempInvalid(false)}
                  />
                  {lidTempInvalid && (
                    <div className="flex-70">
                      <Text Tag="p" size={14} className="text-danger">
                        Lid temperature should be between {MIN_LID_TEMP} -{MAX_LID_TEMP} °C.
                      </Text>
                    </div>
                  )}
                </FormGroup>
              </Col>
            </Row>
            <Center className="text-center p-0 m-0 pt-5">
              <Button
                onClick={addClickHandler}
                color="primary"
                disabled={lidTempInvalid || volumeInvalid || !isFormValid}
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
