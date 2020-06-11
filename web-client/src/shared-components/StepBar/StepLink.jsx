import styled from "styled-components";
import { NavLink } from "reactstrap";

export const StepLink = styled(NavLink)`
	font-size: 16px;
	line-height: 24px;
	color: #707070;
	padding: 8px 0;
	border-radius: 4px;

	&:hover,
	&:focus {
		color: #707070;
	}
`;
