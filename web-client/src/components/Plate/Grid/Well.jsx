import React from "react";
import PropTypes from "prop-types";
import classNames from "classnames";
import styled from "styled-components";

const Well = (props) => {
  const {
    id,
    className,
    status,
    isRunning,
    isSelected,
    taskInitials,
    onClickHandler
  } = props;
  const wellClassnames = classNames(className, {
    running: isRunning,
    selected: isSelected
  });

  return (
    <StyledWell
      id={id}
      // isRunning={isDisabled}
      isSelected={isSelected}
      // isDisabled={isDisabled}
      status={status}
      className={wellClassnames}
      onClick={onClickHandler}
    >
      {taskInitials}
    </StyledWell>
  );
};

const getBackgroundColor = ({ isSelected, isRunning, isDisabled }) => {
  if (isSelected && !isRunning) {
    return "#aedbd5";
  }

  if (isDisabled) {
    return "gray";
  }

  return "#ffffff";
};

const StyledWell = styled.div`
  background-color: ${(props) => getBackgroundColor(props)};
  position: relative;
  display: flex;
  align-items: center;
  justify-content: center;
  width: 50px;
  height: 50px;
  font-size: 20px;
  line-height: 24px;
  color: #666666;
  cursor: pointer;
  border: ${(props) =>
    props.isSelected && props.isRunning
      ? "2px solid #707070"
      : "1px solid #aeaeae"};
  border-radius: 8px;
  margin: 0 24px 48px 0;
  padding: 18px 4px 4px;
  box-shadow: ${(props) =>
    props.isSelected && props.isRunning ? "0 3px 6px #00000029" : ""};
  opacity: ${(props) => (props.isDisabled ? "0.2" : "1")};
  pointer-events: ${(props) => (props.isDisabled ? "none" : "auto")};

  &.selected {
    &:active,
    &:active:focus {
      background-color: #aedbd5;
    }
  }

  &.running {
    &:active,
    &:active:focus {
      background-color: #ffffff;
      border: 2px solid #707070;
      box-shadow: 0 3px 6px #00000029;
    }
  }

  &:active,
  &:active:focus {
    background-color: #aedbd5;
  }

  &:focus {
    outline: none;
  }

  &::before {
    content: "";
    position: absolute;
    top: 0;
    right: 0;
    left: 0;
    height: 14px;
    border-radius: 6px 6px 0 0;
    background-color: ${(props) => props.status};
  }
`;

Well.propTypes = {
  id: PropTypes.number,
  className: PropTypes.string,
  status: PropTypes.string,
  taskInitials: PropTypes.string,
  isSelected: PropTypes.bool,
  isRunning: PropTypes.bool,
  onClickHandler: PropTypes.func.isRequired,
  isDisabled: PropTypes.bool
};

Well.defaultProps = {
  className: "",
  status: "",
  taskInitials: "",
  isSelected: false,
  isRunning: false,
  isDisabled: false
};

export default Well;
