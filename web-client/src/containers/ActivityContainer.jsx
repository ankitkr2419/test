import React, { useEffect, useState } from "react";
import ActivityComponent from "components/ActivityLog";
import { useSelector, useDispatch } from "react-redux";
import {
  activityLogInitiated,
} from "action-creators/activityLogActionCreators";
import { toast } from "react-toastify";
import { ROUTES, TOAST_MESSAGE } from "appConstants";
import { useHistory } from "react-router";
import {
  createExperiment,
  createExperimentSucceeded,
  fetchExperiments,
} from "action-creators/experimentActionCreators";
import { loginReset } from "action-creators/loginActionCreators";

const ActivityContainer = () => {
  const dispatch = useDispatch();
  const history = useHistory();

  //get login reducer details
  const loginReducer = useSelector((state) => state.loginReducer);
  const loginReducerData = loginReducer.toJS();
  let activeDeckObj = loginReducerData?.decks.find((deck) => deck.isActive);
  const { token } = activeDeckObj;

  //get activity logs from reducer
  const activityLogReducer = useSelector((state) => state.activityLogReducer);
  const activityLogData = activityLogReducer.toJS();
  const { activityLogs } = activityLogData;

  const createExperimentReducer = useSelector(
    (state) => state.createExperimentReducer
  );
  const createExperimentReducerData = createExperimentReducer.toJS();

  //search activity by experiment name
  const [searchText, setSearchText] = useState("");

  // reset reducers if activity log tab is clicked
  useEffect(() => {
    dispatch(loginReset());
  }, []);

  //get activity list api call
  useEffect(() => {
    dispatch(activityLogInitiated(token));
  }, []);

  const onSearchTextChanged = (text) => {
    setSearchText(text);
  };

  const expandLogHandler = (experimentDetails) => {
    // Call experiment success action to populate reducer.
    dispatch(
      createExperimentSucceeded({ ...experimentDetails, isExpanded: true })
    );
    history.push(ROUTES.plate);
  };

  return (
    <ActivityComponent
      experiments={activityLogs}
      searchText={searchText}
      onSearchTextChanged={onSearchTextChanged}
      expandLogHandler={expandLogHandler}
    />
  );
};

ActivityContainer.propTypes = {};

export default ActivityContainer;
