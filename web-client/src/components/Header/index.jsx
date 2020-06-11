import styled from "styled-components";

const Header = styled.header`
	position: relative;
	display: flex;
	align-items: center;
	height: 80px;
	background: white 0% 0% no-repeat padding-box;
	padding: 16px 48px;
	box-shadow: 0 4px 16px #00000029;
	z-index: 1;

	.btn-exit {
		color: #707070;

		i {
			font-size: 32px;
		}
	}
`;

export default Header;
