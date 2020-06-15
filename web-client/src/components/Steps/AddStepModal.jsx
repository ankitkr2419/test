import React, { useState } from "react";
import { Modal, ModalBody } from "core-components/Modal";
import { Row, Col } from "core-components/Grid";
import Form from "core-components/Form";
import FormGroup from "core-components/FormGroup";
import Label from "core-components/Label";
import Input from "core-components/Input";
import Button from "core-components/Button";
import Icon from "shared-components/Icon";
import Text from "shared-components/Text";
import ButtonGroup from "shared-components/ButtonGroup";
import ButtonIcon from "shared-components/ButtonIcon";

const AddStepModal = (props) => {
	const [stageModal, setStepModal] = useState(false);
	const toggleStepModal = () => setStepModal(!stageModal);

	return (
		<>
			<Button color="primary" isIcon onClick={toggleStepModal}>
				<Icon size={40} name="plus-2" />
			</Button>
			<Modal isOpen={stageModal} toggle={toggleStepModal} centered>
				<ModalBody>
					<Text tag="h4" className="modal-title">
						Add Step
					</Text>
					<ButtonIcon
						position="absolute"
						placement="right"
						top={24}
						right={32}
						size={32}
						name="cross"
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
										name="target_temperature"
										id="target_temperature"
										placeholder="Type here"
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
										name="hold_time"
										id="hold_time"
										placeholder="Type here"
									/>
									<Label>unit seconds</Label>
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
									<Label>boolean flag</Label>
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
