import React, { useState } from "react";
import PropTypes from "prop-types";
import { Card, CardBody } from "core-components";
import { ButtonGroup, Center, Text } from "shared-components";
import LoginForm from "./LoginForm";
import OperatorLoginModalContainer from "containers/OperatorLoginModalContainer";
import {
  Button,
  Form,
  FormGroup,
  FormError,
  Input,
  Label,
  Row,
  Col,
} from "core-components";
import { LoginFormRtpcr } from "./LoginFormRtpcr";
import { toast } from "react-toastify";

const LoginComponent = (props) => {
  const { loginBtnHandler, forgotHandler } = props;

  const [username, setUsername] = useState("");
  const [password, setPassword] = useState("");

  const handleEmailChange = (e) => {
    setUsername(e.target.value);
  };

  const handlePasswordChange = (e) => {
    setPassword(e.target.value);
  };

  const handleLoginButtonClick = () => {
    if (username !== "" && password !== "") {
      loginBtnHandler({ username, password });
    } else {
      toast.error("Username or password can not be empty");
    }
  };

  return (
    <div className="login-content">
      <Text Tag="h1" size={40} className="text-center text-secondary mb-5">
        Compact 32-RTPCR
      </Text>
      <Card className="card-login">
        <CardBody className="d-flex flex-column justify-content-center p-4">
          <Text
            Tag="h4"
            size={24}
            className="text-center text-primary mt-3 mb-4"
          >
            Welcome
          </Text>
          <Form className="pt-2">
            <LoginFormRtpcr className="mx-auto">
              <FormGroup
                row
                className="has-border-left d-flex align-items-center"
              >
                <Label for="login" sm={3}>
                  Login
                </Label>
                <Col sm={9}>
                  <Input
                    type="text"
                    name="username"
                    id="username"
                    placeholder="Type here"
                    onChange={handleEmailChange}
                    value={username}
                    // invalid={authData.email.invalid}
                  />
                  <FormError>Incorrect username</FormError>
                </Col>
              </FormGroup>

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
                    onChange={handlePasswordChange}
                    value={password}
                    // invalid={authData.password.invalid}
                  />
                  <FormError>Incorrect password</FormError>
                </Col>
              </FormGroup>

              <Center className="my-3">
                <Button color="primary" onClick={handleLoginButtonClick}>
                  Login
                </Button>
              </Center>

              <Center>
                <a href="!#" className="link" onClick={forgotHandler}>
                  Forgot username or password?
                </a>
              </Center>
            </LoginFormRtpcr>
          </Form>
        </CardBody>
      </Card>
    </div>
  );
};

export default React.memo(LoginComponent);
