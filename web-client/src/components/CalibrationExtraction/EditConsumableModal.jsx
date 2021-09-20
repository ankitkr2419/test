import React from "react";
import PropTypes from "prop-types";
import {
  Button,
  FormGroup,
  Row,
  Col,
  Input,
  Label,
  Modal,
  ModalBody,
} from "core-components";
import { ButtonGroup, ButtonIcon, Text } from "shared-components";
import {
  checkConsumableFieldIsInvalid,
  isConsumableModalBtnDisabled,
} from "./helpers";

const EditConsumableModal = (props) => {
  const { isUpdate, isOpen, handleModalBtn, handleCrossBtn, formik } = props;

  const { id, name, description, distance } = formik.values;

  const handleOnChange = ({ name, value }) => {
    formik.setFieldValue(`${name}.value`, value);
  };

  const handleOnBlur = ({ name, value }) => {
    const isInvalid = checkConsumableFieldIsInvalid(name, value);
    formik.setFieldValue(`${name}.isInvalid`, isInvalid);
  };

  const handleOnFocus = ({ name }) => {
    formik.setFieldValue(`${name}.isInvalid`, false);
  };

  return (
    <Modal isOpen={isOpen} centered size="lg">
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
          <Col sm>
            <FormGroup>
              <Label for="id" className="font-weight-bold">
                ID
              </Label>
              <Input
                type="number"
                name="id"
                id="id"
                placeholder={"Type here"}
                value={id.value}
                onChange={(e) => handleOnChange(e.target)}
                onBlur={(e) => handleOnBlur(e.target)}
                onFocus={(e) => handleOnFocus(e.target)}
                disabled={isUpdate === true}
              />
              {id.isInvalid && (
                <Text Tag="p" size={12} className="text-danger px-2 mb-0">
                  Enter valid id
                </Text>
              )}
            </FormGroup>
          </Col>

          <Col sm>
            <FormGroup>
              <Label for="name" className="font-weight-bold">
                Name
              </Label>
              <Input
                type="text"
                name="name"
                id="name"
                placeholder={`Type Here`}
                value={name.value}
                onChange={(e) => handleOnChange(e.target)}
                onBlur={(e) => handleOnBlur(e.target)}
                onFocus={(e) => handleOnFocus(e.target)}
                disabled={isUpdate === true}
              />

              {name.isInvalid && (
                <Text Tag="p" size={12} className="text-danger px-2 mb-0">
                  Enter valid name
                </Text>
              )}
            </FormGroup>
          </Col>

          <Col sm>
            <FormGroup>
              <Label for="description" className="font-weight-bold">
                Description
              </Label>
              <Input
                type="text"
                name="description"
                id="description"
                placeholder="Type here"
                value={description.value}
                onChange={(e) => handleOnChange(e.target)}
                onBlur={(e) => handleOnBlur(e.target)}
                onFocus={(e) => handleOnFocus(e.target)}
              />
              {description.isInvalid && (
                <Text Tag="p" size={12} className="text-danger px-2 mb-0">
                  Enter valid description
                </Text>
              )}
            </FormGroup>
          </Col>

          <Col sm>
            <FormGroup>
              <Label for="distance" className="font-weight-bold">
                Distance
              </Label>
              <Input
                type="number"
                name="distance"
                id="distance"
                placeholder="Type here"
                value={distance.value}
                onChange={(e) => handleOnChange(e.target)}
                onBlur={(e) => handleOnBlur(e.target)}
                onFocus={(e) => handleOnFocus(e.target)}
              />
              {distance.isInvalid && (
                <Text Tag="p" size={12} className="text-danger px-2 mb-0">
                  Enter valid distance
                </Text>
              )}
            </FormGroup>
          </Col>
        </Row>
        <ButtonGroup className="text-center p-0 m-0 pt-5">
          {isUpdate === false && (
            <Button
              onClick={() => handleModalBtn(formik.values)}
              color="primary"
              disabled={isConsumableModalBtnDisabled(formik.values)}
            >
              Add
            </Button>
          )}
          {isUpdate === true && (
            <Button
              onClick={() => handleModalBtn(formik.values)}
              color="primary"
              disabled={isConsumableModalBtnDisabled(formik.values)}
            >
              Update
            </Button>
          )}
        </ButtonGroup>
      </ModalBody>
    </Modal>
  );
};

export default EditConsumableModal;
