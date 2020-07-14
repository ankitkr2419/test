import React from 'react';
import { Input } from 'reactstrap';
import styled from 'styled-components';

const StyledInput = styled(Input)`
	padding: 4px 10px;
`;

// On focus of input field select the complete text if present
// This makes it easy to clear
const selectText = (event) => event.target.select();

const CustomInput = (props) => (
	<StyledInput {...props} onFocus={selectText}>
		{props.children}
	</StyledInput>);

CustomInput.propTypes = {};

export default CustomInput;
