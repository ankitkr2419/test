import React, { useEffect, useState } from "react";
import ActivityComponent from "components/ActivityLog";
import { useSelector, useDispatch } from "react-redux";
import { activityLogInitiated } from "action-creators/activityLogActionCreators";

const ActivityContainer = () => {
  const dispatch = useDispatch();

  //get login reducer details
  const loginReducer = useSelector((state) => state.loginReducer);
  const loginReducerData = loginReducer.toJS();
  let activeDeckObj = loginReducerData?.decks.find((deck) => deck.isActive);
  const { token } = activeDeckObj;

  //get activity logs from reducer
  const activityLogReducer = useSelector((state) => state.activityLogReducer);
  const activityLogData = activityLogReducer.toJS();
  const { activityLogs } = activityLogData;

  //search activity by experiment name
  const [searchText, setSearchText] = useState("");

  //get activity list api call
  useEffect(() => {
    dispatch(activityLogInitiated(token));
  }, []);

  const onSearchTextChanged = (text) => {
    setSearchText(text);
  };

  return (
    <ActivityComponent
      experiments={activityLogs}
      searchText={searchText}
      onSearchTextChanged={onSearchTextChanged}
    />
  );
};

ActivityContainer.propTypes = {};

export default ActivityContainer;
