import React from "react";
import PropTypes from "prop-types";
import styled from "styled-components";
import { Switch } from "core-components";

const StyledWellGridHeader = styled.header`
  display: flex;
  height: 40px;
  align-items: center;
`;

const WellGridHeader = ({
  wells,
  className,
  isGroupSelectionOn,
  toggleMultiSelectOption,
  experimentStatus,
  isExpanded,
}) => {
  const filledWells = wells
    .toJS()
    .filter((wellObj) => wellObj && wellObj.isWellFilled === true);

  return (
    <StyledWellGridHeader className={className}>
      <Switch
        id="selection"
        name="selection"
        label="Group Selection to view on graph"
        checked={isGroupSelectionOn}
        onChange={toggleMultiSelectOption}
        disabled={
          filledWells.length === 0 ||
          (experimentStatus === null && isExpanded === false)
        }
      />
    </StyledWellGridHeader>
  );
};

WellGridHeader.propTypes = {
  className: PropTypes.string,
  isGroupSelectionOn: PropTypes.bool.isRequired,
  toggleMultiSelectOption: PropTypes.func.isRequired,
};

WellGridHeader.defaultProps = {
  className: "",
};

export default WellGridHeader;
