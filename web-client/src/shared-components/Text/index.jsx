import React from 'react';
import PropTypes from 'prop-types';

/**
 * Text component can be used for multi purpose by passing tag as prop
 * e.g tag can be span.
 * @param {*} props
 */

const Text = ({
	Tag, onClick, className, children,
}) => (
	<Tag onClick={onClick} className={className}>
		{children}
	</Tag>
);

Text.propTypes = {
	Tag: PropTypes.string,
	className: PropTypes.string,
	onClick: PropTypes.func,
	children: PropTypes.oneOfType([PropTypes.element, PropTypes.string]),
};

Text.defaultProps = {
	Tag: 'p',
};

export default Text;
