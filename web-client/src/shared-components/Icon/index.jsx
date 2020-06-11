import React from "react";
import styled from "styled-components";

const StyledIcon = styled.i`
  font-size: 24px;
	line-height: 1.1875;
`;

//* Important: Refer "_fonts.scss" for icon names
const Icon = (props) => {
  return(
    <StyledIcon className={`icon-${props.name}`} />
  )
};

export default Icon;