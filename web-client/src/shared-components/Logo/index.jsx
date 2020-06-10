import React from "react";
import styled from "styled-components";
import { Link } from "react-router-dom";
import imgLogo from "assets/images/mylab-logo.png";
import imgSymbol from "assets/images/mylab-symbol.png";

const StyledLogo = styled(Link)`
	display: flex;
	width: ${(props) => (props.sm ? "58px" : "150px")};
	height: 48px;
	align-items: center;
	justify-content: center;
	margin: ${(props) => (props.sm ? "0 12px" : "")};
`;

const Logo = (props) => {
	return (
		<StyledLogo {...props} to="/" className="logo">
			{props.sm ? (
				<img src={imgSymbol} alt="Mylab" className="h-100" />
			) : (
				<img src={imgLogo} alt="Mylab" className="w-100" />
			)}
		</StyledLogo>
	);
};

Logo.propTypes = {};

export default Logo;