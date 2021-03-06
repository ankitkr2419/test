import React from "react";
import PropTypes from "prop-types";
import { Icon } from "shared-components";
import { Button } from "core-components";
import { ButtonBarBox, PrevBtn } from "./Styles";
import { useHistory } from "react-router";

const ButtonBar = (props) => {
  const {
    leftBtnLabel,
    rightBtnLabel,
    handleLeftBtn,
    handleRightBtn,
    isRightBtnDisabled,
    isLeftBtnDisabled,
    btnBarClassname,
    isRTPCR,
    backBtnHandler,
    handleBackToRecipeList,
    pageReset,
  } = props;

  const history = useHistory();

  const handleBackBtn = () => {
    if (isRTPCR) {
      backBtnHandler();
      return;
    }
    history.goBack();
  };

  return (
    <ButtonBarBox
      className={`d-flex justify-content-start align-items-center ${btnBarClassname}`}
    >
      {!pageReset && (
        <PrevBtn onClick={handleBackBtn}>
          <Icon name="angle-left" size={30} />
        </PrevBtn>
      )}

      {pageReset && (
        <PrevBtn onClick={handleBackToRecipeList}>
          <Icon name="angle-left" size={30} />
        </PrevBtn>
      )}

      {leftBtnLabel && (
        <Button
          onClick={handleLeftBtn}
          color="outline-primary"
          disabled={isLeftBtnDisabled}
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
          disabled={isRightBtnDisabled}
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
  isLeftBtnDisabled: false,
  isRightBtnDisabled: false,
  pageReset: false,
};

export default ButtonBar;
