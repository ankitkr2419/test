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
  senseAndHitInitiated,
} from "action-creators/calibrationActionCreators";

const HomingWithSenseAndHitComponent = (props) => {
  const { formik, setSenseHitHomingModel, motorSelected, setMotorSelected } =
    props;
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
    const deck =
      name === DECKNAME.DeckA ? DECKNAME.DeckAShort : DECKNAME.DeckBShort;
    motorSelected.forEach((motorNumber) => {
      const body = {
        motor_number: motorNumber,
      };
      dispatch(senseAndHitInitiated(token, deck, body));
    });
    setMotorSelected([]);
    const deckName =
      name === DECKNAME.DeckA ? DECKNAME.DeckAShort : DECKNAME.DeckBShort;
    dispatch(homingActionInitiated({ deckName: deckName, token: token }));
    setIsProgressBarVisible(!isProgressBarVisible);
    setModalBtnDisable(!modalBtnIsDisabled);
  };

  const handleCompleteBtn = () => {
    dispatch(hideHomingModal());

    // reset abort status of UV, PID, Shaker and Heater
    dispatch(abortReset());

    // reset local states of modals
    setIsProgressBarVisible(false);
    setModalBtnDisable(false);
    setSenseHitHomingModel(false);
  };

  const msgDisplay = () => {
    if (isHomingActionCompleted) {
      return "Homing Completed";
    } else {
      return modalBtnIsDisabled
        ? "Homing in Progress"
        : MODAL_MESSAGE.senseAndHitHomingMsg;
    }
  };

  return (
    <MlModal
      isOpen={showHomingModal}
      textBody={msgDisplay()}
      handleSuccessBtn={() =>
        isHomingActionCompleted ? handleCompleteBtn() : homingConfirmation()
      }
      handleCrossBtn={() => {
        setSenseHitHomingModel(false);
        dispatch(hideHomingModal());
      }}
      successBtn={isHomingActionCompleted ? MODAL_BTN.complete : MODAL_BTN.okay}
      failureBtn={MODAL_BTN.cancel}
      showCrossBtn={false}
      progressPercentage={homingAllDeckCompletionPercentage}
      isProgressBarVisible={isProgressBarVisible}
      disabled={modalBtnIsDisabled}
      senseAndHit={modalBtnIsDisabled}
      isWhiteLightBtn={true}
    />
  );
};

export default HomingWithSenseAndHitComponent;
