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
import { logoutInitiated } from "action-creators/loginActionCreators";

const LogoutComponent = (props) => {
  const { handleLogout, setConfirmModal } = props;
  const loginReducer = useSelector((state) => state.loginReducer);
  const loginReducerData = loginReducer.toJS();
  let activeDeckObj =
    loginReducerData && loginReducerData.decks.find((deck) => deck.isActive);
  const { name, token } = activeDeckObj;
  const [logout, setLogout] = useState(false);

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
    if (isHomingActionCompleted && logout) {
      setModalBtnDisable(false);
      handleLogout();
      dispatch(hideHomingModal());

      // reset abort status of UV, PID, Shaker and Heater
      dispatch(abortReset());

      // reset local states of modals
      setIsProgressBarVisible(false);
      setModalBtnDisable(false);
    }
    // setLogout(false);
  }, [isHomingActionCompleted]);

  const LogoutWithHoming = () => {
    setLogout(true);
    const deckName =
      name === DECKNAME.DeckA ? DECKNAME.DeckAShort : DECKNAME.DeckBShort;
    dispatch(homingActionInitiated({ deckName: deckName, token: token }));
    setIsProgressBarVisible(!isProgressBarVisible);
    setModalBtnDisable(!modalBtnIsDisabled);
  };

  const LogoutWithoutHoming = () => {
    handleLogout();
    dispatch(hideHomingModal());
    dispatch(abortReset());
    setIsProgressBarVisible(false);
    setModalBtnDisable(false);
  };

  const cancelFunc = () => {
    setConfirmModal(false);
    dispatch(hideHomingModal());
    dispatch(abortReset());
    setIsProgressBarVisible(false);
    setModalBtnDisable(false);
  };

  return (
    <MlModal
      isOpen={showHomingModal}
      textBody={MODAL_MESSAGE.logoutConformation}
      handleSuccessBtn={() => LogoutWithHoming()}
      handleCrossBtn={() => LogoutWithoutHoming()}
      successBtn={MODAL_BTN.withHoming}
      failureBtn={MODAL_BTN.withoutHoming}
      progressPercentage={homingAllDeckCompletionPercentage}
      isProgressBarVisible={isProgressBarVisible}
      disabled={modalBtnIsDisabled}
      senseAndHit={modalBtnIsDisabled}
      isLogout={true}
      CancelFunc={cancelFunc}
      isWhiteLightBtn={true}
    />
  );
};

export default LogoutComponent;
