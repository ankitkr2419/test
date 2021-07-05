import React from "react";
import PropTypes from "prop-types";
import styled from "styled-components";
import { Switch } from "core-components";

const StyledSelectAllGridHeader = styled.header`
  display: flex;
  height: 40px;
  align-items: center;
  padding: 0 16px 0 26px;
`;

const SelectAllGridHeader = ({
  className,
  isAllWellsSelected,
  toggleAllWellSelectedOption,
}) => (
  <StyledSelectAllGridHeader className={className}>
    <Switch
      id="selection"
      name="selection"
      label="Select all wells"
      checked={isAllWellsSelected}
      onChange={toggleAllWellSelectedOption}
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
