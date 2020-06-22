
import PropTypes from "prop-types";
import { Card } from "reactstrap";
import styled from "styled-components";

const StyledCard = styled(Card)`
	height: ${(props) => (props.default ? "" : "528px")};
	background: ${(props) =>
		props.default
			? "#F4F4F4 0% 0% no-repeat padding-box"
			: "#fafafa 0% 0% no-repeat padding-box"};
	box-shadow: ${(props) => (props.default ? "" : "0 3px 16px #0000000f")};
	border: 1px solid ${(props) => (props.default ? "#e6e6e6" : "#e5e5e5")};
	border-radius: ${(props) => (props.default ? "24px" : "36px")};
`;

StyledCard.propTypes = {
	default: PropTypes.bool,
};

StyledCard.defaultProps = {
	default: false,
};

export default StyledCard;
