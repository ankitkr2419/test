import React from 'react';
import PropTypes from "prop-types";
import styled from 'styled-components';

const StyledCoordinateItem = styled.li`
	display: flex;
	align-items: center;
	justify-content: center;
	font-size: 18px;
	line-height: 22px;
	color: #999999;
	border: 1px solid transparent;
`;

const CoordinateItem = ({ className, coordinateValue }) => {
	return (
		<StyledCoordinateItem className={`coordinate-item ${className}`}>
			{coordinateValue}
		</StyledCoordinateItem>
	);
};

CoordinateItem.propTypes = {
	className: PropTypes.string,
	coordinateValue: PropTypes.string,
};

CoordinateItem.defaultProps = {
	className: "",
	coordinateValue: "",
};

export default CoordinateItem;