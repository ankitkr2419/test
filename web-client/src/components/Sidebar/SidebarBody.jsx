import React from 'react';
import PropTypes from 'prop-types';
import styled from 'styled-components';

const StyledSidebarBody = styled.div`
	background: #fafafa 0% 0% no-repeat padding-box;
	padding: 48px;
	text-align: center;
	border: 1px solid #e5e5e5;
	border-left: 0 none;
	border-radius: 0px 24px 24px 0px;
`;

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
