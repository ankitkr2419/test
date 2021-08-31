import React, { useEffect } from "react";
import { useDispatch, useSelector } from "react-redux";
import {
  calibrationInitiated,
  updateCalibrationInitiated,
} from "action-creators/calibrationActionCreators";
import CalibrationComponent from "components/Calibration";

const CalibrationRtpcrContainer = () => {
  const dispatch = useDispatch();

  //get login reducer details
  const loginReducer = useSelector((state) => state.loginReducer);
  const loginReducerData = loginReducer.toJS();
  let activeDeckObj = loginReducerData?.decks.find((deck) => deck.isActive);
  const { token } = activeDeckObj;

  //get calibration configurations from reducer
  const calibrationReducer = useSelector((state) => state.calibrationReducer);
  const calibrationReducerData = calibrationReducer.toJS();
  const { configs } = calibrationReducerData;

  //get calibration configurations from reducer
  const updateCalibrationReducer = useSelector((state) => state.updateCalibrationReducer);
  const updateCalibrationReducerData = updateCalibrationReducer.toJS();
  const { isLoading, error } = updateCalibrationReducerData;

  //initially populate with previous data
  useEffect(() => {
    if (token) {
      dispatch(calibrationInitiated(token)); 
    }
  }, [dispatch, token])

  //after update API call is successful populate fields with new data
  useEffect(() => {
    if (token && error === false && isLoading === false) {
      dispatch(calibrationInitiated(token));
    }
  }, [dispatch, isLoading, error, token]);

  const saveButtonClickHandler = (configData) => {
    let data = {
      receiver_name: configData.name,
      receiver_email: configData.email,
      room_temperature: configData.roomTemperature,
      homing_time: configData.homingTime,
      no_of_homing_cycles: configData.noOfHomingCycles,
      cycle_time: configData.cycleTime,
      pid_temperature: configData.pidTemperature,
      pid_minutes: configData.pidMinutes,
    };

    dispatch(updateCalibrationInitiated({ token, data }));
  };

  return (
    <CalibrationComponent
      configs={configs}
      saveButtonClickHandler={saveButtonClickHandler}
    />
  );
};

export default React.memo(CalibrationRtpcrContainer);
