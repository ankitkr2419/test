import React from "react";
import PropTypes from "prop-types";
import styled from "styled-components";
import Icon from "shared-components/Icon";

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
`;

const ButtonIcon = props => {

	return (
		<StyledButtonIcon {...props}>
			<Icon size={props.size} name={props.name} />
		</StyledButtonIcon>
	)
}

ButtonIcon.propTypes = {
	position: PropTypes.string,
	placement: PropTypes.oneOf(["left", "right"]),
	top: PropTypes.number,
	right: PropTypes.number,
	left: PropTypes.number,
	isShadow: PropTypes.bool,
	name: PropTypes.string.isRequired,
	size: PropTypes.number,
};

ButtonIcon.defaultProps = {
	isShadow: false,
	size: 24
};

export default ButtonIcon;