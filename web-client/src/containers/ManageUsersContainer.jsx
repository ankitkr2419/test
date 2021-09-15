import React from "react";
import { useDispatch, useSelector } from "react-redux";
import { useHistory } from "react-router";

import { APP_TYPE, ROUTES } from "appConstants";

import {
  createUserInitiated,
  deleteUserInitiated,
  updateUserInitiated,
} from "action-creators/usersActionCreators";
import ManageUsersComponent from "components/ManageUsersComponent";

const ManageUsersContainer = () => {
  const history = useHistory();
  const dispatch = useDispatch();

  //get login reducer details
  const loginReducer = useSelector((state) => state.loginReducer);
  const loginReducerData = loginReducer.toJS();
  let activeDeckObj = loginReducerData?.decks.find((deck) => deck.isActive);
  const { token, isLoggedIn } = activeDeckObj;

  //get app type
  const appInfoReducer = useSelector((state) => state.appInfoReducer);
  const appInfoData = appInfoReducer.toJS();
  const app = appInfoData?.appInfo?.app;

  //redirect if not logged in
  if (!isLoggedIn) {
    if (app === APP_TYPE.EXTRACTION) {
      history.push(ROUTES.landing);
    } else if (app === APP_TYPE.RTPCR) {
      history.push(ROUTES.splashScreen);
    }
  }

  const handleCreateUser = (userData) => {
    dispatch(createUserInitiated(token, userData));
  };

  const handleDeleteUser = (username) => {
    dispatch(deleteUserInitiated(token, username));
  };

  const handleUpdateUser = (userData) => {
    let oldUsername = userData.oldUsername;
    let updatedUserData = {
      username: userData.username,
      password: userData.password,
      role: userData.role,
    };

    dispatch(updateUserInitiated(token, oldUsername, updatedUserData));
  };

  return (
    <ManageUsersComponent
      handleCreateUser={handleCreateUser}
      handleDeleteUser={handleDeleteUser}
      handleUpdateUser={handleUpdateUser}
    />
  );
};

export default React.memo(ManageUsersContainer);
