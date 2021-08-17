import React from "react";
import PropTypes from "prop-types";
import styled from "styled-components";
import { Switch } from "core-components";
import { EXPERIMENT_STATUS } from "appConstants";

const StyledSelectAllGridHeader = styled.header`
  display: flex;
  height: 40px;
  align-items: center;
`;

const SelectAllGridHeader = ({
  className,
  isAllWellsSelected,
  toggleAllWellSelectedOption,
  experimentStatus,
  isExpanded,
}) => (
  <StyledSelectAllGridHeader className={className}>
    <Switch
      id="selectAll"
      name="selectAll"
      label="Select all wells"
      checked={isAllWellsSelected}
      onChange={() => toggleAllWellSelectedOption(isAllWellsSelected)}
      disabled={
        experimentStatus === EXPERIMENT_STATUS.success ||
        experimentStatus === EXPERIMENT_STATUS.running ||
        experimentStatus === EXPERIMENT_STATUS.stopped ||
        isExpanded === true
      }
    />
  </StyledSelectAllGridHeader>
);

SelectAllGridHeader.propTypes = {
  className: PropTypes.string,
  isAllWellsSelected: PropTypes.bool.isRequired,
  toggleAllWellSelectedOption: PropTypes.func.isRequired,
};

SelectAllGridHeader.defaultProps = {
  className: "",
};

export default SelectAllGridHeader;
