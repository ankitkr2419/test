import React from 'react';
import styled from 'styled-components';
import { DropdownMenu } from 'reactstrap';

const StyledDropdownMenu = styled(DropdownMenu)`
	background: #fafafa 0% 0% no-repeat padding-box;
	box-shadow: 0 3px 32px #0000000f;
	border: 1px solid #e5e5e5;
	border-radius: 0 0 24px 24px;
	margin-top: 12px;
	padding: 0;

	&::after {
		content: '';
		position: absolute;
		top: -6px;
		right: 14px;
		border-left: 6px solid transparent;
		border-right: 6px solid transparent;
		border-bottom: 6px solid #fafafa;
	}

	&::before {
		content: '';
		position: absolute;
		top: -8px;
		right: 12px;
		border-left: 8px solid transparent;
		border-right: 8px solid transparent;
		border-bottom: 8px solid #e5e5e5;
	}
`;

const CustomDropdownMenu = (props) => {
	return <StyledDropdownMenu {...props} />;
};

CustomDropdownMenu.propTypes = {};

export default CustomDropdownMenu;
