import styled from 'styled-components';
import PropTypes from 'prop-types';

const GraphCard = styled.div`
	width: ${(props) => props.width + `px`};
	height: ${(props) => props.height + `px`};
	background: #ffffff 0% 0% no-repeat padding-box;
	border: 1px solid #707070;
	padding: 8px;
	margin: 0 0 24px 0;
`;

export default GraphCard;

GraphCard.PropTypes = {
	width: PropTypes.number,
	height: PropTypes.number,
};

GraphCard.defaultProps = {
	width: 830,
	height: 344,
};
