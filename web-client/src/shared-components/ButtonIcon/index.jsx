import React from 'react';
import PropTypes from 'prop-types';
import { Icon } from 'shared-components';
import { StyledButtonIcon } from './StyledButtonIcon';

const ButtonIcon = (props) => (
	<StyledButtonIcon {...props}>
		<Icon size={props.size} name={props.name} />
	</StyledButtonIcon>
);

ButtonIcon.propTypes = {
	position: PropTypes.string,
	placement: PropTypes.oneOf(['left', 'right']),
	top: PropTypes.number,
	right: PropTypes.number,
	left: PropTypes.number,
	isShadow: PropTypes.bool,
	name: PropTypes.string.isRequired,
	size: PropTypes.number,
	id: PropTypes.string,
};

ButtonIcon.defaultProps = {
	isShadow: false,
	size: 24,
};

export default ButtonIcon;
