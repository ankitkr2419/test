import React from "react";
import { BrowserRouter as Router, Switch } from "react-router-dom";
import { withRouter } from "react-router";

import RouteWithSubRoutes from "RouteHelper";
import { ROOT_URL_PATH } from "./constants";
import AppLayoutContainer from 'containers/AppLayoutContainer';
import LoginContainer from "containers/LoginContainer";
import PlateContainer from "containers/PlateContainer";
import ActivityContainer from "containers/ActivityContainer";
import TemplateLayout from "layouts/TemplateLayout";

export const routeDefinitions = [
	{
		path: ROOT_URL_PATH,
		component: AppLayoutContainer,
		key: 1,
		routes: [
			{
				path: `${ROOT_URL_PATH}templates`,
				exact: true,
				component: TemplateLayout,
				key: 2,
			},
			{
				path: `${ROOT_URL_PATH}login`,
				exact: true,
				component: LoginContainer,
				key: 2,
			},
			{
				path: `${ROOT_URL_PATH}plate`,
				exact: true,
				component: PlateContainer,
				key: 2,
			},
			{
				path: `${ROOT_URL_PATH}activity`,
				exact: true,
				component: ActivityContainer,
				key: 2,
			},
		],
	},
];

const Routes = props => {
  const routesList = props.routesList || routeDefinitions;
  return (
    <Router>
      <Switch>
        {
          routesList.map(route => {
            return <RouteWithSubRoutes key={ route.key } { ...route } />
          })
        }
      </Switch>
    </Router>
  );
};

export default withRouter(Routes);
