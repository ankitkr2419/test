import React from "react";
import PropTypes from "prop-types";
import styled from "styled-components";
import { ButtonIcon, Text } from "shared-components";

const getOpacityValue = (props) => {
  const { isDisabled, isSelected } = props;

  if (isDisabled) {
    return "0.2";
  } else if (isSelected) {
    return "1";
  }
  return "0.5";
};

const StyledSampleTarget = styled.div`
  width: 220px;
  height: 38px;
  display: flex;
  align-items: center;
  background: #ffffff 0% 0% no-repeat padding-box;
  font-size: 18px;
  line-height: 22px;
  color: #707070;
  border: 1px solid #e5e5e5;
  border-radius: 8px;
  box-shadow: 0 3px 6px #0000000b;
  margin: 0 auto 2px;
  opacity: ${(props) => getOpacityValue(props)};

  button {
    color: #999999;
  }
`;

const SampleTarget = (props) => {
  const { label, isSelected, isDisabled, onClickHandler } = props;

  return (
    <StyledSampleTarget
      onClick={onClickHandler}
      isSelected={isSelected}
      isDisabled={isDisabled}
    >
      <Text className="m-0 px-2">{label}</Text>
      {isSelected ? (
        <ButtonIcon name="cross" size={28} className="ml-auto" />
      ) : null}
    </StyledSampleTarget>
  );
};

SampleTarget.propTypes = {
  label: PropTypes.string.isRequired,
  onClickHandler: PropTypes.func.isRequired,
  isSelected: PropTypes.bool,
  isDisabled: PropTypes.bool
};

SampleTarget.defaultProps = {
  isSelected: false
};

export default SampleTarget;
