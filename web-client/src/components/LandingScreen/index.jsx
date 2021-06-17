import React, { useEffect, useState } from "react";
import { useDispatch, useSelector } from "react-redux";
import { VideoCard, MlModal } from "shared-components";

import { MODAL_MESSAGE, MODAL_BTN } from "appConstants";
import { homingActionInitiated } from "action-creators/homingActionCreators";
import { LandingScreen } from "./LandingScreen";

const LandingScreenComponent = (props) => {
  const dispatch = useDispatch();
  const homingReducer = useSelector((state) => state.homingReducer);

  const { isHomingActionCompleted, homingAllDeckCompletionPercentage } =
    homingReducer;

  const [isProgressBarVisible, setIsProgressBarVisible] = useState(false);
  const [isOpen, setIsOpen] = useState(true);
  const [disabled, setDisabled] = useState(false);

  const homingConfirmation = () => {
    dispatch(homingActionInitiated());
    setIsProgressBarVisible(!isProgressBarVisible);
    setDisabled(!disabled);
  };

  useEffect(() => {
    if (isHomingActionCompleted) {
      setDisabled(false);
    }
  }, [isHomingActionCompleted]);

  return (
    <LandingScreen>
      <div className="landing-content">
        <VideoCard />
        <MlModal
          isOpen={isOpen && !isHomingActionCompleted}
          textBody={MODAL_MESSAGE.homingConfirmation}
          handleSuccessBtn={() =>
            isHomingActionCompleted ? setIsOpen(false) : homingConfirmation()
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
