import React, { useEffect, useState } from "react";
import { useDispatch, useSelector } from "react-redux";
import { MlModal } from "shared-components";
import {
  hideHomingModal,
  homingActionInitiated,
} from "action-creators/homingActionCreators";
import { DECKNAME, MODAL_BTN, MODAL_MESSAGE } from "appConstants";
import {
  abortReset,
  runPidReset,
} from "action-creators/calibrationActionCreators";

const HomingModal = (props) => {
  const loginReducer = useSelector((state) => state.loginReducer);
  const loginReducerData = loginReducer.toJS();
  let activeDeckObj =
    loginReducerData && loginReducerData.decks.find((deck) => deck.isActive);
  const { name, token } = activeDeckObj;

  const dispatch = useDispatch();

  const [isProgressBarVisible, setIsProgressBarVisible] = useState(false);
  const [modalBtnIsDisabled, setModalBtnDisable] = useState(false);

  const homingReducer = useSelector((state) => state.homingReducer);

  const {
    showHomingModal,
    isHomingActionCompleted,
    homingAllDeckCompletionPercentage,
  } = homingReducer;

  // if homing is successfully completed then enable complete btn
  useEffect(() => {
    if (isHomingActionCompleted) {
      setModalBtnDisable(false);
    }
  }, [isHomingActionCompleted]);

  const homingConfirmation = () => {
    const deckName =
      name === DECKNAME.DeckA ? DECKNAME.DeckAShort : DECKNAME.DeckBShort;
    dispatch(homingActionInitiated({ deckName: deckName, token: token }));
    setIsProgressBarVisible(!isProgressBarVisible);
    setModalBtnDisable(!modalBtnIsDisabled);
  };

  const handleCompleteBtn = () => {
    dispatch(hideHomingModal());

    // reset abort status of PID, Shaker and Heater
    dispatch(abortReset());

    // reset local states
    setIsProgressBarVisible(false);
    setModalBtnDisable(false);
  };

  return (
    <MlModal
      isOpen={showHomingModal}
      textBody={MODAL_MESSAGE.homingConfirmation}
      handleSuccessBtn={() =>
        isHomingActionCompleted ? handleCompleteBtn() : homingConfirmation()
      }
      successBtn={isHomingActionCompleted ? MODAL_BTN.complete : MODAL_BTN.okay}
      showCrossBtn={false}
      progressPercentage={homingAllDeckCompletionPercentage}
      isProgressBarVisible={isProgressBarVisible}
      disabled={modalBtnIsDisabled}
    />
  );
};

export default HomingModal;
