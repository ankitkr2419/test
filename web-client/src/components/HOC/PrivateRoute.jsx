import React from 'react';
import { useSelector } from 'react-redux';
import { Redirect } from 'react-router';

const privateRoute = Component => (props) => {
	const loginReducer = useSelector(state => state.loginReducer);
	const {
		isUserLoggedIn,
		isLoginTypeAdmin,
		isLoginTypeOperator,
	} = loginReducer.toJS();

	if (isUserLoggedIn === false) {
		return <Redirect to="/login" />;
	}

	return (
		<Component
			{...props}
			isLoginTypeAdmin={isLoginTypeAdmin}
			isLoginTypeOperator={isLoginTypeOperator}
		/>
	);
};
export default privateRoute;
