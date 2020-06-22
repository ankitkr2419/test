import React from 'react';
import { BrowserRouter as Router, Switch, Redirect } from 'react-router-dom';
import RouteWithSubRoutes from 'RouteHelper';
import { useSelector } from 'react-redux';
import AppHeader from 'components/AppHeader';

import '../assets/scss/default.scss';
/**
 * AppLayoutContainer Will contain routes(content), headers, sub-headers, notification etc.
 * @param {*} props
 */
const AppLayoutContainer = (props) => {
	const { routes } = props;
	const loginReducer  = useSelector(state => state.loginReducer);

	return (
		<Router>
			<AppHeader isUserLoggedIn={loginReducer.get('isUserLoggedIn')}/>
			<section className="ml-content">
				<Switch>
					<Redirect exact from="/" to="/login" />
					{routes.map(route => <RouteWithSubRoutes key={route.key} {...route} />)}
				</Switch>
			</section>
		</Router>
	);
};

AppLayoutContainer.propTypes = {};

export default AppLayoutContainer;
