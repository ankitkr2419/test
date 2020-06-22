import { fromJS } from 'immutable';
import loginActions from 'actions/loginActions';

const loginInitialState = fromJS({
	isLoading: true,
	isUserLoggedIn: false,
	isLoginTypeAdmin: false,
	isLoginTypeOperator: false,
	isError: false,
});

export const loginReducer = (state = loginInitialState, action) => {
	switch (action.type) {
	case loginActions.loginInitiated:
		return state.merge({
			isLoading: true,
			isUserLoggedIn: false,
			isAdmin: false,
			isError: false,
		});
	case loginActions.successAction:
		return state.merge({
			isLoading: false,
			isUserLoggedIn: true,
			isAdmin: true,
			isError: false,
		});
	case loginActions.failureAction:
		return state.merge({
			isLoading: true,
			isUserLoggedIn: false,
			isAdmin: false,
			isError: true,
		});
	case loginActions.setLoginTypeAsOperator:
		return state.merge({
			isLoginTypeOperator: true,
			isUserLoggedIn: true,
			isAdmin: false,
		});
	case loginActions.loginReset:
		return loginInitialState;
	default:
		return state;
	}
};
