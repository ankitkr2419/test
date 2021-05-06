import React from "react";

import styled from "styled-components";
import PropTypes from "prop-types";
import { Icon } from "shared-components";
import { Button } from "core-components";

const ButtonBarBox = styled.div`
  width: 93.82%;
  height: 3.25rem;
  background-color: #fff;
  z-index: 2;
  border-radius: 2rem 0 0 2rem;
  padding: 0.5rem 4.938rem 0.5rem 2.375rem;
  box-shadow: 0px 3px 16px rgba(0, 0, 0, 0.06);
  position: absolute;
  right: 0;
  // bottom: 3rem;
  top: 29rem;
  > button {
    width: 160px;
    &:hover,
    &:focus {
      color: #ffffff !important;
      > i {
        color: #ffffff !important;
      }
    }
    > i {
      color: #f38220;
    }
  }
`;
const PrevBtn = styled.div`
  min-width: inherit;
  border: 0;
  box-shadow: none;
  color: #f38220;
`;

const ButtonBar = (props) => {
  const { leftBtnLabel, rightBtnLabel, handleLeftBtn, handleRightBtn } = props;
  return (
    <ButtonBarBox className="d-flex justify-content-start align-items-center mt-5">
      <PrevBtn>
        <Icon name="angle-left" size={30} />
      </PrevBtn>

      {leftBtnLabel && (
        <Button
          onClick={handleLeftBtn}
          color="outline-primary"
          className="ml-auto text-dark"
          size="md"
        >
          {" "}
          {leftBtnLabel === "Add Process" && (
            <Icon size={20} name="plus-2" className="mb-0 p-0" />
          )}
          {leftBtnLabel}
        </Button>
      )}

      {rightBtnLabel && (
        <Button
          onClick={handleRightBtn}
          color="primary"
          className={leftBtnLabel ? "ml-4" : "ml-auto"}
          size="md"
        >
          {" "}
          {rightBtnLabel}
        </Button>
      )}
    </ButtonBarBox>
  );
};

ButtonBar.propTypes = {
  isUserLoggedIn: PropTypes.bool,
};

ButtonBar.defaultProps = {
  isUserLoggedIn: false,
};

export default ButtonBar;
