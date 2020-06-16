import React from "react";
import { Input } from "reactstrap";
import styled from "styled-components";

const StyledInput = styled(Input)`
	padding: 4px 10px;
`;

const CustomInput = (props) => {
	return <StyledInput {...props}>{props.children}</StyledInput>;
};

CustomInput.propTypes = {};

export default CustomInput;
