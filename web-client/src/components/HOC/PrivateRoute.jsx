import { ROUTES } from "appConstants";
import React, { useEffect } from "react";
import { useSelector } from "react-redux";
import { Redirect } from "react-router";

const privateRoute = (Component) => (props) => {
  const loginReducer = useSelector((state) => state.loginReducer);
  const loginReducerData = loginReducer.toJS();
  const activeDeckObj =
    loginReducerData && loginReducerData.decks.find((deck) => deck.isActive);
  const { isLoggedIn, isAdmin } = activeDeckObj;

  if (!isLoggedIn) {
    return <Redirect to={`/${ROUTES.login}`} />;
  }

  return (
    <Component
      {...props}
      isLoginTypeAdmin={isAdmin}
      isLoginTypeOperator={!isAdmin}
    />
  );
};
export default privateRoute;
