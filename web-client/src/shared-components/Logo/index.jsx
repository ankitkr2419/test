import React from "react";
import styled from "styled-components";
import { Link } from "react-router-dom";
import PropTypes from "prop-types";
import imgLogo from "assets/images/mylab-logo.png";
import imgSymbol from "assets/images/mylab-symbol.png";

const StyledLogo = styled(Link)`
  display: flex;
  width: ${(props) => (props.size === "sm" ? "58px" : "150px")};
  height: 48px;
  align-items: center;
  justify-content: center;
  margin: ${(props) => (props.size === "sm" ? "0 12px" : "")};
`;

const Logo = (props) => {
  const { isSmall } = props;
  const size = isSmall ? "sm" : "lg";

  return (
    <StyledLogo size={size} to="/" className="logo">
      {isSmall ? (
        <img src={imgSymbol} alt="Mylab" className="h-100" />
      ) : (
        <img src={imgLogo} alt="Mylab" className="w-100" />
      )}
    </StyledLogo>
  );
};

Logo.propTypes = {
  isSmall: PropTypes.bool,
};

Logo.defaultProps = {
  isSmall: false,
};

export default Logo;
