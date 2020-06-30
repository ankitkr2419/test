import styled from 'styled-components';
import { NavLink } from 'react-router-dom';

const StyledNavLink = styled(NavLink).attrs({ className: 'nav-link' })`
	font-size: 16px;
	line-height: 18px;
	color: #707070;
	padding: 8px;

	&:hover {
		color: #707070;
	}

	&.active {
		position: relative;

		&::before {
			content: '';
			position: absolute;
			background: #666666 0% 0% no-repeat padding-box;
			left: 0;
			right: 0;
			bottom: -4px;
			height: 4px;
			border-radius: 2px;
		}
	}
`;

export default StyledNavLink;
