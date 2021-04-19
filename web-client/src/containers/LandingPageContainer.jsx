import React from "react";

import LandingScreenComponent from "components/LandingScreen";
import { useSelector } from "react-redux";

const LandingPageContainer = (props) => {
  const loginReducer = useSelector(
    (state) => state.loginReducer
  );

  const loginReducerData = loginReducer.toJS()
  let activeDeckObj = loginReducerData && loginReducerData.decks.find(deck => deck.isActive)
  let deckName  = activeDeckObj ? activeDeckObj.name : ''

  return <LandingScreenComponent deckName={deckName} />;
};

LandingPageContainer.propTypes = {};

export default LandingPageContainer;
