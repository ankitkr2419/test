import React from "react";
import { HashRouter as Router, Switch } from "react-router-dom";
import { withRouter } from "react-router";

import RouteWithSubRoutes from "RouteHelper";
import AppLayoutContainer from "containers/AppLayoutContainer";
import PlateContainer from "containers/PlateContainer";
import ActivityContainer from "containers/ActivityContainer";
import TemplateLayout from "layouts/TemplateLayout";
import PrivateRoute from "components/HOC/PrivateRoute";
import SplashScreenContainer from "containers/SplashScreenContainer";
import LandingPageContainer from "containers/LandingPageContainer";
import RecipeListingContainer from "containers/RecipeListingContainer";
import AllSetContainer from "containers/AllSetContainer";
import LabwareContainer from "containers/LabwareContainer";
import ProcessListingContainer from "containers/ProcessListingContainer";
import { ROOT_URL_PATH, ROUTES } from "./appConstants";
import SelectProcessContainer from "containers/SelectProcessContainer";
import PiercingContainer from "containers/PiercingContainer";
import TipPickupContainer from "containers/TipPickupContainer";
import AspireDispenseContainer from "containers/AspireDispenseContainer";
import ShakingContainer from "containers/ShakingContainer";
import HeatingContainer from "containers/HeatingContainer";
import MagnetContainer from "containers/MagnetContainer";
import DelayContainer from "containers/DelayContainer";
import TipDiscardContainer from "containers/TipDiscardContainer";

export const routeDefinitions = [
  {
    path: ROOT_URL_PATH,
    component: AppLayoutContainer,
    key: 1,
    routes: [
      {
        path: `${ROOT_URL_PATH}splashscreen`,
        exact: true,
        component: SplashScreenContainer,
        key: 6,
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
      {
        path: `${ROOT_URL_PATH}recipe-listing`,
        exact: true,
        component: RecipeListingContainer,
        key: 9,
      },
      {
        path: `${ROOT_URL_PATH}labware`,
        exact: true,
        component: LabwareContainer,
        key: 10,
      },
      {
        path: `${ROOT_URL_PATH}all-set`,
        exact: true,
        component: AllSetContainer,
        key: 11,
      },
      {
        path: `${ROOT_URL_PATH}select-process`,
        exact: true,
        component: SelectProcessContainer,
        key: 18,
      },
      {
        path: `${ROOT_URL_PATH}piercing`,
        exact: true,
        component: PiercingContainer,
        key: 13,
      },
      {
        path: `${ROOT_URL_PATH}shaking`,
        exact: true,
        component: ShakingContainer,
        key: 14,
      },
      {
        path: `${ROOT_URL_PATH}heating`,
        exact: true,
        component: HeatingContainer,
        key: 15,
      },
      // {
      //   path: `${ROOT_URL_PATH}tip-position`,
      //   exact: true,
      //   component: TipPositionContainer,
      //   key: 16,
      // },
      {
        path: `${ROOT_URL_PATH}magnet`,
        exact: true,
        component: MagnetContainer,
        key: 17,
      },
      {
        path: `${ROOT_URL_PATH}tip-discard`,
        exact: true,
        component: TipDiscardContainer,
        key: 18,
      },
      {
        path: `${ROOT_URL_PATH}aspire-dispense`,
        exact: true,
        component: AspireDispenseContainer,
        key: 19,
      },
      {
        path: `${ROOT_URL_PATH}tip-pickup`,
        exact: true,
        component: TipPickupContainer,
        key: 20,
      },
      {
        path: `${ROOT_URL_PATH}${ROUTES.processListing}`,
        exact: true,
        component: ProcessListingContainer,
        key: 22,
      },
      {
        path: `${ROOT_URL_PATH}delay`,
        exact: true,
        component: DelayContainer,
        key: 21,
      },
    ],
  },
];

const Routes = (props) => {
  const routesList = props.routesList || routeDefinitions;
  return (
    <Router>
      <Switch>
        {routesList.map((route) => (
          <RouteWithSubRoutes key={route.key} {...route} />
        ))}
      </Switch>
    </Router>
  );
};

export default withRouter(Routes);
