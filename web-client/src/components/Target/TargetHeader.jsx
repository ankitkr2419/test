import React from 'react';
import PropTypes from 'prop-types';
import styled from 'styled-components';
import { Text, ButtonIcon } from 'shared-components';

const StyledTargetHeader = styled.header`
	padding: 0 48px 16px 88px;

	h6.text-default {
		line-height: 22px;
	}

	.text-truncate-multi-line {
		height: 50px;
		overflow: hidden;
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
			{isLoginTypeAdmin === true && (
				<div className='d-flex'>
					<ButtonIcon name='pencil' size={28} className='px-0 border-0' />
					<Text
						size={14}
						className='flex-25 text-default text-truncate-multi-line font-weight-light mb-0 pl-3 pr-2 py-1'
					>
						Template Name
					</Text>
					<Text
						size={14}
						className='flex-100 text-default text-truncate-multi-line font-weight-light mb-0 px-2 py-1'
					>
						Lorem ipsum, dolor sit amet consectetur adipisicing elit. Facere,
						quo autem? Veniam eligendi sit cum! Ratione pariatur dolorem impedit
						dolorum perferendis temporibus, quam veritatis ducimus nesciunt,
						nihil blanditiis eius nobis.
					</Text>
				</div>
			)}
		</StyledTargetHeader>
	);
};

TargetHeader.propTypes = {
	isLoginTypeAdmin: PropTypes.bool.isRequired,
	isLoginTypeOperator: PropTypes.bool.isRequired,
};

export default TargetHeader;
