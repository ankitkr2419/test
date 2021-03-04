import React from 'react';
import { HashRouter as Router, Switch } from 'react-router-dom';
import { withRouter } from 'react-router';

import RouteWithSubRoutes from 'RouteHelper';
import AppLayoutContainer from 'containers/AppLayoutContainer';
import LoginContainer from 'containers/LoginContainer';
import PlateContainer from 'containers/PlateContainer';
import ActivityContainer from 'containers/ActivityContainer';
import TemplateLayout from 'layouts/TemplateLayout';
import PrivateRoute from 'components/HOC/PrivateRoute';
import SplashScreenContainer from 'containers/SplashScreenContainer';
import LandingPageContainer from 'containers/LandingPageContainer';
import { ROOT_URL_PATH } from './appConstants';

export const routeDefinitions = [
	{
		path: ROOT_URL_PATH,
		component: AppLayoutContainer,
		key: 1,
		routes: [
			{
				path: `${ROOT_URL_PATH}login`,
				exact: true,
				component: LoginContainer,
				key: 2,
			},
			{
				path: `${ROOT_URL_PATH}templates`,
				exact: true,
				component: PrivateRoute(TemplateLayout),
				key: 3,
			},
			{
				path: `${ROOT_URL_PATH}plate`,
				exact: true,
				component: PrivateRoute(PlateContainer),
				key: 4,
			},
			{
				path: `${ROOT_URL_PATH}activity`,
				exact: true,
				component: ActivityContainer,
				key: 5,
			},
			{
				path: `${ROOT_URL_PATH}splashscreen`,
				exact: true,
				component: SplashScreenContainer,
				key: 6,
			},
			{
				path: `${ROOT_URL_PATH}landing`,
				exact: true,
				component: LandingPageContainer,
				key: 7,
			},
			{
				path: `${ROOT_URL_PATH}deckcard`,
				exact: true,
				component: LandingPageContainer,
				key: 8,
			},
		],
	},
];

const Routes = (props) => {
	const routesList = props.routesList || routeDefinitions;
	return (
		<Router>
			<Switch>
				{
					routesList.map(route => <RouteWithSubRoutes key={ route.key } { ...route } />)
				}
			</Switch>
		</Router>
	);
};

export default withRouter(Routes);
