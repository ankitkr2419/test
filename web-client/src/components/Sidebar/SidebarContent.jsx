import React from 'react';
import PropTypes from 'prop-types';
import { StyledSidebarContent } from './Styles';

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
