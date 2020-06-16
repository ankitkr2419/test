import React from "react";
import PropTypes from "prop-types";
import styled from "styled-components";

const StyledIcon = styled.i`
	display: inline-block;
	font-size: ${(props) => props.size}px;
	line-height: 1;
	vertical-align: middle;
`;

//* Important: Refer "_fonts.scss" for icon names
const Icon = (props) => {
  return(
    <StyledIcon className={`icon-${props.name}`} size={props.size} />
  )
};


Icon.propTypes = {
  name: PropTypes.string.isRequired,
  size: PropTypes.number,
};


Icon.defaultProps = {
	size: 24,
};

export default Icon;