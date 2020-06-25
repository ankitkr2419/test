import React from 'react';
import styled, { css } from 'styled-components';
import PropTypes from 'prop-types';
import { Button } from 'reactstrap';
import classNames from 'classnames';

const StyledButton = styled(Button)`
	width: ${(props) => (props.size === "sm" ? "" : "202px")};
	min-width: ${(props) => (props.size === "sm" ? "84px" : "")};
	height: ${(props) => (props.size === "sm" ? "32px" : "40px")};
	font-size: 16px;
	line-height: 19px;
	font-weight: ${(props) => (props.outline ? "normal" : "bold")};
	padding: ${(props) => (props.size === "sm" ? "4px 24px" : "10px 20px")};
	border-radius: ${(props) => (props.icon === "true" ? "8px" : "27px")};
	box-shadow: ${(props) => (props.outline ? "none" : "0 2px 6px #00000029")};
	border-width: ${(props) => (props.outline ? "2px" : "")};

	${(props) =>
		props.icon === "true" &&
		css`
			display: flex;
			align-items: center;
			justify-content: center;
			width: 50px;
			padding: 0 4px;
		`};
`;

const CustomButtom = (props) => {
	const classes = classNames(props.className, { active: props.active });

	return (
		<StyledButton
			{...props}
			className={classes}
			icon={props.icon.toString()}
		>
			{props.children}
		</StyledButton>
	);
};

CustomButtom.propTypes = {
	icon: PropTypes.bool,
	active: PropTypes.bool,
};

CustomButtom.defaultProps = {
	icon: false,
	active: false,
};

export default CustomButtom;
