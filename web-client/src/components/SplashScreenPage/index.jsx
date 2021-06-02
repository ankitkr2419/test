import React from "react";
import { useHistory } from "react-router-dom";
import { ImageIcon } from "shared-components";
import { ROUTES } from "appConstants";
import CirclelogoIcon from "assets/images/mylab-logo-with-circle.png";
import { SplashScreen } from './SplashScreen';

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

export default React.memo(SplashScreenComponent);
