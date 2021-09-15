import React, { useState } from 'react';

import PropTypes from 'prop-types';
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
} from 'core-components';
import { Center, Text, ButtonIcon, MlModal } from 'shared-components';
import { OperatorLoginForm } from './OperatorLoginForm';
import { MODAL_BTN, MODAL_MESSAGE } from "appConstants";

const OperatorLoginModal = (props) => {

	const { 
		operatorLoginModalOpen,
		toggleOperatorLoginModal,
		handleEmailChange,
		handlePasswordChange,
		handleLoginButtonClick,
		authData
	} = props;

	//const { confirmationText, isOpen, confirmationClickHandler } = props;
  const [showForgetModal, setForgetModalVisibility] = useState(false);
 
	const forgotHandler = (e) => {
    e.preventDefault();
    toggleForgotModal();
  };

  const toggleForgotModal = () => {
    setForgetModalVisibility(!showForgetModal);
  };

	// Operator Login Modal
	return (
		<>
		{/* Operator Login Modal */}
		<Modal 
			isOpen={operatorLoginModalOpen} 
			toggle={toggleOperatorLoginModal} 
			centered 
			size="md"
			>
			<ModalBody>
				{showForgetModal && (
					<MlModal
						isOpen={showForgetModal}
						textBody={MODAL_MESSAGE.forgotPasswordMsg}
						successBtn={MODAL_BTN.okay}
						handleSuccessBtn={toggleForgotModal}
						handleCrossBtn={toggleForgotModal}
					/>
				)}
				<Text Tag="h4" size={24} className="text-center text-primary mt-3 mb-4">
					Welcome
				</Text>
				<ButtonIcon
					position="absolute"
					placement="right"
					top={4}
					right={8}
					size={36}
					name="cross"
					onClick={toggleOperatorLoginModal}
					className="border-0"
				/>
				<Form>
					<OperatorLoginForm className="col-11 mx-auto">
						<Row>
							<Col>
								<FormGroup row className="has-border-left d-flex align-items-center" >
									<Label for='login' sm={3}>Login</Label>
									<Col sm={9}>
										<Input
											type='text'
											name='username'
											id='username'
											placeholder='Type here'
											onChange={handleEmailChange}
											value={authData.email.value}
											invalid={authData.email.invalid}
										/>
										<FormError>Incorrect username</FormError>
									</Col>
								</FormGroup>
							</Col>
						</Row>
					
						<Row>
							<Col>
								<FormGroup row className="has-border-left d-flex align-items-center">
									<Label for='password' sm={3}>Password</Label>
									<Col sm={9}>
										<Input
											type='password'
											name='password'
											id='password'
											placeholder='Type here'
											onChange={handlePasswordChange}
											value={authData.password.value}
											invalid={authData.password.invalid}
										/>
										<FormError>Incorrect password</FormError>
									</Col>
								</FormGroup>
							</Col>
						</Row>
					
						<Center className='my-3'>
							<Button 
								color='primary'
								onClick={handleLoginButtonClick}
							>
								Login
							</Button>
						</Center>
					
						<Center>
							<a href="!#" className="link" onClick={forgotHandler}>
								Forgot username or password?
							</a> 
						</Center>
					</OperatorLoginForm>
				</Form>
			</ModalBody>
		</Modal>
		</>
	);
};

OperatorLoginModal.propTypes = {
	confirmationText: PropTypes.string,
	isOpen: PropTypes.bool,
	confirmationClickHandler: PropTypes.func,
};

OperatorLoginModal.defaultProps = {
	confirmationText: 'Are you sure you want to Exit?',
	isOpen: false,
};

export default React.memo(OperatorLoginModal);
