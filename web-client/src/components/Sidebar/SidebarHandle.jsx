import React from 'react';
import PropTypes from 'prop-types';
import styled from 'styled-components';
import Icon from 'shared-components/Icon';

export const Shadow = styled.div`
	position: absolute;
	top: 0;
	left: 0;
	height: 172px;
	width: 60px;
	border-radius: 0 40px 40px 0;
	box-shadow: 0 2px 6px #00000029;
	z-index: 1;

	&::after {
		content: '';
		position: absolute;
		top: 50%;
		transform: translate(0%, -50%);
		width: 20px;
		height: 184px;
		background-color: #aedbd5;
		left: -8px;
		z-index: 2;
	}
`;

const StyledSidebarHandle = styled.button`
	position: absolute;
	top: 50%;
	right: -48px;
	transform: translate(0%, -50%);
	background-color: #aedbd5;
	border: 0 none;
	height: 172px;
	width: 60px;
	border-radius: 0 40px 40px 0;
	color: #ffffff;
	padding: 4px;
	z-index: 1;

	i {
		position: absolute;
		top: 50%;
		left: 50%;
		transform: translate(-50%, -50%);
	}
`;

const SidebarHandle = ({ icon, size, clickHandler }) => (
	<StyledSidebarHandle
		onClick={clickHandler}
		className='sidebar-handle'
		icon={icon}
		size={size}
	>
		<Icon name={icon} size={size} />
		<Shadow />
	</StyledSidebarHandle>
);

SidebarHandle.propTypes = {
	icon: PropTypes.string.isRequired,
	size: PropTypes.number,
	clickHandler: PropTypes.func.isRequired,
};

SidebarHandle.defaultProps = {
	size: 24,
};

export default SidebarHandle;
