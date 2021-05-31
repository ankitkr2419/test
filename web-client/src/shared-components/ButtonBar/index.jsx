import React from "react";
import PropTypes from "prop-types";
import { Icon } from "shared-components";
import { Button } from "core-components";
import { ButtonBarBox, PrevBtn } from "./Styles";
import { useHistory } from "react-router";

const ButtonBar = (props) => {
  const { leftBtnLabel, rightBtnLabel, handleLeftBtn, handleRightBtn } = props;
  const history = useHistory();

  const handleBackBtn = () => {
    history.goBack();
  };

  return (
    <ButtonBarBox className="d-flex justify-content-start align-items-center mt-5">
      <PrevBtn onClick={handleBackBtn}>
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
