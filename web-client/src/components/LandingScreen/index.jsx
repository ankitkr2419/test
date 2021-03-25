import React, { useState } from "react";
import { useDispatch } from "react-redux";
// import ConfirmationModal from "components/modals/ConfirmationModal";

// import SearchBox from 'shared-components/SearchBox';
// import ButtonBar from 'shared-components/ButtonBar';
import AppFooter from "components/AppFooter";
import { MODAL_MESSAGE, MODAL_BTN } from "appConstants";
import { homingActionInitiated } from "action-creators/homingActionCreators";
// import TimeModal from "components/modals/TimeModal";
import { VideoCard, MlModal } from "shared-components";

const LandingScreenComponent = (props) => {
  const [homingStatus, setHomingStatus] = useState(true);
  const dispatch = useDispatch();

  const homingConfirmation = () => {
    dispatch(homingActionInitiated());
    setHomingStatus(!homingStatus);
  };

  return (
    <div className="ml-content">
      <div className="landing-content">
        <VideoCard />

        <MlModal
          isOpen={homingStatus}
          textBody={MODAL_MESSAGE.homingConfirmation}
          handleSuccessBtn={homingConfirmation}
          successBtn={MODAL_BTN.okay}
          showCrossBtn={false}
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
