import React, { useEffect } from 'react';
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
} from 'core-components';
import {
	ButtonGroup, ButtonIcon, Text,
} from 'shared-components';

const AddStepModal = (props) => {
	const  {
		isCreateStepModalVisible,
		toggleCreateStepModal,
		stepFormState,
		updateStepFormStateWrapper,
		isFormValid,
		addClickHandler,
		resetFormValues,
	} = props;

	const {
		rampRate,
		targetTemperature,
		holdTime,
		dataCapture,
	} = stepFormState;

	// eslint-disable-next-line arrow-body-style
	useEffect(() => {
		return () => {
			resetFormValues();
		};
	}, [resetFormValues]);

	const onChangeHandler = ({ target: { name, value } }) => {
		updateStepFormStateWrapper(name, value);
	};

	return (
		<>
			<Modal isOpen={isCreateStepModalVisible} toggle={toggleCreateStepModal} centered size="lg">
				<ModalBody>
					<Text tag="h4" className="modal-title text-center text-truncate font-weight-bold">
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
									type="text"
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
									type="text"
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
									type="time"
									name="holdTime"
									id="hold_time"
									placeholder="mm:ss"
									value={holdTime}
									onChange={onChangeHandler}
								/>
								<Label>mm:ss</Label>
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
										updateStepFormStateWrapper(event.target.name, event.target.checked);
									}}
									className="mr-2"
									checked={dataCapture}
								/>
							</FormGroup>
						</Col>
					</Row>
					<ButtonGroup className="text-center p-0 m-0 pt-5">
						<Button onClick={addClickHandler} color="primary" disabled={isFormValid === false}>
								Add
						</Button>
					</ButtonGroup>
				</ModalBody>
			</Modal>
		</>
	);
};

AddStepModal.propTypes = {};

export default AddStepModal;
