import React, { useEffect, useState } from "react";
import { useHistory } from "react-router";
import { useDispatch, useSelector } from "react-redux";
import CalibrationExtractionComponent from "components/CalibrationExtraction";
import { logoutInitiated } from "action-creators/loginActionCreators";
import { runPid } from "action-creators/calibrationActionCreators";
import { DECKNAME } from "appConstants";

const CalibrationExtractionContainer = () => {
  const dispatch = useDispatch();
  const history = useHistory();

  const [showConfirmationModal, setConfirmModal] = useState(false);

  //get login reducer details
  const loginReducer = useSelector((state) => state.loginReducer);
  const loginReducerData = loginReducer.toJS();
  let activeDeckObj = loginReducerData?.decks.find((deck) => deck.isActive);
  const { name, token } = activeDeckObj;

  const pidProgessReducer = useSelector((state) => state.pidProgessReducer);
  const pidProgessReducerData = pidProgessReducer.toJS();
  const { progressStatus, deckName, progress, remainingTime, totalTime } =
    pidProgessReducerData;

  console.log("pidProgessReducerData: ", pidProgessReducerData);

  const handleLogout = () => {
    dispatch(
      logoutInitiated({ deckName: name, token: token, showToast: true })
    );
  };

  const handleBtnClick = () => {
    const deckName =
      name === DECKNAME.DeckA ? DECKNAME.DeckAShort : DECKNAME.DeckBShort;
    dispatch(runPid(token, deckName));
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
      handleBtnClick={handleBtnClick}
      showConfirmationModal={showConfirmationModal}
      progressData={pidProgessReducerData}
      deckName={name}
    />
  );
};

export default React.memo(CalibrationExtractionContainer);
