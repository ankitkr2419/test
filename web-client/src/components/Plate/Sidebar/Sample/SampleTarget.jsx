import React from 'react';
import PropTypes from 'prop-types';
import styled from 'styled-components';
import { ButtonIcon, Text } from 'shared-components';

const StyledSampleTarget = styled.div`
	width: 220px;
	height: 42px;
	display: flex;
	align-items: center;
	background: #ffffff 0% 0% no-repeat padding-box;
	font-size: 18px;
	line-height: 22px;
	color: #707070;
	border: 1px solid #e5e5e5;
	border-radius: 8px;
	box-shadow: 0 3px 6px #0000000b;
	margin: 0 auto 8px;
	padding: 1px;
	opacity: ${(props) => (props.isSelected ? '1' : '0.5')};

	button {
		color: #999999;
	}
`;

const SampleTarget = ({ label, isSelected, onClickHandler }) => (
	<StyledSampleTarget onClick={onClickHandler} isSelected={isSelected}>
		<Text className='m-0 px-3'>{label}</Text>
		{isSelected ? (
			<ButtonIcon name='cross' size={28} className='ml-auto' />
		) : null}
	</StyledSampleTarget>
);

SampleTarget.propTypes = {
	label: PropTypes.string.isRequired,
	onClickHandler: PropTypes.func.isRequired,
	isSelected: PropTypes.bool,
};

SampleTarget.defaultProps = {
	isSelected: false,
};

export default SampleTarget;
