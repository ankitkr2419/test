import React from "react";
import PropTypes from "prop-types";
import { Switch } from "core-components";
import { Text } from "shared-components";

const WhiteLight = (props) => {
  const { isLightOn, handleWhiteLightClick } = props;

  return (
    <div className="mt-3 ml-auto">
      <Switch
        id="whiteLightSwitch"
        name="whiteLightSwitch"
        checked={isLightOn}
        onChange={handleWhiteLightClick}
      />
      <Text className="text-default pl-1" size={12}>
        White Light
      </Text>
    </div>
  );
};

WhiteLight.propTypes = {
  isLightOn: PropTypes.bool,
};

export default WhiteLight;
