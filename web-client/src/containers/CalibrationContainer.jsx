import React, { useEffect } from "react";
import { ROUTES } from "appConstants";
import { useHistory } from "react-router";
import { useDispatch, useSelector } from "react-redux";
import {
  calibrationInitiated,
  updateCalibrationInitiated,
} from "action-creators/calibrationActionCreators";
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

  const saveBtnClickHandler = (configData) => {
    let data = {
      room_temperature: configData.roomTemperature,
      homing_time: configData.homingTime,
      no_of_homing_cycles: configData.noOfHomingCycles,
      cycle_time: configData.cycleTime,
    };
    dispatch(updateCalibrationInitiated({ token, data }));
  };

  if (!isLoggedIn) {
    history.push(ROUTES.splashScreen);
  }

  return (
    <CalibrationComponent
      configs={configs}
      saveBtnClickHandler={saveBtnClickHandler}
    />
  );
};

export default CalibrationContainer;
