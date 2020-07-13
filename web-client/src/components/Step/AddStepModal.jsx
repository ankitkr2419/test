import React, { useEffect } from 'react';
import PropTypes from 'prop-types';
import {
	Button,
	FormGroup,
	Row,
	Col,
	Input,
	Label,
	Modal,
	ModalBody,
	CheckBox,
	FormError,
} from 'core-components';
import { ButtonGroup, ButtonIcon, Text } from 'shared-components';
import { validateHoldTime, validateRampRate, validateTargetTemperature } from './stepHelper';

const AddStepModal = (props) => {
	const {
		isCreateStepModalVisible,
		toggleCreateStepModal,
		stepFormState,
		updateStepFormStateWrapper,
		isFormValid,
		addClickHandler,
		resetFormValues,
		saveClickHandler,
		stageType,
	} = props;

	const {
		stepId,
		rampRate,
		targetTemperature,
		holdTime,
		dataCapture,
		holdTimeError,
		rampRateError,
		targetTemperatureError,
	} = stepFormState;

	// stageId will be present when we are updating stage
	const isUpdateForm = stepId !== null;
	// If stageType is hold column size will be 4 or else will be 3
	const colSize = stageType === 'hold' ? 4 : 3;
	// eslint-disable-next-line arrow-body-style
	useEffect(() => {
		return () => {
			resetFormValues();
		};
		// eslint-disable-next-line
	}, []);

	const onChangeHandler = ({ target: { name, value } }) => {
		updateStepFormStateWrapper(name, value);
	};

	const onHoldTimeBlurHandler = () => {
		if (validateHoldTime(holdTime) === null) {
			updateStepFormStateWrapper('holdTimeError', true);
		}
	};

	const onHoldTimeFocusHandler = () => {
		updateStepFormStateWrapper('holdTimeError', false);
	};

	const onRampRateBlurHandler = () => {
		if (validateRampRate(rampRate) === false) {
			updateStepFormStateWrapper('rampRateError', true);
		}
	};

	const onRampRateFocusHandler = () => {
		updateStepFormStateWrapper('rampRateError', false);
	};

	const onTargetTemperatureBlurHandler = () => {
		if (validateTargetTemperature(targetTemperature) === false) {
			updateStepFormStateWrapper('targetTemperatureError', true);
		}
	};

	const onTargetTemperatureFocusHandler = () => {
		updateStepFormStateWrapper('targetTemperatureError', false);
	};

	return (
		<>
			<Modal
				isOpen={isCreateStepModalVisible}
				toggle={toggleCreateStepModal}
				centered
				size='lg'
			>
				<ModalBody>
					<Text
						tag='h4'
						className='modal-title text-center text-truncate font-weight-bold'
					>
						Add Step
					</Text>
					<ButtonIcon
						position='absolute'
						placement='right'
						top={24}
						right={32}
						size={32}
						name='cross'
						onClick={toggleCreateStepModal}
					/>
					<Row form className='mb-5 pb-5'>
						<Col sm={colSize}>
							<FormGroup>
								<Label for='ramp_rate' className='font-weight-bold'>
									Ramp Rate
								</Label>
								<Input
									type='number'
									name='rampRate'
									id='ramp_rate'
									placeholder='0.5 - 6'
									value={rampRate}
									onChange={onChangeHandler}
									onBlur={onRampRateBlurHandler}
									onFocus={onRampRateFocusHandler}
									invalid={rampRateError}
								/>
								<Label>unit °C</Label>
								<FormError>Invalid ramp rate</FormError>
							</FormGroup>
						</Col>
						<Col sm={colSize}>
							<FormGroup>
								<Label for='target_temperature' className='font-weight-bold'>
									Target Temperature
								</Label>
								<Input
									type='number'
									name='targetTemperature'
									id='target_temperature'
									placeholder='22 - 120'
									value={targetTemperature}
									onChange={onChangeHandler}
									onBlur={onTargetTemperatureBlurHandler}
									onFocus={onTargetTemperatureFocusHandler}
									invalid={targetTemperatureError}
								/>
								<Label>unit °C</Label>
								<FormError>Invalid target temperature</FormError>
							</FormGroup>
						</Col>
						<Col sm={colSize}>
							<FormGroup>
								<Label for='hold_time' className='font-weight-bold'>
									Hold Time
								</Label>
								<Input
									type='text'
									name='holdTime'
									id='hold_time'
									placeholder='seconds'
									value={holdTime}
									onBlur={onHoldTimeBlurHandler}
									onFocus={onHoldTimeFocusHandler}
									onChange={onChangeHandler}
									invalid={holdTimeError}
								/>
								<FormError>Invalid hold time</FormError>
							</FormGroup>
						</Col>
						{/* If the stage type is hold don't show datacapture checkbox */}
						{stageType !== 'hold' &&
						(<Col sm={colSize}>
							<FormGroup>
								<Label for='data_capture' className='font-weight-bold'>
									Data Capture
								</Label>
								<CheckBox
									name='dataCapture'
									id='dataCapture'
									onChange={(event) => {
										updateStepFormStateWrapper(
											event.target.name,
											event.target.checked,
										);
									}}
									className='mr-2 ml-3 py-2'
									checked={dataCapture}
								/>
							</FormGroup>
						</Col>)}
					</Row>
					<ButtonGroup className='text-center p-0 m-0 pt-5'>
						{isUpdateForm === false && (
							<Button
								onClick={addClickHandler}
								color='primary'
								disabled={isFormValid === false}
							>
								Add
							</Button>
						)}
						{isUpdateForm === true && (
							<Button
								onClick={saveClickHandler}
								color='primary'
								disabled={isFormValid === false}
							>
								Save
							</Button>
						)}
					</ButtonGroup>
				</ModalBody>
			</Modal>
		</>
	);
};

AddStepModal.propTypes = {
	isCreateStepModalVisible: PropTypes.bool.isRequired,
	toggleCreateStepModal: PropTypes.func.isRequired,
	stepFormState: PropTypes.object.isRequired,
	updateStepFormStateWrapper: PropTypes.func.isRequired,
	isFormValid: PropTypes.bool.isRequired,
	addClickHandler: PropTypes.func.isRequired,
	resetFormValues: PropTypes.func.isRequired,
	saveClickHandler: PropTypes.func.isRequired,
	stageType: PropTypes.string.isRequired,
};

export default AddStepModal;
