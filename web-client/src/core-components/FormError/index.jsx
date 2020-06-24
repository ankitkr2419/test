import React from 'react';
import styled from 'styled-components';
import { FormFeedback } from 'reactstrap';

const StyledFormError = styled(FormFeedback)`
	position: absolute;
	bottom: -18px;
	text-align: right;
	overflow: hidden;
	text-overflow: ellipsis;
	white-space: nowrap;
	padding: 0 8px;
`;

const FormError = props => {
  return (
    <StyledFormError {...props}>
      {props.children}
    </StyledFormError>
  );
};

export default FormError;