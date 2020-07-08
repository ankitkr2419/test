import React, { useState, useEffect } from 'react';
import { Button } from 'core-components';
import PropTypes from 'prop-types';
import {
	StyledUl,
	StyledLi,
	CustomButton,
	Center,
	Text,
} from 'shared-components';
import imgNoTemplate from 'assets/images/no-template-available.svg';
import CreateTemplateModal from './CreateTemplateModal';


const TemplateComponent = (props) => {
	const {
		templates,
		createTemplate,
		deleteTemplate,
		updateTemplate,
		updateSelectedWizard,
		updateTemplateID,
		isLoginTypeOperator,
		isLoginTypeAdmin,
		createExperiment,
		isTemplateEdited,
		setIsTemplateEdited,
		templateID,
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
	// Local state to store template name
	const [selectedTemplateId, setSelectedTemplateId] = useState(null);

	// helper method to toggle create template modal
	const toggleCreateTemplateModal = () => {
		setCreateTemplateModalVisibility(!isCreateTemplateModalVisible);
	};

	const autofillNameDescription = () => {
		const template = templates.find(ele => ele.get('id') === templateID);
		setTemplateName(template.get('name'));
		setTemplateDescription(template.get('description'));
	};

	useEffect(() => {
		if (isTemplateEdited === true && isCreateTemplateModalVisible === false) {
			autofillNameDescription();
			toggleCreateTemplateModal();
		}
	}, [isTemplateEdited]);

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
			if (isTemplateEdited) {
				updateTemplate(templateID, {
					description: templateDescription,
					name: templateName,
				});
				setIsTemplateEdited(false);
			} else {
				createTemplate({
					description: templateDescription,
					name: templateName,
				});
			}
			toggleCreateTemplateModal();
		}
		// TODO show error notification
	};

	const deleteClickHandler = () => {
		// Delete api call
		deleteTemplate(selectedTemplateId);
	};

	const editClickHandler = () => {
		// Updates template id to templateState maintain over templateLayout
		updateTemplateID(selectedTemplateId);
		// navigate to next wizard
		updateSelectedWizard('target');
	};

	const onTemplateButtonClickHandler = ({ id: templateId, description }) => {
		// if its admin save template id and show edit & delete options on button
		// if its operator save template id an navigate to target wizard

		// set selected template id to local state for maintaining active state of button
		setSelectedTemplateId(templateId);
		if (isLoginTypeOperator === true) {
			// make api call to save experiments
			createExperiment({
				template_id: templateId,
				description,
			});
			// Updates template id to templateState maintain over templateLayout
			updateTemplateID(templateId);
			// navigate to next wizard
			// updateSelectedWizard('target');
		}
	};

	const resetFormValues = () => {
		setTemplateDescription('');
		setTemplateName('');
	};

	return (
		<div className="d-flex flex-100 flex-column p-4 mt-3">
			{templates.size === 0 && (
				<Center className="no-template-wrap">
					<img
						src={imgNoTemplate}
						alt="No templates available"
						className="img-no-template"
					/>
					<Text className="d-flex justify-content-center" Tag="p">
            No templates available
					</Text>
				</Center>
			)}

			{/* templates size check before iteration */}
			{templates.size !== 0 && (
				<StyledUl>
					{templates.map(template => (
						<StyledLi key={template.get('id')}>
							<CustomButton
								title={template.get('name')}
								isActive={template.get('id') === selectedTemplateId}
								isEditable={
									isLoginTypeAdmin === true && isLoginTypeOperator === false
								}
								onButtonClickHandler={() => {
									onTemplateButtonClickHandler(template.toJS());
								}}
								onEditClickHandler={editClickHandler}
								isDeletable={
									isLoginTypeAdmin === true && isLoginTypeOperator === false
								}
								onDeleteClickHandler={deleteClickHandler}
							/>
						</StyledLi>
					))}
				</StyledUl>
			)}
			<Center className="mb-5">
				{isLoginTypeAdmin === true && (
					<Button color="primary" onClick={toggleCreateTemplateModal}>
            Create New
					</Button>
				)}
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
					isTemplateEdited={isTemplateEdited}
					setIsTemplateEdited={setIsTemplateEdited}
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
	isLoginTypeOperator: PropTypes.bool.isRequired,
	isLoginTypeAdmin: PropTypes.bool.isRequired,
};

export default React.memo(TemplateComponent);
