import React, { useEffect } from "react";
import {
  HashRouter as Router,
  Switch,
  Redirect,
  useHistory
} from "react-router-dom";
import RouteWithSubRoutes from "RouteHelper";
import { useSelector, useDispatch } from "react-redux";
import AppHeader from "components/AppHeader";
import styled from "styled-components";

import "../assets/scss/default.scss";
import { fetchActiveWells } from "action-creators/activeWellActionCreators";
import { getActiveLoadedWellFlag } from "selectors/activeWellSelector";
import { connectSocket } from "web-socket";
import ModalContainer from "./ModalContainer";
import { useLocation } from "react-router-dom";
import AppFooter from "components/AppFooter";
import { APP_TYPE, ROUTES } from "appConstants";

export const CardOverlay = styled.div`
  position: absolute;
  // display: none;
  width: 100%;
  height: 100%;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background-color: rgba(0, 0, 0, 0.28);
  z-index: 3;
  cursor: pointer;
`;

/**
 * AppLayoutContainer Will contain routes(content), headers, sub-headers, notification etc.
 * @param {*} props
 */
const AppLayoutContainer = (props) => {
  const { routes } = props;
  const dispatch = useDispatch();
  const history = useHistory();
  const loginReducer = useSelector((state) => state.loginReducer);
  const loginReducerData = loginReducer.toJS();
  const activeDeckObj =
    loginReducerData && loginReducerData.decks.find((deck) => deck.isActive);

  const isActiveWellDataLoaded = useSelector(getActiveLoadedWellFlag);
  const socketReducer = useSelector((state) => state.socketReducer);
  const isOpened = socketReducer.get("isOpened");

  const appInfoReducer = useSelector((state) => state.appInfoReducer);
  const appInfoData = appInfoReducer.toJS();
  const app = appInfoData?.appInfo?.app;

  // connect to websocket on mount
  useEffect(() => {
    // if connection is already opened
    if (isOpened === false) {
      connectSocket(dispatch);
    }
  }, [isOpened, dispatch]);

  useEffect(() => {
    if (
      loginReducer.get("isLoginTypeOperator") === true &&
      isActiveWellDataLoaded === false
    ) {
      dispatch(fetchActiveWells());
    }
  }, [loginReducer, isActiveWellDataLoaded, dispatch]);

  // if app is undefined, then we redirect to splashscreen
  useEffect(() => {
    if (!app) {
      history.push(ROUTES.splashScreen);
    }
  }, []);

  const location = useLocation();

  //recipe reducer
  const recipeActionReducer = useSelector((state) => state.recipeActionReducer);
  let recipeActionReducerData = recipeActionReducer.decks.find(
    (deckObj) => deckObj.name === activeDeckObj.name
  );

  //cleanUp reducer
  const cleanUpReducer = useSelector((state) => state.cleanUpReducer);
  let cleanUpReducerData = cleanUpReducer.decks.find(
    (deckObj) => deckObj.name === activeDeckObj.name
  );

  return (
    <Router>
      {(cleanUpReducerData.showCleanUp ||
        recipeActionReducerData.showProcess) && <CardOverlay />}
      {location.pathname === "/splashscreen" ? null : (
        <AppHeader
          isPlateRoute={loginReducer.get("isPlateRoute")}
          isUserLoggedIn={activeDeckObj.isLoggedIn} //{loginReducer.get("isUserLoggedIn")}
          isLoginTypeAdmin={activeDeckObj.isAdmin} //{loginReducer.get("isLoginTypeAdmin")}
          isLoginTypeOperator={!activeDeckObj.isAdmin} //{loginReducer.get("isLoginTypeOperator")}
          isLoginTypeEngineer={activeDeckObj.isEngineer}
          isTemplateRoute={loginReducer.get("isTemplateRoute")}
          token={activeDeckObj.token}
          deckName={activeDeckObj.name}
          app={app}
        />
      )}
      {/* Modal container will helps in displaying error/info/warning through modal */}
      <ModalContainer />
      <section className="ml-content">
        <Switch>
          <Redirect exact from="/" to={`/${ROUTES.splashScreen}`} />
          {routes.map((route) => (
            <RouteWithSubRoutes key={route.key} {...route} />
          ))}
        </Switch>
      </section>
      {/**dont show appFooter in rtpcr flow, login and splashscreen */}
      {location.pathname === `/${ROUTES.splashScreen}` ||
      location.pathname === `/${ROUTES.login}` ||
      app === APP_TYPE.RTPCR ? null : (
        <AppFooter />
      )}
    </Router>
  );
};

AppLayoutContainer.propTypes = {};

export default AppLayoutContainer;
