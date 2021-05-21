import React from 'react';
import PropTypes from 'prop-types';
import Icon from 'shared-components/Icon';
import { Shadow, StyledSidebarHandle } from './Styles';

const SidebarHandle = ({
	icon, size, clickHandler, isDisabled,
}) => (
	<StyledSidebarHandle
		onClick={clickHandler}
		className='sidebar-handle'
		icon={icon}
		size={size}
		disabled={isDisabled}
	>
		<Icon name={icon} size={size} />
		<Shadow />
	</StyledSidebarHandle>
);

SidebarHandle.propTypes = {
	icon: PropTypes.string.isRequired,
	size: PropTypes.number,
	clickHandler: PropTypes.func.isRequired,
};

SidebarHandle.defaultProps = {
	size: 24,
};

export default SidebarHandle;
