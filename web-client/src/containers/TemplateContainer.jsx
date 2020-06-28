import React, { useEffect } from 'react';
import PropTypes from 'prop-types';
import { useDispatch, useSelector } from 'react-redux';
import TemplateComponent from 'components/Template';
import {
	fetchTemplates,
	createTemplate as createTemplateAction,
	deleteTemplate as deleteTemplateAction,
	addTemplateReset,
	deleteTemplateReset,
} from 'action-creators/templateActionCreators';

import {
	createExperiment as createExperimentAction,
	createExperimentReset,
} from 'action-creators/experimentActionCreators';

const TemplateContainer = (props) => {
	const {
		isLoginTypeOperator, isLoginTypeAdmin, updateSelectedWizard, updateTemplateID,
	} = props;
	const dispatch = useDispatch();
	// reading templates from redux
	const templates = useSelector(state => state.listTemplatesReducer);
	// isTemplateCreated = true means template created successfully
	const { isTemplateCreated } = useSelector(
		state => state.createTemplateReducer,
	);
	// isTemplateDeleted = true means template deleted successfully
	const { isTemplateDeleted } = useSelector(
		state => state.deleteTemplateReducer,
	);

	// isTemplateDeleted = true means template deleted successfully
	const { isExperimentSaved, id } = useSelector(
		state => state.createExperimentReducer,
	);

	useEffect(() => {
		// Once we create template will fetch updated template list
		if (isTemplateCreated === true) {
			dispatch(addTemplateReset());
			dispatch(fetchTemplates());
		}
	}, [isTemplateCreated, dispatch]);

	useEffect(() => {
		// Once we delete template will fetch updated template list
		if (isTemplateDeleted === true) {
			dispatch(deleteTemplateReset());
			dispatch(fetchTemplates());
		}
	}, [isTemplateDeleted, dispatch]);

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
	}, [updateSelectedWizard, dispatch, isExperimentSaved, id]);

	const createTemplate = (template) => {
		// creating template though api
		dispatch(createTemplateAction(template));
	};

	const deleteTemplate = (templateID) => {
		// deleting template though api
		dispatch(deleteTemplateAction(templateID));
	};

	/**
	 * createExperiment belongs to operator flow
	 */
	const createExperiment = (experimentBody) => {
		// console.log('experimentBody: ', experimentBody);
		dispatch(createExperimentAction(experimentBody));
	};

	return (
		<TemplateComponent
			// Extracting list before passing down to component reference=>Immutable
			templates={templates.get('list')}
			createTemplate={createTemplate}
			deleteTemplate={deleteTemplate}
			updateSelectedWizard={updateSelectedWizard}
			updateTemplateID={updateTemplateID}
			isLoginTypeAdmin={isLoginTypeAdmin}
			isLoginTypeOperator={isLoginTypeOperator}
			createExperiment={createExperiment}
		/>
	);
};

TemplateContainer.propTypes = {
	updateSelectedWizard: PropTypes.func.isRequired,
	updateTemplateID: PropTypes.func.isRequired,
};

export default React.memo(TemplateContainer);
