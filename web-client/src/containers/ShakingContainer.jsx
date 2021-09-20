import React from "react";
import { useLocation } from "react-router";

import ShakingComponent from "components/Shaking";
import ShakerComponent from "components/CalibrationExtraction/ShakerComponent";
import { ROUTES } from "appConstants";

const ShakingContainer = (props) => {
  const location = useLocation();
  if (location.pathname === `/${ROUTES.calibration}/shaker`) {
    return <ShakerComponent />;
  }

  return <ShakingComponent />;
};

ShakingContainer.propTypes = {};

export default ShakingContainer;
