import React, { useEffect, useState } from "react";
import { useHistory } from "react-router";
import { useDispatch, useSelector } from "react-redux";
import CalibrationExtractionComponent from "components/CalibrationExtraction";
import { logoutInitiated } from "action-creators/loginActionCreators";
import {
  motorInitiated,
  runPid,
} from "action-creators/calibrationActionCreators";
import { DECKNAME } from "appConstants";
import { useFormik } from "formik";
import { formikInitialState } from "components/CalibrationExtraction/helpers";

const CalibrationExtractionContainer = () => {
  const dispatch = useDispatch();
  const history = useHistory();

  const [showConfirmationModal, setConfirmModal] = useState(false);

  const formik = useFormik({
    initialValues: formikInitialState,
    enableReinitialize: true,
  });

  //get login reducer details
  const loginReducer = useSelector((state) => state.loginReducer);
  const loginReducerData = loginReducer.toJS();
  let activeDeckObj = loginReducerData?.decks.find((deck) => deck.isActive);
  const { name, token } = activeDeckObj;

  const pidProgessReducer = useSelector((state) => state.pidProgessReducer);
  const pidProgessReducerData = pidProgessReducer.toJS();

  const handleLogout = () => {
    dispatch(
      logoutInitiated({ deckName: name, token: token, showToast: true })
    );
  };

  const handlePidBtn = () => {
    const deckName =
      name === DECKNAME.DeckA ? DECKNAME.DeckAShort : DECKNAME.DeckBShort;
    dispatch(runPid(token, deckName));
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

  const toggleConfirmModal = () => setConfirmModal(!showConfirmationModal);

  //api call to get configurations
  // useEffect(() => {
  //   if (token) {
  //     //TODO initial api's if required
  //   }
  // }, [dispatch, token]);

  return (
    <CalibrationExtractionComponent
      toggleConfirmModal={toggleConfirmModal}
      handleLogout={handleLogout}
      handleBtnClick={handlePidBtn}
      handleMotorBtn={handleMotorBtn}
      showConfirmationModal={showConfirmationModal}
      progressData={pidProgessReducerData}
      deckName={name}
      formik={formik}
    />
  );
};

export default React.memo(CalibrationExtractionContainer);
