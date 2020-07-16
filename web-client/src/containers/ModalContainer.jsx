import React from 'react';
import ErrorModal from 'components/modals/ErrorModal';
import { useSelector, useDispatch } from 'react-redux';
import modalActions from 'actions/modalActions';

/**
 * ModalContainer can we used to show error/info through modal.
 */
export default function ModalContainer() {
	const dispatch = useDispatch();
	const modalReducer = useSelector(state => state.modalReducer);
	const modalType = modalReducer.get('modalType');

	const hideModal = () => {
		dispatch({
			type: modalActions.hideModal,
		});
	};

	if (modalType === modalActions.errorModal) {
		return (
			<ErrorModal
				isOpen={modalReducer.get('isModalVisible')}
				hideModal={hideModal}
				message={modalReducer.get('message')}
			/>
		);
	}

	return null;
}
