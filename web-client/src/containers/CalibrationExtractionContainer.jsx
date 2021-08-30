import React, { useEffect, useState } from "react";
import { useHistory } from "react-router";
import { useDispatch, useSelector } from "react-redux";
import CalibrationExtractionComponent from "components/CalibrationExtraction";
import {
  deckBlockInitiated,
  logoutInitiated,
} from "action-creators/loginActionCreators";
import {
  abortPid,
  commonDetailsInitiated,
  fetchPidInitiated,
  motorInitiated,
  runPid,
  updateCommonDetailsInitiated,
  updatePidInitiated,
} from "action-creators/calibrationActionCreators";
import { DECKNAME, PID_STATUS } from "appConstants";
import { useFormik } from "formik";
import {
  formikInitialState,
  formikToArray,
} from "components/CalibrationExtraction/helpers";

const CalibrationExtractionContainer = () => {
  const dispatch = useDispatch();

  const [showConfirmationModal, setConfirmModal] = useState(false);

  const formik = useFormik({
    initialValues: formikInitialState,
    enableReinitialize: true,
  });

  //get login reducer details
  const loginReducer = useSelector((state) => state.loginReducer);
  const loginReducerData = loginReducer.toJS();
  let activeDeckObj = loginReducerData?.decks.find((deck) => deck.isActive);
  const { isAdmin, name, token } = activeDeckObj;

  const heaterReducer = useSelector((state) => state.heaterProgressReducer);
  const heaterProgressReducerData = heaterReducer.toJS();
  const { data } = heaterProgressReducerData;

  const pidProgessReducer = useSelector((state) => state.pidProgessReducer);
  const pidProgessReducerData = pidProgessReducer.toJS();
  const progressData = pidProgessReducerData.decks.find(
    (deckObj) => deckObj.deckName === name
  );

  const pidReducer = useSelector((state) => state.pidReducer);
  const pidReducerData = pidReducer.toJS();
  const { pidStatus, pidData, isPidUpdateApi } = pidReducerData;

  const commonDetailsReducer = useSelector(
    (state) => state.commonDetailsReducer
  );
  const commonDetailsReducerData = commonDetailsReducer.toJS();
  const { isUpdateApi, details } = commonDetailsReducerData;

  // fetch pidDetails API (pidTemp, pidMinutes) called initially
  useEffect(() => {
    dispatch(fetchPidInitiated(token));
  }, []);

  // fetch commonDetails (name, email, roomTemp) API called initially
  useEffect(() => {
    dispatch(commonDetailsInitiated(token));
  }, []);

  useEffect(() => {
    if (
      commonDetailsReducerData.error === false &&
      commonDetailsReducerData.isLoading === false
    ) {
      // populate formik data with fetched values
      if (isUpdateApi === false) {
        const { receiver_name, receiver_email, room_temperature } = details;
        formik.setFieldValue("name.value", receiver_name);
        formik.setFieldValue("email.value", receiver_email);
        formik.setFieldValue("roomTemperature.value", room_temperature);
      } else {
        // fetch updated data after updation
        dispatch(commonDetailsInitiated(token));
      }
    }
  }, [
    commonDetailsReducerData.error,
    commonDetailsReducerData.isLoading,
    isUpdateApi,
  ]);

  useEffect(() => {
    if (pidReducerData.error === false && pidReducerData.isLoading === false) {
      if (isPidUpdateApi === false) {
        // populate formik data with fetched values
        formik.setFieldValue("pidTemperature.value", pidData.pid_temperature);
      } else {
        // fetch updated data after updation
        dispatch(fetchPidInitiated(token));
      }
    }
  }, [pidReducerData.error, pidReducerData.isLoading, isPidUpdateApi]);

  /**another deck must be blocked**/
  useEffect(() => {
    dispatch(deckBlockInitiated({ deckName: name }));
  }, []);

  const handleLogout = () => {
    dispatch(
      logoutInitiated({ deckName: name, token: token, showToast: true })
    );
  };

  const handlePidBtn = () => {
    const deckName =
      name === DECKNAME.DeckA ? DECKNAME.DeckAShort : DECKNAME.DeckBShort;
    if (pidStatus === PID_STATUS.running) {
      // dispatch abort API if progressing
      dispatch(abortPid(token, deckName));
    } else {
      // dispatch run PID progressing
      dispatch(runPid(token, deckName));
    }
  };

  const handleMotorBtn = (e) => {
    e.preventDefault();

    const { motorNumber, direction, distance } = formik.values;

    const body = {
      deck: name === DECKNAME.DeckA ? DECKNAME.DeckAShort : DECKNAME.DeckBShort,
      motor_number: motorNumber.value,
      direction: direction.value,
      distance: distance.value,
    };

    dispatch(motorInitiated(token, body));
  };

  const handleSaveDetailsBtn = (data) => {
    const { name, email, roomTemperature } = data;
    const requestBody = {
      receiver_name: name.value,
      receiver_email: email.value,
      room_temperature: roomTemperature.value,
    };
    dispatch(updateCommonDetailsInitiated({ token, data: requestBody }));
  };

  const handlePidUpdateBtn = (pidData) => {
    const { pidTemperature } = pidData;
    const requestBody = {
      pid_temperature: pidTemperature.value,
      pid_minutes: 30, // will be removed in future
      micro_lit_pulses: 25, // will be removed in future
      shaker_steps_per_revolution: 800, // will be removed in future
    };
    dispatch(updatePidInitiated(token, requestBody));
  };

  const toggleConfirmModal = () => setConfirmModal(!showConfirmationModal);

  const handleTipesTubesButton = (e) => {
    e.preventDefault();

    let { allowedPositions } = formik.values;
    let arrayOfAllowedPositions = formikToArray(allowedPositions);
    
    //TODO api call under dev
  };

  return (
    <CalibrationExtractionComponent
      toggleConfirmModal={toggleConfirmModal}
      handleLogout={handleLogout}
      handleBtnClick={handlePidBtn}
      handleMotorBtn={handleMotorBtn}
      handleSaveDetailsBtn={handleSaveDetailsBtn}
      handlePidUpdateBtn={handlePidUpdateBtn}
      showConfirmationModal={showConfirmationModal}
      heaterData={data}
      progressData={progressData}
      pidStatus={pidStatus}
      deckName={name}
      formik={formik}
      isAdmin={isAdmin}
      handleTipesTubesButton={handleTipesTubesButton}
    />
  );
};

export default React.memo(CalibrationExtractionContainer);
