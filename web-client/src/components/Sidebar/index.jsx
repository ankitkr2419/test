import React from "react";
import PropTypes from "prop-types";
import styled from "styled-components";
import SidebarContent from "./SidebarContent";
import SidebarBody from "./SidebarBody";
import SidebarHandle from "./SidebarHandle";

const StyledSidebar = styled.aside`
	display: flex;
	flex-direction: column;
	position: absolute;
	min-width: 408px;
	top: 0;
	bottom: 0;
	left: -332px;
	padding: 0 56px 0 0;
	overflow: hidden;
	transition: left 1s ease;
	z-index: 0;

	&.open {
		left: 0;
	}

	&.close {
		left: -100%;
	}
`;

const Sidebar = ({ className, handleIcon, handleIconSize, children }) => {
	return (
		<StyledSidebar className={`sidebar sidebar-${className}`}>
			<SidebarContent className="flex-100">
				<SidebarHandle
					icon={handleIcon}
					size={handleIconSize}
				/>
				<SidebarBody className="flex-100">{children}</SidebarBody>
			</SidebarContent>
		</StyledSidebar>
	);
};

Sidebar.propTypes = {
	className: PropTypes.string,
	handleIcon: PropTypes.string.isRequired,
	handleIconSize: PropTypes.number,
};

Sidebar.defaultProps = {
	className: "",
	handleIconSize: 24,
};

export default Sidebar;
