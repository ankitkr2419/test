import React, { useState } from 'react';
import { Button } from 'core-components';
import PropTypes from 'prop-types';
import {
	StyledUl,
	StyledLi,
	CustomButton,
	Center,
	Text,
} from 'shared-components';
import CreateTemplateModal from './CreateTemplateModal';

const TemplateComponent = (props) => {
	const {
		templates,
		createTemplate,
		deleteTemplate,
		updateSelectedWizard,
		updateTemplateID,
	} = props;

	// Local state to manage create template modal
	const [
		isCreateTemplateModalVisible,
		setCreateTemplateModalVisibility,
	] = useState(false);
	// Local state to store template description
	const [templateDescription, setTemplateDescription] = useState('');
	// Local state to store template name
	const [templateName, setTemplateName] = useState('');

	// helper method to toggle create template modal
	const toggleCreateTemplateModal = () => {
		setCreateTemplateModalVisibility(!isCreateTemplateModalVisible);
	};

	// Validate create template form
	const validateTemplateForm = () => {
		if (templateDescription !== '' && templateName !== '') {
			return true;
		}
		return false;
	};

	const addClickHandler = () => {
		if (validateTemplateForm()) {
			// Create new template rest api call.
			createTemplate({
				description: templateDescription,
				name: templateName,
			});
			toggleCreateTemplateModal();
		}
		// TODO show error notification
	};

	const deleteClickHandler = (templateID) => {
		deleteTemplate(templateID);
	};

	const editClickHandler = (templateID) => {
		updateTemplateID(templateID);
		updateSelectedWizard('target');
	};

	const resetFormValues = () => {
		setTemplateDescription('');
		setTemplateName('');
	};

	return (
		<div className="d-flex flex-100 flex-column p-4 mt-3">
			{templates.size === 0 && (
				<Text className="d-flex justify-content-center" Tag="h4">
          No templates available
				</Text>
			)}
			<StyledUl>
				{/* templates size check before iteration */}
				{templates.size !== 0
          && templates.map(template => (
					<StyledLi key={template.get('id')}>
						<CustomButton
							title={template.get('name')}
							isEditable
							onEditClickHandler={() => {
								editClickHandler(template.get('id'));
							}}
							isDeletable
							onDeleteClickHandler={() => {
								deleteClickHandler(template.get('id'));
							}}
						/>
					</StyledLi>
          ))}
			</StyledUl>
			<Center className="text-center">
				{/*
          TODO Handle login flow when operator
          <Button color="primary">Next</Button>
        */}
				<Button color="primary" onClick={toggleCreateTemplateModal}>
          Create New
				</Button>
			</Center>
			{isCreateTemplateModalVisible && (
				<CreateTemplateModal
					isCreateTemplateModalVisible={isCreateTemplateModalVisible}
					toggleCreateTemplateModal={toggleCreateTemplateModal}
					templateDescription={templateDescription}
					setTemplateDescription={setTemplateDescription}
					templateName={templateName}
					setTemplateName={setTemplateName}
					addClickHandler={addClickHandler}
					isFormValid={validateTemplateForm()}
					resetFormValues={resetFormValues}
				/>
			)}
		</div>
	);
};

TemplateComponent.propTypes = {
	templates: PropTypes.shape({}).isRequired,
	createTemplate: PropTypes.func.isRequired,
	deleteTemplate: PropTypes.func.isRequired,
	updateSelectedWizard: PropTypes.func.isRequired,
	updateTemplateID: PropTypes.func.isRequired,
};

export default React.memo(TemplateComponent);
