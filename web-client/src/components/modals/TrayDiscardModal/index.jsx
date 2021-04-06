import React, { useState } from 'react';

import PropTypes from 'prop-types';
import styled from 'styled-components';
import { 
	Modal, 
	ModalBody, 
	Button,
	CheckBox
} from 'core-components';
import { Center, Text, ButtonIcon, ImageIcon } from 'shared-components';
import alertIcon from 'assets/images/alertIcon.svg';
import collectAndEmptyTrayImage from 'assets/images/collect-and-empty-tray.jpg';
import doneThumbsUpImage from 'assets/images/done-thumbs-up-image.svg';
//For Tray Discard Modal
// Need to toggle this class for gray scale effect
const TrayDiscardSection = styled.div`
.status-box{
	&.gray-scale-box{
	filter: grayscale(1);
	}
}
`;


const TrayDiscardModal = (props) => {
	//const { confirmationText, isOpen, confirmationClickHandler } = props;

	// const toggleModal = () => {};

	//Tray Discard Modal
	const [trayDiscardModal, setTrayDiscardModal] = useState(false);
	const toggleTrayDiscardModal = () => setTrayDiscardModal(!trayDiscardModal);
	return (
		<>
		{/* Tray Discard Modal */}
		<Button color="secondary" onClick={toggleTrayDiscardModal} className="ml-2 border-primary">Discard Tray</Button>
				<Modal isOpen={trayDiscardModal} toggle={toggleTrayDiscardModal} centered size="md">
				<ModalBody className="p-0">
				<div className="d-flex justify-content-center align-items-center modal-heading">
						<Text className="mb-0 title font-weight-bold">Deck B</Text>
						<ButtonIcon
							position="absolute"
							placement="right"
							top={0}
							right={16}
							size={36}
							name="cross"
							onClick={toggleTrayDiscardModal}
							className="border-0"
						/>
					</div>
					<TrayDiscardSection className="gray-scale-box d-flex justify-content-center align-items-center">
					<Center className="mt-4">
						<ImageIcon src={alertIcon} alt="alert icon not available" className='mb-4' />
						<Text Tag="h5" size={18} className="text-center mx-5 mb-0">
							Tray will be discarded from Machine!
						</Text>
						<Text Tag="h5" size={18} className="text-center mx-5 mb-4">
							You can collect and empty the tray.
						</Text>
						<ImageIcon src={collectAndEmptyTrayImage} alt="" className='mb-4 mx-auto d-block' />

						<Button color="primary" size="sm">Continue to Discard</Button>
						<div className="status-box my-5">
							<ImageIcon src={doneThumbsUpImage} alt="" className='mb-4' />
							<CheckBox
								id='done'
								name='done'
								label='Successfully emptied the discarded tray & inserted back.'
								className='mb-5'
							/>
							<Button color="primary">Yes</Button>
						</div>
					</Center>
					</TrayDiscardSection>
			</ModalBody>
		</Modal>
		</>
	);
};

TrayDiscardModal.propTypes = {
	//confirmationText: PropTypes.string,
	isOpen: PropTypes.bool,
	confirmationClickHandler: PropTypes.func,
};

TrayDiscardModal.defaultProps = {
	//confirmationText: 'Are you sure you want to Exit?',
	isOpen: false,
};

export default TrayDiscardModal;
