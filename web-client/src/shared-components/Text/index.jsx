import React from 'react';
import PropTypes from 'prop-types';
import styled from 'styled-components';

/**
 * Text component can be used for multi purpose by passing tag as prop
 * e.g tag can be span.
 * @param {*} props
 */

const StyledText = ({ Tag, onClick, className, size, children }) => (
	<Tag size={size} onClick={onClick} className={className}>
		{children}
	</Tag>
);

const Text = styled(StyledText)`
	font-size: ${(props) => props.size}px;
`;

Text.propTypes = {
	Tag: PropTypes.oneOf([
		'h1',
		'h2',
		'h3',
		'h4',
		'h5',
		'h6',
		'p',
		'span',
		'label',
	]),
	className: PropTypes.string,
	onClick: PropTypes.func,
	children: PropTypes.any,
	size: PropTypes.number,
};

Text.defaultProps = {
	Tag: 'p',
	size: 16,
};

export default Text;
