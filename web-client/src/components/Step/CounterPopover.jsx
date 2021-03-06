/* eslint-disable no-undef */
// Added the above line to avoid the document undefine error
//  in simulateOutSideClick function
import React from 'react';
import PropTypes from 'prop-types';
import {
	Popover,
	PopoverBody,
	Input,
	Button,
	FormGroup,
} from 'core-components';
import { ButtonIcon, Text, Icon } from 'shared-components';
import { validateRepeatCount } from 'components/Stage/stageHelper';
import { MIN_REPEAT_COUNT, MAX_REPEAT_COUNT } from 'components/Stage/stageConstants';
import { StyledCounterPopover } from './StyledCounterPopover';

const CounterPopover = (props) => {
	const {
		className,
		cycleRepeatCount,
		repeatCounterState,
		updateRepeatCounterStateWrapper,
		saveRepeatCount,
		...rest
	} = props;

	const {
		repeatCount,
		repeatCountError,
	} = repeatCounterState;

	// helper function to close the popover
	const simulateOutSideClick = () => document.body.click();

	// close repeat count popver
	const closeCounterPopover = () => {
		simulateOutSideClick();
		// set the repeat count stored in local state with repeat count value over server
		updateRepeatCounterStateWrapper('repeatCount', cycleRepeatCount);
		// reset repeat count error flag
		updateRepeatCounterStateWrapper('repeatCountError', false);
	};

	const onSaveClickHandler = () => {
		// call save the repeat count helper function only if there is a change
		if (parseInt(repeatCount, 10) !== cycleRepeatCount) {
			saveRepeatCount(repeatCount);
		}
		// close the popover
		simulateOutSideClick();
	};

	// repeat count change handler
	const onRepeatCountChangeHandler = ({ target: { name, value } }) => {
		updateRepeatCounterStateWrapper(name, value);
	};

	// reset repeatCountError to false stored in repeatCounter local state
	const onRepeatCountFocusHandler = () => {
		updateRepeatCounterStateWrapper('repeatCountError', false);
	};

	// set repeatCountError true stored in repeatCounter local state
	const onRepeatCountBlurHandler = () => {
		if (validateRepeatCount(repeatCount) === false) {
			updateRepeatCounterStateWrapper('repeatCountError', true);
		}
	};

	return (
		<StyledCounterPopover className={className} {...rest}>
			<ButtonIcon size={24} name='pencil' id='PopoverCounter' />
			<Text size={16} className='d-flex align-items-center mb-0 mx-1 p-1'>
				{cycleRepeatCount}
			</Text>
			<Popover
				trigger='legacy'
				target='PopoverCounter'
				placement='top'
				popperClassName='popover-counter'
			>
				<PopoverBody className='d-flex'>
					<FormGroup className='mb-0'>
						<Input
							type='number'
							id='counter'
							name='repeatCount'
							placeholder='Enter Count'
							className='flex-100'
							value={repeatCount}
							onChange={onRepeatCountChangeHandler}
							invalid={repeatCountError}
							onFocus={onRepeatCountFocusHandler}
							onBlur={onRepeatCountBlurHandler}
						/>
						<Text Tag='p' size={11} className={`${repeatCountError && 'text-danger'} px-2 mb-0`}>
								Enter count between {MIN_REPEAT_COUNT} to {MAX_REPEAT_COUNT}
						</Text>
					</FormGroup>
					<Button
						icon
						outline
						color={`${repeatCountError === true ? 'secondary' : 'primary'}`}
						className='rounded-circle ml-3'
						disabled={repeatCountError === true}
						onClick={onSaveClickHandler}
					>
						<Icon size={32} name='check' />
					</Button>
					<Button
						icon
						outline
						color='secondary'
						className='rounded-circle ml-3'
						onClick={closeCounterPopover}
					>
						<Icon size={32} name='cross' />
					</Button>
				</PopoverBody>
			</Popover>
		</StyledCounterPopover>
	);
};

CounterPopover.propTypes = {
	cycleRepeatCount: PropTypes.number.isRequired,
	repeatCounterState: PropTypes.object.isRequired,
	updateRepeatCounterStateWrapper: PropTypes.func.isRequired,
	saveRepeatCount: PropTypes.func.isRequired,
};

export default CounterPopover;
