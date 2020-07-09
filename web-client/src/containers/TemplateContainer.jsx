import React, { useEffect } from 'react';
import PropTypes from 'prop-types';
import { useDispatch, useSelector } from 'react-redux';
import TemplateComponent from 'components/Template';
import {
	fetchTemplates,
	createTemplate as createTemplateAction,
	deleteTemplate as deleteTemplateAction,
	updateTemplate as updateTemplateAction,
	addTemplateReset,
	deleteTemplateReset,
	updateTemplateReset,
} from 'action-creators/templateActionCreators';

import {
	createExperiment as createExperimentAction,
	createExperimentReset,
} from 'action-creators/experimentActionCreators';
import { getIsExperimentSaved } from 'selectors/experimentSelector';

const TemplateContainer = (props) => {
	const {
		isLoginTypeOperator,
		isLoginTypeAdmin,
		updateSelectedWizard,
		updateTemplateID,
		isTemplateEdited,
		setIsTemplateEdited,
		selectedTemplateID,
	} = props;
	const dispatch = useDispatch();
	// reading templates from redux
	const templates = useSelector(state => state.listTemplatesReducer);
	// isTemplateCreated = true means template created successfully
	const { isTemplateCreated, response  } = useSelector(
		state => state.createTemplateReducer,
	);
	// isTemplateDeleted = true means template deleted successfully
	const { isTemplateDeleted } = useSelector(
		state => state.deleteTemplateReducer,
	);

	const { isTemplateUpdated } = useSelector(
		state => state.updateTemplateReducer,
	);
	// isTemplateDeleted = true means experiment created successfully
	const isExperimentSaved = useSelector(getIsExperimentSaved);

	useEffect(() => {
		// Once we create template will fetch updated template list
		if (isTemplateCreated === true) {
			// update the templateId in templateState maintained in templateLayout with created Id
			updateTemplateID(response.id);
			// navigate to next wizard
			updateSelectedWizard('target');
			dispatch(addTemplateReset());
		}
	}, [isTemplateCreated, dispatch, response, updateSelectedWizard, updateTemplateID]);

	useEffect(() => {
		// Once we delete template will fetch updated template list
		if (isTemplateDeleted === true) {
			dispatch(deleteTemplateReset());
			dispatch(fetchTemplates());
		}
	}, [isTemplateDeleted, dispatch]);

	useEffect(() => {
		if (isTemplateUpdated === true) {
			dispatch(updateTemplateReset());
			dispatch(fetchTemplates());
		}
	}, [isTemplateUpdated, dispatch]);

	useEffect(() => {
		// getting templates through api.
		dispatch(fetchTemplates());
	}, [dispatch]);

	/**
	 * if login type is operator
	 * once he select template will create experiment by calling experiment post rest call
	 * once its successful will navigate to target-operator wizard
	 */
	useEffect(() => {
		if (isExperimentSaved === true) {
			updateSelectedWizard('target-operator');
			dispatch(createExperimentReset());
		}
	}, [updateSelectedWizard, dispatch, isExperimentSaved]);

	const createTemplate = (template) => {
		// creating template though api
		dispatch(createTemplateAction(template));
	};

	const updateTemplate = (templateID, template) => {
		dispatch(updateTemplateAction(templateID, template));
	};

	const deleteTemplate = (templateID) => {
		// deleting template though api
		dispatch(deleteTemplateAction(templateID));
	};

	/**
	 * createExperiment belongs to operator flow
	 */
	const createExperiment = (experimentBody) => {
		dispatch(createExperimentAction(experimentBody));
	};

	return (
		<TemplateComponent
			// Extracting list before passing down to component reference=>Immutable
			templates={templates.get('list')}
			createTemplate={createTemplate}
			deleteTemplate={deleteTemplate}
			updateTemplate={updateTemplate}
			updateSelectedWizard={updateSelectedWizard}
			updateTemplateID={updateTemplateID}
			templateID={selectedTemplateID}
			isLoginTypeAdmin={isLoginTypeAdmin}
			isLoginTypeOperator={isLoginTypeOperator}
			createExperiment={createExperiment}
			isTemplateEdited={isTemplateEdited}
			setIsTemplateEdited={setIsTemplateEdited}
		/>
	);
};

TemplateContainer.propTypes = {
	updateSelectedWizard: PropTypes.func.isRequired,
	updateTemplateID: PropTypes.func.isRequired,
};

export default React.memo(TemplateContainer);
