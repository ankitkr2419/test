import React, { useState } from "react";
import {
	Button as IconButton,
	ModalBody,
	Form,
	Row,
	Col,
	FormGroup,
	Label,
	Input,
} from "reactstrap";
import Modal from "core-components/Modal";
import Icon from "shared-components/Icon";
import Text from "shared-components/Text";
import ButtonClose from "shared-components/ButtonClose";
import Button from "core-components/Button";
import ButtonGroup from "shared-components/ButtonGroup";

const AddStepModal = (props) => {
	const [stageModal, setStepModal] = useState(false);
	const toggleStepModal = () => setStepModal(!stageModal);

	return (
		<>
			<IconButton
				color="primary"
				className="btn-plus p-0"
				onClick={toggleStepModal}
			>
				<Icon size="40" name="plus-2" />
			</IconButton>
			<Modal isOpen={stageModal} toggle={toggleStepModal} centered>
				<ModalBody>
					<Text tag="h4" className="modal-title">
						Add Step
					</Text>
					<ButtonClose
						position="absolute"
						placement="right"
						top="24"
						right="32"
						onClick={toggleStepModal}
					/>
					<Form>
						<Row form className="mb-5 pb-5">
							<Col sm={3}>
								<FormGroup>
									<Label for="ramp_rate" className="font-weight-bold">
										Ramp Rate
									</Label>
									<Input
										type="text"
										name="ramp_rate"
										id="ramp_rate"
										placeholder="Type here"
									/>
									<Text tag="label">unit °C</Text>
								</FormGroup>
							</Col>
							<Col sm={3}>
								<FormGroup>
									<Label for="target_temperature" className="font-weight-bold">
										Target Temperature
									</Label>
									<Input
										type="text"
										name="target_temperature"
										id="target_temperature"
										placeholder="Type here"
									/>
									<Text tag="label">unit °C</Text>
								</FormGroup>
							</Col>
							<Col sm={3}>
								<FormGroup>
									<Label for="hold_time" className="font-weight-bold">
										Hold Time
									</Label>
									<Input
										type="text"
										name="hold_time"
										id="hold_time"
										placeholder="Type here"
									/>
									<Text tag="label">unit seconds</Text>
								</FormGroup>
							</Col>
							<Col sm={3}>
								<FormGroup>
									<Label for="data_capture" className="font-weight-bold">
										Data Capture
									</Label>
									<Input
										type="text"
										name="data_capture"
										id="data_capture"
										placeholder="Type here"
									/>
									<Text tag="label">boolean flag</Text>
								</FormGroup>
							</Col>
						</Row>
						<ButtonGroup className="text-center p-0 m-0 pt-5">
							<Button color="primary" disabled>
								Add
							</Button>
						</ButtonGroup>
					</Form>
				</ModalBody>
			</Modal>
		</>
	);
};

AddStepModal.propTypes = {};

export default AddStepModal;
