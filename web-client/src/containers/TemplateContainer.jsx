import React, { useEffect } from 'react';
import { Redirect } from 'react-router';
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

const TemplateContainer = (props) => {
	const { updateSelectedWizard, updateTemplateID } = props;
	const dispatch = useDispatch();
	// reading templates from redux
	const templates = useSelector(state => state.listTemplatesReducer);
	// isTemplateCreated = true means template created successfully
	const { isTemplateCreated } = useSelector(state => state.createTemplateReducer);
	// isTemplateDeleted = true means template deleted successfully
	const { isTemplateDeleted } = useSelector(state => state.deleteTemplateReducer);
	// loginReducer will listen to get logged in state with type of use logged in
	const loginReducer = useSelector(state => state.loginReducer);
	const { isUserLoggedIn, isLoginTypeOperator } = loginReducer.toJS();

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

	const createTemplate = (template) => {
		// creating template though api
		dispatch(createTemplateAction(template));
	};

	const deleteTemplate = (templateID) => {
		// deleting template though api
		dispatch(deleteTemplateAction(templateID));
	};

	if (isUserLoggedIn === false) {
		return <Redirect to='/login'/>;
	}

	return (
		<TemplateComponent
			// Extracting list before passing down to component reference=>Immutable
			templates={templates.get('list')}
			createTemplate={createTemplate}
			deleteTemplate={deleteTemplate}
			updateSelectedWizard={updateSelectedWizard}
			updateTemplateID={updateTemplateID}
			isLoginTypeOperator={isLoginTypeOperator}
		/>
	);
};

TemplateContainer.propTypes = {
	updateSelectedWizard: PropTypes.func.isRequired,
	updateTemplateID: PropTypes.func.isRequired,
};

export default TemplateContainer;
