import React, { useEffect } from 'react';
import TemplateComponent from 'components/Template';
import { useDispatch, useSelector } from 'react-redux';
import { fetchTemplates } from 'action-creators/templateActionCreators';

const TemplateContainer = (props) => {
	const dispatch = useDispatch();
	const templates = useSelector(state => state.listTemplatesReducer);

	useEffect(() => {
		dispatch(fetchTemplates());
	}, [dispatch]);

	return <TemplateComponent templates={templates.get('list')}/>;
};

TemplateContainer.propTypes = {};

export default TemplateContainer;
