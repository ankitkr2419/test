import React from 'react';
import PropTypes from 'prop-types';
import { Link } from 'react-router-dom';
import {
	Button,
	Form,
	FormGroup,
	FormError,
	Input,
	Label,
	Card,
	CardBody,
} from 'core-components';
import { Text, Center } from 'shared-components';

const LoginForm = (props) => {
	const {
		username, setUsername, password, setPassword, loginAsAdmin, isLoginError,
	} = props;

	const isFormDataPresent = username !== '' && password !== '';

	const onSubmit = (event) => {
		event.preventDefault();
		const data = new FormData(event.target);

		loginAsAdmin({
			username: data.get('username'),
			password: data.get('password'),
		});
	};

	return (
		<div className="flex-100">
			<Text tag="h6" className="text-center font-weight-bold text-default pt-5">
        Admin Login
			</Text>
			<Card default>
				<CardBody className="p-4">
					<Form onSubmit={onSubmit} className="pt-3 pb-2">
						<FormGroup>
							<Label for="username">Username</Label>
							<Input
								type="text"
								name="username"
								id="username"
								placeholder="Type here"
								value={username}
								onChange={(event) => {
									setUsername(event.target.value);
								}}
								invalid={isLoginError}
							/>
							<FormError>Incorrect username</FormError>
						</FormGroup>
						<FormGroup>
							<Label for="password">Password</Label>
							<Input
								type="password"
								name="password"
								id="password"
								placeholder="Type here"
								value={password}
								onChange={(event) => {
									setPassword(event.target.value);
								}}
								invalid={isLoginError}
							/>
							<FormError>Incorrect password</FormError>
						</FormGroup>
						<Center className="py-4 mb-3">
							<Button disabled={isFormDataPresent === false} color="primary" className="mx-auto">
                Login
							</Button>
						</Center>
						<Center>
							<Link to="/">Forgot username or password?</Link>
						</Center>
					</Form>
				</CardBody>
			</Card>
		</div>
	);
};

LoginForm.propTypes = {
	username: PropTypes.string.isRequired,
	setUsername: PropTypes.func.isRequired,
	password: PropTypes.string.isRequired,
	setPassword: PropTypes.func.isRequired,
	loginAsAdmin: PropTypes.func.isRequired,
	isLoginError: PropTypes.bool.isRequired,
};

export default LoginForm;
