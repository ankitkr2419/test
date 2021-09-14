import React, { useState } from "react";
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
  Select,
} from "core-components";
import { Center, Text, ButtonIcon } from "shared-components";

import { OperatorLoginForm } from "../../OperatorLoginModal/OperatorLoginForm";

const DeleteUserModal = (props) => {
  const { isOpen, toggleModal, handleDeleteUser } = props;

  const [username, setUsername] = useState("");

  const onDeleteUser = () => {
    if (username !== "") {
      handleDeleteUser(username);
    }
  };

  return (
    <Modal isOpen={isOpen} centered size="md">
      <ModalBody>
        <Text Tag="h4" size={24} className="text-center text-primary mt-3 mb-4">
          Delete User
        </Text>
        <ButtonIcon
          position="absolute"
          placement="right"
          top={4}
          right={8}
          size={36}
          name="cross"
          onClick={toggleModal}
          className="border-0"
        />

        <Form>
          <OperatorLoginForm className="col-11 mx-auto">
            <Row>
              <Col>
                <FormGroup
                  row
                  className="has-border-left d-flex align-items-center"
                >
                  <Label for="login" sm={3}>
                    Username
                  </Label>
                  <Col sm={9}>
                    <Input
                      type="text"
                      name="username"
                      id="username"
                      placeholder="Type here"
                      onChange={(e) => setUsername(e.target.value)}
                      value={username}
                    />
                    <FormError>Incorrect username</FormError>
                  </Col>
                </FormGroup>
              </Col>
            </Row>

            <Center className="my-3">
              <Button color="primary" onClick={onDeleteUser}>
                Delete
              </Button>
            </Center>
          </OperatorLoginForm>
        </Form>
      </ModalBody>
    </Modal>
  );
};

DeleteUserModal.propTypes = {
  isOpen: PropTypes.bool.isRequired,
  toggleModal: PropTypes.func.isRequired,
  handleDeleteUser: PropTypes.func.isRequired,
};

export default React.memo(DeleteUserModal);
