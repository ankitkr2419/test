import React from "react";
import "./Header.scss";
import styled from "styled-components";
import imgLogo from "assets/images/mylab-logo.png";

const Logo = styled.div`
  display: flex;
	width: 150px;
  height: 48px;
  align-items: center;
  justify-content: center;
`;

const StyledHeader = styled.header`
	position: relative;
	display: flex;
	height: 80px;
	background: white 0% 0% no-repeat padding-box;
	padding: 16px 48px;
	box-shadow: 0 4px 16px #00000029;
	z-index: 1;
`;

const Header = (props) => {
	return (
		<StyledHeader className="ml-header">
			<Logo className="logo">
				<img src={imgLogo} alt="Mylab" className="w-100" />
			</Logo>
		</StyledHeader>
	);
};

Header.propTypes = {};

export default Header;
