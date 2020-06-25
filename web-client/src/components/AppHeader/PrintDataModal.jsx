import React, { useState } from 'react';
import { ButtonIcon, Text, Center, Icon } from 'shared-components';
import { Modal, ModalBody, Button, CheckBox } from 'core-components';

const PrintDataModal = (props) => {
	const [printDataModal, setPrintDataModal] = useState(false);
	const togglePrintDataModal = () => setPrintDataModal(!printDataModal);
	return (
		<>
			<ButtonIcon
				size={34}
				name='print'
				className='mx-2'
				onClick={togglePrintDataModal}
			/>
			<Modal
				isOpen={printDataModal}
				toggle={togglePrintDataModal}
				centered
				size='lg'
			>
				<ModalBody>
					<Text
						Tag='h4'
						className='modal-title text-center text-truncate font-weight-bold d-flex justify-content-center align-items-center'
					>
						<Icon size={32} name='print' className='mr-2' />
						Print Data for samples
					</Text>
					<ButtonIcon
						position='absolute'
						placement='right'
						top={24}
						right={32}
						size={32}
						name='cross'
						onClick={togglePrintDataModal}
					/>
					<div className='mb-5 pb-3'>
						<CheckBox
							id='sample1'
							name='sample1'
							label='ID-xx-xxx'
							className='mb-3'
						/>
						<CheckBox
							id='sample2'
							name='sample2'
							label='ID-xx-xxx'
							className='mb-3'
						/>
						<CheckBox
							id='sample3'
							name='sample3'
							label='ID-xx-xxx'
							className='mb-3'
						/>
						<CheckBox
							id='sample4'
							name='sample4'
							label='ID-xx-xxx'
							className='mb-3'
						/>
						<CheckBox
							id='sample5'
							name='sample5'
							label='ID-xx-xxx'
							className='mb-3'
						/>
					</div>
					<Center>
						<Button color='primary' className='mb-2'>
							Print
						</Button>
					</Center>
				</ModalBody>
			</Modal>
		</>
	);
};

PrintDataModal.propTypes = {};

export default PrintDataModal;
