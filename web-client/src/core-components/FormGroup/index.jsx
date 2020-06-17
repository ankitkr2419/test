import React from "react";
import { FormGroup } from "reactstrap";
import styled from "styled-components";

const StyledFormGroup = styled(FormGroup)`
	position: relative;
`;

const CustomFormGroup = (props) => {
	return <StyledFormGroup {...props}>{props.children}</StyledFormGroup>;
};

CustomFormGroup.propTypes = {};

export default CustomFormGroup;
