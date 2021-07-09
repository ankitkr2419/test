import React, { useEffect } from "react";
import { HashRouter as Router, Switch, Redirect } from "react-router-dom";
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
  const loginReducer = useSelector((state) => state.loginReducer);
  const isActiveWellDataLoaded = useSelector(getActiveLoadedWellFlag);
  const socketReducer = useSelector((state) => state.socketReducer);
  const isOpened = socketReducer.get("isOpened");

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

  const location = useLocation();

  const loginReducerData = loginReducer.toJS();
  let activeDeckObj = loginReducerData?.decks.find((deck) => deck.isActive);

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
          isUserLoggedIn={activeDeckObj.isLoggedIn}
          isLoginTypeAdmin={activeDeckObj.isAdmin}
          isLoginTypeOperator={!activeDeckObj.isAdmin}
          isTemplateRoute={loginReducer.get("isTemplateRoute")}
          currentDeckName={activeDeckObj.name}
        />
      )}
      {/* Modal container will helps in displaying error/info/warning through modal */}
      <ModalContainer />
      <section className="ml-content">
        <Switch>
          <Redirect exact from="/" to="/splashscreen" />
          {routes.map((route) => (
            <RouteWithSubRoutes key={route.key} {...route} />
          ))}
        </Switch>
      </section>
      {location.pathname === "/splashscreen" ? null : <AppFooter />}
    </Router>
  );
};

AppLayoutContainer.propTypes = {};

export default AppLayoutContainer;
