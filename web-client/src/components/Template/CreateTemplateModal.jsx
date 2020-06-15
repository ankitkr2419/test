import React, { useState } from "react";
import {
	Form,
	FormGroup,
	Label,
	Input,
} from "reactstrap";
import { Modal, ModalBody } from "core-components/Modal";
import { Row, Col} from "core-components/Grid";
import Text from "shared-components/Text";
import ButtonClose from "shared-components/ButtonClose";
import Button from "core-components/Button";
import ButtonGroup from "shared-components/ButtonGroup";

const CreateTemplateModal = (props) => {
	const [createTemplateModal, setCreateTemplateModal] = useState(false);
	const toggleCreateTemplateModal = () => setCreateTemplateModal(!createTemplateModal);

	return (
		<>
			<Button color="primary" onClick={toggleCreateTemplateModal}>
				Create New
			</Button>
			<Modal
				isOpen={createTemplateModal}
				toggle={toggleCreateTemplateModal}
				centered
			>
				<ModalBody>
					<Text tag="h4" className="modal-title">
						Create New Template
					</Text>
					<ButtonClose
						position="absolute"
						placement="right"
						top="24"
						right="32"
						onClick={toggleCreateTemplateModal}
					/>
					<Form>
						<Row form className="mb-5 pb-5">
							<Col sm={3}>
								<FormGroup>
									<Label for="template_name" className="font-weight-bold">
										Template Name
									</Label>
									<Input
										type="text"
										name="template_name"
										id="template_name"
										placeholder="Type here"
									/>
								</FormGroup>
							</Col>
							<Col sm={9}>
								<FormGroup>
									<Label
										for="template_description"
										className="font-weight-bold"
									>
										Description
									</Label>
									<Input
										type="text"
										name="template_description"
										id="template_description"
										placeholder="Type here"
									/>
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

CreateTemplateModal.propTypes = {};

export default CreateTemplateModal;
