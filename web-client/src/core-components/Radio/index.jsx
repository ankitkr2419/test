import React from 'react';
import { CustomInput } from "reactstrap";
import styled from "styled-components";

const StyledRadioButton = styled(CustomInput)`
	padding-left: ${(props) => (props.label ? "40px" : "24px")};

	.custom-control-input:checked ~ .custom-control-label::after {
		left: ${(props) => (props.label ? "-36px" : "-20px")};
	}

	.custom-control-label::before {
		left: ${(props) => (props.label ? "-40px" : "-24px")};
	}
`;

const Radio = props => {
  return(
    <StyledRadioButton type="radio" {...props} />
  );
};

export default Radio;