import React, { useEffect } from 'react';
import PropTypes from 'prop-types';
import {
	Button,
	Form,
	FormGroup,
	Row,
	Col,
	Input,
	Label,
	Modal,
	ModalBody,
	Select,
} from 'core-components';
import { Center, ButtonIcon, Text } from 'shared-components';
import { stageTypeOptions } from './stageConstants';

const AddStageModal = (props) => {
	const {
		toggleCreateStageModal,
		isCreateStageModalVisible,
		addClickHandler,
		isFormValid,
		resetModalState,
		saveClickHandler,
		stageFormStateJS,
		updateStageFormStateWrapper,
	} = props;

	const { stageId, stageType, stageRepeatCount } = stageFormStateJS;

	const isRepeatCountDisabled = stageType && stageType.value === 'hold';

	// stageId will be present when we are updating stage
	const isUpdateForm = stageId !== null;

	// eslint-disable-next-line arrow-body-style
	useEffect(() => {
		return () => {
			// reset from state
			resetModalState();
		};
		// eslint-disable-next-line
	}, []);

	const onChangeHandler = ({ target: { name, value } }) => {
		updateStageFormStateWrapper(name, value);
	};

	return (
		<>
			<Modal
				isOpen={isCreateStageModalVisible}
				toggle={toggleCreateStageModal}
				centered
				size='lg'
			>
				<ModalBody>
					<Text
						tag='h4'
						className='modal-title text-center text-truncate font-weight-bold'
					>
						Add Stage
					</Text>
					<ButtonIcon
						position='absolute'
						placement='right'
						top={24}
						right={32}
						size={32}
						name='cross'
						onClick={toggleCreateStageModal}
					/>
					<Form>
						<Row form className='mb-5 pb-5'>
							<Col sm={6}>
								<FormGroup>
									<Label for='stageType' className='font-weight-bold'>
										Stage type
									</Label>
									<Select
										options={stageTypeOptions}
										onChange={(selectedStageType) => {
											updateStageFormStateWrapper(
												'stageType',
												selectedStageType
											);
										}}
										value={stageType}
									/>
								</FormGroup>
							</Col>
							<Col sm={6}>
								<FormGroup>
									<Label for='repeatCount' className='font-weight-bold'>
										Repeat Count
									</Label>
									<Input
										type='number'
										min='0'
										name='stageRepeatCount'
										id='stage'
										placeholder='Type here'
										value={stageRepeatCount}
										onChange={onChangeHandler}
										disabled={isRepeatCountDisabled}
									/>
								</FormGroup>
							</Col>
						</Row>
						<Center className='text-center p-0 m-0 pt-5'>
							{isUpdateForm === false && (
								<Button
									color='primary'
									className='mb-3'
									onClick={addClickHandler}
									disabled={isFormValid === false}
								>
									Add
								</Button>
							)}
							{isUpdateForm === true && (
								<Button
									color='primary'
									className='mb-3'
									onClick={saveClickHandler}
									disabled={isFormValid === false}
								>
									Save
								</Button>
							)}
						</Center>
					</Form>
				</ModalBody>
			</Modal>
		</>
	);
};

AddStageModal.propTypes = {
	toggleCreateStageModal: PropTypes.func.isRequired,
	isCreateStageModalVisible: PropTypes.bool.isRequired,
	addClickHandler: PropTypes.func.isRequired,
	isFormValid: PropTypes.bool.isRequired,
	resetModalState: PropTypes.func.isRequired,
	saveClickHandler: PropTypes.func.isRequired,
	stageFormStateJS: PropTypes.object.isRequired,
	updateStageFormStateWrapper: PropTypes.func.isRequired,
};

export default AddStageModal;
