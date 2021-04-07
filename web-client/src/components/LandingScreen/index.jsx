import React, { useState } from "react";
import { useDispatch, useSelector } from "react-redux";

import AppFooter from "components/AppFooter";
import { MODAL_MESSAGE, MODAL_BTN } from "appConstants";
import { homingActionInitiated } from "action-creators/homingActionCreators";
import { operatorLoginReset } from "action-creators/operatorLoginModalActionCreators";
import { VideoCard, MlModal } from "shared-components";
import styled from "styled-components";

const LandingScreen = styled.div`
  .landing-content {
    padding: 2.313rem 4.5rem;
    &::after {
      height: 9.125rem;
    }
  }
`;

const LandingScreenComponent = () => {
  const dispatch = useDispatch();
  const homingReducer = useSelector((state) => state.homingReducer);

  // dispatch(operatorLoginReset());

  const {
    isHomingActionCompleted,
    homingAllDeckCompletionPercentage,
  } = homingReducer;

  // const [homingStatus, setHomingStatus] = useState(true);
  const [isProgressBarVisible, setIsProgressBarVisible] = useState(false);

  const homingConfirmation = () => {
    dispatch(homingActionInitiated());
    setIsProgressBarVisible(!isProgressBarVisible);
  };

  return (
    <LandingScreen>
      <div className="landing-content">
        <VideoCard />
        <MlModal
          isOpen={!isHomingActionCompleted}
          textBody={MODAL_MESSAGE.homingConfirmation}
          handleSuccessBtn={homingConfirmation}
          successBtn={MODAL_BTN.okay}
          showCrossBtn={false}
          progressPercentage={homingAllDeckCompletionPercentage}
          isProgressBarVisible={isProgressBarVisible}
        />
      </div>
      <AppFooter loginBtn={true} />
    </LandingScreen>
  );
};

LandingScreenComponent.propTypes = {};

export default LandingScreenComponent;
