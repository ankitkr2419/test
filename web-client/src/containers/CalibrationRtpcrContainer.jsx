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
      pid_temperature: configData.pidTemperature,
      pid_minutes: configData.pidMinutes,
    };
    dispatch(updateCalibrationInitiated({ token, data }));
    //populate with new data
    dispatch(calibrationInitiated(token));
  };

  return (
    <CalibrationComponent
      configs={configs}
      saveBtnClickHandler={saveBtnClickHandler}
    />
  );
};

export default React.memo(CalibrationRtpcrContainer);
