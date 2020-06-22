import React, { useState } from 'react';
import PropTypes from 'prop-types';
import { Card, CardBody, Button } from 'core-components';
import { ButtonGroup } from 'shared-components';
import LoginForm from './LoginForm';

const LoginComponent = (props) => {
	const {
		isAdminFormVisible,
		setIsAdminFormVisibility,
		loginAsOperator,
		loginAsAdmin,
		isLoginError,
	} = props;

	const [username, setUsername] = useState('');
	const [password, setPassword] = useState('');

	return (
		<div className="login-content">
			<ButtonGroup>
				<Button
					color="primary"
					onClick={() => {
						setIsAdminFormVisibility(true);
					}}
					className="btn-secondary mr-4"
				>
          Admin
				</Button>
				<Button color="primary" className="btn-secondary mr-4">
          Supervisor
				</Button>
			</ButtonGroup>
			<Card className="card-login">
				<CardBody className="d-flex scroll-y">
					<div className="flex-100">
						<h1 className="card-title">Compact 32</h1>
						{isAdminFormVisible === false && (
							<Button
								onClick={loginAsOperator}
								color="primary"
								className="mx-auto"
							>
                Login as Operator
							</Button>
						)}
						{isAdminFormVisible === true && (
							<Button
								onClick={() => {
									setIsAdminFormVisibility(false);
								}}
								color="primary"
								className="btn-secondary mr-4"
							>
                Back
							</Button>
						)}
					</div>
					{isAdminFormVisible && (
						<LoginForm
							username={username}
							setUsername={setUsername}
							password={password}
							setPassword={setPassword}
							loginAsAdmin={loginAsAdmin}
							isLoginError={isLoginError}
						/>
					)}
				</CardBody>
			</Card>
		</div>
	);
};

LoginComponent.propTypes = {
	isAdminFormVisible: PropTypes.bool.isRequired,
	setIsAdminFormVisibility: PropTypes.func.isRequired,
	loginAsOperator: PropTypes.func.isRequired,
	loginAsAdmin: PropTypes.func.isRequired,
	isLoginError: PropTypes.bool.isRequired,
};

export default React.memo(LoginComponent);
