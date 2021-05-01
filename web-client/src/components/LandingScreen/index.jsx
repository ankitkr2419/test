import React, { useState } from "react";
import { useDispatch, useSelector } from "react-redux";

import styled from "styled-components";
import AppFooter from "components/AppFooter";
import { VideoCard, MlModal } from "shared-components";

import { MODAL_MESSAGE, MODAL_BTN } from "appConstants";
import { homingActionInitiated } from "action-creators/homingActionCreators";

const LandingScreen = styled.div`
  .landing-content {
    padding: 2.313rem 4.5rem;
    &::after {
      height: 9.125rem;
    }
  }
`;

const LandingScreenComponent = (props) => {
  let { deckName }  = props;
  const dispatch = useDispatch();
  const homingReducer = useSelector((state) => state.homingReducer);

  const {
    isHomingActionCompleted,
    homingAllDeckCompletionPercentage,
  } = homingReducer;

  const [isProgressBarVisible, setIsProgressBarVisible] = useState(false);

  const homingConfirmation = () => {
    dispatch(homingActionInitiated());
    setIsProgressBarVisible(!isProgressBarVisible);
  };

  return (
    <LandingScreen>
      <div className="landing-content">
        <VideoCard />
        {/* <MlModal
          isOpen={!isHomingActionCompleted}
          textBody={MODAL_MESSAGE.homingConfirmation}
          handleSuccessBtn={homingConfirmation}
          successBtn={MODAL_BTN.okay}
          showCrossBtn={false}
          progressPercentage={homingAllDeckCompletionPercentage}
          isProgressBarVisible={isProgressBarVisible}
        /> */}
      </div>
      <AppFooter />
    </LandingScreen>
  );
};

LandingScreenComponent.propTypes = {};

export default LandingScreenComponent;
