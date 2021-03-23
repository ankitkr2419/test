import React, { useState } from "react";
import { useDispatch } from "react-redux";
import ReusableModal from "shared-components/ReusableModal";
// import ConfirmationModal from "components/modals/ConfirmationModal";

// import SearchBox from 'shared-components/SearchBox';
// import ButtonBar from 'shared-components/ButtonBar';
import AppFooter from "components/AppFooter";
import { MODAL_MESSAGE, MODAL_BTN } from "appConstants";
import { homingActionInitiated } from "action-creators/homingActionCreators";
// import TimeModal from "components/modals/TimeModal";
import { VideoCard } from "shared-components/index";

const LandingScreenComponent = (props) => {
  const [homingStatus, setHomingStatus] = useState(true);
  const dispatch = useDispatch();

  const homingConfirmation = (isConfirmed) => {
    setHomingStatus(isConfirmed);
    if (isConfirmed) {
      dispatch(homingActionInitiated());
    }
  };
  return (
    <div className="ml-content">
      <div className="landing-content">

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
