import React, { useEffect } from "react";
import { ROUTES } from "appConstants";
import { useHistory } from "react-router";
import { useSelector } from "react-redux";

const CalibrationContainer = () => {
  const history = useHistory();

  //get login reducer details
  const loginReducer = useSelector((state) => state.loginReducer);
  const loginReducerData = loginReducer.toJS();
  let activeDeckObj = loginReducerData?.decks.find((deck) => deck.isActive);
  const { isLoggedIn } = activeDeckObj;

  //api call to get configurations 
  useEffect(() => {
    //TODO
  }, []);

  if (!isLoggedIn) {
    history.push(ROUTES.splashScreen);
  }

  return <div>Container...</div>;
  // return <CalibrationComponent />;//TODO
};

export default CalibrationContainer;
