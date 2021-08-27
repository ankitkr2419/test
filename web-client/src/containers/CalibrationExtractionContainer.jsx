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
  motorInitiated,
  runPid,
  updateCommonDetailsInitiated,
} from "action-creators/calibrationActionCreators";
import { DECKNAME, PID_STATUS } from "appConstants";
import { useFormik } from "formik";
import { formikInitialState } from "components/CalibrationExtraction/helpers";

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

  const pidProgessReducer = useSelector((state) => state.pidProgessReducer);
  const pidProgessReducerData = pidProgessReducer.toJS();
  const progressData = pidProgessReducerData.decks.find(
    (deckObj) => deckObj.deckName === name
  );

  const pidReducer = useSelector((state) => state.pidReducer);
  const { pidStatus } = pidReducer.toJS();

  const commonDetailsReducer = useSelector(
    (state) => state.commonDetailsReducer
  );
  const commonDetailsReducerData = commonDetailsReducer.toJS();
  const { error, isLoading, isUpdateApi, details } = commonDetailsReducerData;

  // fetch API called initially
  useEffect(() => {
    dispatch(commonDetailsInitiated(token));
  }, []);

  useEffect(() => {
    // populate formik data with fetched values
    if (error === false && isLoading === false && isUpdateApi === false) {
      const { receiver_name, receiver_email, room_temperature } = details;
      formik.setFieldValue("name.value", receiver_name);
      formik.setFieldValue("email.value", receiver_email);
      formik.setFieldValue("roomTemperature.value", room_temperature);
    }

    // API call to fetch new data after updation
    if (error === false && isLoading === false && isUpdateApi === true) {
      dispatch(commonDetailsInitiated(token));
    }
  }, [error, isLoading, isUpdateApi]);

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

  const toggleConfirmModal = () => setConfirmModal(!showConfirmationModal);

  return (
    <CalibrationExtractionComponent
      toggleConfirmModal={toggleConfirmModal}
      handleLogout={handleLogout}
      handleBtnClick={handlePidBtn}
      handleMotorBtn={handleMotorBtn}
      handleSaveDetailsBtn={handleSaveDetailsBtn}
      showConfirmationModal={showConfirmationModal}
      progressData={progressData}
      pidStatus={pidStatus}
      deckName={name}
      formik={formik}
      isAdmin={isAdmin}
    />
  );
};

export default React.memo(CalibrationExtractionContainer);
