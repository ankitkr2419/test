import { fromJS } from 'immutable';

// const action types
export const templateModalActions = {
	SET_TEMPLATE_DESCRIPTION: 'SET_TEMPLATE_DESCRIPTION',
	SET_TEMPLATE_NAME: 'SET_TEMPLATE_NAME',
	SET_IS_TEMPLATE_EDITED: 'SET_IS_TEMPLATE_EDITED',
	RESET_TEMPLATE_MODAL: 'RESET_TEMPLATE_MODAL',
	TOGGLE_TEMPLATE_MODAL_VISIBLE: 'TOGGLE_TEMPLATE_MODAL_VISIBLE',
};

// Initial state wrap with fromJS for immutability
export const templateModalInitialState = fromJS({
	templateDescription: '',
	templateName: '',
	isCreateTemplateModalVisible: false,
	isTemplateEdited: false,
});

const templateModalReducer = (state, action) => {
	switch (action.type) {
	case templateModalActions.SET_TEMPLATE_NAME:
		return state.setIn(['templateName'], action.value);
	case templateModalActions.SET_TEMPLATE_DESCRIPTION:
		return state.setIn(['templateDescription'], action.value);
	case templateModalActions.TOGGLE_TEMPLATE_MODAL_VISIBLE:
		return state.update('isCreateTemplateModalVisible', value => !value);
	case templateModalActions.SET_IS_TEMPLATE_EDITED:
		return state.setIn(['isTemplateEdited'], true);
	case templateModalActions.RESET_TEMPLATE_MODAL:
		return templateModalInitialState;
	default:
		throw new Error('Invalid action type');
	}
};

export default templateModalReducer;
