import React, { useState } from "react";
import {
	Button as IconButton,
	Form,
	FormGroup,
	Label,
	Input
} from "reactstrap";
import { Modal, ModalBody } from "core-components/Modal";
import { Row, Col } from "core-components/Grid";
import Icon from "shared-components/Icon";
import Text from "shared-components/Text";
import ButtonClose from "shared-components/ButtonClose";
import Select from "core-components/Select";
import Button from "core-components/Button";
import ButtonGroup from "shared-components/ButtonGroup";

const AddStageModal = props => {
  
  const [stageModal, setStageModal] = useState(false);
  const toggleStageModal = () => setStageModal(!stageModal);

  return (
		<>
			<IconButton
				color="primary"
				className="btn-plus p-0"
				onClick={toggleStageModal}
			>
				<Icon size={40} name="plus-2" />
			</IconButton>
			<Modal isOpen={stageModal} toggle={toggleStageModal} centered>
				<ModalBody>
					<Text tag="h4" className="modal-title">
						Add Stage
					</Text>
					<ButtonClose
						position="absolute"
						placement="right"
						top="24"
						right="32"
						onClick={toggleStageModal}
					/>
					<Form>
						<Row form className="mb-5 pb-5">
							<Col sm={4}>
								<FormGroup>
									<Label for="stage" className="font-weight-bold">
										Stage
									</Label>
									<Input
										type="text"
										name="stage"
										id="stage"
										placeholder="Type here"
									/>
								</FormGroup>
							</Col>
							<Col sm={4}>
								<FormGroup>
									<Label for="stageType" className="font-weight-bold">
										Stage type
									</Label>
									<Select />
								</FormGroup>
							</Col>
							<Col sm={3}>
								<FormGroup>
									<Label for="repeatCount" className="font-weight-bold">
										Repeat Count
									</Label>
									<Select />
								</FormGroup>
							</Col>
						</Row>
						<ButtonGroup className="text-center p-0 m-0 pt-5">
							<Button color="primary" disabled>Add</Button>
						</ButtonGroup>
					</Form>
				</ModalBody>
			</Modal>
		</>
	);
};

AddStageModal.propTypes = {};

export default AddStageModal;