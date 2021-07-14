import React, { useEffect } from "react";
import { ROUTES } from "appConstants";
import { useHistory } from "react-router";
import { useDispatch, useSelector } from "react-redux";
import { calibrationInitiated } from "action-creators/calibrationActionCreators";
import CalibrationComponent from "components/Calibration";

const CalibrationContainer = () => {
  const dispatch = useDispatch();
  const history = useHistory();

  //get login reducer details
  const loginReducer = useSelector((state) => state.loginReducer);
  const loginReducerData = loginReducer.toJS();
  let activeDeckObj = loginReducerData?.decks.find((deck) => deck.isActive);
  const { token, isLoggedIn } = activeDeckObj;

  //get calibration configurations from reducer
  const calibrationReducer = useSelector((state) => state.calibrationReducer);
  const calibrationReducerData = calibrationReducer.toJS();
  const { configs } = calibrationReducerData;

  //api call to get configurations
  useEffect(() => {
    if (token) {
      dispatch(calibrationInitiated(token));
    }
  }, [dispatch, token]);

  if (!isLoggedIn) {
    history.push(ROUTES.splashScreen);
  }

  return <CalibrationComponent configs={configs} />;
};

export default CalibrationContainer;
