import React, { useEffect, useState } from "react";
import { useDispatch, useSelector } from "react-redux";
import { VideoCard, MlModal } from "shared-components";

import {
  MODAL_MESSAGE,
  MODAL_BTN,
  ROUTES,
  CREDS_FOR_HOMING,
} from "appConstants";
import {
  homingActionInitiated,
  hideHomingModal,
} from "action-creators/homingActionCreators";
import { LandingScreen } from "./LandingScreen";
import { Redirect } from "react-router";
import { login, logoutInitiated } from "action-creators/loginActionCreators";

const LandingScreenComponent = (props) => {
  const dispatch = useDispatch();
  const homingReducer = useSelector((state) => state.homingReducer);
  const loginReducer = useSelector((state) => state.loginReducer);
  const loginReducerData = loginReducer.toJS();
  let activeDeckObj =
    loginReducerData && loginReducerData.decks.find((deck) => deck.isActive);

  const { isLoading, isLoggedInForHoming, tokenForHoming } = loginReducerData;
  let { isLoggedIn, error } = activeDeckObj ? activeDeckObj : {};

  const {
    showHomingModal,
    isHomingActionCompleted,
    homingAllDeckCompletionPercentage,
  } = homingReducer;

  const [isProgressBarVisible, setIsProgressBarVisible] = useState(false);
  const [isOpen, setIsOpen] = useState(true);
  const [disabled, setDisabled] = useState(false);

  const homingConfirmation = () => {
    dispatch(homingActionInitiated({ token: tokenForHoming }));
    setIsProgressBarVisible(!isProgressBarVisible);
    setDisabled(!disabled);
  };

  const handleCompleteBtn = () => {
    setIsOpen(false);
    dispatch(hideHomingModal());
  };

  useEffect(() => {
    if (!isHomingActionCompleted) {
      dispatch(login(CREDS_FOR_HOMING));
    }
  }, []);

  useEffect(() => {
    if (isHomingActionCompleted && isLoggedInForHoming && !isLoading) {
      dispatch(
        logoutInitiated({
          deckName: "",
          token: tokenForHoming,
          showToast: false,
        })
      );
    }
  }, [isHomingActionCompleted]);

  useEffect(() => {
    if (error === false && isHomingActionCompleted) {
      setDisabled(false);
    }
  }, [error, isHomingActionCompleted]);

  /**
   * if user logged in, go to recipeListing page
   */
  if (isLoggedIn && !error) {
    return <Redirect to={`/${ROUTES.recipeListing}`} />;
  }

  return (
    <LandingScreen>
      <div className="landing-content">
        <VideoCard />
        <MlModal
          isOpen={isOpen && showHomingModal}
          textBody={MODAL_MESSAGE.homingConfirmation}
          handleSuccessBtn={() =>
            isHomingActionCompleted ? handleCompleteBtn() : homingConfirmation()
          }
          successBtn={
            isHomingActionCompleted ? MODAL_BTN.complete : MODAL_BTN.okay
          }
          showCrossBtn={false}
          progressPercentage={homingAllDeckCompletionPercentage}
          isProgressBarVisible={isProgressBarVisible}
          disabled={disabled}
        />
      </div>
    </LandingScreen>
  );
};

LandingScreenComponent.propTypes = {};

export default React.memo(LandingScreenComponent);
