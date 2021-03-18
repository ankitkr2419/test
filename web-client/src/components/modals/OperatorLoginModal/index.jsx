import React from 'react';

import PropTypes from 'prop-types';
import styled from 'styled-components';
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
import { Center, Text, ButtonIcon } from 'shared-components';

//For Operator Login Form
const OperatorLoginForm = styled.div`
.has-border-left{
	.form-control{
		border-left:7px solid #F38220;
	}
}
.link{
	color:#3C70FF;
}
`;


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

	// const toggleModal = () => {};
	// Operator Login Modal
	// const [operatorLoginModal, setOperatorLoginModal] = useState(false);
	// const toggleOperatorLoginModal = () => setOperatorLoginModal(!operatorLoginModal);
	return (
		<>
		{/* Operator Login Modal */}
		{/* <Button color="primary" onClick={toggleOperatorLoginModal}>Operator Login</Button> */}
		<Modal isOpen={operatorLoginModalOpen} toggle={toggleOperatorLoginModal} centered size="sm">
			<ModalBody>
				<Text Tag="h4" size="24" className="text-center text-primary mb-4">
					Welcome
				</Text>
				<ButtonIcon
					position="absolute"
					placement="right"
					top={16}
					right={16}
					size={36}
					name="cross"
					onClick={toggleOperatorLoginModal}
					className="border-0"
				/>
				<Form>
					<OperatorLoginForm className="col-10 mx-auto">
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
							<a href="!#" className="link">
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

export default OperatorLoginModal;
