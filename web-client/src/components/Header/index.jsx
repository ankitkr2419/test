import React from "react";
import "./Header.scss";
import styled from "styled-components";
import imgLogo from "../../../assets/images/mylab-logo.png";

const Logo = styled.div`
  display: flex;
	width: 150px;
  height: 48px;
  align-items: center;
  justify-content: center;
`;

const Header = (props) => {
	return (
		<header className="ml-header">
			<Logo className="logo">
				<img src={imgLogo} alt="Mylab" className="w-100" />
			</Logo>
		</header>
	);
};

Header.propTypes = {};

export default Header;
