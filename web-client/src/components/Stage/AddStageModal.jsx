import React, { useState } from "react";
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
} from "core-components";
import Icon from "shared-components/Icon";
import Text from "shared-components/Text";
import ButtonGroup from "shared-components/ButtonGroup";
import ButtonIcon from "shared-components/ButtonIcon";

const AddStageModal = props => {
  
  const [stageModal, setStageModal] = useState(false);
  const toggleStageModal = () => setStageModal(!stageModal);

  return (
		<>
			<Button color="primary" isIcon onClick={toggleStageModal}>
				<Icon size={40} name="plus-2" />
			</Button>
			<Modal isOpen={stageModal} toggle={toggleStageModal} centered size="lg">
				<ModalBody>
					<Text
						tag="h4"
						className="modal-title text-center text-truncate font-weight-bold"
					>
						Add Stage
					</Text>
					<ButtonIcon
						position="absolute"
						placement="right"
						top={24}
						right={32}
						size={32}
						name="cross"
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

AddStageModal.propTypes = {};

export default AddStageModal;