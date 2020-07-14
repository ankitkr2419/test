import React from 'react';
import PropTypes from 'prop-types';
import styled from 'styled-components';
import { InputGroup } from 'reactstrap';

const StyledInputGroup = styled(InputGroup)``;

const CustomInputGroup = ({ children }) => {
	return <StyledInputGroup>{children}</StyledInputGroup>;
};

CustomInputGroup.propTypes = {
	children: PropTypes.element,
};

export default CustomInputGroup;
