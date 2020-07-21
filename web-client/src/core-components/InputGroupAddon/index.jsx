import React from 'react';
import PropTypes from 'prop-types';
import styled from 'styled-components';
import { InputGroupAddon } from 'reactstrap';

const StyledInputGroupAddon = styled(InputGroupAddon)``;

const CustomInputGroupAddon = (props) => {
	const { children, ...rest } = props;
	return <StyledInputGroupAddon {...rest}>{children}</StyledInputGroupAddon>;
};

CustomInputGroupAddon.propTypes = {
	children: PropTypes.element,
};

export default CustomInputGroupAddon;
