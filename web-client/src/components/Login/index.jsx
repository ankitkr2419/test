import React from 'react';
import styled from 'styled-components';
import { Card, CardBody } from "core-components";
import { Text } from "shared-components";
import LoginForm from './LoginForm';

const StyledLogin = styled.div``;

const Login = props => {
  return (
		<StyledLogin className="flex-100">
			<Text tag="h6" className="text-center font-weight-bold text-default pt-5">
				Admin Login
			</Text>
			<Card default>
				<CardBody className="p-4">
					<LoginForm />
				</CardBody>
			</Card>
		</StyledLogin>
	);
};

Login.propTypes = {};

export default Login;