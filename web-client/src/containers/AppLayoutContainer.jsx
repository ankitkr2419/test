import React from "react";
import { BrowserRouter as Router, Switch, Redirect } from "react-router-dom";
import RouteWithSubRoutes from "RouteHelper";
import Header from "components/Header";
import Logo from "shared-components/Logo";
import "../assets/scss/default.scss";
import {
	Nav,
	NavItem,
	NavLink,
} from "core-components/Nav";

const AppLayoutContainer = (props) => {
  // AppLayoutContainer Will contain headers, sub-headers, notification etc.
  const { routes } = props;

  return (
		<Router>
			<Header>
				<Logo isSmall />
				<Nav className="mx-3">
					<NavItem>
						<NavLink to="/templates">Template</NavLink>
					</NavItem>
					<NavItem>
						<NavLink to="/plate">Plate</NavLink>
					</NavItem>
					<NavItem>
						<NavLink to="/activity">Activity Log</NavLink>
					</NavItem>
				</Nav>
			</Header>
			<section className="ml-content">
				<Switch>
					<Redirect exact from='/' to='/login' />
					{routes.map((route) => {
						return <RouteWithSubRoutes key={route.key} {...route} />;
					})}
				</Switch>
			</section>
		</Router>
	);
};

AppLayoutContainer.propTypes = {};

export default AppLayoutContainer;
