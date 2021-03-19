import React, { useState } from 'react';

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

//For Enter Time Form
const EnterTimeForm = styled.div`
.row-small-gutter {
    margin-left: -10px !important;
    margin-right: -10px !important;
}

.row-small-gutter > * {
    padding-left: 10px !important;
    padding-right: 10px !important;
}
label{
	font-size:0.813rem;
	line-height:0.938rem;
}
`;


const TimeModal = (props) => {
	const { 
		// confirmationText, 
		// isOpen, 
		// confirmationClickHandler,
		buttonText,
	 } = props;

	// const toggleModal = () => {};
	// Operator Login Modal
	const [timeModal, setTimeModal] = useState(false);
	const toggleTimeModal = () => setTimeModal(!timeModal);
	return (
		<>
		{/* Operator Login Modal */}
		  <Button color="primary" onClick={toggleTimeModal}>{buttonText}</Button>
				<Modal isOpen={timeModal} toggle={toggleTimeModal} centered size="sm">
				<ModalBody>
					<div className="d-flex justify-content-center align-items-center flex-column h-100 py-4">
					<Text Tag="h5" size="20" className="text-center font-weight-bold mb-4">
					Enter Time Here
					</Text>
					<ButtonIcon
						position="absolute"
						placement="right"
						top={16}
						right={16}
						size={36}
						name="cross"
						onClick={toggleTimeModal}
						className="border-0"
					/>
					<Form>
					<EnterTimeForm className="col-11 mx-auto">
						<Row>
							<Col>
								<FormGroup row className="d-flex align-items-center justify-content-center row-small-gutter" >
									<Col sm={4}>
										<Input
											type='text'
											name='hours'
											id='hours'
											placeholder='Type here'
											value=""
										/>
										<Label for='hours' className="font-weight-bold">Hours</Label>
										<FormError>Incorrect Hours</FormError>
									</Col>
									<Col sm={4}>
										<Input
											type='text'
											name='minutes'
											id='minutes'
											placeholder='Type here'
											value=""
										/>
										<Label for='minutes' className="font-weight-bold">Minutes</Label>
										<FormError>Incorrect Minutes</FormError>
									</Col>
									<Col sm={4}>
										<Input
											type='text'
											name='seconds'
											id='seconds'
											placeholder='Type here'
											value=""
										/>
										<Label for='minutes' className="font-weight-bold">Seconds</Label>
										<FormError>Incorrect Seconds</FormError>
									</Col>
								</FormGroup>
							</Col>
						</Row>
						<Center className='my-3'>
							<Button color='primary'>
							 Next
							</Button>
						</Center>
					</EnterTimeForm>
					</Form>
					</div>
			</ModalBody>
		</Modal>
		</>
	);
};

TimeModal.propTypes = {
	confirmationText: PropTypes.string,
	isOpen: PropTypes.bool,
	confirmationClickHandler: PropTypes.func,
};

TimeModal.defaultProps = {
	confirmationText: 'Are you sure you want to Exit?',
	isOpen: false,
};

export default TimeModal;
