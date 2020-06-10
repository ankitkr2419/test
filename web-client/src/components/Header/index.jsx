import React from "react";
import styled from "styled-components";
import Logo from "shared-components/Logo";

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
			<Logo isSmall />
		</StyledHeader>
	);
};

Header.propTypes = {};

export default Header;
