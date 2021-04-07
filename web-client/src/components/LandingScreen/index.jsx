import React, { useState } from "react";
import { useDispatch, useSelector } from "react-redux";
// import ConfirmationModal from "components/modals/ConfirmationModal";

import AppFooter from "components/AppFooter";
import { MODAL_MESSAGE, MODAL_BTN } from "appConstants";
import { homingActionInitiated } from "action-creators/homingActionCreators";
// import TimeModal from "components/modals/TimeModal";
import { VideoCard, MlModal, Loader } from "shared-components";
import styled from "styled-components";
const LandingScreen = styled.div`
  .landing-content{
    padding: 2.313rem 4.5rem;
    &::after{
      height:9.125rem;
    }
  }
`;
const LandingScreenComponent = () => {
  const [homingStatus, setHomingStatus] = useState(true);
  const dispatch = useDispatch();

  const homingReducer = useSelector((state) => state.homingReducer);
  const { isHomingActionCompleted } = homingReducer;

  const homingConfirmation = () => {
    dispatch(homingActionInitiated());
    setHomingStatus(!homingStatus);
  };

  return (
    <LandingScreen>
      <div className="landing-content">
        {isHomingActionCompleted && <Loader />}
        <VideoCard />
        <MlModal
          isOpen={homingStatus}
          textBody={MODAL_MESSAGE.homingConfirmation}
          handleSuccessBtn={homingConfirmation}
          successBtn={MODAL_BTN.okay}
          showCrossBtn={false}
        />
      </div>
      <AppFooter loginBtn={true} />
    </LandingScreen>
  );
};

LandingScreenComponent.propTypes = {};

export default LandingScreenComponent;
