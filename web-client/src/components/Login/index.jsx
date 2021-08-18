import React, { useState } from 'react';
import PropTypes from 'prop-types';
import { Card, CardBody, Button } from 'core-components';
import { ButtonGroup, Text } from 'shared-components';
import LoginForm from './LoginForm';
import OperatorLoginModalContainer from "containers/OperatorLoginModalContainer";

const LoginComponent = (props) => {
	const {
		isAdminFormVisible,
		setIsAdminFormVisibility,
		operatorLoginModalOpen,
    toggleOperatorLoginModal,
		deckName,
		loginAsAdmin,
		isLoginError,
	} = props;

	const [username, setUsername] = useState('');
	const [password, setPassword] = useState('');

	return (
		<div className='login-content'>
			<ButtonGroup>
				<Button
					color='secondary'
					onClick={() => {
						setIsAdminFormVisibility(true);
					}}
					className='mr-4'
					active={isAdminFormVisible}
				>
					Admin
				</Button>
				<Button color='secondary' className='mr-4'>
					Supervisor
				</Button>
			</ButtonGroup>
			<Card className='card-login'>
				<CardBody className='d-flex flex-column px-5 py-4 scroll-y'>
					<div className='d-flex align-items-center mb-3 pt-3'>
						<h1 className='text-default font-weight-normal mb-0 px-3'>
							Compact 32
						</h1>
						{isAdminFormVisible && (
							<Text
								tag='h6'
								className='flex-50 text-default text-center font-weight-bold mb-0 ml-auto px-3'
							>
								Admin Login
							</Text>
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
					{operatorLoginModalOpen && (
            <OperatorLoginModalContainer
              operatorLoginModalOpen={operatorLoginModalOpen}
              toggleOperatorLoginModal={toggleOperatorLoginModal}
              deckName={deckName}
            />
          )}
					<div className='mt-auto mb-5 pb-4'>
						{isAdminFormVisible === false && (
							<Button onClick={toggleOperatorLoginModal} color='primary'>
								Login as Operator
							</Button>
						)}
						{isAdminFormVisible === true && (
							<Button
								onClick={() => {
									setIsAdminFormVisibility(false);
								}}
								color='secondary'
							>
								Back
							</Button>
						)}
					</div>
				</CardBody>
			</Card>
		</div>
	);
};

LoginComponent.propTypes = {
	isAdminFormVisible: PropTypes.bool.isRequired,
	setIsAdminFormVisibility: PropTypes.func.isRequired,
	loginAsAdmin: PropTypes.func.isRequired,
	isLoginError: PropTypes.bool.isRequired,
};

export default React.memo(LoginComponent);
