import React from 'react';
import PropTypes from 'prop-types';
import {
	Button,
	Form,
	FormGroup,
	FormError,
	Input,
	Label,
	Card,
	CardBody,
	Row,
	Col,
} from 'core-components';

const LoginForm = (props) => {
	const {
		username,
		setUsername,
		password,
		setPassword,
		loginAsAdmin,
		isLoginError,
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
		<Card default className='mb-5'>
			<CardBody className='px-5 py-4'>
				<Form onSubmit={onSubmit}>
					<Row>
						<Col>
							<FormGroup>
								<Label for='username'>Username</Label>
								<Input
									type='text'
									name='username'
									id='username'
									placeholder='Type here'
									value={username}
									onChange={(event) => {
										setUsername(event.target.value);
									}}
									invalid={isLoginError}
								/>
								<FormError>Incorrect username</FormError>
							</FormGroup>
						</Col>
						<Col>
							<FormGroup>
								<Label for='password'>Password</Label>
								<Input
									type='password'
									name='password'
									id='password'
									placeholder='Type here'
									value={password}
									onChange={(event) => {
										setPassword(event.target.value);
									}}
									invalid={isLoginError}
								/>
								<FormError>Incorrect password</FormError>
							</FormGroup>
						</Col>
					</Row>
					<div className='text-right pt-4 pb-1 mb-3'>
						<Button disabled={isFormDataPresent === false} color='primary'>
							Login
						</Button>
					</div>
				</Form>
			</CardBody>
		</Card>
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
