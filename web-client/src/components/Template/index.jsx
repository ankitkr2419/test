import React, { useState } from 'react';
import { Button } from 'core-components';
import PropTypes from 'prop-types';
import {
	StyledUl,
	StyledLi,
	CustomButton,
	Center,
	Text,
	ImageIcon,
} from 'shared-components';
import imgNoTemplate from 'assets/images/no-template-available.svg';

const TemplateComponent = (props) => {
	const {
		templates,
		deleteTemplate,
		updateSelectedWizard,
		updateTemplateID,
		isLoginTypeOperator,
		isLoginTypeAdmin,
		createExperiment,
		toggleTemplateModal,
	} = props;

	// Local state to store template name
	const [selectedTemplateId, setSelectedTemplateId] = useState(null);

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

	return (
		<div className="d-flex flex-100 flex-column p-4 mt-3">
			{templates.size === 0 && (
				<Center className="no-template-wrap">
					<ImageIcon
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
					<Button color="primary" onClick={toggleTemplateModal}>
            Create New
					</Button>
				)}
			</Center>
		</div>
	);
};

TemplateComponent.propTypes = {
	templates: PropTypes.shape({}).isRequired,
	deleteTemplate: PropTypes.func.isRequired,
	updateSelectedWizard: PropTypes.func.isRequired,
	updateTemplateID: PropTypes.func.isRequired,
	isLoginTypeOperator: PropTypes.bool.isRequired,
	isLoginTypeAdmin: PropTypes.bool.isRequired,
};

export default React.memo(TemplateComponent);
