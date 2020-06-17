import React, { useEffect } from 'react';
import PropTypes from 'prop-types';
import { useDispatch, useSelector } from 'react-redux';
import TemplateComponent from 'components/Template';
import {
	fetchTemplates,
	createTemplate as createTemplateAction,
	deleteTemplate as deleteTemplateAction,
} from 'action-creators/templateActionCreators';

const TemplateContainer = (props) => {
	const { updateSelectedWizard, updateTemplateID } = props;
	const dispatch = useDispatch();
	// reading templates from redux
	const templates = useSelector(state => state.listTemplatesReducer);

	useEffect(() => {
		// getting templates through api.
		dispatch(fetchTemplates());
	}, [dispatch]);

	const createTemplate = (template) => {
		// creating template though api
		dispatch(createTemplateAction(template));
	};

	const deleteTemplate = (templateID) => {
		// deleting template though api
		dispatch(deleteTemplateAction(templateID));
	};

	return (
		<TemplateComponent
			// Extracting list before passing down to component ref. Immutable
			templates={templates.get('list')}
			createTemplate={createTemplate}
			deleteTemplate={deleteTemplate}
			updateSelectedWizard={updateSelectedWizard}
			updateTemplateID={updateTemplateID}
		/>
	);
};

TemplateContainer.propTypes = {
	updateSelectedWizard: PropTypes.func.isRequired,
	updateTemplateID: PropTypes.func.isRequired,
};

export default TemplateContainer;
