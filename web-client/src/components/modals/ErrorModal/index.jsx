import React from 'react';
import { Button, ModalBody } from 'core-components';
import { Center, Text, ImageIcon } from 'shared-components';
import alertIcon from 'assets/images/alertIcon.svg';
import { StyledModal } from './StyledModal';

const ErrorModal = ({ isOpen, hideModal, message }) => (
	<StyledModal isOpen={isOpen} toggle={hideModal} centered size="sm">
		<ModalBody className="d-flex flex-column align-items-center justify-content-center">
			<Center>
				<ImageIcon src={alertIcon} alt="alert icon not available" className='mb-4' />
				<Text tag="span" className="mb-4 text-default">
					{message}
				</Text>

				<Button onClick={hideModal} color="primary">
          Ok
				</Button>
			</Center>
		</ModalBody>
	</StyledModal>
);

export default ErrorModal;
