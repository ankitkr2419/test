import { fromJS } from 'immutable';
import modalActions from 'actions/modalActions';

const modalInitialState = fromJS({
	isModalVisible: false,
	modalType: null,
	message: null,
});

export const modalReducer = (state = modalInitialState, action) => {
	switch (action.type) {
	case modalActions.successModal:
		return state.merge({
			isModalVisible: true,
			modalType: modalActions.successModal,
			message: action.payload.message,
		});
	case modalActions.errorModal:
		return state.merge({
			isModalVisible: true,
			modalType: modalActions.errorModal,
			message: action.payload.message,
		});
	case modalActions.hideModal:
		return modalInitialState.merge({
			isModalVisible: false,
		});
	default:
		return state;
	}
};
