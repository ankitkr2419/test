import React from 'react';
import { Popover, PopoverBody, Input, Button } from 'core-components';
import { ButtonIcon, Text, Icon } from 'shared-components';
import styled from 'styled-components';

const StyledCounterPopover = styled.div`
	display: flex;
`;

const CounterPopover = ({ className, ...rest }) => {
	return (
		<StyledCounterPopover className={className} {...rest}>
			<ButtonIcon size={24} name='pencil' id='PopoverCounter' />
			<Text size={16} className='d-flex align-items-center mb-0 mx-1 p-1'>
				50
			</Text>
			<Popover
				trigger='legacy'
				target='PopoverCounter'
				placement='top'
				popperClassName='popover-counter'
			>
				<PopoverBody className='d-flex'>
					<Input
						type='number'
						id='counter'
						name='counter'
						className='flex-100'
					/>
					<Button icon color='primary' className='rounded-circle ml-3'>
						<Icon size={32} name='check' />
					</Button>
					<Button icon color='secondary' className='rounded-circle ml-3'>
						<Icon size={32} name='cross' />
					</Button>
				</PopoverBody>
			</Popover>
		</StyledCounterPopover>
	);
};

CounterPopover.propTypes = {};

export default CounterPopover;
