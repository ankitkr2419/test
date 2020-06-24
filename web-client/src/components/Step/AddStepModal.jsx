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
import { validateHoldTime } from './stepHelper';

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
	} = props;

	const {
		stepId,
		rampRate,
		targetTemperature,
		holdTime,
		dataCapture,
		holdTimeError,
	} = stepFormState;

	// stageId will be present when we are updating stage
	const isUpdateForm = stepId !== null;

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

	return (
		<>
			<Modal
				isOpen={isCreateStepModalVisible}
				toggle={toggleCreateStepModal}
				centered
				size="lg"
			>
				<ModalBody>
					<Text
						tag="h4"
						className="modal-title text-center text-truncate font-weight-bold"
					>
            Add Step
					</Text>
					<ButtonIcon
						position="absolute"
						placement="right"
						top={24}
						right={32}
						size={32}
						name="cross"
						onClick={toggleCreateStepModal}
					/>
					<Row form className="mb-5 pb-5">
						<Col sm={3}>
							<FormGroup>
								<Label for="ramp_rate" className="font-weight-bold">
                  Ramp Rate
								</Label>
								<Input
									type="number"
									min="-273.15"
									max="1000"
									name="rampRate"
									id="ramp_rate"
									placeholder="Type here"
									value={rampRate}
									onChange={onChangeHandler}
								/>
								<Label>unit °C</Label>
							</FormGroup>
						</Col>
						<Col sm={3}>
							<FormGroup>
								<Label for="target_temperature" className="font-weight-bold">
                  Target Temperature
								</Label>
								<Input
									type="number"
									min="-273.15"
									max="1000"
									name="targetTemperature"
									id="target_temperature"
									placeholder="Type here"
									value={targetTemperature}
									onChange={onChangeHandler}
								/>
								<Label>unit °C</Label>
							</FormGroup>
						</Col>
						<Col sm={3}>
							<FormGroup>
								<Label for="hold_time" className="font-weight-bold">
                  Hold Time
								</Label>
								<Input
									type="text"
									name="holdTime"
									id="hold_time"
									placeholder="mm:ss"
									value={holdTime}
									onBlur={onHoldTimeBlurHandler}
									onFocus={onHoldTimeFocusHandler}
									onChange={onChangeHandler}
									invalid={holdTimeError}
								/>
								<FormError>Invalid hold time</FormError>
							</FormGroup>
						</Col>
						<Col sm={3}>
							<FormGroup>
								<Label for="data_capture" className="font-weight-bold">
                  Data Capture
								</Label>
								<CheckBox
									name="dataCapture"
									id="dataCapture"
									onChange={(event) => {
										updateStepFormStateWrapper(
											event.target.name,
											event.target.checked,
										);
									}}
									className="mr-2 ml-3"
									checked={dataCapture}
								/>
							</FormGroup>
						</Col>
					</Row>
					<ButtonGroup className="text-center p-0 m-0 pt-5">
						{isUpdateForm === false && (
							<Button
								onClick={addClickHandler}
								color="primary"
								disabled={isFormValid === false}
							>
                Add
							</Button>
						)}
						{isUpdateForm === true && (
							<Button
								onClick={saveClickHandler}
								color="primary"
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
};

export default AddStepModal;
