import React from 'react';
import PropTypes from 'prop-types';
import styled from 'styled-components';
import { Text } from 'shared-components';

const StyledTargetHeader = styled.header`
	padding: 0 48px 20px 88px;

	h6.text-default {
		line-height: 22px;
	}
`;

const TargetHeader = (props) => {
	const { isLoginTypeAdmin, isLoginTypeOperator } = props;
	return (
		<StyledTargetHeader className='target-header'>
			{isLoginTypeOperator === true && (
				<Text
					Tag='h6'
					size={18}
					className='text-default font-weight-light mb-0'
				>
					Template Name
				</Text>
			)}
			{isLoginTypeAdmin === true && ''}
		</StyledTargetHeader>
	);
};

TargetHeader.propTypes = {
	isLoginTypeAdmin: PropTypes.bool.isRequired,
	isLoginTypeOperator: PropTypes.bool.isRequired,
};

export default TargetHeader;
