import React, { useState } from 'react';
import { ButtonIcon, Text, Center, Icon } from 'shared-components';
import { Modal, ModalBody, Button, CheckBox } from 'core-components';

const ExportDataModal = (props) => {
	const [exportDataModal, setExportDataModal] = useState(false);
	const toggleExportDataModal = () => setExportDataModal(!exportDataModal);
	return (
		<>
			<ButtonIcon
				size={34}
				name='external-link'
				className='mx-2'
				onClick={toggleExportDataModal}
			/>
			<Modal
				isOpen={exportDataModal}
				toggle={toggleExportDataModal}
				centered
				size='lg'
			>
				<ModalBody>
					<Text
						Tag='h4'
						className='modal-title text-center text-truncate font-weight-bold d-flex justify-content-center align-items-center'
					>
						<Icon size={32} name='external-link' className='mr-2' />
						Export Data for samples
					</Text>
					<ButtonIcon
						position='absolute'
						placement='right'
						top={24}
						right={32}
						size={32}
						name='cross'
						onClick={toggleExportDataModal}
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
							Export
						</Button>
					</Center>
				</ModalBody>
			</Modal>
		</>
	);
};

ExportDataModal.propTypes = {};

export default ExportDataModal;
