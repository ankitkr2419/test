import React from 'react';
import { CustomInput } from "reactstrap";
import styled from "styled-components";

const StyledCheckBox = styled(CustomInput)`
	padding-left: ${(props) => (props.label ? "40px" : "24px")};

	.custom-control-input:checked ~ .custom-control-label::after {
		background-image: none;
		background-color: #aedbd5;
		left: ${(props) => (props.label ? "-36px" : "-20px")};
	}

	.custom-control-input:focus ~ .custom-control-label::before {
		box-shadow: none;
		border-color: #999999;
	}

	.custom-control-label {
		&::before {
			width: 24px;
			height: 24px;
			background-color: white;
			border: 1px solid #999999;
			top: 0;
			left: ${(props) => (props.label ? "-40px" : "-24px")};
			border-radius: 4px;
		}

		&::after {
			border-radius: 2px;
		}
	}

	.custom-control-input:checked ~ .custom-control-label::before,
	.custom-control-input:not(:disabled):active ~ .custom-control-label::before {
		background-color: white;
		border-color: #999999;
	}
`;

const CheckBox = props => {
  return(
    <StyledCheckBox type="checkbox" {...props} />
  );
};

export default CheckBox;