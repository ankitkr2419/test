import React from "react";
import { APP_TYPE, ROUTES } from "appConstants";
import { useHistory } from "react-router";
import { useSelector } from "react-redux";
import CalibrationRtpcrContainer from "./CalibrationRtpcrContainer";
import CalibrationExtractionContainer from "./CalibrationExtractionContainer";

const CalibrationContainer = () => {
  const history = useHistory();

  //get login reducer details
  const loginReducer = useSelector((state) => state.loginReducer);
  const loginReducerData = loginReducer.toJS();
  let activeDeckObj = loginReducerData?.decks.find((deck) => deck.isActive);
  const { isLoggedIn } = activeDeckObj;

  //get app type
  const appInfoReducer = useSelector((state) => state.appInfoReducer);
  const appInfoData = appInfoReducer.toJS();
  const app = appInfoData?.appInfo?.app;

  //redirect if not logged in
  if (!isLoggedIn) {
    if (app === APP_TYPE.EXTRACTION) {
      history.push(ROUTES.landing);
    } else if (app === APP_TYPE.RTPCR) {
      history.push(ROUTES.splashScreen);
    }
  }

  return (
    <>
      {/** extraction flow */}
      {app === APP_TYPE.EXTRACTION && <CalibrationExtractionContainer />}

      {/** rtpcr flow */}
      {app === APP_TYPE.RTPCR && <CalibrationRtpcrContainer />}
    </>
  );
};

export default CalibrationContainer;
