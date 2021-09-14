import React, { useEffect, useState } from "react";
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

const UpdateUserModal = (props) => {
  const { isOpen, toggleModal, roleOptions, handleUpdateUser } = props;

  const [oldUsername, setOldUsername] = useState("");
  const [username, setUsername] = useState("");
  const [password, setPassword] = useState("");
  const [role, setRole] = useState(roleOptions[1] || null);

  useEffect(() => {
    setUsername(oldUsername);
  }, [oldUsername]);

  const onUpdateUser = () => {
    if (
      oldUsername !== "" &&
      username !== "" &&
      password !== "" &&
      role?.value !== null
    ) {
      handleUpdateUser({ oldUsername, username, password, role: role.value });
    }
  };

  return (
    <Modal isOpen={isOpen} centered size="md">
      <ModalBody>
        <Text Tag="h4" size={24} className="text-center text-primary mt-3 mb-4">
          Update User
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
                    Old Username
                  </Label>
                  <Col sm={9}>
                    <Input
                      type="text"
                      name="oldUsername"
                      id="oldUsername"
                      placeholder="Type here"
                      onChange={(e) => setOldUsername(e.target.value)}
                      value={oldUsername}
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
                  <Label for="login" sm={3}>
                    New Username
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
              <Button color="primary" onClick={onUpdateUser}>
                Update
              </Button>
            </Center>
          </OperatorLoginForm>
        </Form>
      </ModalBody>
    </Modal>
  );
};

UpdateUserModal.propTypes = {
  isOpen: PropTypes.bool.isRequired,
  toggleModal: PropTypes.func.isRequired,
  roleOptions: PropTypes.array.isRequired,
  handleUpdateUser: PropTypes.func.isRequired,
};

export default React.memo(UpdateUserModal);
