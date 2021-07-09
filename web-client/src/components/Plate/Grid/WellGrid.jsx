import React from "react";
import PropTypes from "prop-types";
import styled from "styled-components";

const StyledWellGrid = styled.div`
  display: flex;
  flex-wrap: wrap;
`;

const WellGrid = ({ className, children }) => {
  return <StyledWellGrid className={className}>{children}</StyledWellGrid>;
};

WellGrid.propTypes = {
  className: PropTypes.string
};

WellGrid.defaultProps = {
  className: ""
};

export default WellGrid;
