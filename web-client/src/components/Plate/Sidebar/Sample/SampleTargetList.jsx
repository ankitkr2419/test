import React from "react";
import PropTypes from "prop-types";
import styled from "styled-components";
import SampleTarget from "./SampleTarget";
import { Text } from "shared-components";

const StyledSampleTargetList = styled.div`
  flex: 1;
  margin: 0 0 24px;

  .list-title {
    color: #707070;
  }
`;

const SampleTargetList = ({ list, onTargetClickHandler, isDisabled }) => (
  <StyledSampleTargetList>
    <Text
      size={14}
      className={`list-title mb-1 px-2 font-weight-bold ${
        isDisabled ? "disabled" : ""
      }`}
    >
      Selected Targets
    </Text>
    {list.map((ele, index) => (
      <SampleTarget
        key={ele.get("target_id")}
        onClickHandler={() => onTargetClickHandler(index)}
        label={ele.get("target_name")}
        isSelected={ele.get("isSelected")}
        isDisabled={isDisabled}
      />
    ))}
  </StyledSampleTargetList>
);

SampleTargetList.propTypes = {
  list: PropTypes.object.isRequired,
  onTargetClickHandler: PropTypes.func.isRequired,
  isDisabled: PropTypes.bool
};

export default SampleTargetList;
