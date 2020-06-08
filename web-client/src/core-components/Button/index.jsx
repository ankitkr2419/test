import styled, {css} from "styled-components";
import { Button } from "reactstrap";

const StyledButton = styled(Button)`
	width: 202px;
	height: 40px;
	font-size: 16px;
	line-height: 19px;
	font-weight: bold;
	padding: 10px 20px;
	border-radius: 27px;
	box-shadow: 0 2px 6px #00000029;

	${(props) =>
		props.icon &&
		css`
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
		`};

	&.btn-primary {
		color: white;
	}

	&.btn-secondary {
		color: #666666;
	}
`;

export default StyledButton;