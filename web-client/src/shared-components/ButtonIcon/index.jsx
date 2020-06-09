import styled from "styled-components";

const ButtonIcon = styled.button`
	width: 42px;
	height: 42px;
	display: flex;
	justify-content: center;
	align-items: center;
	background-color: transparent;
	color: white !important;
	font-weight: normal;
	padding: 4px;
	border: 1px solid white;
	border-radius: 50%;
	box-shadow: 0 2px 6px #00000029;
	position: ${(props) => props.position};
	left: ${(props) =>
		(props.position === "absolute" || props.position === "fixed") &&
		props.placement === "left"
			? `${props.left}px`
			: ""};
	right: ${(props) =>
		(props.position === "absolute" || props.position === "fixed") &&
		props.placement === "right"
			? `${props.right}px`
			: ""};
`;

export default ButtonIcon;