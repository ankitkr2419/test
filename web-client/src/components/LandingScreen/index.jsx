import React, { useState } from "react";
import { useDispatch, useSelector } from "react-redux";

import ReusableModal from "shared-components/ReusableModal";
import AppFooter from "components/AppFooter";
import { MODAL_MESSAGE, MODAL_BTN } from "appConstants";
import { homingActionInitiated } from "action-creators/homingActionCreators";
// import TimeModal from "components/modals/TimeModal";
import { VideoCard, Loader } from "shared-components/index";
const LandingScreenComponent = (props) => {
  const [homingStatus, setHomingStatus] = useState(true);
  const dispatch = useDispatch();

  const homingReducer = useSelector((state) => state.homingReducer);
  const { isHomingActionCompleted } = homingReducer;

  const homingConfirmation = (isConfirmed) => {
    setHomingStatus(isConfirmed);
    if (isConfirmed) {
      dispatch(homingActionInitiated());
      setHomingStatus(false);
    }
  };

  return (
    <div className="ml-content">
      <div className="landing-content">
        {isHomingActionCompleted && <Loader />}
        <VideoCard />
        <ReusableModal
          isOpen={homingStatus}
          toggleModal={homingConfirmation}
          textBody={MODAL_MESSAGE.homingConfirmation}
          clickHandler={homingConfirmation}
          primaryBtn={MODAL_BTN.okay}
        />
      </div>
      <AppFooter
        loginBtn={true}
        // deckName="Deck A"
        // showProcess={true}
        // showCleanUp={true}
      />
    </div>
  );
};

LandingScreenComponent.propTypes = {};

export default LandingScreenComponent;
