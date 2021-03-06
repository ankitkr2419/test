export const createTemplateActions =  {
	createAction: 'CREATE_TEMPLATE_INITIATED',
	successAction: 'CREATE_TEMPLATE_SUCCEEDED',
	failureAction: 'CREATE_TEMPLATE_FAILURE',
	createTemplateReset: 'CREATE_TEMPLATE_RESET',
};

//finish & save template creation process
export const finishCreateTemplateActions =  {
	createAction: 'FINISH_CREATE_TEMPLATE_INITIATED',
	successAction: 'FINISH_CREATE_TEMPLATE_SUCCEEDED',
	failureAction: 'FINISH_CREATE_TEMPLATE_FAILURE',
	createTemplateReset: 'FINISH_CREATE_TEMPLATE_RESET',
};

export const listTemplateActions =  {
	listAction: 'FETCH_TEMPLATES_INITIATED',
	successAction: 'FETCH_TEMPLATES_SUCCEEDED',
	failureAction: 'FETCH_TEMPLATES_FAILURE',
};

export const updateTemplateActions =  {
	updateAction: 'UPDATE_TEMPLATE_INITIATED',
	successAction: 'UPDATE_TEMPLATE_SUCCEEDED',
	failureAction: 'UPDATE_TEMPLATE_FAILURE',
	updateTemplateReset: 'UPDATE_TEMPLATE_RESET'
};

export const deleteTemplateActions =  {
	deleteAction: 'DELETE_TEMPLATE_INITIATED',
	successAction: 'DELETE_TEMPLATE_SUCCEEDED',
	failureAction: 'DELETE_TEMPLATE_FAILURE',
	deleteTemplateReset: 'DELETE_TEMPLATE_RESET',
};
