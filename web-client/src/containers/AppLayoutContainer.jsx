import React from "react";
import { BrowserRouter as Router, Switch } from "react-router-dom";
import RouteWithSubRoutes from "RouteHelper";
import Header from "components/Header";
import "../assets/scss/default.scss";

const AppLayoutContainer = (props) => {
  // AppLayoutContainer Will contain headers, sub-headers, notification etc.
  const { routes } = props;

  return (
		<Router>
			<Header />
			<section className="ml-content">
				<Switch>
					{/* 
            TODO redirect to home page 
            <Redirect exact from='/' to='/templates' />
          */}
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
