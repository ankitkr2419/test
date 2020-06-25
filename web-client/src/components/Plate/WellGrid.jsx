import React from 'react';
import PropTypes from "prop-types";
import styled from 'styled-components';

const StyledWellGrid = styled.div`
	display: flex;
	flex-wrap: wrap;
	flex-basis: 0;
	flex-grow: 1;
	min-width: 0;
	max-width: 100%;
	padding: 16px 0 0 16px;
	border: 1px solid #cbcbcb;
	max-width: 894px;
	max-height: 578px;
`;

const WellGrid = ({ className, children }) => {
	return <StyledWellGrid className={className}>{children}</StyledWellGrid>;
};


WellGrid.propTypes = {
	className: PropTypes.string,
};

WellGrid.defaultProps = {
	className: "",
};


export default WellGrid;