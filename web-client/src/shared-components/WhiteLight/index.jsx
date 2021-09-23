import React from "react";
import PropTypes from "prop-types";
import { Switch } from "core-components";
import { Text } from "shared-components";

const WhiteLight = (props) => {
  const { isLightOn, handleWhiteLightClick } = props;

  return (
    <div className="d-flex ml-auto">
      <Text className="text-default my-1" size={18}>
        White Light
      </Text>
      <Switch
        className="ml-3"
        id="whiteLightSwitch"
        name="whiteLightSwitch"
        checked={isLightOn}
        onChange={handleWhiteLightClick}
      />
    </div>
  );
};

WhiteLight.propTypes = {
  isLightOn: PropTypes.bool,
};

export default WhiteLight;
