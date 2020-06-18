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
} from 'core-components';
import { ButtonGroup, ButtonIcon, Text } from 'shared-components';

const CreateTemplateModal = (props) => {
	const {
		isCreateTemplateModalVisible,
		toggleCreateTemplateModal,
		templateDescription,
		setTemplateDescription,
		templateName,
		setTemplateName,
		addClickHandler,
		isFormValid,
		resetFormValues,
	} = props;

	// eslint-disable-next-line arrow-body-style
	useEffect(() => {
		return () => {
			resetFormValues();
		};
		// eslint-disable-next-line
		}, []);

	return (
		<>
			<Modal
				isOpen={isCreateTemplateModalVisible}
				toggle={toggleCreateTemplateModal}
				centered
				size="lg"
			>
				<ModalBody>
					<Text
						tag="h4"
						className="modal-title text-center text-truncate font-weight-bold"
					>
            Create New Template
					</Text>
					<ButtonIcon
						position="absolute"
						placement="right"
						top={24}
						right={32}
						size={32}
						name="cross"
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
										value={templateName}
										onChange={(event) => {
											setTemplateName(event.target.value);
										}}
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
										value={templateDescription}
										onChange={(event) => {
											setTemplateDescription(event.target.value);
										}}
									/>
								</FormGroup>
							</Col>
						</Row>
						<ButtonGroup className="text-center p-0 m-0 pt-5">
							<Button
								onClick={addClickHandler}
								color="primary"
								disabled={isFormValid === false}
							>
                Add
							</Button>
						</ButtonGroup>
					</Form>
				</ModalBody>
			</Modal>
		</>
	);
};

CreateTemplateModal.propTypes = {
	isCreateTemplateModalVisible : PropTypes.bool.isRequired,
	toggleCreateTemplateModal : PropTypes.func.isRequired,
	templateDescription : PropTypes.string.isRequired,
	setTemplateDescription : PropTypes.func.isRequired,
	templateName : PropTypes.string.isRequired,
	setTemplateName : PropTypes.func.isRequired,
	addClickHandler : PropTypes.func.isRequired,
	isFormValid : PropTypes.bool.isRequired,
};

export default CreateTemplateModal;
