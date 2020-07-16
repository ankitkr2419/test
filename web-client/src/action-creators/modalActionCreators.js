const { default: modalActions } = require('actions/modalActions');

export const showSuccessModal = message => ({
	type: modalActions.successModal,
	payload: {
		message,
	},
});

export const showErrorModal = message => ({
	type: modalActions.errorModal,
	payload: {
		message,
	},
});

export const showWarningModal = message => ({
	type: modalActions.warningModal,
	payload: {
		message,
	},
});
