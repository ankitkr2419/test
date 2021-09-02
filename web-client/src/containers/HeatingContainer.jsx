import React from "react";
import { useLocation } from "react-router";
import { ROUTES } from "appConstants";

import HeatingComponent from "components/Heating";
import HeaterComponent from "components/CalibrationExtraction/HeaterComponent";

const HeatingContainer = (props) => {
  const location = useLocation();
  if (location.pathname === `/${ROUTES.calibration}/heater`) {
    return <HeaterComponent />;
  }

  return <HeatingComponent />;
};

HeatingContainer.propTypes = {};

export default HeatingContainer;
