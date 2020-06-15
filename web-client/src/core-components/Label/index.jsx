import React from "react";
import { Label } from "reactstrap";
import styled from "styled-components";

const StyledLabel = styled(Label)`
	font-size: 14px;
	line-height: 16px;
	font-weight: normal;
	color: #666666;
	padding: 0 12px;
`;

const CustomLabel = (props) => {
	return <StyledLabel {...props}>{props.children}</StyledLabel>;
};

CustomLabel.propTypes = {};

export default CustomLabel;
