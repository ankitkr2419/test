import React from "react";

import PropTypes from "prop-types";
import { Icon } from "shared-components";
import { ActionBtn } from './Styles';

const ActionButton = (props) => {
  const { label, icon, disabled } = props;

  return (
    <>
      <ActionBtn
        className="d-flex justify-content-center align-items-center"
        disabled={disabled}
        style={disabled ? {opacity: "0.4", pointerEvents:"none"}:{}}
      >
        <div className="semi-circle-outter-box">
          <div className="semi-circular-button">
            <Icon name={icon} size={18} />
            <div className="btn-label row flex-nowrap font-weight-bold">
              {label} 
            </div>
          </div>
        </div>
      </ActionBtn>
    </>
  );
};

ActionButton.propTypes = {
  isUserLoggedIn: PropTypes.bool,
  disabled: PropTypes.bool,
};

ActionButton.defaultProps = {
  isUserLoggedIn: false,
  disabled: false,
};

export default React.memo(ActionButton);
