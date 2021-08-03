import React, { useEffect, useState } from "react";
import ActivityComponent from "components/ActivityLog";
import { useSelector, useDispatch } from "react-redux";
import {
  activityLogInitiated,
  mailReportInitiated,
} from "action-creators/activityLogActionCreators";
import { toast } from "react-toastify";
import { TOAST_MESSAGE } from "appConstants";

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

  //get status of mail from reducer
  const mailReportReducer = useSelector((state) => state.mailReportReducer);
  const { isLoading, error } = mailReportReducer.toJS();

  //search activity by experiment name
  const [searchText, setSearchText] = useState("");

  //get activity list api call
  useEffect(() => {
    dispatch(activityLogInitiated(token));
  }, []);

  // check if mail is sent or not and show toast msg acc.
  useEffect(() => {
    if (isLoading === false && error === true) {
      toast.success(TOAST_MESSAGE.sendingMailSuccess);
    } else if (isLoading === false && error === false) {
      toast.success(TOAST_MESSAGE.sendingMailSuccess);
    }
  }, [isLoading, error]);

  const onSearchTextChanged = (text) => {
    setSearchText(text);
  };

  const mailActivityReportHandler = () => {
    // dispatch(mailReportInitiated({ report, token })); //API call to send an email
  };

  return (
    <ActivityComponent
      experiments={activityLogs}
      searchText={searchText}
      onSearchTextChanged={onSearchTextChanged}
      mailActivityReportHandler={mailActivityReportHandler}
    />
  );
};

ActivityContainer.propTypes = {};

export default ActivityContainer;
