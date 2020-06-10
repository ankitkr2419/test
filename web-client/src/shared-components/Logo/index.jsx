import React from "react";
import styled from "styled-components";
import imgLogo from "assets/images/mylab-logo.png";
import { Link } from "react-router-dom";

const StyledLogo = styled(Link)`
	display: flex;
	width: 150px;
	height: 48px;
	align-items: center;
	justify-content: center;
`;

const Logo = (props) => {
	return (
    <StyledLogo to="/" className="logo">
      <img src={imgLogo} alt="Mylab" className="w-100" />
    </StyledLogo>
	);
};

Logo.propTypes = {};

export default Logo;