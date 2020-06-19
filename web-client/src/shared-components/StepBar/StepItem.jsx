import styled from "styled-components";
import { NavItem } from "reactstrap";
import PropTypes from "prop-types";
import bgStep from "../../assets/images/steps.svg";

export const StepItem = styled(NavItem)`
	width: 185px;
	height: 40px;
	background: transparent url(${bgStep}) no-repeat;
	background-size: auto 84px;
	background-position: -22px -19px;
	margin: 0;
	padding: 0 23px;
	text-align: center;
	opacity: ${(props) => (props.disabled ? "0.53" : "")};
	pointer-events: ${(props) => (props.disabled ? "none" : "")};

	+ .nav-item {
		margin: 0 0 0 -16px;
	}
`;

StepItem.passProps = false;

StepItem.propTypes = {
	disabled: PropTypes.bool,
};

StepItem.defaultProps = {
	disabled: false
};