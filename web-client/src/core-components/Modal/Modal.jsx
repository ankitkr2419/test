import React from "react";
import styled from "styled-components";
import { Modal } from "reactstrap";

const StyledModal = styled(Modal)`
	.modal-content {
		background: #fafafa 0% 0% no-repeat padding-box;
		border: 1px solid #e5e5e5;
		box-shadow: 0 3px 16px #0000000f;
	}

	.modal-title {
		color: #666666;
		margin: 16px 0 64px;
	}
`;

const CustomModal = (props) => {
	return <StyledModal {...props}>{props.children}</StyledModal>;
};

CustomModal.propTypes = {};

export default CustomModal;
