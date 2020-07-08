import React from 'react';
import PropTypes from 'prop-types';
import { Button, Modal, ModalBody } from 'core-components';
import { Center, Text, ButtonGroup } from 'shared-components';

const ConfirmationModal = (props) => {
	const { confirmationText, isOpen, confirmationClickHandler } = props;

	const toggleModal = () => {};

	return (
		<Modal isOpen={isOpen} toggle={toggleModal} centered size="md">
			<ModalBody>
				<Text tag="h4" className="text-center text-truncate font-weight-bold">
					{confirmationText}
				</Text>

				<Center className="text-center p-0 m-0 pt-5">
					<ButtonGroup>
						<Button
							onClick={() => {
								confirmationClickHandler(true);
							}}
							color="primary"
							className="mr-4"
						>
              Yes
						</Button>
						<Button
							color="secondary"
							className="mr-4"
							onClick={() => {
								confirmationClickHandler(false);
							}}
						>
              No
						</Button>
					</ButtonGroup>
				</Center>
			</ModalBody>
		</Modal>
	);
};

ConfirmationModal.propTypes = {
	confirmationText: PropTypes.string,
	isOpen: PropTypes.bool,
	confirmationClickHandler: PropTypes.func,
};

ConfirmationModal.defaultProps = {
	confirmationText: 'Are you sure you want to Exit?',
	isOpen: false,
};

export default ConfirmationModal;
