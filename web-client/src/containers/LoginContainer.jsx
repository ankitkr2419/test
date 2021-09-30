import React from "react";
import { Redirect } from "react-router";
import LoginComponent from "components/Login";
import { useDispatch, useSelector } from "react-redux";
import { login } from "action-creators/loginActionCreators";
import { ROUTES } from "appConstants";

const LoginContainer = () => {
  const dispatch = useDispatch();

  const loginReducer = useSelector((state) => state.loginReducer);
  const loginReducerData = loginReducer.toJS();
  let activeDeckObj = loginReducerData?.decks.find((deck) => deck.isActive);
  const { isLoggedIn, name, isEngineer } = activeDeckObj; //name refers to deckName

  const loginBtnHandler = (data) => {
    let { username, password } = data;
    // let role = username ? username.split("@").pop().split(".")[0] : undefined;
    let email = username; //.split("@")[0];
    //TODO remove comments once tested properly
    dispatch(login({ email, password, deckName: name, showToast: true }));
  };

  // redirection once logged in
  if (isLoggedIn === true) {
    return <Redirect to={isEngineer ? ROUTES.calibration : "/templates"} />;
  }

  return <LoginComponent loginBtnHandler={loginBtnHandler} />;
};

export default LoginContainer;
