import React from 'react';
import PropTypes from 'prop-types';
import styled from 'styled-components';

const StyledCoordinate = styled.ul`
	display: flex;
	flex-direction: ${(props) =>
		props.direction === 'horizontal' ? '' : 'column'};
	margin: ${(props) => (props.direction === 'horizontal' ? '0 0 0 26px' : '0')};
	padding: ${(props) =>
		props.direction === 'horizontal' ? '0 17px' : '17px 0'};
	list-style: none;
`;

const Coordinate = ({ className, direction, children }) => {
	return (
		<StyledCoordinate
			direction={direction}
			className={`coordinate -${direction} ${className}`}
		>
			{children}
		</StyledCoordinate>
	);
};

Coordinate.propTypes = {
	className: PropTypes.string,
	direction: PropTypes.oneOf(['vertical', 'horizontal']),
};

Coordinate.defaultProps = {
	className: '',
	direction: 'vertical',
};

export default Coordinate;
