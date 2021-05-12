import React from 'react';
import PropTypes from 'prop-types';
import { StyledSidebarBody } from './Styles';

const SidebarBody = ({ className, children }) => {
	return (
		<StyledSidebarBody className={`sidebar-body scroll-y ${className}`}>
			{children}
		</StyledSidebarBody>
	);
};

SidebarBody.propTypes = {
	className: PropTypes.string,
	children: PropTypes.any,
};

SidebarBody.defaultProps = {
	className: '',
};

export default SidebarBody;
