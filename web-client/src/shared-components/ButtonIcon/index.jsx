import React from "react";
import PropTypes from "prop-types";
import styled from "styled-components";

const StyledButtonIcon = styled.button`
	width: 40px;
	height: 40px;
	display: flex;
	justify-content: center;
	align-items: center;
	background-color: transparent;
	color: #707070;
	font-weight: normal;
	padding: 4px;
	border: 1px solid white;
	border-radius: 50%;
	box-shadow: ${(props) => (props.isShadow ? "0 2px 6px #00000020" : "")};
	position: ${(props) => props.position};
	top: ${(props) =>
		(props.position === "absolute" || props.position === "fixed")
			? `${props.top}px`
			: ""};
	right: ${(props) =>
		(props.position === "absolute" || props.position === "fixed") &&
		props.placement === "right"
			? `${props.right}px`
			: ""};
	left: ${(props) =>
		(props.position === "absolute" || props.position === "fixed") &&
		props.placement === "left"
			? `${props.left}px`
			: ""};

	i {
		font-size: 28px;
		line-height: 1;
	}
`;

const ButtonIcon = props => {

	return (
		<StyledButtonIcon {...props}>
			{props.children}
		</StyledButtonIcon>
	)
}

ButtonIcon.propTypes = {
	position: PropTypes.string,
	placement: PropTypes.oneOf(["left", "right"]),
	top: PropTypes.string,
	right: PropTypes.string,
	left: PropTypes.string,
	isShadow: PropTypes.bool,
};

ButtonIcon.defaultProps = {
	isShadow: false,
};

export default ButtonIcon;