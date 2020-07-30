import React from 'react';
import PropTypes from 'prop-types';
import {
	Button,
	FormGroup,
	Row,
	Col,
	Input,
	InputGroupWithAddonText,
	Label,
	Modal,
	ModalBody,
	CheckBox,
} from 'core-components';
import { ButtonGroup, ButtonIcon, Text } from 'shared-components';
import {
	validateHoldTime,
	validateRampRate,
	validateTargetTemperature,
} from './stepHelper';
import {
	MIN_RAMP_RATE,
	MAX_RAMP_RATE,
	MAX_TARGET_TEMPERATURE,
	MIN_TARGET_TEMPERATURE,
	CYCLE_STAGE,
	HOLD_STAGE,
} from './stepConstants';

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
		cycleRepeatCount,
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
	const colSize = stageType === HOLD_STAGE ? 4 : 3;

	const onChangeHandler = ({ target: { name, value } }) => {
		// set rampRate/targetTemperature/holdTime with its value in stepForm local state
		updateStepFormStateWrapper(name, value);
	};

	const onHoldTimeBlurHandler = () => {
		if (validateHoldTime(holdTime) === null) {
			// set holdTimeError flag to true maintained over stepForm local state
			updateStepFormStateWrapper('holdTimeError', true);
		}
	};

	const onHoldTimeFocusHandler = () => {
		// reset holdTimeError flag to false maintained over stepForm local state
		updateStepFormStateWrapper('holdTimeError', false);
	};

	const onRampRateBlurHandler = () => {
		if (validateRampRate(rampRate) === false) {
			// set rampRateError flag to true maintained over stepForm local state
			updateStepFormStateWrapper('rampRateError', true);
		}
	};

	const onRampRateFocusHandler = () => {
		// reset rampRateError flag to false maintained over stepForm local state
		updateStepFormStateWrapper('rampRateError', false);
	};

	const onTargetTemperatureBlurHandler = () => {
		if (validateTargetTemperature(targetTemperature) === false) {
			// set targetTemperatureError flag to true maintained over stepForm local state
			updateStepFormStateWrapper('targetTemperatureError', true);
		}
	};

	const onTargetTemperatureFocusHandler = () => {
		// reset targetTemperatureError flag to false maintained over stepForm local state
		updateStepFormStateWrapper('targetTemperatureError', false);
	};

	return (
		<>
			<Modal
				isOpen={isCreateStepModalVisible}
				toggle={toggleCreateStepModal}
				onExit={resetFormValues}
				centered
				size='lg'
			>
				<ModalBody>
					<Text
						tag='h4'
						size={24}
						className='modal-title text-center text-truncate text-capitalize font-weight-bold'
					>
						Add Step - {stageType}{' '}
						{/* If its cycle stage then show repeat count in header */}
						{stageType === CYCLE_STAGE ? `(Repeat count - ${cycleRepeatCount})` : ''}
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
					<Row form className='mb-3 pb-3'>
						<Col sm={colSize}>
							<FormGroup>
								<Label for='ramp_rate' className='font-weight-bold'>
									Ramp Rate
								</Label>
								<InputGroupWithAddonText addonText='unit °C'>
									<Input
										type='number'
										name='rampRate'
										id='ramp_rate'
										placeholder={`${MIN_RAMP_RATE} - ${MAX_RAMP_RATE}`}
										value={rampRate}
										onChange={onChangeHandler}
										onBlur={onRampRateBlurHandler}
										onFocus={onRampRateFocusHandler}
										invalid={rampRateError}
									/>
								</InputGroupWithAddonText>
								<Text
									Tag='p'
									size={12}
									className={`${rampRateError && 'text-danger'} px-2 mb-0`}
								>
									Enter value between {MIN_RAMP_RATE} to {MAX_RAMP_RATE}
								</Text>
							</FormGroup>
						</Col>
						<Col sm={colSize}>
							<FormGroup>
								<Label for='target_temperature' className='font-weight-bold'>
									Target Temperature
								</Label>
								<InputGroupWithAddonText addonText='unit °C'>
									<Input
										type='number'
										name='targetTemperature'
										id='target_temperature'
										placeholder={`${MIN_TARGET_TEMPERATURE} - ${MAX_TARGET_TEMPERATURE}`}
										value={targetTemperature}
										onChange={onChangeHandler}
										onBlur={onTargetTemperatureBlurHandler}
										onFocus={onTargetTemperatureFocusHandler}
										invalid={targetTemperatureError}
									/>
								</InputGroupWithAddonText>
								<Text
									Tag='p'
									size={12}
									className={`${
										targetTemperatureError && 'text-danger'
									} px-2 mb-0`}
								>
									Enter value between {MIN_TARGET_TEMPERATURE} to {MAX_TARGET_TEMPERATURE}
								</Text>
							</FormGroup>
						</Col>
						<Col sm={colSize}>
							<FormGroup>
								<Label for='hold_time' className='font-weight-bold'>
									Hold Time
								</Label>
								<InputGroupWithAddonText addonText='unit sec'>
									<Input
										type='number'
										name='holdTime'
										id='hold_time'
										placeholder='seconds'
										value={holdTime}
										onBlur={onHoldTimeBlurHandler}
										onFocus={onHoldTimeFocusHandler}
										onChange={onChangeHandler}
										invalid={holdTimeError}
									/>
								</InputGroupWithAddonText>
								{holdTimeError && (
									<Text Tag='p' size={12} className='text-danger px-2 mb-0'>
										Invalid Hold time
									</Text>
								)}
							</FormGroup>
						</Col>
						{/* If the stage type is hold don't show datacapture checkbox */}
						{stageType !== HOLD_STAGE && (
							<Col sm={colSize}>
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
							</Col>
						)}
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
	cycleRepeatCount: PropTypes.number.isRequired,
};

export default AddStepModal;
