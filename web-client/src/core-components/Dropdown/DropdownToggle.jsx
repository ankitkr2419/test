import React from 'react';
import PropTypes from 'prop-types';
import styled from "styled-components";
import { DropdownToggle } from 'reactstrap';
import { Icon } from 'shared-components';

const StyledDropdownToggle = styled(DropdownToggle)`
	background-color: ${(props) => (props.icon ? "" : "transparent")};
	border-color: ${(props) => (props.icon ? "" : "transparent")};
	padding: ${(props) => (props.icon ? "6px 12px" : "4px")};

	&:hover,
	&:not(:disabled):not(.disabled):active,
	&:focus,
	&.focus {
		background-color: ${(props) => (props.icon ? "" : "transparent")};
		border-color: ${(props) => (props.icon ? "" : "transparent")};
	}

	&:focus,
	&.focus,
	&:not(:disabled):not(.disabled):active:focus {
		box-shadow: ${(props) => (props.icon ? "" : "none")};
	}

	i {
		color: #71b1a8;
	}
`;

const CustomDropdownToggle = (props) => {
	const { icon, name, size, children, ...rest } = props;
	return (
		<StyledDropdownToggle {...rest}>
			{icon ? <Icon name={name} size={size} /> : children}
		</StyledDropdownToggle>
	);
};

CustomDropdownToggle.propTypes = {
  icon: PropTypes.bool,
	name: PropTypes.string,
	size: PropTypes.number,
};

CustomDropdownToggle.defaultProps = {
  icon: false,
	size: 24,
};

export default CustomDropdownToggle;