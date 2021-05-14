import React from "react";

import LandingScreenComponent from "components/LandingScreen";
import { useSelector } from "react-redux";
import { Redirect } from "react-router";
import { ROUTES } from "../appConstants";

const LandingPageContainer = (props) => {
  const loginReducer = useSelector((state) => state.loginReducer);

  const loginReducerData = loginReducer.toJS()
  let activeDeckObj = loginReducerData && loginReducerData.decks.find(deck => deck.isActive)
  let deckName  = activeDeckObj ? activeDeckObj.name : ''
  let { isLoggedIn, error } = activeDeckObj ? activeDeckObj : {};
 
  /**
   * if user logged in, go to recipeListing page
   */
  if (isLoggedIn && !error) {
    return <Redirect to={`/${ROUTES.recipeListing}`} />;
  }

  return <LandingScreenComponent deckName={deckName} />;
};

LandingPageContainer.propTypes = {};

export default LandingPageContainer;
