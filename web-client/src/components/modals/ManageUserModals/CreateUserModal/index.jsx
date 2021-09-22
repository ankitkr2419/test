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

const CreateUserModal = (props) => {
  const { isOpen, toggleModal, roleOptions, handleCreateUser } = props;

  const [username, setUsername] = useState("");
  const [password, setPassword] = useState("");
  const [role, setRole] = useState(roleOptions[1] || null);

  const onCreateUser = () => {
    if (username !== "" && password !== "" && role?.value !== null) {
      handleCreateUser({ username, password, role: role.value });
    }
  };

  return (
    <Modal isOpen={isOpen} centered size="md">
      <ModalBody>
        <Text Tag="h4" size={24} className="text-center text-primary mt-3 mb-4">
          Create New User
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

            <Row>
              <Col>
                <FormGroup
                  row
                  className="has-border-left d-flex align-items-center"
                >
                  <Label for="password" sm={3}>
                    Password
                  </Label>
                  <Col sm={9}>
                    <Input
                      type="password"
                      name="password"
                      id="password"
                      placeholder="Type here"
                      onChange={(e) => setPassword(e.target.value)}
                      value={password}
                    />
                    <FormError>Incorrect password</FormError>
                  </Col>
                </FormGroup>
              </Col>
            </Row>

            <Row>
              <Col>
                <FormGroup
                  row
                  className="has-border-left d-flex align-items-center"
                >
                  <Label sm={3}>Role</Label>

                  <Col sm={9}>
                    <Select
                      placeholder="Select Role"
                      className=""
                      options={roleOptions}
                      value={role}
                      onChange={(value) => setRole(value)}
                    />
                  </Col>
                </FormGroup>
              </Col>
            </Row>

            <Center className="my-3">
              <Button color="primary" onClick={onCreateUser}>
                Create
              </Button>
            </Center>
          </OperatorLoginForm>
        </Form>
      </ModalBody>
    </Modal>
  );
};

CreateUserModal.propTypes = {
  isOpen: PropTypes.bool.isRequired,
  toggleModal: PropTypes.func.isRequired,
  roleOptions: PropTypes.array.isRequired,
  handleCreateUser: PropTypes.func.isRequired,
};

export default React.memo(CreateUserModal);
