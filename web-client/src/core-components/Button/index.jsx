import styled from "styled-components";
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

	&.btn-primary {
		color: white;
	}

	&.btn-secondary {
		color: #666666;
	}
`;

export default StyledButton;