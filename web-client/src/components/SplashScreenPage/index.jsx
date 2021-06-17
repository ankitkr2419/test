import React from "react";
import { useHistory } from "react-router-dom";
import { ImageIcon } from "shared-components";
import CirclelogoIcon from "assets/images/mylab-logo-with-circle.png";
import { SplashScreen } from './SplashScreen';
import { toast } from "react-toastify";

const SplashScreenComponent = (props) => {
  const { redirectionPath } = props;
  const history = useHistory();

  const redirectToLandingPage = () => {
    if(redirectionPath){
      return history.push(redirectionPath);
    } else {
      toast.warn('Unknown application type!');
    }
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
