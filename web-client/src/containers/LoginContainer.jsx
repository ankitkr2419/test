import React, { useState } from 'react';
import { Redirect } from 'react-router';

import LoginComponent from 'components/Login';
import { useDispatch, useSelector } from 'react-redux';
import { loginAsOperator as loginAsOperatorAction, login } from 'action-creators/loginActionCreators';
import { connectSocket } from 'web-socket';

const LoginContainer = () => {
	const loginReducer = useSelector(state => state.loginReducer);
	const { isUserLoggedIn, isError } = loginReducer.toJS();

	// isError in case user enters wrong credentials
	// TODO Extract error message from response once api implemented

	// local state to handle admin form visibility
	const [isAdminFormVisible, setIsAdminFormVisibility] = useState(false);
	const dispatch = useDispatch();

	const loginAsOperator = () => {
		// TODO check if connection is already opened
		connectSocket(dispatch);
		dispatch(loginAsOperatorAction());
	};

	const loginAsAdmin = (data) => {
		dispatch(login(data));
	};

	// redirection to admin once logged in
	if (isUserLoggedIn === true) {
	// if (isUserLoggedIn === true && isSocketConnected === true) {
		return <Redirect to='/templates'/>;
	}

	return (
		<LoginComponent
			isAdminFormVisible={isAdminFormVisible}
			setIsAdminFormVisibility={setIsAdminFormVisibility}
			loginAsOperator={loginAsOperator}
			loginAsAdmin={loginAsAdmin}
			isLoginError={isError}
		/>
	);
};

export default LoginContainer;
