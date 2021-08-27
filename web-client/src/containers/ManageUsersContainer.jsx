import React, { useEffect, useState } from "react";
import { useDispatch, useSelector } from "react-redux";

import {
  createUserInitiated,
  deleteUserInitiated,
  updateUserInitiated,
} from "action-creators/usersActionCreators";
import ManageUsersComponent from "components/ManageUsersComponent";

const ManageUsersContainer = () => {
  const dispatch = useDispatch();

  //get login reducer details
  const loginReducer = useSelector((state) => state.loginReducer);
  const loginReducerData = loginReducer.toJS();
  let activeDeckObj = loginReducerData?.decks.find((deck) => deck.isActive);
  const { token } = activeDeckObj;

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
