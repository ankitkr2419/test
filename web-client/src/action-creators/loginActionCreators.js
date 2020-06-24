import loginActions from 'actions/loginActions';

export const login = body => ({
	type: loginActions.loginInitiated,
	payload: {
		body,
	},
});

export const loginFailed = errorResponse => ({
	type: loginActions.failureAction,
	payload: {
		...errorResponse,
		error: true,
	},
});

export const loginReset = () => ({
	type: loginActions.loginReset,
});

export const loginAsOperator = () => ({
	type: loginActions.setLoginTypeAsOperator,
});
