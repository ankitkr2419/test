import React from 'react';
import PropTypes from 'prop-types';
import styled from 'styled-components';
import { PopoverHeader } from 'reactstrap';

const StyledPopoverHeader = styled(PopoverHeader)`
	height: 36px;
	line-height: 1.2;
	font-weight: normal;
	background-color: ${props => props.status || '#aedbd5'} !important;
`;

const CustomPopoverHeader = (props) => {
	const { children, ...rest } = props;

	return <StyledPopoverHeader {...rest}>{children}</StyledPopoverHeader>;
};

CustomPopoverHeader.propTypes = {
	children: PropTypes.any,
};

CustomPopoverHeader.defaultProps = {};

export default CustomPopoverHeader;
