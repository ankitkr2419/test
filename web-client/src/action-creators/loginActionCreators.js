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

export const setActiveDeck = deckName => ({
	type: loginActions.setActiveDeck,
	payload: {
		deckName
	}
})

export const loginReset = (deckName) => ({
	type: loginActions.loginReset,
	payload: {
		deckName
	}
});

export const loginAsOperator = () => ({
	type: loginActions.setLoginTypeAsOperator,
});

export const setIsPlateRoute = isPlateRoute => ({
	type: loginActions.setIsPlateRoute,
	payload: {
		isPlateRoute,
	},
});

export const setIsTemplateRoute = isTemplateRoute => ({
	type: loginActions.setIsTemplateRoute,
	payload: {
		isTemplateRoute,
	},
});
