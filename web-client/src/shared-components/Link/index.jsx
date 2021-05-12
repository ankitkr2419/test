import React from 'react';
import PropTypes from 'prop-types';
import { StyledLink } from './StyledLink';

const CustomLink = (props) => {
	const { icon, ...rest } = props;
	return (
		<StyledLink icon={icon.toString()} {...rest}>
			{props.children}
		</StyledLink>
	);
};

CustomLink.propTypes = {
	icon: PropTypes.bool,
};

CustomLink.defaultProps = {
	icon: false,
};

export default CustomLink;
