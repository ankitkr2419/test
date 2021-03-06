import React from 'react';
import PropTypes from 'prop-types';
import classNames from 'classnames';
import styled from 'styled-components';
import SidebarContent from './SidebarContent';
import SidebarBody from './SidebarBody';
import SidebarHandle from './SidebarHandle';

const Sidebar = ({
	className,
	bodyClassName,
	handleIcon,
	handleIconSize,
	children,
	isOpen,
	isClose,
	toggleSideBar,
	isDisabled,
}) => {
	const classes = classNames(
		'sidebar',
		`sidebar-${className}`,
		{ open: isOpen },
		{ close: isClose }
	);
	return (
		<StyledSidebar className={classes}>
			<SidebarHandle
				clickHandler={toggleSideBar}
				icon={handleIcon}
				size={handleIconSize}
				isDisabled={isDisabled}
			/>
			<SidebarContent className='flex-100'>
				<SidebarBody className={`flex-100 ${bodyClassName}`}>
					{children}
				</SidebarBody>
			</SidebarContent>
		</StyledSidebar>
	);
};

const StyledSidebar = styled.aside`
	display: flex;
	flex-direction: column;
	position: absolute;
	min-width: 408px;
	top: 0;
	bottom: 0;
	left: -388px;
	z-index: 0;
	transition: left 1s ease;

	&.open {
		left: 0;
		z-index: 1;
	}

	&.close {
		left: -100%;
	}
`;

Sidebar.propTypes = {
	className: PropTypes.string,
	bodyClassName: PropTypes.string,
	handleIconSize: PropTypes.number,
	isOpen: PropTypes.bool,
	isClose: PropTypes.bool,
	isDisabled: PropTypes.bool,
	children: PropTypes.any,
	handleIcon: PropTypes.string.isRequired,
	toggleSideBar: PropTypes.func.isRequired,
};

Sidebar.defaultProps = {
	className: '',
	bodyClassName: '',
	handleIconSize: 24,
	isOpen: false,
	isClose: false,
};

export default Sidebar;
