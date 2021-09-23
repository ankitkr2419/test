import React from "react";
import PropTypes from "prop-types";
import { Switch } from "core-components";
import { ImageIcon, Text } from "shared-components";
import bulb from "assets/images/bulb.svg";
import { StyledWhiteLightComponent } from "./Style";

const WhiteLight = (props) => {
  const { isLightOn, handleWhiteLightClick } = props;

  return (
    <StyledWhiteLightComponent className="d-flex ml-auto">
      <ImageIcon className="bulb" src={bulb} />
      <Text className="text-default my-1 px-3" size={18}>
        White Light
      </Text>
      <Switch
        className="ml-3"
        id="whiteLightSwitch"
        name="whiteLightSwitch"
        checked={isLightOn}
        onChange={handleWhiteLightClick}
      />
    </StyledWhiteLightComponent>
  );
};

WhiteLight.propTypes = {
  isLightOn: PropTypes.bool,
};

export default WhiteLight;
