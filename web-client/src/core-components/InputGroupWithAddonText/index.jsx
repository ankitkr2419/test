import React from 'react';
import PropTypes from 'prop-types';
import styled from 'styled-components';
import { Text, Icon } from 'shared-components';

const StyledInputGroupWithAddonText = styled.div`
	position: relative;

	input {
		padding-right: ${(props) => props.paddingRight}px;
	}

	span,
	i {
		position: absolute;
		top: 50%;
		right: 8px;
		transform: translate(0, -50%);
		color: #707070;
		opacity: 0.45;
	}
`;

const InputGroupWithAddonText = (props) => {
	const { addonText, addonIcon, addonSize, paddingRight, children } = props;
	return (
		<StyledInputGroupWithAddonText paddingRight={paddingRight}>
			{children}
			{addonText ? (
				<Text Tag='span' size={addonSize}>
					{addonText}
				</Text>
			) : null}
			{addonIcon ? <Icon name={addonIcon} size={addonSize} /> : null}
		</StyledInputGroupWithAddonText>
	);
};

InputGroupWithAddonText.propTypes = {
	children: PropTypes.element,
	addonText: PropTypes.string,
	addonIcon: PropTypes.string,
	addonSize: PropTypes.number,
	paddingRight: PropTypes.number,
};

InputGroupWithAddonText.defaultProps = {
	paddingRight: 24,
};

export default InputGroupWithAddonText;
