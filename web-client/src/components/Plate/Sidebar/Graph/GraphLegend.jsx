import React from 'react';
import PropTypes from 'prop-types';
import styled from 'styled-components';

const StyledGraphLegend = styled.div`
	width: 14px;
	height: 14px;
	border-radius: 50%;
	background-color: ${(props) => props.color};
	margin: 0 8px 0 4px;
`;

const GraphLegend = (props) => {
	return <StyledGraphLegend color={props.color} />;
};

GraphLegend.propTypes = {
	color: PropTypes.string,
};

GraphLegend.defaultProps = {
	color: '#e5e5e5',
};

export default GraphLegend;
