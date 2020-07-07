import React from 'react';
import PropTypes from 'prop-types';
import styled from 'styled-components';

const StyledSidebarContent = styled.div`
	display: flex;
	flex-direction: column;
	position: relative;
	background: #aedbd5 0% 0% no-repeat padding-box;
	padding: 20px 20px 20px 0;
	box-shadow: 0 2px 6px #00000029;
	overflow: hidden;
`;

const SidebarContent = ({ className, children }) => {
	return (
		<StyledSidebarContent className={`sidebar-content ${className}`}>
			{children}
		</StyledSidebarContent>
	);
};

SidebarContent.propTypes = {
	className: PropTypes.string,
};

SidebarContent.defaultProps = {
	className: '',
};

export default SidebarContent;
