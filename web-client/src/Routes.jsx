import React from "react";
import { BrowserRouter as Router, Switch } from "react-router-dom";
import { withRouter } from "react-router";

import RouteWithSubRoutes from "./RouteHelper";

import AppLayoutContainer from './containers/AppLayoutContainer';
import { ROOT_URL_PATH } from "./constants";

export const routeDefinitions = [
  {
    path: ROOT_URL_PATH,
    component: AppLayoutContainer,
    key: 1,
    routes: []
  }
 ];;

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
