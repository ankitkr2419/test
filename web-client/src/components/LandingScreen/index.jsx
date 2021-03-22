import React, { useState } from "react";
import { useDispatch } from "react-redux";
import { ImageIcon, Center, Icon } from "shared-components";
import ReusableModal from "shared-components/ReusableModal";
// import ConfirmationModal from "components/modals/ConfirmationModal";

// import SearchBox from 'shared-components/SearchBox';
// import ButtonBar from 'shared-components/ButtonBar';
import imgNoTemplate from "assets/images/video-thumbnail-poster.jpg";
import styled from "styled-components";
import AppFooter from "components/AppFooter";
import { MODAL_MESSAGE, MODAL_BTN } from "appConstants";
import { CardBody, Card } from "core-components";
import { homingActionInitiated } from "action-creators/homingActionCreators";
import TimeModal from "components/modals/TimeModal";

const VideoPlayButton = styled.button`
  color: #7c7976;
  background-color: transparent;
  border: 0;
  outline: none;
  position: absolute;
  top: 50%;
  left: 50%;
  transform: translate(-50%, -50%);
`;

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
        <Card className="card-video">
          <CardBody className="d-flex flex-column p-0">
            <Center className="video-thumbnail-wrapper">
              <ImageIcon
                src={imgNoTemplate}
                alt="No templates available"
                className="img-video-thumbnail"
              />
              <VideoPlayButton>
                <Icon name="play" size={124} />
              </VideoPlayButton>
            </Center>
          </CardBody>
        </Card>
        <ReusableModal
          isOpen={homingStatus}
          toggleModal={homingConfirmation}
          textHead={MODAL_MESSAGE.homingConfirmation}
          clickHandler={homingConfirmation}
          primaryBtn={MODAL_BTN.okay}
        />
      </div>
      <AppFooter />
    </div>
  );
};

LandingScreenComponent.propTypes = {};

export default LandingScreenComponent;
