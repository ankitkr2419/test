import React from "react";

import styled from "styled-components";
import PropTypes from "prop-types";
import { Icon } from "shared-components";



const ActionBtn = styled.button`
  position: absolute;
  top: 0;
  right: 0;
  z-index: 2;
  display: block;
  outline: none;
  border: 0;
  background: transparent;
  .semi-circle-outter-box {
    width: 3.875rem;
    height: 2.5rem;
    background-color: #b3d9d0;
    border-bottom-left-radius: 5.5rem;
    border-bottom-right-radius: 5.5rem;
    box-shadow: 0px -3px 6px rgb(0 0 0 / 31%);
    display: flex;
    justify-content: center;
    align-items: center;
  }
  .semi-circular-button {
    width: 2.75rem;
    height: 2.75rem;
    border-radius: 50%;
    border: 1px solid #f0801d;
    background-color: #f38220;
    z-index: 1;
    text-decoration: none;
    margin-top: -1.375rem;
    display: flex;
    justify-content: center;
    align-items: center;
    color: #fff;
    position: relative;
    .btn-label {
      position: absolute;
      bottom: -2rem;
      left: 0;
      right: 0;
      text-align: center;
      display: flex;
      justify-content: center;
      align-items: center;
      font-size: 0.75rem;
      line-height: 0.875rem;
      color: #3c3c3c;
    }
    .icon-play,
    .icon-resume {
      margin-left: 0.25rem;
    }
  }
`;
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
