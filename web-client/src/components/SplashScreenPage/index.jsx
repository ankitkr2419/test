import React from "react";
import { useHistory } from "react-router-dom";

import styled from "styled-components";
import { ImageIcon } from "shared-components";
import { ROUTES } from "appConstants";

import CirclelogoIcon from "assets/images/mylab-logo-with-circle.png";

const SplashScreen = styled.div`
  background: url("/images/logo-bg.svg") left -4.875rem top -5.5rem no-repeat,
    url("/images/honey-bees-bg.svg") right -1.75rem bottom -1.5rem no-repeat;
  .circle-image {
    margin-right: 14.313rem;
    margin-left: auto;
  }
  cursor: pointer;
`;

const SplashScreenComponent = () => {
  const history = useHistory();

  const redirectToLandingPage = () => {
    return history.push(ROUTES.landing);
  };

  return (
    <SplashScreen
      className="splash-screen-content h-100 py-0 bg-white d-flex justify-content-center align-items-center"
      onClick={redirectToLandingPage}
    >
      <div className="circle-image">
        <ImageIcon src={CirclelogoIcon} alt="My Lab" />
      </div>
    </SplashScreen>
  );
};

export default SplashScreenComponent;
