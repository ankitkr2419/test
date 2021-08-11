import React from "react";
import styled from "styled-components";
import PropTypes from "prop-types";
import ButtonIcon from "shared-components/ButtonIcon";
import Text from "shared-components/Text";

const StyledTemplate = styled.div`
  position: relative;
  width: ${(props) => (props.isActive ? "315px" : "220px")};
  height: ${(props) => (props.isActive ? "60px" : "44px")};
  background: ${(props) => (props.isActive ? "#aedbd5" : "#ffffff")};
  display: flex;
  justify-content: center;
  align-items: center;
  font-size: 18px;
  line-height: 22px;
  color: ${(props) => (props.isActive ? "#ffffff" : "#707070")};
  font-weight: ${(props) => (props.isActive ? "bold" : "")};
  box-shadow: 0px 3px 16px #0000000b;
  border: 1px solid #e5e5e5;
  border-radius: 8px;
  padding: ${(props) => (props.isActive ? "8px 74px" : "8px 16px")};
  overflow: hidden;
  cursor: pointer;
`;

const EditButton = ({ onClickHandler }) => (
  <ButtonIcon
    position="absolute"
    placement="left"
    left={16}
    size={28}
    name="pencil"
    isShadow
    className="text-reset"
    onClick={onClickHandler}
  />
);

const DeleteButton = ({ onClickHandler }) => (
  <ButtonIcon
    position="absolute"
    placement="right"
    right={16}
    isShadow
    size={28}
    name="trash"
    className="text-reset"
    onClick={onClickHandler}
  />
);

export const CustomButton = (props) => {
  const {
    isEditable,
    isDeletable,
    title,
    onEditClickHandler,
    onDeleteClickHandler,
    onButtonClickHandler,
    isActive,
    ...rest
  } = props;

  return (
    <StyledTemplate
      onClick={onButtonClickHandler}
      {...rest}
      isActive={isActive}
    >
      {isEditable && isActive && (
        <EditButton onClickHandler={onEditClickHandler} />
      )}

      <Text Tag="span" onClick={onButtonClickHandler} className="text-truncate">
        {title}
      </Text>

      {isDeletable && isActive && (
        <DeleteButton onClickHandler={onDeleteClickHandler} />
      )}
    </StyledTemplate>
  );
};

CustomButton.propTypes = {
  title: PropTypes.string.isRequired,
  isActive: PropTypes.bool.isRequired,
  onButtonClickHandler: PropTypes.func.isRequired,
  isEditable: PropTypes.bool,
  onEditClickHandler: (props, propName, componentName) => {
    if (
      props.isEditable === true &&
      (props[propName] === undefined || typeof props[propName] !== "function")
    ) {
      return new Error("Please provide a onEditClickHandler function!");
    }
  },
  isDeletable: PropTypes.bool,
  onDeleteClickHandler: (props, propName, componentName) => {
    if (
      props.isDeletable === true &&
      (props[propName] === undefined || typeof props[propName] !== "function")
    ) {
      return new Error("Please provide a onDeleteClickHandler function!");
    }
  },
};

CustomButton.defaultProps = {
  isActive: false,
  isEditable: false,
  isDeletable: false,
};
